package view_domains

import "github.com/labstack/echo/v4"

type MasterProvinceViewInterface interface {
	BaseCRUDViewInterface
}

type MasterRegencyViewInterface interface {
	BaseCRUDViewInterface
}

type MasterCategoryViewInterface interface {
	BaseCRUDViewInterface
}

type MasterQuestionViewInterface interface {
	BaseCRUDViewInterface
	AssignQuestion(c echo.Context) error
	UnassignQuestion(c echo.Context) error
}

type MasterPaymentConfigViewInterface interface {
	SetDPPercentage(c echo.Context) error
	Set3TerminFirst(c echo.Context) error
	Set3TerminSecond(c echo.Context) error
	GetConfigs(c echo.Context) error
}
