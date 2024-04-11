package domains

import "github.com/labstack/echo/v4"

type RouteInterface interface {
	RegisterRoutes(router *echo.Echo)
}
