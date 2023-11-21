package routes

import "github.com/labstack/echo/v4"

func RegisterAllRoutes(server *echo.Echo) {
	// groups
	authGroup := server.Group("")
	// register
	registerAuthRoute(authGroup)
}
