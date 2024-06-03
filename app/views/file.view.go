package views

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
)

type fileView struct{}

func NewFileView() view_domains.FileViewInterface {
	return &fileView{}
}

func (fileView) GetFile(c echo.Context) error {
	model := c.Param("model")
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	field := c.Param("field")
	ctx := c.Request().Context()

	alias := fmt.Sprintf("%s.%s", model, field)
	res, err := providers.ResourceProvider.CreateResource(alias)
	if err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	file, mime, err := providers.ServiceProvider.FileService.GetFile(ctx, uint(id), res)
	if err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	return c.Blob(http.StatusOK, mime, file)
}
