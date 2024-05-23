package view_domains

import "github.com/labstack/echo/v4"

type TransactionViewInterface interface {
	BaseCRUDViewInterface
	UploadMOU(c echo.Context) error
}
