package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type transactionRoute struct{}

func (transactionRoute) RegisterRoutes(router *echo.Echo) {
	transRoute := router.Group("/transactions")
	transRoute.POST("/", providers.ViewProvider.TransactionView.Create)
	transRoute.GET("/", providers.ViewProvider.TransactionView.Read)
	transRoute.GET("/:id/", providers.ViewProvider.TransactionView.Find)
	transRoute.PATCH("/:id/", providers.ViewProvider.TransactionView.Update)
	transRoute.DELETE("/:id/", providers.ViewProvider.TransactionView.Delete)
	transRoute.POST("/:id/upload-mou/", providers.ViewProvider.TransactionView.UploadMOU)
}

func init() {
	addRoute(transactionRoute{})
}
