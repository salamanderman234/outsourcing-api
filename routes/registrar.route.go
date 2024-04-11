package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains"
)

var routes = []domains.RouteInterface{}

func RegisterAllRoutes(router *echo.Echo) {
	for _, route := range routes {
		route.RegisterRoutes(router)
	}
}

func addRoute(route domains.RouteInterface) {
	routes = append(routes, route)
}
