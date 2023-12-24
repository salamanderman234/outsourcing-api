package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerOrderRoutes(group *echo.Group) {
	group.POST("/", domains.ViewRegistry.OrderView.MakeServiceOrder)
}
