package routes

import (
	"github.com/labstack/echo/v4"
	custom_middlewares "github.com/salamanderman234/outsourcing-api/app/middlewares"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type paymentRoute struct{}

func (paymentRoute) RegisterRoutes(router *echo.Echo) {
	// province
	// provinceRoute := router.Group("/payments")
	router.POST("/transactions/:id/pay/", providers.ViewProvider.PaymentView.Pay, custom_middlewares.MustVerifyUser)
	router.POST("/payment-notif/", providers.ViewProvider.PaymentView.AfterPayHook, custom_middlewares.MustVerifyUser)
}

func init() {
	addRoute(paymentRoute{})
}
