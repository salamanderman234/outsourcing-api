package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerServiceRoute(group *echo.Group) {
	serviceGroup := group.Group("services")
	serviceGroup.GET("/", domains.ViewRegistry.ServiceView.Read)
	serviceGroup.POST("/", domains.ViewRegistry.ServiceView.Create)
	serviceGroup.PATCH("/", domains.ViewRegistry.ServiceView.Update)
	serviceGroup.DELETE("/", domains.ViewRegistry.ServiceView.Delete)

	serviceItemGroup := group.Group("service-items")
	serviceItemGroup.GET("/", domains.ViewRegistry.ServiceItemView.Read)
	serviceItemGroup.POST("/", domains.ViewRegistry.ServiceItemView.Create)
	serviceItemGroup.PATCH("/", domains.ViewRegistry.ServiceItemView.Update)
	serviceItemGroup.DELETE("/", domains.ViewRegistry.ServiceItemView.Delete)
}
