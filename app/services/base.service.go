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

func baseCreateFunc(
	ctx context.Context,
	policy policies.Policy,
	data domains.ModelInterface,
	form any,
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policy.Create(claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	if err := helpers.Validator.Validate(form); err != nil {
		return err
	}
	if err := helpers.Translator.TranslateStruct(form, data); err != nil {
		return err
	}
	data.SetUpdatedEmail(claims.Email)
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
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policy.Find(id, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	err := providers.RepoProvider.BaseRepo.Find(ctx, id, result)
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
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policies.MasterPolicy.Delete(id, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	err := providers.RepoProvider.BaseRepo.Delete(ctx, []uint{id}, model)
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
) error {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policy.Update(id, claims) {
		helpers.Logger.Warning(
			fmt.Sprintf("(Forbidden) User: %s", claims.Email),
		)
		return types.ErrForbiden
	}
	if err := helpers.Validator.Validate(form); err != nil {
		return err
	}
	if err := helpers.Translator.TranslateStruct(form, data); err != nil {
		return err
	}
	err := providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, data)
	if err != nil {
		return err
	}
	data.SetID(id)
	return nil
}
