package services

import (
	"context"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

// province master
type masterProvinceService struct{}

func NewMasterProvinceService() service_domains.MasterProvinceServiceInterface {
	return &masterProvinceService{}
}

func (masterProvinceService) Create(ctx context.Context,
	data forms.MasterProviceCreateForm,
) (models.Province, error) {
	var province models.Province
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &province, data)
	return province, err
}
func (masterProvinceService) Read(ctx context.Context,
	q string, page uint,
) ([]models.Province, *types.Pagination, error) {
	var results []models.Province
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "province", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Province{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (masterProvinceService) Find(ctx context.Context, id uint) (models.Province, error) {
	var province models.Province
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &province)
	return province, err
}
func (masterProvinceService) Update(ctx context.Context,
	id uint, data forms.MasterProviceUpdateForm,
) (uint, models.Province, error) {
	var province models.Province
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &province, data)
	return id, province, err
}
func (masterProvinceService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Province{})
	return id, err
}

// end of province master

// regency master
type regencyMasterService struct{}

func NewRegencyMasterService() service_domains.MasterRegencyServiceInterface {
	return &regencyMasterService{}
}
func (regencyMasterService) Create(
	ctx context.Context,
	data forms.MasterRegencyCreateForm,
) (models.Regency, error) {
	var regency models.Regency
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &regency, data)
	return regency, err
}
func (regencyMasterService) Read(ctx context.Context,
	q string,
	page uint,
) ([]models.Regency, *types.Pagination, error) {
	var results []models.Regency
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "regency", Operator: "LIKE", Str: q},
		},
		Query: q,
		Model: &models.Regency{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (regencyMasterService) Find(ctx context.Context, id uint) (models.Regency, error) {
	var regency models.Regency
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &regency)
	return regency, err
}
func (regencyMasterService) Update(ctx context.Context,
	id uint,
	data forms.MasterRegencyUpdateForm,
) (uint, models.Regency, error) {
	var regency models.Regency
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &regency, data)
	return id, regency, err
}
func (regencyMasterService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Regency{})
	return id, err
}

// end of regency master

// category master
type categoryMasterService struct{}

func NewCategoryMasterService() service_domains.MasterCategoryServiceInterface {
	return &categoryMasterService{}
}
func (categoryMasterService) Create(
	ctx context.Context,
	data forms.MasterCategoryCreateForm,
) (models.Category, error) {

	var category models.Category
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &category, data)
	return category, err
}
func (categoryMasterService) Read(ctx context.Context, q string, page uint) ([]models.Category, *types.Pagination, error) {
	var results []models.Category
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "category_name", Operator: "LIKE", Str: q},
			{Field: "description", Operator: "LIKE", Str: q, IsOr: true},
		},
		Query: q,
		Model: &models.Category{},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.MasterPolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (categoryMasterService) Find(ctx context.Context, id uint) (models.Category, error) {
	var category models.Category
	err := baseFindFunc(ctx, policies.MasterPolicy{}, id, &category)
	return category, err
}
func (categoryMasterService) Update(ctx context.Context, id uint, data forms.MasterCategoryUpdateForm) (uint, models.Category, error) {
	var category models.Category
	err := baseUpdateFunc(ctx, policies.MasterPolicy{}, id, &category, data)
	return id, category, err
}
func (categoryMasterService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.Category{})
	return id, err
}

// end of category master
