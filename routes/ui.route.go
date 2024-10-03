package routes

import (
	"github.com/labstack/echo/v4"
)

type uiRoute struct{}

func (uiRoute) RegisterRoutes(router *echo.Echo) {

}

func init() {
	addRoute(uiRoute{})
}
