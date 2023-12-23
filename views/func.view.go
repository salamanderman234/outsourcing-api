package views

import (
	"context"
	"errors"
	"math"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

func readFile(c echo.Context, formName string) (*multipart.FileHeader, error) {
	file, err := c.FormFile(formName)
	if errors.Is(err, http.ErrMissingFile) {
		return nil, nil
	}
	if err != nil {
		return nil, domains.ErrGetMultipartFormData
	}
	return file, nil
}

type callCreateServFunc func(ctx context.Context) (domains.Entity, error)

func basicCreateView(c echo.Context, data any, callFunc callCreateServFunc) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "created",
	}
	if err := c.Bind(data); err != nil {
		msg := "invalid request"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	created, err := callFunc(ctx)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: created,
	}
	return c.JSON(http.StatusCreated, resp)
}

type callUpdateServFunc func(ctx context.Context, id uint) (int, any, error)

func basicUpdateView(c echo.Context, data any, callFunc callUpdateServFunc) error {
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
	if err := c.Bind(data); err != nil {
		msg := "invalid request"
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	aff, updated, err := callFunc(ctx, id)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		ID:       &id,
		Affected: &aff,
		Updated:  updated,
	}
	return c.JSON(http.StatusOK, resp)
}

type callDeleteServFunc func(ctx context.Context, id uint) (int, int, error)

func basicDeleteView(c echo.Context, callFunc callDeleteServFunc) error {
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
	deleted, aff, err := callFunc(ctx, id)
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

type callReadServFunc func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error)

func basicReadView(c echo.Context, callFunc callReadServFunc) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	var query string
	var page uint
	var order string
	withPagination := -1
	desc := -1
	echo.QueryParamsBinder(c).
		Uint("id", &id).
		Uint("page", &page).
		String("q", &query).
		String("order-by", &order).
		Int("desc", &desc).
		Int("with-pagination", &withPagination)
	if desc == -1 {
		desc = 1
	}
	if withPagination == -1 {
		withPagination = 1
	}
	datas, pagination, err := callFunc(ctx, id, query, uint(math.Max(float64(1), float64(page))), order, desc > 0, withPagination > 0)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	bodyResp := domains.DataBodyResponse{
		Pagination: pagination,
	}
	data, ok := datas.(domains.Entity)
	if ok {
		bodyResp.Data = data
	} else {
		bodyResp.Datas = datas
	}
	resp.Body = bodyResp
	return c.JSON(http.StatusOK, resp)
}
