package views

import (
	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type authView struct{}

func NewAuthView() view_domains.AuthViewInterface {
	return &authView{}
}

func (authView) Login(c echo.Context) error {
	var loginForm forms.LoginForm
	ctx := c.Request().Context()
	if err := c.Bind(&loginForm); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	user, token, err := providers.ServiceProvider.AuthService.Login(ctx, loginForm)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Data: map[string]any{
			"token": token,
			"user":  user,
		},
		Error: err,
	})
	return c.JSON(status, resp)
}
func (authView) Register(c echo.Context) error {
	var role enums.UserRolesEnum
	var form forms.UserRegisterForm
	ctx := c.Request().Context()
	role = enums.UserRolesEnum(c.Param("role"))
	switch string(role) {
	case string(enums.AdminUserRole):
		role = enums.AdminUserRole
		form.AdminProfile = &forms.AdminProfileRegisterForm{}
	case string(enums.EmployeeUserRole):
		role = enums.EmployeeUserRole
		form.EmployeeProfile = &forms.EmployeeProfileRegisterForm{}
	case string(enums.SupervisorUserRole):
		role = enums.SupervisorUserRole
		form.SupervisorProfile = &forms.SupervisorProfileRegisterForm{}
	case string(enums.ServiceUserRole):
		role = enums.ServiceUserRole
		form.ServiceUserProfile = &forms.ServiceUserProfileRegisterForm{}
	default:
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: types.ErrRouteNotFound,
		})
		return c.JSON(status, resp)
	}
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	user, token, err := providers.ServiceProvider.AuthService.RegisterUser(ctx, form, role)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Data: map[string]any{
			"token":           token,
			"registered_user": user,
		},
		Error: err,
	})
	return c.JSON(status, resp)
}
func (authView) ChangePassword(c echo.Context) error {
	var form forms.ChangePasswordForm
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	err := providers.ServiceProvider.AuthService.ForgotPassword(ctx, form)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (authView) ResetPassword(c echo.Context) error {
	var form forms.ResetPasswordForm
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	err := providers.ServiceProvider.AuthService.ResetPassword(ctx, form)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (authView) VerifyUser(c echo.Context) error {
	var form forms.VerifyUserForm
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	err := providers.ServiceProvider.AuthService.VerifyEmail(ctx, form)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
