package routes

import "github.com/labstack/echo/v4"

func RegisterAllRoutes(server *echo.Echo) {
	// groups
	authGroup := server.Group("")
	masterGroup := server.Group("/master/")
	// register
	registerAuthRoute(authGroup)
	registerMasterRoutes(masterGroup)
}
