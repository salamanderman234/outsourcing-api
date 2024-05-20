package view_domains

import "github.com/labstack/echo/v4"

type ApplicationServiceViewInterface interface {
	BaseCRUDViewInterface
	AddRequiredItem(c echo.Context) error
	AddAdditionalItem(c echo.Context) error
}

type ApplicationPackageServiceViewInterface interface {
	BaseCRUDViewInterface
}
