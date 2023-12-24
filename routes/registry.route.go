package routes

import "github.com/labstack/echo/v4"

func RegisterAllRoutes(server *echo.Echo) {
	// groups
	authGroup := server.Group("")
	masterGroup := server.Group("/master/")
	serviceGroup := server.Group("")
	orderGroup := server.Group("/service-orders")
	// register
	registerAuthRoute(authGroup)
	registerMasterRoutes(masterGroup)
	registerServiceRoute(serviceGroup)
	registerOrderRoutes(orderGroup)
}
