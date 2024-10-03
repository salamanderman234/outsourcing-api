package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type fileRoute struct{}

func (fileRoute) RegisterRoutes(router *echo.Group) {
	router.GET("/resource/:model/:id/:field/", providers.ViewProvider.FileView.GetFile)
}
func init() {
	addAPIRoute(fileRoute{})
}
