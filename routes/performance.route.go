package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type performanceRoute struct{}

func (performanceRoute) RegisterRoutes(router *echo.Group) {
	// province
	pRoute := router.Group("/performances")
	pRoute.POST("/form/", providers.ViewProvider.PerformanceView.CreateForm)
	pRoute.GET("/form/:id/", providers.ViewProvider.PerformanceView.GetForm)
	pRoute.DELETE("/form/:id/", providers.ViewProvider.PerformanceView.DeleteForm)
	pRoute.POST("/submit/", providers.ViewProvider.PerformanceView.SubmitAnswer)
	pRoute.GET("/", providers.ViewProvider.PerformanceView.GetEmployeePerformances)
}

func init() {
	addAPIRoute(performanceRoute{})
}
