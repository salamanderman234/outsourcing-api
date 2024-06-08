package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/providers"
)

type complaintRoute struct{}

func (complaintRoute) RegisterRoutes(router *echo.Echo) {
	r := router.Group("/complaints")
	r.POST("/", providers.ViewProvider.ComplaintView.Create)
	r.GET("/:id/", providers.ViewProvider.ComplaintView.Find)
	r.GET("/", providers.ViewProvider.ComplaintView.Read)
	r.PATCH("/:id/", providers.ViewProvider.ComplaintView.Update)
	r.DELETE("/:id/", providers.ViewProvider.ComplaintView.Delete)
	// reply
	r.POST("/replies/", providers.ViewProvider.ComplaintView.Reply)
	r.PATCH("/replies/:id/", providers.ViewProvider.ComplaintView.UpdateReply)
	r.DELETE("/replies/:id/", providers.ViewProvider.ComplaintView.DeleteReply)
}

func init() {
	addRoute(complaintRoute{})
}
