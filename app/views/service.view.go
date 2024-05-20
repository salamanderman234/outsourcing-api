package views

import (
	"context"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

// service
type applicationServiceView struct{}

func NewApplicationServiceView() view_domains.ApplicationServiceViewInterface {
	return &applicationServiceView{}
}

func (applicationServiceView) AddRequiredItem(c echo.Context) error {
	var form forms.RequiredItemAddForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.AppServiceService.AddRequiredItems(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}

func (applicationServiceView) AddAdditionalItem(c echo.Context) error {
	var form forms.AdditionalItemServiceAddForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.AppServiceService.AddAdditionalItems(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}

func (applicationServiceView) Create(c echo.Context) error {
	var form forms.ServiceCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.AppServiceService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (applicationServiceView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.AppServiceService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (applicationServiceView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.AppServiceService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (applicationServiceView) Update(c echo.Context) error {
	var form forms.ServiceUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.AppServiceService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (applicationServiceView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.AppServiceService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

// end of service
// package service
type applicationPackageServiceView struct{}

func NewApplicationPackageServiceView() view_domains.ApplicationPackageServiceViewInterface {
	return &applicationPackageServiceView{}
}
func (applicationPackageServiceView) Create(c echo.Context) error {
	var form forms.PackageCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.AppPackageServiceService.Create(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (applicationPackageServiceView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.AppPackageServiceService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}
func (applicationPackageServiceView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.AppPackageServiceService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (applicationPackageServiceView) Update(c echo.Context) error {
	var form forms.PackageUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.AppPackageServiceService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (applicationPackageServiceView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.AppPackageServiceService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

// end of package service
