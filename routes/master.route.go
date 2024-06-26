package routes

import (
	"github.com/labstack/echo/v4"
	custom_middlewares "github.com/salamanderman234/outsourcing-api/app/middlewares"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type masterRoute struct{}

func (masterRoute) RegisterRoutes(router *echo.Group) {
	// province
	provinceRoute := router.Group("/provinces")
	provinceRoute.POST("/", providers.ViewProvider.MasterProvinceView.Create)
	provinceRoute.GET("/", providers.ViewProvider.MasterProvinceView.Read)
	provinceRoute.GET("/:id/", providers.ViewProvider.MasterProvinceView.Find)
	provinceRoute.PATCH("/:id/", providers.ViewProvider.MasterProvinceView.Update)
	provinceRoute.DELETE("/:id/", providers.ViewProvider.MasterProvinceView.Delete)
	// regency
	regencyRoute := router.Group("/regencies")
	regencyRoute.POST("/", providers.ViewProvider.MasterRegencyView.Create)
	regencyRoute.GET("/", providers.ViewProvider.MasterRegencyView.Read)
	regencyRoute.GET("/:id/", providers.ViewProvider.MasterRegencyView.Find)
	regencyRoute.PATCH("/:id/", providers.ViewProvider.MasterRegencyView.Update)
	regencyRoute.DELETE("/:id/", providers.ViewProvider.MasterRegencyView.Delete)
	// regency
	categoryRoute := router.Group("/categories")
	categoryRoute.POST("/", providers.ViewProvider.MasterCategoryView.Create)
	categoryRoute.GET("/", providers.ViewProvider.MasterCategoryView.Read)
	categoryRoute.GET("/:id/", providers.ViewProvider.MasterCategoryView.Find)
	categoryRoute.PATCH("/:id/", providers.ViewProvider.MasterCategoryView.Update)
	categoryRoute.DELETE("/:id/", providers.ViewProvider.MasterCategoryView.Delete)
	// user
	userRoute := router.Group("/users")
	userRoute.GET("/:role/", providers.ViewProvider.UserView.Read)
	userRoute.GET("/:id/", providers.ViewProvider.UserView.Find)
	userRoute.PATCH("/:id/", providers.ViewProvider.UserView.Update, custom_middlewares.MustVerifyUser)
	// question
	questionRoute := router.Group("/questions")
	questionRoute.GET("/", providers.ViewProvider.MasterQuestionView.Read)
	questionRoute.POST("/", providers.ViewProvider.MasterQuestionView.Create)
	questionRoute.GET("/:id/", providers.ViewProvider.MasterQuestionView.Find)
	questionRoute.PATCH("/:id/", providers.ViewProvider.MasterQuestionView.Update)
	questionRoute.DELETE("/:id/", providers.ViewProvider.MasterQuestionView.Delete)
	questionRoute.POST("/assign/", providers.ViewProvider.MasterQuestionView.AssignQuestion)
	questionRoute.PATCH("/unassign/", providers.ViewProvider.MasterQuestionView.UnassignQuestion)
	// payment config
	paymentConfigRoute := router.Group("/payment_config")
	paymentConfigRoute.PATCH("/set_dp_percentage/", providers.ViewProvider.MasterPaymentConfigView.SetDPPercentage)
	paymentConfigRoute.PATCH("/set_3_termin_first_percentage/", providers.ViewProvider.MasterPaymentConfigView.Set3TerminFirst)
	paymentConfigRoute.PATCH("/set_3_termin_second_percentage/", providers.ViewProvider.MasterPaymentConfigView.Set3TerminSecond)
}

func init() {
	addAPIRoute(masterRoute{})
}
