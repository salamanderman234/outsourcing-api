package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains"
)

type masterRoute struct{}

func (masterRoute) RegisterRoutes(router *echo.Echo) {
	// province
	provinceRoute := router.Group("/provinces")
	provinceRoute.POST("/", domains.ViewRegistry.MasterProvinceView.Create)
	provinceRoute.GET("/", domains.ViewRegistry.MasterProvinceView.Read)
	provinceRoute.GET("/:id/", domains.ViewRegistry.MasterProvinceView.Find)
	provinceRoute.PATCH("/:id/", domains.ViewRegistry.MasterProvinceView.Update)
	provinceRoute.DELETE("/:id/", domains.ViewRegistry.MasterProvinceView.Delete)
	// regency
	regencyRoute := router.Group("/regencies")
	regencyRoute.POST("/", domains.ViewRegistry.MasterRegencyView.Create)
	regencyRoute.GET("/", domains.ViewRegistry.MasterRegencyView.Read)
	regencyRoute.GET("/:id/", domains.ViewRegistry.MasterRegencyView.Find)
	regencyRoute.PATCH("/:id/", domains.ViewRegistry.MasterRegencyView.Update)
	regencyRoute.DELETE("/:id/", domains.ViewRegistry.MasterRegencyView.Delete)
}

func init() {
	addRoute(masterRoute{})
}
