package view_domains

import "github.com/labstack/echo/v4"

type UIViewInterface interface {
	Dashboard(c echo.Context) error
	CategoryIndex(c echo.Context) error
}
