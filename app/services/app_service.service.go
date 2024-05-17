package services

import (
	"context"
	"fmt"

	service_domains "github.com/salamanderman234/outsourcing-api/app/domains/services"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/models"
	"github.com/salamanderman234/outsourcing-api/app/policies"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
)

// application service
type applicationServiceService struct{}

func NewApplicationServiceService() service_domains.ApplicationServiceServiceInterface {
	return &applicationServiceService{}
}

func (applicationServiceService) Create(ctx context.Context,
	data forms.ServiceCreateForm) (models.Service, error) {

	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	if !policies.ServicePolicy.Create(claims) {
		helpers.Logger.Warning(fmt.Sprintf("(Forbidden) User: %s", claims.Email))
		return models.Service{}, types.ErrForbiden
	}
	if err := helpers.Validator.Validate(data); err != nil {
		return models.Service{}, err
	}
	var service models.Service
	if err := helpers.Translator.TranslateStruct(data, &service); err != nil {
		return service, err
	}

	for index, item := range service.RequiredItems {
		if item.Quantity == nil {
			zero := uint(0)
			item.Quantity = &zero
		}
		if item.PricePerItem == nil {
			zero := uint(0)
			item.PricePerItem = &zero
		}
		if service.EtcPrice == nil {
			zero := uint(0)
			service.EtcPrice = &zero
		}
		subTotal := (*item.Quantity) * (*item.PricePerItem)
		service.RequiredItems[index].Subtotal = &subTotal
		subTotal += *service.EtcPrice
		service.EtcPrice = &subTotal
	}
	if service.EtcPrice == nil {
		zero := uint(0)
		service.EtcPrice = &zero
	}
	if service.Discount == nil {
		zero := uint(0)
		service.Discount = &zero
	}
	totalPrice := (*service.EtcPrice) + (*service.EmployeePrice) + (*service.ServicePrice)
	totalPrice -= (*service.Discount)
	service.TotalPrice = &totalPrice

	services := []models.Service{
		service,
	}
	err := providers.RepoProvider.BaseRepo.Create(ctx, services)
	if err != nil {
		return service, err
	}
	service = services[0]
	return service, nil
}
func (applicationServiceService) Read(
	ctx context.Context,
	q string,
	page uint,
) ([]models.Service, *types.Pagination, error) {
	var results []models.Service
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "service_name", Operator: "LIKE", Str: q},
			{Field: "description", Operator: "LIKE", Str: q, IsOr: true},
			{Field: "includes", Operator: "LIKE", Str: q, IsOr: true},
		},
		Query:    q,
		Model:    &models.Service{},
		Preloads: []string{"Category", "AdditionalItems", "RequiredItems"},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.ServicePolicy,
		params,
		&results,
	)
	return results, pagination, err
}
func (applicationServiceService) Find(ctx context.Context, id uint) (models.Service, error) {
	var service models.Service
	err := baseFindFunc(ctx, &policies.ServicePolicy, id, &service)
	return service, err
}
func (applicationServiceService) Update(
	ctx context.Context,
	id uint,
	data forms.ServiceUpdateForm,
) (uint, models.Service, error) {

	var service models.Service
	err := baseUpdateFunc(ctx, &policies.ServicePolicy, id, &service, data)
	return id, service, err
}
func (applicationServiceService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.ServicePolicy, id, &models.Service{})
	return id, err
}

// end of application service
// start package service
type applicationPackageService struct{}

func NewApplicationPackageService() service_domains.ApplicationPackageInterface {
	return &applicationPackageService{}
}
func (applicationPackageService) Create(ctx context.Context,
	data forms.PackageServiceCreateForm) (models.Package, error) {

	return models.Package{}, nil
}
func (applicationPackageService) Read(ctx context.Context,
	q string, page uint,
) ([]models.Package, *types.Pagination, error) {

	var results []models.Package
	params := types.DBSearchParams{
		Page:           page,
		WithPagination: page > 0,
		Params: []types.WhereQuery{
			{Field: "package_name", Operator: "LIKE", Str: q},
			{Field: "description", Operator: "LIKE", Str: q},
			{Field: "includes", Operator: "LIKE", Str: q},
		},
		Query:    q,
		Model:    &models.Package{},
		Preloads: []string{"Services"},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.ServicePolicy,
		params,
		&results,
	)
	return results, pagination, err
}
func (applicationPackageService) Find(ctx context.Context, id uint) (models.Package, error) {
	var pack models.Package
	err := baseFindFunc(ctx, &policies.ServicePolicy, id, &pack)
	return pack, err
}
func (applicationPackageService) Update(
	ctx context.Context,
	id uint,
	data forms.PackageUpdateForm,
) (uint, models.Package, error) {
	var pack models.Package
	err := baseUpdateFunc(ctx, &policies.ServicePolicy, id, &pack, data)
	return id, pack, err
}
func (applicationPackageService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.ServicePolicy, id, &models.Package{})
	return id, err
}

// end of package service
