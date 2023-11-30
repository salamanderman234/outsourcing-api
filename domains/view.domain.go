package domains

import "github.com/labstack/echo/v4"

// ----- AUTH VIEW -----
type BasicAuthView interface {
	Login(c echo.Context) error
	RegisterAdmin(c echo.Context) error
	RegisterUserService(c echo.Context) error
	RegisterEmployee(c echo.Context) error
	RegisterSupervisor(c echo.Context) error
	Verify(c echo.Context) error
	Refresh(c echo.Context) error
}

// ----- END OF AUTH VIEW -----
