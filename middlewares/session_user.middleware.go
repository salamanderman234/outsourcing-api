package middlewares

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

func RetrieveSessionUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return next(c)
		}
		token = strings.ReplaceAll(token, "Bearer ", "")
		ctx := c.Request().Context()
		resp := domains.BasicResponse{}
		payload, err := domains.ServiceRegistry.AuthServ.Check(ctx, token)
		if err != nil {
			status, msg, errBody := helpers.HandleError(err)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		}
		user, err := domains.ServiceRegistry.UserService.Find(ctx, payload.ID)
		if errors.Is(err, domains.ErrRecordNotFound) {
			status, msg, errBody := helpers.HandleError(err)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		} else if err != nil {
			status, msg, errBody := helpers.HandleError(domains.ErrInvalidToken)
			resp.Message = msg
			resp.Body = *errBody
			return c.JSON(status, resp)
		}
		c.Set("user", user)
		return next(c)
	}
}
