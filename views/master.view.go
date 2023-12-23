package views

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

// --> Category
type categoryView struct{}

func NewCategoryView() domains.CategoryView {
	return &categoryView{}
}
func (categoryView) Create(c echo.Context) error {
	var data domains.CategoryCreateForm
	createCallFunc := func(ctx context.Context) (domains.Entity, error) {
		file, err := readFile(c, "icon")
		if err != nil {
			return nil, err
		}
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  file,
		}
		return domains.ServiceRegistry.CategoryServ.Create(ctx, data, iconFileMap)
	}
	return basicCreateView(c, &data, createCallFunc)
}
func (categoryView) GetIcon(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))
	category, _, err := domains.ServiceRegistry.CategoryServ.Read(ctx, uint(id), "", 1, "", true, false)
	if errors.Is(err, domains.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	categoryEntity := category.(domains.CategoryEntity)
	path := categoryEntity.Icon
	return c.File(path)
}
func (categoryView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.CategoryServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (categoryView) Update(c echo.Context) error {
	var data domains.CategoryUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		file, err := readFile(c, "icon")
		if err != nil {
			return 0, nil, err
		}
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  file,
		}
		return domains.ServiceRegistry.CategoryServ.Update(ctx, id, data, iconFileMap)
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (categoryView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		return domains.ServiceRegistry.CategoryServ.Delete(ctx, id)
	}
	return basicDeleteView(c, deleteCallFunc)
}
