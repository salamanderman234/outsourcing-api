package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type authRoute struct{}

func (authRoute) RegisterRoutes(router *echo.Echo) {
	router.POST("/login/", providers.ViewProvider.AuthView.Login)
	router.POST("/:role/register/", providers.ViewProvider.AuthView.Register)
	router.POST("/forgot/", providers.ViewProvider.AuthView.ChangePassword)
	router.POST("/reset/", providers.ViewProvider.AuthView.ResetPassword)
	router.POST("/verify/", providers.ViewProvider.AuthView.VerifyUser)
}

func init() {
	addRoute(authRoute{})
}
