package views

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// --> Category
type categoryView struct{}

func NewCategoryView() domains.ServiceCategoryView {
	return &categoryView{}
}
func (categoryView) Create(c echo.Context) error {
	var data domains.CategoryCreateForm
	createCallFunc := func(ctx context.Context) (any, error) {
		file, err := readFile(c, "icon")
		if err != nil {
			return nil, err
		}
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  file,
		}
		return domains.ServiceRegistry.CategoryServ.Create(ctx, data, iconFileMap)
	}
	return basicCreateView(c, &data, createCallFunc)
}
func (categoryView) GetIcon(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	category, _, err := domains.ServiceRegistry.CategoryServ.Read(ctx, uint(id), "", 1, "", true)
	if errors.Is(err, domains.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	categoryEntity := category.(domains.CategoryEntity)
	path := categoryEntity.Icon
	return c.File(path)
}
func (categoryView) Read(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	var query string
	var page uint
	var order string
	desc := -1
	echo.QueryParamsBinder(c).
		Uint("id", &id).
		Uint("page", &page).
		String("q", &query).
		String("order-by", &order).
		Int("desc", &desc)

	if desc == -1 {
		desc = 1
	}
	datas, pagination, err := domains.ServiceRegistry.CategoryServ.Read(ctx,
		id, query, uint(math.Max(float64(1), float64(page))), order, desc > 0,
	)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	bodyResp := domains.DataBodyResponse{
		Pagination: pagination,
	}
	data, ok := datas.(domains.CategoryEntity)
	if ok {
		bodyResp.Data = data
	} else {
		bodyResp.Datas = datas
	}
	resp.Body = bodyResp
	return c.JSON(http.StatusOK, resp)
}
func (categoryView) Update(c echo.Context) error {
	var data domains.CategoryUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		file, err := readFile(c, "icon")
		if err != nil {
			return 0, nil, err
		}
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  file,
		}
		return domains.ServiceRegistry.CategoryServ.Update(ctx, id, data, iconFileMap)
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (categoryView) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	if err := echo.QueryParamsBinder(c).Uint("id", &id).BindError(); err != nil || id == 0 {
		msg := "id is required, must be an unsigned integer and cannot be 0"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	deleted, aff, err := domains.ServiceRegistry.CategoryServ.Delete(ctx, id)
	deletedID := uint(deleted)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		ID:       &deletedID,
		Affected: &aff,
	}
	return c.JSON(http.StatusOK, resp)
}
