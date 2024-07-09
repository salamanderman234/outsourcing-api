package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type authRoute struct{}

func (authRoute) RegisterRoutes(router *echo.Group) {
	login := router.POST("/login/", providers.ViewProvider.AuthView.Login)
	login.Name = "auth.login"
	register := router.POST("/:role/register/", providers.ViewProvider.AuthView.Register)
	register.Name = "auth.register"
	changePass := router.POST("/forgot/", providers.ViewProvider.AuthView.ChangePassword)
	changePass.Name = "auth.change-password"
	resetPass := router.POST("/reset/", providers.ViewProvider.AuthView.ResetPassword)
	resetPass.Name = "auth.reset-password"
	verifyUser := router.GET("/:user_id/verify/", providers.ViewProvider.AuthView.VerifyUser)
	verifyUser.Name = "auth.verify"
	router.POST("/:id/send_verify/", providers.ViewProvider.AuthView.SendVerify)
}

func init() {
	addAPIRoute(authRoute{})
}
