package views

import (
	"strconv"

	"github.com/labstack/echo/v4"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/providers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
)

type paymentView struct{}

func NewPaymentView() view_domains.PaymentViewInterface {
	return &paymentView{}
}

func (paymentView) Pay(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()

	token, redir, err := providers.ServiceProvider.MidtransService.CreatePaymentFromTransaction(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.CreateAction,
		Error:  err,
		Data: map[string]string{
			"midtrans_token": token,
			"redirect_url":   redir,
		},
	})
	return c.JSON(status, resp)
}
func (paymentView) AfterPayHook(c echo.Context) error {
	var form forms.PaymentNotifForm
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(types.ResponseParams{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	err := providers.ServiceProvider.MidtransService.AfterPaymentAction(ctx, form)
	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		Action: enums.UpdateAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}
