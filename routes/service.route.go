package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type serviceRoute struct{}

func (serviceRoute) RegisterRoutes(router *echo.Echo) {
	// service
	serviceRoute := router.Group("/services")
	serviceRoute.POST("/", providers.ViewProvider.ApplicationServiceView.Create)
	serviceRoute.GET("/", providers.ViewProvider.ApplicationServiceView.Read)
	serviceRoute.GET("/:id/", providers.ViewProvider.ApplicationServiceView.Find)
	serviceRoute.PATCH("/:id/", providers.ViewProvider.ApplicationServiceView.Update)
	serviceRoute.DELETE("/:id/", providers.ViewProvider.ApplicationServiceView.Delete)
}

func init() {
	addRoute(serviceRoute{})
}
