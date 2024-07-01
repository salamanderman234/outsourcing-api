package views

import (
	"context"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type performanceView struct{}

func NewPerformanceView() view_domains.PerformanceViewInterface {
	return &performanceView{}
}

func (performanceView) CreateForm(c echo.Context) error {
	form := forms.PerformanceCreateForm{}
	callback := func(ctx context.Context) (any, error) {
		return providers.ServiceProvider.PerformanceService.CreateForm(ctx, form.PlacementID)
	}
	return baseCreateFunc(c, &form, callback)
}
func (performanceView) GetForm(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.PerformanceService.GetForm(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (performanceView) DeleteForm(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (uint, error) {
		return id, providers.ServiceProvider.PerformanceService.DeleteForm(ctx, id)
	}
	return baseDeleteFunc(c, callback)
}
func (performanceView) SubmitAnswer(c echo.Context) error {
	form := forms.PerformanceSubmitForm{}
	callback := func(ctx context.Context) (any, error) {
		return nil, providers.ServiceProvider.PerformanceService.SubmitAnswer(ctx, form)
	}
	return baseCreateFunc(c, &form, callback)
}
func (performanceView) GetEmployeePerformances(c echo.Context) error {
	var filter forms.PerformanceFilter
	if err := c.Bind(&filter); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Action: enums.ReadAction,
			Error:  err,
		})
		return c.JSON(status, resp)
	}
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		data, err := providers.ServiceProvider.PerformanceService.GetEmployeePerformances(
			ctx,
			filter.EmployeeID,
			&filter.From,
			&filter.To,
		)
		return data, nil, err
	}
	return baseReadFunc(c, callback)
}
