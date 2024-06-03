package custom_middlewares

import (
	"context"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/app/types/enums"
	"github.com/salamanderman234/outsourcing-api/configs"
)

func RetrieveAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return next(c)
		}
		token = strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
		if token == "" {
			status, resp := helpers.Response.CreateResponse(types.ResponseParams{
				Error: types.ErrForbiden,
			})
			return c.JSON(status, resp)
		}
		payload, err := helpers.JWT.VerifyToken(token)
		if err != nil {
			status, resp := helpers.Response.CreateResponse(types.ResponseParams{
				Error: types.ErrForbiden,
			})
			return c.JSON(status, resp)
		}
		email := payload.Email
		role := payload.Role
		tokenType := payload.Subject
		if tokenType != string(enums.AccessTokenType) {
			status, resp := helpers.Response.CreateResponse(types.ResponseParams{
				Error: types.ErrForbiden,
			})
			return c.JSON(status, resp)
		}
		if email != "" {
			uri := c.Request().RequestURI
			helpers.Logger.Info(
				fmt.Sprintf("(Session) User (%s) %s is attempting to access %s", role, email, uri),
			)
		}

		c.Set(string(configs.VarConfig.AccessContextName), payload)
		c.SetRequest(c.Request().WithContext(
			context.WithValue(
				c.Request().Context(),
				configs.VarConfig.AccessContextName,
				payload),
		),
		)
		return next(c)
	}
}
