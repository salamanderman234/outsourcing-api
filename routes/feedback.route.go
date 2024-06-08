package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type feedbackRoute struct{}

func (feedbackRoute) RegisterRoutes(router *echo.Echo) {
	r := router.Group("/feedbacks")
	r.POST("/", providers.ViewProvider.FeedbackView.Create)
	r.GET("/:id/", providers.ViewProvider.FeedbackView.Find)
	r.GET("/", providers.ViewProvider.FeedbackView.Read)
	r.PATCH("/:id/", providers.ViewProvider.FeedbackView.Update)
	r.DELETE("/:id/", providers.ViewProvider.FeedbackView.Delete)
}

func init() {
	addRoute(feedbackRoute{})
}
