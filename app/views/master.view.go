package views

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/enums"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	view_domains "github.com/salamanderman234/outsourcing-api/app/domains/views"
	"github.com/salamanderman234/outsourcing-api/app/forms"
	"github.com/salamanderman234/outsourcing-api/app/helpers"
)

// province master
type masterProvinceView struct{}

func NewMasterProvinceView() view_domains.MasterProvinceViewInterface {
	return &masterProvinceView{}
}

func (masterProvinceView) Create(c echo.Context) error {
	var form forms.MasterProviceCreateForm
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	result, err := domains.ServiceRegistry.MasterProvinceService.Create(ctx, form)
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.CreateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (masterProvinceView) Read(c echo.Context) error {
	q := c.QueryParam("query")
	idParam := c.QueryParam("page")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	results, pagination, err := domains.ServiceRegistry.MasterProvinceService.Read(ctx, q, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action:     enums.ReadAction,
		Error:      err,
		Datas:      results,
		Pagination: pagination,
	})
	return c.JSON(status, resp)
}
func (masterProvinceView) Find(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	result, err := domains.ServiceRegistry.MasterProvinceService.Find(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.ReadAction,
		Error:  err,
		Data:   result,
	})
	return c.JSON(status, resp)
}
func (masterProvinceView) Update(c echo.Context) error {
	var form forms.MasterProviceUpdateForm
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	_, result, err := domains.ServiceRegistry.MasterProvinceService.Update(ctx, uint(id), form)
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.UpdateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (masterProvinceView) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	_, err := domains.ServiceRegistry.MasterProvinceService.Delete(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.DeleteAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}

// end of province master

// regency master
type regencyMasterView struct{}

func NewRegencyMasterView() view_domains.MasterRegencyViewInterface {
	return &regencyMasterView{}
}
func (regencyMasterView) Create(c echo.Context) error {
	var form forms.MasterRegencyCreateForm
	ctx := c.Request().Context()
	if err := (&echo.DefaultBinder{}).BindBody(c, &form); err != nil {
		status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	result, err := domains.ServiceRegistry.MasterRegencyService.Create(ctx, form)
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.CreateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (regencyMasterView) Read(c echo.Context) error {
	q := c.QueryParam("query")
	idParam := c.QueryParam("page")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	results, pagination, err := domains.ServiceRegistry.MasterRegencyService.Read(ctx, q, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action:     enums.ReadAction,
		Error:      err,
		Datas:      results,
		Pagination: pagination,
	})
	return c.JSON(status, resp)
}
func (regencyMasterView) Find(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	result, err := domains.ServiceRegistry.MasterRegencyService.Find(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.ReadAction,
		Error:  err,
		Data:   result,
	})
	return c.JSON(status, resp)
}
func (regencyMasterView) Update(c echo.Context) error {
	var form forms.MasterRegencyUpdateForm
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	if err := c.Bind(&form); err != nil {
		status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
			Error: err,
		})
		return c.JSON(status, resp)
	}
	_, result, err := domains.ServiceRegistry.MasterRegencyService.Update(ctx, uint(id), form)
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.UpdateAction,
		Data:   result,
		Error:  err,
	})
	return c.JSON(status, resp)
}
func (regencyMasterView) Delete(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	ctx := c.Request().Context()
	_, err := domains.ServiceRegistry.MasterRegencyService.Delete(ctx, uint(id))
	status, resp := helpers.Response.CreateResponse(responses.ResponseConfig{
		Action: enums.DeleteAction,
		Error:  err,
	})
	return c.JSON(status, resp)
}

// end of regency master
