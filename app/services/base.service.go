package services

import (
	"context"
	"fmt"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type beforeFunc func() error

func baseCreateFunc(
	ctx context.Context,
	policy policies.Policy,
	data domains.ModelInterface,
	form any,
	beforeCreate ...beforeFunc,
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policy.Create(claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	data.SetUpdatedEmail(claims.Email)
	if err := helpers.Validator.Validate(form); err != nil {
		return err
	}
	if err := helpers.Translator.TranslateStruct(form, data); err != nil {
		return err
	}
	data.SetUpdatedEmail(claims.Email)
	for _, fun := range beforeCreate {
		err := fun()
		if err != nil {
			return err
		}
	}
	err := providers.RepoProvider.BaseRepo.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func baseFindFunc(
	ctx context.Context,
	policy policies.Policy,
	id uint,
	result domains.ModelInterface,
	preloads ...string,
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, result, preloads...)
	if !policy.Find(result, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	if err != nil {
		return err
	}
	return nil
}

func baseReadFunc(
	ctx context.Context,
	policy policies.Policy,
	params types.DBSearchParams,
	results any,
) (*types.Pagination, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policy.ReadAll(claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return nil, types.ErrForbiden
	}

	pagination, err := providers.RepoProvider.BaseRepo.ReadAll(
		ctx,
		results,
		params,
	)
	if err != nil {
		return nil, err
	}
	return pagination, nil
}

func baseDeleteFunc(
	ctx context.Context,
	policy policies.Policy,
	id uint,
	model domains.ModelInterface,
	beforeDelete ...beforeFunc,
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, model)
	if err != nil {
		return err
	}
	if !policy.Delete(model, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	for _, fun := range beforeDelete {
		err := fun()
		if err != nil {
			return err
		}
	}
	err = providers.RepoProvider.BaseRepo.Delete(ctx, []uint{id}, model)
	if err != nil {
		return err
	}
	return nil
}

func baseUpdateFunc(
	ctx context.Context,
	policy policies.Policy,
	id uint,
	data domains.ModelInterface,
	form any,
	beforeUpdate ...beforeFunc,
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, data)
	if err != nil {
		return err
	}
	if !policy.Update(data, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	data.SetUpdatedEmail(claims.Email)
	if err := helpers.Validator.Validate(form); err != nil {
		return err
	}
	if err := helpers.Translator.TranslateStruct(form, data); err != nil {
		return err
	}
	for _, fun := range beforeUpdate {
		err := fun()
		if err != nil {
			return err
		}
	}
	err = providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, data)
	if err != nil {
		return err
	}
	data.SetID(id)
	return nil
}
