package views

import (
	"context"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type orderView struct{}

func NewOrderView() domains.ServiceOrderView {
	return &orderView{}
}

func (orderView) MakeServiceOrder(c echo.Context) error {
	var data domains.ServiceOrderForm
	user := c.Get("user").(domains.UserEntity)
	createCallFunc := func(ctx context.Context) (domains.Entity, error) {
		return domains.ServiceRegistry.OrderServ.MakeOrder(ctx, user, data)
	}
	return basicCreateView(c, &data, createCallFunc)
}
func (orderView) CancelServiceOrder(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	if err := echo.QueryParamsBinder(c).Uint("id", &id).BindError(); err != nil || id == 0 {
		status, msg, errBody := helpers.HandleError(domains.ErrMissingId)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	ok, err := domains.ServiceRegistry.OrderServ.CancelOrder(ctx, domains.UserEntity{}, id)
	if !ok {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
func (orderView) ListOrder(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	var status string
	var serviceUserId uint
	var page uint
	var order string
	withPagination := -1
	desc := -1
	echo.QueryParamsBinder(c).
		Uint("id", &id).
		Uint("page", &page).
		Uint("service-user-id", &serviceUserId).
		String("status", &status).
		String("order-by", &order).
		Int("desc", &desc).
		Int("with-pagination", &withPagination)
	if desc == -1 {
		desc = 1
	}
	if withPagination == -1 {
		withPagination = 1
	}
	var err error
	var datas any
	var pagination *domains.Pagination
	if id != 0 {
		datas, err = domains.ServiceRegistry.OrderServ.DetailOrder(ctx, domains.UserEntity{}, id)
	} else {
		datas, pagination, err = domains.ServiceRegistry.OrderServ.ListOrder(
			ctx, domains.UserEntity{},
			serviceUserId, status, uint(math.Max(float64(1), float64(page))),
			order, desc > 0, withPagination > 0,
		)
	}

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
func (orderView) UploadMOU(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	user := c.Get("user")
	if user == nil {
		status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	userCon, ok := user.(domains.UserEntity)
	if !ok {
		status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	var orderId uint
	if err := echo.QueryParamsBinder(c).Uint("id", &orderId).BindError(); err != nil {
		status, msg, errBody := helpers.HandleError(domains.ErrMissingId)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	if orderId == 0 {
		status, msg, errBody := helpers.HandleError(domains.ErrMissingId)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	file, err := readFile(c, "mou")
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	iconFileMap := domains.EntityFileMap{
		Field: "mou",
		File:  file,
	}
	ok, err = domains.ServiceRegistry.OrderServ.UploadMOU(ctx, userCon, orderId, iconFileMap)
	if !ok {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: map[string]string{
			"message": "success",
		},
	}
	return c.JSON(http.StatusOK, resp)
}
func (orderView) UserListOrder(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	user := c.Get("user")
	if user == nil {
		status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	userCon, ok := user.(domains.UserEntity)
	if !ok {
		status, msg, errBody := helpers.HandleError(domains.ErrInvalidAccess)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	var status string
	var page uint
	var order string
	withPagination := -1
	desc := -1
	echo.QueryParamsBinder(c).
		Uint("page", &page).
		String("status", &status).
		String("order-by", &order).
		Int("desc", &desc).
		Int("with-pagination", &withPagination)
	if desc == -1 {
		desc = 1
	}
	if withPagination == -1 {
		withPagination = 1
	}

	var pagination *domains.Pagination
	datas, pagination, err := domains.ServiceRegistry.OrderServ.ListOrder(
		ctx, domains.UserEntity{},
		userCon.ID, status, uint(math.Max(float64(1), float64(page))),
		order, desc > 0, withPagination > 0,
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
	bodyResp.Datas = datas
	resp.Body = bodyResp
	return c.JSON(http.StatusOK, resp)
}
func (orderView) UpdateStatusServiceOrder(c echo.Context) error {
	var data domains.ServiceOrderUpdateStatusForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		return domains.ServiceRegistry.OrderServ.UpdateOrderStatus(ctx,
			domains.UserEntity{}, id, data,
		)
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
