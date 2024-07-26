package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type serviceRoute struct{}

func (serviceRoute) RegisterRoutes(router *echo.Group) {
	// service
	serviceRoute := router.Group("/services")
	serviceRoute.POST("/", providers.ViewProvider.ApplicationServiceView.Create)
	serviceRoute.GET("/", providers.ViewProvider.ApplicationServiceView.Read)
	serviceRoute.GET("/:id/", providers.ViewProvider.ApplicationServiceView.Find)
	serviceRoute.PATCH("/:id/", providers.ViewProvider.ApplicationServiceView.Update)
	serviceRoute.DELETE("/:id/", providers.ViewProvider.ApplicationServiceView.Delete)
	serviceRoute.POST("/add-required-item/", providers.ViewProvider.ApplicationServiceView.AddRequiredItem)
	serviceRoute.POST("/add-additional-item/", providers.ViewProvider.ApplicationServiceView.AddAdditionalItem)
	serviceRoute.DELETE("/remove-required-item/:id/", providers.ViewProvider.ApplicationServiceView.RemoveRequiredItem)
	serviceRoute.DELETE("/remove-additional-item/:id/", providers.ViewProvider.ApplicationServiceView.RemoveAdditionalItem)
	// package
	packageRoute := router.Group("/packages")
	packageRoute.POST("/", providers.ViewProvider.ApplicationPackageServiceView.Create)
	packageRoute.GET("/", providers.ViewProvider.ApplicationPackageServiceView.Read)
	packageRoute.GET("/:id/", providers.ViewProvider.ApplicationPackageServiceView.Find)
	packageRoute.PATCH("/:id/", providers.ViewProvider.ApplicationPackageServiceView.Update)
	packageRoute.DELETE("/:id/", providers.ViewProvider.ApplicationPackageServiceView.Delete)
}

func init() {
	addAPIRoute(serviceRoute{})
}
