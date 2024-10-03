package view_domains

import "github.com/labstack/echo/v4"

type AuthViewInterface interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	ChangePassword(c echo.Context) error
	ResetPassword(c echo.Context) error
	VerifyUser(c echo.Context) error
	SendVerify(c echo.Context) error
}
