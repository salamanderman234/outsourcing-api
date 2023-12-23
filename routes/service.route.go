package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerServiceRoute(group *echo.Group) {
	serviceCategoryGroup := group.Group("services")
	serviceCategoryGroup.GET("/", domains.ViewRegistry.ServiceView.Read)
	serviceCategoryGroup.POST("/", domains.ViewRegistry.ServiceView.Create)
	serviceCategoryGroup.PATCH("/", domains.ViewRegistry.ServiceView.Update)
	serviceCategoryGroup.DELETE("/", domains.ViewRegistry.ServiceView.Delete)
}
