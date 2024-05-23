package view_domains

import "github.com/labstack/echo/v4"

type PaymentViewInterface interface {
	Pay(c echo.Context) error
	AfterPayHook(c echo.Context) error
}
