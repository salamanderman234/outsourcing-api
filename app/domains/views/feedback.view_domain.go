package view_domains

import (
	"github.com/labstack/echo/v4"
)

type FeedbackViewInterface interface {
	BaseCRUDViewInterface
}

type ComplaintViewInterface interface {
	Create(c echo.Context) error
	Find(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Reply(c echo.Context) error
	UpdateReply(c echo.Context) error
	DeleteReply(c echo.Context) error
	Delete(c echo.Context) error
}

type PerformanceViewInterface interface {
	CreateForm(c echo.Context) error
	GetForm(c echo.Context) error
	DeleteForm(c echo.Context) error
	SubmitAnswer(c echo.Context) error
	GetEmployeePerformances(c echo.Context) error
}
