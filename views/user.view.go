package views

import (
	"context"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
	"github.com/salamanderman234/outsourcing-api/helpers"
)

// ----- USER VIEW -----
type userView struct{}

func NewUserView() domains.UserView {
	return &userView{}
}
func (userView) Create(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"message": "not found",
	})
}
func (userView) Read(c echo.Context) error {
	ctx := c.Request().Context()
	resp := domains.BasicResponse{
		Message: "ok",
	}
	var id uint
	echo.QueryParamsBinder(c).
		Uint("id", &id)
	if id == 0 {
		status, msg, errBody := helpers.HandleError(domains.ErrMissingId)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	user, err := domains.ServiceRegistry.UserServ.Find(ctx, id)
	if err != nil {
		status, msg, errBody := helpers.HandleError(err)
		resp.Message = msg
		resp.Body = *errBody
		return c.JSON(status, resp)
	}
	resp.Body = domains.DataBodyResponse{
		Data: user,
	}
	return c.JSON(http.StatusOK, resp)
}
func (userView) Update(c echo.Context) error {
	var data domains.UserEditForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		aff, updated, err := domains.ServiceRegistry.UserServ.Update(ctx, id, data)
		return int(aff), updated, err
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (userView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		deletedId, aff, err := domains.ServiceRegistry.UserServ.Delete(ctx, id)
		return int(deletedId), int(aff), err
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF USER VIEW -----
// ----- SERVICE USER VIEW -----
type serviceUserView struct{}

func NewServiceUserView() domains.UserServiceView {
	return &serviceUserView{}
}
func (serviceUserView) Create(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"message": "not found",
	})
}
func (serviceUserView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.ServiceUserServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (serviceUserView) Update(c echo.Context) error {
	var data domains.ServiceUserUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		aff, updated, err := domains.ServiceRegistry.ServiceUserServ.Update(ctx, id, data)
		return int(aff), updated, err
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (serviceUserView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		deletedId, aff, err := domains.ServiceRegistry.ServiceUserServ.Delete(ctx, id)
		return int(deletedId), int(aff), err
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF SERVICE USER VIEW -----
// ----- EMPLOYEE VIEW -----
type employeeView struct{}

func NewEmployeeView() domains.EmployeeView {
	return &employeeView{}
}
func (employeeView) Create(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"message": "not found",
	})
}
func (employeeView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.EmployeeServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (employeeView) Update(c echo.Context) error {
	var data domains.EmployeeUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		aff, updated, err := domains.ServiceRegistry.EmployeeServ.Update(ctx, id, data)
		return int(aff), updated, err
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (employeeView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		deletedId, aff, err := domains.ServiceRegistry.EmployeeServ.Delete(ctx, id)
		return int(deletedId), int(aff), err
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF EMPLOYEE VIEW -----
// ----- SUPERVISOR VIEW -----
type supervisorView struct{}

func NewSupervisorView() domains.SupervisorView {
	return &supervisorView{}
}
func (supervisorView) Create(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"message": "not found",
	})
}
func (supervisorView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.SupervisorServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (supervisorView) Update(c echo.Context) error {
	var data domains.SupervisorUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		aff, updated, err := domains.ServiceRegistry.SupervisorServ.Update(ctx, id, data)
		return int(aff), updated, err
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (supervisorView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		deletedId, aff, err := domains.ServiceRegistry.SupervisorServ.Delete(ctx, id)
		return int(deletedId), int(aff), err
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF SUPERVISOR VIEW -----
// ----- ADMIN VIEW -----
type adminView struct{}

func NewAdminView() domains.AdminView {
	return &adminView{}
}
func (adminView) Create(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"message": "not found",
	})
}
func (adminView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.AdminServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (adminView) Update(c echo.Context) error {
	var data domains.AdminUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		aff, updated, err := domains.ServiceRegistry.AdminServ.Update(ctx, id, data)
		return int(aff), updated, err
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (adminView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		deletedId, aff, err := domains.ServiceRegistry.AdminServ.Delete(ctx, id)
		return int(deletedId), int(aff), err
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF ADMIN VIEW -----
