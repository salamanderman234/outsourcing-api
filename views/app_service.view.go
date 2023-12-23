package views

import (
	"context"
	"math"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

// ----- SERVICE ITEM VIEW -----
type serviceItemView struct{}

func NewServiceItemView() domains.ServiceItemView {
	return &serviceItemView{}
}
func (serviceItemView) Create(c echo.Context) error {
	var data domains.ServiceItemCreateForm
	createCallFunc := func(ctx context.Context) (domains.Entity, error) {
		return domains.ServiceRegistry.ServiceItemServ.Create(ctx, data)
	}
	return basicCreateView(c, &data, createCallFunc)
}

func (serviceItemView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.ServiceItemServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (serviceItemView) Update(c echo.Context) error {
	var data domains.ServiceItemUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		return domains.ServiceRegistry.ServiceItemServ.Update(ctx, id, data)
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (serviceItemView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		return domains.ServiceRegistry.ServiceItemServ.Delete(ctx, id)
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF SERVICE ITEM VIEW -----
// ----- SERVICE VIEW -----
type partialServiceView struct{}

func NewPartialServiceView() domains.PartialServiceView {
	return &partialServiceView{}
}
func (partialServiceView) Create(c echo.Context) error {
	var data domains.PartialServiceCreateForm
	createCallFunc := func(ctx context.Context) (domains.Entity, error) {
		fileIcon, err := readFile(c, "icon")
		if err != nil {
			return nil, err
		}
		fileImage, _ := readFile(c, "image")
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  fileIcon,
		}
		imageFileMap := domains.EntityFileMap{
			Field: "image",
			File:  fileImage,
		}
		return domains.ServiceRegistry.ServiceServ.Create(ctx, data, iconFileMap, imageFileMap)
	}
	return basicCreateView(c, &data, createCallFunc)
}

func (partialServiceView) Read(c echo.Context) error {
	callFunc := func(c context.Context, id uint, query string, page uint, orderBy string, desc bool, withPagination bool) (any, *domains.Pagination, error) {
		return domains.ServiceRegistry.ServiceServ.Read(c,
			id, query, uint(math.Max(float64(1), float64(page))), orderBy, desc, withPagination,
		)
	}
	return basicReadView(c, callFunc)
}
func (partialServiceView) Update(c echo.Context) error {
	var data domains.PartialServiceUpdateForm
	updateCallFunc := func(ctx context.Context, id uint) (int, any, error) {
		fileIcon, err := readFile(c, "icon")
		if err != nil {
			return 0, nil, err
		}
		fileImage, _ := readFile(c, "image")
		iconFileMap := domains.EntityFileMap{
			Field: "icon",
			File:  fileIcon,
		}
		imageFileMap := domains.EntityFileMap{
			Field: "image",
			File:  fileImage,
		}
		return domains.ServiceRegistry.ServiceServ.Update(ctx, id, data, iconFileMap, imageFileMap)
	}
	return basicUpdateView(c, &data, updateCallFunc)
}
func (partialServiceView) Delete(c echo.Context) error {
	deleteCallFunc := func(ctx context.Context, id uint) (int, int, error) {
		return domains.ServiceRegistry.ServiceServ.Delete(ctx, id)
	}
	return basicDeleteView(c, deleteCallFunc)
}

// ----- END OF SERVICE VIEW -----
// ----- SERVICE PACKAGE VIEW -----
// type servicePackageView struct {}
// func NewServicePackageView() domains.ServicePackageView {
// 	return &servicePackageView{}
// }
// func(servicePackageView) Create(c echo.Context) error {
// 	var data domains.ServicePackageCreateForm
// 	createCallFunc := func(ctx context.Context) (domains.Entity, error) {
// 		fileIcon, err := readFile(c, "icon")
// 		if err != nil {
// 			return nil, err
// 		}
// 		fileImage, _ := readFile(c, "image")
// 		iconFileMap := domains.EntityFileMap{
// 			Field: "icon",
// 			File:  fileIcon,
// 		}
// 		imageFileMap := domains.EntityFileMap{
// 			Field: "image",
// 			File:  fileImage,
// 		}
// 		return domains.ServiceRegistry.ServiceServ.Create(ctx, data, iconFileMap, imageFileMap)
// 	}
// 	return basicCreateView(c, &data, createCallFunc)
// }
// ----- END OF SERVICE PACKAGE VIEW -----
