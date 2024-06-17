package views

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type uiView struct {
}

func NewUIView() view_domains.UIViewInterface {
	return &uiView{}
}

func (uiView) Dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "index", map[string]any{})
}
func (uiView) CategoryIndex(c echo.Context) error {
	ctx := c.Request().Context()
	q := c.QueryParam("q")
	page := c.QueryParam("page")
	pageInt, _ := strconv.Atoi(page)
	if pageInt == 0 {
		pageInt = 1
	}
	results, pagination, err := providers.ServiceProvider.MasterCategoryService.Read(ctx, q, uint(pageInt))
	if err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Action: enums.ReadAction,
			Error:  err,
		})
		return c.Render(status, "errors.error", map[string]any{
			"code":    status,
			"message": resp.GetMsg(),
		})

	}
	return c.Render(http.StatusOK, "masters.category.index", map[string]any{
		"datas":      results,
		"pagination": pagination,
	})
}
