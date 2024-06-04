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

func (applicationServiceService) AddRequiredItems(
	ctx context.Context,
	form forms.RequiredItemAddForm,
) (models.RequiredItemService, error) {
	var requiredItem models.RequiredItemService
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &requiredItem, form)
	return requiredItem, err
}

func (applicationServiceService) AddAdditionalItems(
	ctx context.Context,
	form forms.AdditionalItemServiceAddForm,
) (models.AdditionalItemService, error) {
	var additioanlItem models.AdditionalItemService
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &additioanlItem, form)
	return additioanlItem, err
}

func (applicationServiceService) Create(ctx context.Context,
	data forms.ServiceCreateForm) (models.Service, error) {

	claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
	policy := policies.ServicePolicy{}.Create(claims)
	if !policy {
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

	// file
	mainImage := data.MainImage
	if mainImage != "" {
		res, err := providers.ResourceProvider.CreateResource("services.main_image")
		if err != nil {
			return service, err
		}
		result, err := providers.ServiceProvider.FileService.UploadFile(ctx, mainImage, res)
		if err != nil {
			return service, err
		}
		service.MainImage = &result
	}
	icon := data.Icon
	if icon != "" {
		res, err := providers.ResourceProvider.CreateResource("services.icon")
		if err != nil {
			return service, err
		}
		result, err := providers.ServiceProvider.FileService.UploadFile(ctx, icon, res)
		if err != nil {
			return service, err
		}
		service.Icon = &result
	}

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
		policies.ServicePolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (applicationServiceService) Find(ctx context.Context, id uint) (models.Service, error) {
	var service models.Service
	err := baseFindFunc(ctx, policies.ServicePolicy{}, id, &service, "RequiredItems", "AdditionalItems", "Category")
	return service, err
}
func (applicationServiceService) Update(
	ctx context.Context,
	id uint,
	data forms.ServiceUpdateForm,
) (uint, models.Service, error) {

	var service models.Service
	err := baseUpdateFunc(ctx, policies.ServicePolicy{}, id, &service, data)
	return id, service, err
}
func (applicationServiceService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.ServicePolicy{}, id, &models.Service{})
	return id, err
}

// end of application service
// start package service
type applicationPackageService struct{}

func NewApplicationPackageService() service_domains.ApplicationPackageInterface {
	return &applicationPackageService{}
}
func (applicationPackageService) Create(ctx context.Context,
	data forms.PackageCreateForm) (models.Package, error) {

	var service models.Package
	before := func() error {
		etcP := uint(0)
		empP := uint(0)
		servP := uint(0)
		for index, item := range service.Services {
			serviceID := item.ServiceID
			var serv models.Service
			err := providers.RepoProvider.BaseRepo.Find(ctx, *serviceID, &serv, "RequiredItems", "AdditionalItems", "Category")
			if err != nil {
				return err
			}
			service.Services[index].ServicePrice = serv.ServicePrice
			empPrice := (*serv.EmployeePrice) * (*item.TotalEmployee)
			service.Services[index].EmployeePrice = &empPrice
			service.Services[index].ServicePrice = serv.ServicePrice
			servP += (*serv.ServicePrice)
			empP += empPrice
			etcPrice := *(serv.EtcPrice)

			additionals := item.AdditionalPackageServiceItems

			for y, additional := range additionals {
				var add models.AdditionalItemService
				err := providers.RepoProvider.BaseRepo.Find(ctx, *additional.AdditionalItemServiceID, &add)
				if err != nil {
					return err
				}
				sub := (*additional.Quantity) * (*add.PricePerItem)
				etcPrice += sub
				service.Services[index].AdditionalPackageServiceItems[y].Price = add.PricePerItem
				service.Services[index].AdditionalPackageServiceItems[y].SubTotalPrice = &sub
			}
			service.Services[index].EtcPrice = &etcPrice
			etcP += etcPrice
			sub := (*service.Services[index].EtcPrice) + (*service.Services[index].EmployeePrice) + (*service.Services[index].ServicePrice)
			service.Services[index].SubTotalPrice = &sub

		}
		if service.EtcPrice == nil {
			zero := uint(0)
			service.EtcPrice = &zero
		}
		if service.Discount == nil {
			zero := uint(0)
			service.Discount = &zero
		}
		service.EtcPrice = &etcP
		service.EmployeePrice = &empP
		service.ServicePrice = &servP

		totalPrice := (*service.EtcPrice) + (*service.EmployeePrice) + (*service.ServicePrice)
		totalPrice -= (*service.Discount)
		service.TotalPrice = &totalPrice
		return nil
	}
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &service, data, before)
	return service, err
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
		Query: q,
		Model: &models.Package{},
		Preloads: []string{
			"Services",
			"Services.AdditionalPackageServiceItems",
			"Services.Service",
			"Services.AdditionalPackageServiceItems.AdditionalItemService",
		},
	}

	pagination, err := baseReadFunc(
		ctx,
		policies.ServicePolicy{},
		params,
		&results,
	)
	return results, pagination, err
}
func (applicationPackageService) Find(ctx context.Context, id uint) (models.Package, error) {
	var pack models.Package
	err := baseFindFunc(ctx, &policies.ServicePolicy{}, id, &pack,
		"Services",
		"Services.AdditionalPackageServiceItems",
		"Services.Service",
		"Services.AdditionalPackageServiceItems.AdditionalItemService",
	)
	return pack, err
}
func (applicationPackageService) Update(
	ctx context.Context,
	id uint,
	data forms.PackageUpdateForm,
) (uint, models.Package, error) {
	var pack models.Package
	err := baseUpdateFunc(ctx, &policies.ServicePolicy{}, id, &pack, data)
	return id, pack, err
}
func (applicationPackageService) Delete(ctx context.Context, id uint) (uint, error) {
	err := baseDeleteFunc(ctx, policies.ServicePolicy{}, id, &models.Package{})
	return id, err
}

// end of package service
