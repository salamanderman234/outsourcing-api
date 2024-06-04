package views

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type userView struct{}

func NewUserView() view_domains.UserViewInterface {
	return &userView{}
}

func (userView) Create(c echo.Context) error {
	err := types.ErrRecordNotFound
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (userView) Read(c echo.Context) error {
	regencyParam := c.QueryParam("regency_id")
	regency, _ := strconv.Atoi(regencyParam)
	role := enums.UserRolesEnum(c.Param("role"))
	switch string(role) {
	case string(enums.AdminUserRole):
		role = enums.AdminUserRole
	case string(enums.EmployeeUserRole):
		role = enums.EmployeeUserRole
	case string(enums.SupervisorUserRole):
		role = enums.SupervisorUserRole
	case string(enums.ServiceUserRole):
		role = enums.ServiceUserRole
	default:
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: types.ErrRouteNotFound,
		})
		return c.JSON(status, resp)
	}
	callback := func(ctx context.Context, q string, page uint) (any, *types.Pagination, error) {
		return providers.ServiceProvider.UserService.Read(ctx, q, uint(regency), role, page)
	}
	return baseReadFunc(c, callback)
}
func (userView) Find(c echo.Context) error {
	callback := func(ctx context.Context, id uint) (any, error) {
		return providers.ServiceProvider.UserService.Find(ctx, id)
	}
	return baseFindFunc(c, callback)
}
func (userView) Update(c echo.Context) error {
	var form forms.UserUpdateForm
	callback := func(ctx context.Context, id uint) (uint, any, error) {
		return providers.ServiceProvider.UserService.Update(ctx, id, form)
	}
	return baseUpdateFunc(c, &form, callback)
}
func (userView) Delete(c echo.Context) error {
	err := types.ErrRecordNotFound
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
