package views

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type placementView struct{}

func NewPlacementView() view_domains.PlacementViewInterface {
	return &placementView{}
}

func (placementView) Create(c echo.Context) error {
	var form forms.PlacementCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.PlacementService.CreatePlacement(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (placementView) GetPlacementOrder(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.PlacementService.GetPlacementOrder(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (placementView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.PlacementService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (placementView) Read(c echo.Context) error {
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.PlacementService.Read(ctx, q, page)
	}
	return baseReadFunc(c, callback)
}

func (placementView) Update(c echo.Context) error {
	var form forms.PlacementUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.PlacementService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}

func (placementView) PlaceEmployee(c echo.Context) error {
	var form forms.PlacementDetailEmployeeCreateForm
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.PlacementService.PlaceNewEmployee(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (placementView) CutoffEmployee(c echo.Context) error {
	var form forms.PlacementDetailEmployeeUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		retId, err := providers.ServiceProvider.PlacementService.CutoffEmployeePlacement(ctx, id, form)
		return retId, form, err
	}
	return baseUpdateFunc(c, &form, callback)
}
func (placementView) RemoveEmployee(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.PlacementService.RemoveEmployeePlacement(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
func (placementView) Delete(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return providers.ServiceProvider.PlacementService.Delete(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}

func (placementView) PlacementDetails(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		results, err := providers.ServiceProvider.PlacementService.PlacementDetails(ctx, uint(id))
		return results, nil, err
	}
	return baseReadFunc(c, callback)
}

func (placementView) PlacementEmployeeDetail(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.PlacementService.EmployeePlacementDetail(ctx, id)
	}
	return baseFindFunc(c, callback)
}
