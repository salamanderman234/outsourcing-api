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
type CategoryView interface {
	BasicCrudView
	GetIcon(c echo.Context) error
}

// ----- END OF MASTER DATA -----
// ----- USER -----
type EmployeeView interface {
	BasicCrudView
}
type AdminView interface {
	BasicCrudView
}
type SupervisorView interface {
	BasicCrudView
}
type UserServiceView interface {
	BasicCrudView
}

// ----- END OF USER -----
// ---- APP SERVICE -----
type ServiceItemView interface {
	BasicCrudView
}
type PartialServiceView interface {
	BasicCrudView
}
type ServicePackageView interface {
	BasicCrudView
}

// ---- END OF APP SERVICE ----
// ---- ORDER VIEW ----
type ServiceOrderView interface {
	MakeServiceOrder(c echo.Context) error
	CancelServiceOrder(c echo.Context) error
	ListOrder(c echo.Context) error
	UserListOrder(c echo.Context) error
	UpdateStatusServiceOrder(c echo.Context) error
}

// ---- END OF ORDER VIEW ----
