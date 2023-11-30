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
// -> Basic
type BasicCrudView interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

// ----- MASTER DATA -----
type ServiceCategoryView interface {
	BasicCrudView
}
type DistrictView interface {
	BasicCrudView
}
type SubDistrictView interface {
	BasicCrudView
}
type VillageView interface {
	BasicCrudView
}

//----- END OF MASTER DATA -----
