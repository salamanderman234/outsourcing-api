package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type placementRoute struct{}

func (placementRoute) RegisterRoutes(router *echo.Echo) {
	placementR := router.Group("/placements")
	placementR.POST("/", providers.ViewProvider.PlacementView.Create)
	router.GET("/transactions/:id/placement/", providers.ViewProvider.PlacementView.GetPlacementOrder)
	placementR.GET("/:id/", providers.ViewProvider.PlacementView.Find)
	placementR.GET("/", providers.ViewProvider.PlacementView.Read)
	placementR.GET("/:id/details/", providers.ViewProvider.PlacementView.PlacementDetails)
	placementR.PATCH("/:id/", providers.ViewProvider.PlacementView.Update)
	placementR.DELETE("/:id/", providers.ViewProvider.PlacementView.Delete)
	placementR.POST("/employees/", providers.ViewProvider.PlacementView.PlaceEmployee)
	placementR.GET("/employees/:id/", providers.ViewProvider.PlacementView.PlacementEmployeeDetail)
	placementR.PATCH("/employees/:id/cutoff", providers.ViewProvider.PlacementView.CutoffEmployee)
	placementR.DELETE("/employees/:id/", providers.ViewProvider.PlacementView.RemoveEmployee)
}

func init() {
	addRoute(placementRoute{})
}
