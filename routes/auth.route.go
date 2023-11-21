package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerAuthRoute(group *echo.Group) {
	group.POST("/login", domains.ViewRegistry.AuthView.Login)
	group.GET("/verify", domains.ViewRegistry.AuthView.Verify)
	group.GET("/refresh", domains.ViewRegistry.AuthView.Refresh)
}
