package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains"
)

type authRoute struct{}

func (authRoute) RegisterRoutes(router *echo.Echo) {
	router.POST("/login/", domains.ViewRegistry.AuthView.Login)
	router.POST("/:role/register/", domains.ViewRegistry.AuthView.Register)
	router.POST("/forgot/", domains.ViewRegistry.AuthView.ChangePassword)
	router.POST("/reset/", domains.ViewRegistry.AuthView.ResetPassword)
	router.POST("/verify/", domains.ViewRegistry.AuthView.VerifyUser)
}

func init() {
	addRoute(authRoute{})
}
