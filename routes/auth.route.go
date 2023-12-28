package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerAuthRoute(group *echo.Group) {
	group.POST("/login/", domains.ViewRegistry.AuthView.Login)
	group.POST("/admins/register/", domains.ViewRegistry.AuthView.RegisterAdmin)
	group.POST("/supervisors/register/", domains.ViewRegistry.AuthView.RegisterSupervisor)
	group.POST("/service-users/register/", domains.ViewRegistry.AuthView.RegisterUserService)
	group.POST("/employees/register/", domains.ViewRegistry.AuthView.RegisterEmployee)
	group.GET("/verify/", domains.ViewRegistry.AuthView.Verify)
	group.GET("/refresh/", domains.ViewRegistry.AuthView.Refresh)
}
