package view_domains

import "github.com/labstack/echo/v4"

type FileViewInterface interface {
	GetFile(c echo.Context) error
}
