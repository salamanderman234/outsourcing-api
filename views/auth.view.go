package views

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/configs"
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
		if errors.Is(err, domains.ErrRecordNotFound) {
			err = domains.ErrInvalidCreds
		}
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	cookie := http.Cookie{
		Name:     configs.REFRESH_TOKEN_COOKIE_NAME,
		Value:    tokenPair.Refresh,
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	resp.Body = domains.DataBodyResponse{
		Data: tokenPair,
	}
	return c.JSON(http.StatusOK, resp)
}
func (authView) Register(c echo.Context) error {
	return nil
}
func (authView) Verify(c echo.Context) error {
	resp := domains.BasicResponse{
		Message: "ok",
	}
	ctx := c.Request().Context()
	token := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
	claims, err := domains.AuthServiceRegistry.AuthServ.Check(ctx, token)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: claims,
	}
	return c.JSON(http.StatusOK, resp)
}
func (authView) Refresh(c echo.Context) error {
	resp := domains.BasicResponse{
		Message: "ok",
	}
	ctx := c.Request().Context()
	cookie, err := c.Cookie(configs.REFRESH_TOKEN_COOKIE_NAME)
	if errors.Is(err, echo.ErrCookieNotFound) || cookie == nil {
		resp.Message = domains.ErrInvalidAccess.Error()
		msg := "refresh token cookie is required"
		resp.Body = domains.ErrorBodyResponse{
			Error: &msg,
		}
		return c.JSON(http.StatusForbidden, resp)
	}
	token := cookie.Value
	pair, err := domains.AuthServiceRegistry.AuthServ.Refresh(ctx, token)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: pair,
	}
	return c.JSON(http.StatusOK, resp)
}
