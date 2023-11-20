package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

type authView struct{}

func NewAuthView() domains.BasicAuthView {
	return &authView{}
}

func (authView) Login(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	form := domains.BasicLoginForm{}
	if err := c.Bind(&form); err != nil {
		msg := err.Error()
		payload := domains.ErrorBodyResponse{
			Error: &msg,
		}
		resp.Message = domains.ErrBadRequest.Error()
		resp.Body = payload
		return c.JSON(http.StatusBadRequest, resp)
	}
	tokenPair, err := domains.AuthServiceRegistry.AuthServ.Login(ctx, form, form.Remember)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: tokenPair,
	}
	return c.JSON(http.StatusOK, resp)
}
func (authView) Register(c echo.Context) error {
	return nil
}
func (authView) Verify(c echo.Context) error {
	c.Request().Header.Get("Authorization")
	_, err := helpers.VerifyToken("dsfsd")
	return c.JSON(200, err)
}
func (authView) Refresh(c echo.Context) error {
	return nil
}
