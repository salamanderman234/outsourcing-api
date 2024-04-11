package services

import (
	"context"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	auth_types "github.com/salamanderman234/outsourcing-api/app/domains/types/auth"
	database_types "github.com/salamanderman234/outsourcing-api/app/domains/types/databases"
	custom_errors "github.com/salamanderman234/outsourcing-api/app/domains/types/errors"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/configs"
)

type masterProvinceService struct{}

func NewMasterProvinceService() service_domains.MasterProvinceServiceInterface {
	return &masterProvinceService{}
}

func (masterProvinceService) Create(ctx context.Context,
	data forms.MasterProviceCreateForm,
) (models.Province, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.MasterPolicy.Create(claims) {
		return models.Province{}, custom_errors.ErrForbiden
	}
	if err := helpers.Validator.Validate(data); err != nil {
		return models.Province{}, err
	}
	var province models.Province
	if err := helpers.Translator.TranslateStruct(data, &province); err != nil {
		return models.Province{}, err
	}
	provinces := []models.Province{
		province,
	}
	err := domains.RepoRegistry.BaseRepo.Create(ctx, provinces)
	if err != nil {
		return models.Province{}, err
	}
	province = provinces[0]
	return province, nil
}
func (masterProvinceService) Read(ctx context.Context,
	q string, page uint,
) ([]models.Province, *responses.Pagination, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.MasterPolicy.ReadAll(claims) {
		return []models.Province{}, &responses.Pagination{}, custom_errors.ErrForbiden
	}
	config := database_types.DBSearchConfig{
		Query:          q,
		Page:           page,
		WithPagination: page > 0,
		Params: []database_types.DBSearchParam{
			{Field: "province", Operator: "LIKE"},
		},
		Model: &models.Province{},
	}
	var finalResults []models.Province
	pagination, err := domains.RepoRegistry.BaseRepo.ReadAll(
		ctx,
		&finalResults,
		config,
	)
	if err != nil {
		return finalResults, pagination, err
	}

	return finalResults, pagination, nil
}
func (masterProvinceService) Find(ctx context.Context, id uint) (models.Province, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.MasterPolicy.Find(id, claims) {
		return models.Province{}, custom_errors.ErrForbiden
	}
	var finalResult models.Province
	err := domains.RepoRegistry.BaseRepo.Find(ctx, id, &finalResult)
	if err != nil {
		return models.Province{}, err
	}
	return finalResult, nil
}
func (masterProvinceService) Update(ctx context.Context,
	id uint, data forms.MasterProviceUpdateForm,
) (uint, models.Province, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.MasterPolicy.Update(id, claims) {
		return id, models.Province{}, custom_errors.ErrForbiden
	}
	if err := helpers.Validator.Validate(data); err != nil {
		return id, models.Province{}, err
	}
	var province models.Province
	if err := helpers.Translator.TranslateStruct(data, &province); err != nil {
		return id, province, err
	}
	err := domains.RepoRegistry.BaseRepo.Update(ctx, []uint{id}, province)
	if err != nil {
		return id, province, err
	}
	province.ID = id
	return id, province, nil
}
func (masterProvinceService) Delete(ctx context.Context, id uint) (uint, error) {
	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(auth_types.JWTCLaims)
	if !policies.MasterPolicy.Delete(id, claims) {
		return id, custom_errors.ErrForbiden
	}
	err := domains.RepoRegistry.BaseRepo.Delete(ctx, []uint{id}, models.Province{})
	if err != nil {
		return id, err
	}
	return id, nil
}
