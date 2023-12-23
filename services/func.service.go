package services

import (
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

func basicCreateService(data any, model domains.Model, entity domains.Entity, fun createCallRepoFunc) (domains.Entity, error) {
	if ok, err := helpers.Validate(data); !ok {
		valErr := domains.ErrValidation
		valErr.ValidationErrors = err
		return nil, valErr
	}
	if err := helpers.Convert(data, model); err != nil {
		return nil, err
	}
	result, err := fun()
	if err != nil {
		return nil, err
	}
	if err := helpers.Convert(result, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

type updateCallRepoFunc func(id uint) (int, domains.Model, error)

func basicUpdateService(id uint, data any, model domains.Model, entity domains.Entity, fun updateCallRepoFunc) (int, any, error) {
	if ok, err := helpers.Validate(data); !ok {
		valErr := domains.ErrValidation
		valErr.ValidationErrors = err
		return 0, nil, valErr
	}
	if err := helpers.Convert(data, model); err != nil {
		return 0, nil, err
	}
	aff, result, err := fun(id)
	if err != nil {
		return 0, nil, err
	}
	if err := helpers.Convert(result, entity); err != nil {
		return 0, nil, err
	}
	return int(aff), data, nil
}
