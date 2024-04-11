package view_domains

import "github.com/labstack/echo/v4"

type BaseCRUDViewInterface interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
