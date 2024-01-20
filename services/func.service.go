package services

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/configs"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

func generatePairToken(id uint, username string, role string, profilePic string, remember bool) (domains.TokenPair, error) {
	var accessExpire *jwt.NumericDate
	var refreshExpire *jwt.NumericDate
	result := domains.TokenPair{}

	accessExpire = helpers.GenerateExpireTime(configs.ACCESS_TOKEN_EXPIRE_TIME)
	if !remember {
		refreshExpire = helpers.GenerateExpireTime(configs.REFRESH_TOKEN_EXPIRE_TIME)
	}

	accessClaims := helpers.CreateJWTClaims(id, &username, &role, &profilePic, accessExpire)
	refreshClaims := helpers.CreateJWTClaims(id, nil, nil, nil, refreshExpire)

	accessToken, err := helpers.GenerateToken(accessClaims)
	if err != nil {
		return result, err
	}
	refreshToken, err := helpers.GenerateToken(refreshClaims)
	if err != nil {
		return result, err
	}
	result.Access = accessToken
	result.Refresh = refreshToken
	return result, nil
}

type createCallRepoFunc func() (domains.Model, error)

func basicCreateService(strict bool, c context.Context, data any, model domains.Model, entity domains.Entity, fun createCallRepoFunc) error {
	user, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if strict && !ok {
		return domains.ErrInvalidAccess
	}
	if ok, err := helpers.Validate(data); !ok {
		return err
	}
	if err := helpers.Convert(data, model); err != nil {
		return err
	}
	if ok {
		valid := model.GetPolicy().Create(user, model)
		if !valid {
			return domains.ErrInvalidAccess
		}
	}
	result, err := fun()
	if err != nil {
		return err
	}
	if err := helpers.Convert(result, entity); err != nil {
		return err
	}
	return nil
}

type updateCallRepoFunc func(id uint) (int, domains.Model, error)

func basicUpdateService(strict bool, c context.Context, id uint, data any, model domains.Model, entity domains.Entity, fun updateCallRepoFunc) (int, error) {
	user, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if strict && !ok {
		return 0, domains.ErrInvalidAccess
	}
	if ok, err := helpers.Validate(data); !ok {
		return 0, err
	}
	if err := helpers.Convert(data, model); err != nil {
		return 0, err
	}
	if ok {
		valid := model.GetPolicy().Update(id, user, model)
		if !valid {
			return 0, domains.ErrInvalidAccess
		}
	}
	aff, result, err := fun(id)
	if err != nil {
		return 0, err
	}
	if err := helpers.Convert(result, entity); err != nil {
		return 0, err
	}
	return int(aff), nil
}

type findCallRepoFunc func(id uint) (domains.Model, error)

func basicFindService(strict bool, c context.Context, id uint, model domains.Model, entity domains.Entity, fun findCallRepoFunc) error {
	user, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if strict && !ok {
		return domains.ErrInvalidAccess
	}
	result, err := fun(id)
	if err != nil {
		return err
	}
	if ok {
		valid := model.GetPolicy().Find(id, user, result)
		if !valid {
			return domains.ErrInvalidAccess
		}
	}
	if err := helpers.Convert(result, entity); err != nil {
		return err
	}
	return nil
}

type findHandlerFunc func() (any, error)
type readHandlerFunc func() (any, uint, error)
type convertHandlerFunc func(datas any) error

func basicReadService(
	strict bool,
	c context.Context,
	id uint,
	q string,
	page uint,
	orderBy string,
	isDesc bool,
	withPagination bool,
	entity domains.Entity,
	findFun findHandlerFunc,
	readFun readHandlerFunc,
	conFun convertHandlerFunc,
	model domains.Model,
) (*domains.Pagination, error) {
	var pagination domains.Pagination
	user, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if strict && !ok {
		return nil, domains.ErrInvalidAccess
	}
	if id != 0 {
		result, err := findFun()
		valid := model.GetPolicy().Find(id, user, result)
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, domains.ErrInvalidAccess
		}
		err = helpers.Convert(result, entity)
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	datas, maxPage, err := readFun()
	if ok {
		valid := model.GetPolicy().ReadAll(user, datas)
		if !valid {
			return nil, domains.ErrInvalidAccess
		}
	}
	if err != nil {
		return nil, err
	}
	err = conFun(datas)
	if err != nil {
		return nil, err
	}
	if withPagination {
		queries := helpers.MakeDefaultGetPaginationQueries(q, id, page, orderBy, isDesc, withPagination)
		pagination = helpers.MakePagination(maxPage, uint(page), queries)
		return &pagination, nil
	}
	return nil, nil
}

type deleteCallRepoFunc func() (int, int, error)

func basicDeleteService(strict bool, c context.Context, id uint, fun deleteCallRepoFunc, model domains.Model) (int, int, error) {
	user, ok := c.Value(configs.UserKey).(domains.UserEntity)
	if strict && !ok {
		return 0, 0, domains.ErrInvalidAccess
	}
	idResult, aff, err := fun()
	valid := model.GetPolicy().Delete(id, user, nil)
	if !valid {
		return 0, 0, domains.ErrInvalidAccess
	}
	return idResult, int(aff), err
}
