package services

import (
	"context"
	"fmt"
	"math"

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

func (applicationServiceService) RemoveAdditionalItems(ctx context.Context, id uint) error {
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &models.AdditionalItemService{})
	return err
}

func (applicationServiceService) RemoveRequiredItems(ctx context.Context, id uint) error {
	item := models.RequiredItemService{}
	sub := uint(0)
	serviceId := uint(0)
	before := func() error {
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &item)
		if err != nil {
			return err
		}
		sub = *item.Subtotal
		serviceId = *item.ServiceID
		return nil
	}
	err := baseDeleteFunc(ctx, policies.MasterPolicy{}, id, &item, before)
	if err == nil {
		service := models.Service{}
		err := providers.RepoProvider.BaseRepo.Find(ctx, serviceId, &service)
		if err != nil {
			return err
		}

		curAddtionalPrice := *service.EtcPrice
		curTotal := *service.TotalPrice
		curAddtionalPrice = uint(math.Max(float64(curAddtionalPrice)-float64(sub), 0))
		curTotal = uint(math.Max(float64(curTotal)-float64(sub), 0))

		errX := providers.RepoProvider.BaseRepo.Update(ctx, []uint{service.ID}, &models.Service{
			TotalPrice: &curTotal,
			EtcPrice:   &curAddtionalPrice,
		})
		if errX != nil {
			return errX
		}
	}
	return err
}

func (applicationServiceService) AddRequiredItems(
	ctx context.Context,
	form forms.RequiredItemAddForm,
) (models.RequiredItemService, error) {
	var requiredItem models.RequiredItemService
	serv := models.Service{}
	total := form.PricePerItem * form.Quantity

	before := func() error {
		id := form.ServiceID
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &serv)
		if err != nil {
			return err
		}
		requiredItem.Subtotal = &total
		return nil
	}
	err := baseCreateFunc(ctx, policies.MasterPolicy{}, &requiredItem, form, before)
	if err == nil {
		curAddtionalPrice := *serv.EtcPrice
		curTotal := *serv.TotalPrice
		curAddtionalPrice += total
		curTotal += total

		errX := providers.RepoProvider.BaseRepo.Update(ctx, []uint{serv.ID}, &models.Service{
			TotalPrice: &curTotal,
			EtcPrice:   &curAddtionalPrice,
		})
		if errX != nil {
			return requiredItem, errX
		}
	}
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
		res.SetData(&service)
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
		res.SetData(&service)
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

	var temp models.Service
	var service models.Service
	before := func() error {
		mainImage := data.MainImage
		icon := data.Icon
		if mainImage != nil || icon != nil {
			err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
			if err != nil {
				return nil
			}
		}
		if mainImage != nil {
			res, err := providers.ResourceProvider.CreateResource("services.main_image")
			if err != nil {
				return err
			}
			res.SetData(&service)
			old := temp.MainImage
			result, err := providers.ServiceProvider.FileService.UploadFile(ctx, *mainImage, res)
			if err != nil {
				return err
			}
			if old != nil {
				go providers.ServiceProvider.FileService.DeleteFile(ctx, *old)
			}
			service.MainImage = &result
		}
		if icon != nil {
			res, err := providers.ResourceProvider.CreateResource("services.icon")
			if err != nil {
				return err
			}
			res.SetData(&service)
			old := temp.Icon
			result, err := providers.ServiceProvider.FileService.UploadFile(ctx, *icon, res)
			if err != nil {
				return err
			}
			if old != nil {
				go providers.ServiceProvider.FileService.DeleteFile(ctx, *old)
			}
			service.Icon = &result
		}

		return nil
	}
	err := baseUpdateFunc(ctx, policies.ServicePolicy{}, id, &service, data, before)
	if err == nil {
		employeePrice := uint(0)
		servicePrice := uint(0)
		etcPrice := uint(0)

		if temp.EtcPrice != nil {
			etcPrice = *temp.EtcPrice
		}
		if data.EmployeePrice != nil {
			employeePrice = *data.EmployeePrice
		}
		if data.ServicePrice != nil {
			servicePrice = *data.ServicePrice
		}
		total := etcPrice + employeePrice + servicePrice

		errUpdate := providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, &models.Service{
			TotalPrice: &total,
		})

		if errUpdate != nil {
			return id, service, errUpdate
		}
	}
	return id, service, err
}
func (applicationServiceService) Delete(ctx context.Context, id uint) (uint, error) {
	before := func() error {
		var temp models.Service
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
		if err != nil {
			return nil
		}
		mainImage := temp.MainImage
		icon := temp.Icon
		if mainImage != nil {
			go providers.ServiceProvider.FileService.DeleteFile(ctx, *temp.MainImage)
		}
		if icon != nil {
			go providers.ServiceProvider.FileService.DeleteFile(ctx, *temp.Icon)
		}
		return nil
	}
	err := baseDeleteFunc(ctx, policies.ServicePolicy{}, id, &models.Service{}, before)
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

		mainImage := data.MainImage
		if mainImage != "" {
			res, err := providers.ResourceProvider.CreateResource("packages.main_image")
			if err != nil {
				return err
			}
			res.SetData(&service)
			result, err := providers.ServiceProvider.FileService.UploadFile(ctx, mainImage, res)
			if err != nil {
				return err
			}
			service.MainImage = &result
		}

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
			"Services.AdditionalPackageServiceItems.AdditionalItemService",
			"Services.Service",
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
	var temp models.Package
	var pack models.Package
	before := func() error {
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
		if err != nil {
			return err
		}
		mainImage := data.MainImage
		if mainImage != nil {
			res, err := providers.ResourceProvider.CreateResource("packages.main_image")
			if err != nil {
				return err
			}
			result, err := providers.ServiceProvider.FileService.UploadFile(ctx, *mainImage, res)
			if err != nil {
				return err
			}
			old := temp.MainImage
			if old != nil {
				go providers.ServiceProvider.FileService.DeleteFile(ctx, *old)
			}
			pack.MainImage = &result
		}
		return nil
	}
	err := baseUpdateFunc(ctx, &policies.ServicePolicy{}, id, &pack, data, before)
	if err == nil {
		discount := uint(0)
		servicePrice := uint(0)
		etcPrice := uint(0)
		employeePrice := uint(0)
		if data.Discount != nil {
			discount = *data.Discount
		}
		if temp.ServicePrice != nil {
			servicePrice = *temp.ServicePrice
		}
		if temp.EtcPrice != nil {
			etcPrice = *temp.EtcPrice
		}
		if temp.EmployeePrice != nil {
			employeePrice = *temp.EmployeePrice
		}
		total := (servicePrice + etcPrice + employeePrice) - discount
		errUpdate := providers.RepoProvider.BaseRepo.Update(ctx, []uint{id}, &models.Package{
			TotalPrice: &total,
		})
		if errUpdate != nil {
			return id, pack, errUpdate
		}
	}
	return id, pack, err
}
func (applicationPackageService) Delete(ctx context.Context, id uint) (uint, error) {
	before := func() error {
		var temp models.Package
		err := providers.RepoProvider.BaseRepo.Find(ctx, id, &temp)
		if err != nil {
			return err
		}
		mainImage := temp.MainImage
		if mainImage != nil {
			go providers.ServiceProvider.FileService.DeleteFile(ctx, *mainImage)
		}
		return nil
	}
	err := baseDeleteFunc(ctx, policies.ServicePolicy{}, id, &models.Package{}, before)
	return id, err
}

// end of package service
