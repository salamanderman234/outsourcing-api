package views

import (
	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
)

type placementView struct{}

func NewPlacementView() view_domains.PlacementViewInterface {
	return &placementView{}
}

func (placementView) Create(c echo.Context) error {
	return nil
}
func (placementView) GetPlacementOrder(c echo.Context) error {
	return nil
}
func (placementView) Find(c echo.Context) error {
	return nil
}
func (placementView) Read(c echo.Context) error {
	return nil
}
func (placementView) PlaceEmployee(c echo.Context) error {
	return nil
}
func (placementView) CutoffEmployee(c echo.Context) error {
	return nil
}
func (placementView) RemoveEmployee(c echo.Context) error {
	return nil
}
func (placementView) Delete(c echo.Context) error {
	return nil
}
