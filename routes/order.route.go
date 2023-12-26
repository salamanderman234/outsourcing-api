package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/middlewares"
)

func registerOrderRoutes(group *echo.Group) {
	group.POST("/", domains.ViewRegistry.OrderView.MakeServiceOrder,
		middlewares.RetrieveSessionUser,
		middlewares.OnlyServiceUser,
	)
	group.GET("/", domains.ViewRegistry.OrderView.ListOrder,
		middlewares.RetrieveSessionUser,
	)
}
