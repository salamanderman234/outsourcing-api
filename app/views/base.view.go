package views

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type createCallback func(ctx context.Context) (any, error)

func baseCreateFunc(c echo.Context, form any, callback createCallback) error {
	ctx := c.Request().Context()
	if err := c.Bind(form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	result, err := callback(ctx)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.CreateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}

type readCallback func(ctx context.Context, q string, page uint) (any, *types.Pagination, error)

func baseReadFunc(c echo.Context, callback readCallback) error {
	q := c.QueryParam("query")
	pageParam := c.QueryParam("page")
	page, _ := strconv.Atoi(pageParam)
	ctx := c.Request().Context()
	results, pagination, err := callback(ctx, q, uint(page))
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action:     enums.ReadAction,
		Error:      err,
		Datas:      results,
		Pagination: pagination,
	})
	return c.JSON(status, resp)
}

type findCallback func(ctx context.Context, id uint) (any, error)

func baseFindFunc(c echo.Context, callback findCallback) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	result, err := callback(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.ReadAction,
		Error:  err,
		Data:   result,
	})
	return c.JSON(status, resp)
}

type updateCallback func(ctx context.Context, id uint) (uint, any, error)

func baseUpdateFunc(c echo.Context, form any, callback updateCallback) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	if err := c.Bind(form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	_, result, err := callback(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.UpdateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}

type deleteCallback func(ctx context.Context, id uint) (uint, error)

func baseDeleteFunc(c echo.Context, callback deleteCallback) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	_, err := callback(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.DeleteAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
