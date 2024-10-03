package view_domains

import "github.com/labstack/echo/v4"

type TransactionViewInterface interface {
	BaseCRUDViewInterface
	UploadMOU(c echo.Context) error
	ConfirmTransaction(c echo.Context) error
	AskForMOU(c echo.Context) error
	SetStatus(c echo.Context) error
}
