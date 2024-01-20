package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/middlewares"
)

func registerUserRoute(group *echo.Group) {
	group.Use(middlewares.RetrieveSessionUser)
	userGroup := group
	userGroup.GET("/", domains.ViewRegistry.UserView.Read)
	userGroup.PATCH("/", domains.ViewRegistry.UserView.Update)
	userGroup.DELETE("/", domains.ViewRegistry.UserView.Delete)

	userGroup.GET("/service-users/", domains.ViewRegistry.ServiceUserView.Read)
	userGroup.PATCH("/service-users/", domains.ViewRegistry.ServiceUserView.Update)
	userGroup.DELETE("/service-users/", domains.ViewRegistry.ServiceUserView.Delete)

	userGroup.GET("/employees/", domains.ViewRegistry.EmployeeView.Read)
	userGroup.PATCH("/employees/", domains.ViewRegistry.EmployeeView.Update)
	userGroup.DELETE("/employees/", domains.ViewRegistry.EmployeeView.Delete)

	userGroup.GET("/supervisors/", domains.ViewRegistry.SupervisorView.Read)
	userGroup.PATCH("/supervisors/", domains.ViewRegistry.SupervisorView.Update)
	userGroup.DELETE("/supervisors/", domains.ViewRegistry.SupervisorView.Delete)

	userGroup.GET("/admins/", domains.ViewRegistry.AdminView.Read)
	userGroup.PATCH("/admins/", domains.ViewRegistry.AdminView.Update)
	userGroup.DELETE("/admins/", domains.ViewRegistry.AdminView.Delete)
}
