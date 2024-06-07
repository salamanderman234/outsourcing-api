package view_domains

import (
	"github.com/labstack/echo/v4"
)

type PlacementViewInterface interface {
	Create(c echo.Context) error
	GetPlacementOrder(c echo.Context) error
	Find(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	PlacementDetails(c echo.Context) error
	PlaceEmployee(c echo.Context) error
	PlacementEmployeeDetail(c echo.Context) error
	CutoffEmployee(c echo.Context) error
	RemoveEmployee(c echo.Context) error
}
