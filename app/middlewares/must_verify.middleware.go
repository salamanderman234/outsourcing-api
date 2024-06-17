package custom_middlewares

import (
	"github.com/labstack/echo/v4"
)

func MustVerifyUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// ctx := c.Request().Context()
		// claims, _ := ctx.Value(configs.VarConfig.UserContextName).(types.JWTCLaims)
		// if claims.ID == "" {
		// 	return next(c)
		// }
		// verifiedAt := claims.V
		// if verifiedAt == "" {
		// 	status, resp := helpers.Response.CreateResponse(types.ResponseParams{
		// 		Error: types.ErrNotVerifiedUser,
		// 	})
		// 	return c.JSON(status, resp)
		// }
		return next(c)
	}
}
