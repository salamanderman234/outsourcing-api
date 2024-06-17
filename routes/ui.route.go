package routes

import (
	"github.com/labstack/echo/v4"
	custom_middlewares "github.com/salamanderman234/outsourcing-api/app/middlewares"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type uiRoute struct{}

func (uiRoute) RegisterRoutes(router *echo.Echo) {
	uiRoute := router.Group("", custom_middlewares.MustVerifyUser)
	uiRoute.GET("/", providers.ViewProvider.UIView.Dashboard)
	uiRoute.GET("/categories", providers.ViewProvider.UIView.CategoryIndex)
}

func init() {
	addRoute(uiRoute{})
}
