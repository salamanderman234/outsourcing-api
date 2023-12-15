package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func registerMasterRoutes(group *echo.Group) {
	// service category
	serviceCategoryGroup := group.Group("categories")
	serviceCategoryGroup.GET("/", domains.ViewRegistry.CategoryView.Read)
	serviceCategoryGroup.POST("/", domains.ViewRegistry.CategoryView.Create)
	serviceCategoryGroup.PATCH("/", domains.ViewRegistry.CategoryView.Update)
	serviceCategoryGroup.DELETE("/", domains.ViewRegistry.CategoryView.Delete)
	serviceCategoryGroup.GET("/icon/:id", domains.ViewRegistry.CategoryView.GetIcon)
	// district
	districtGroup := group.Group("districts")
	districtGroup.GET("/", func(c echo.Context) error { return nil })
	districtGroup.POST("/", func(c echo.Context) error { return nil })
	districtGroup.PATCH("/", func(c echo.Context) error { return nil })
	districtGroup.DELETE("/", func(c echo.Context) error { return nil })
	// subdistrict
	subDistrictGroup := group.Group("sub-districts")
	subDistrictGroup.GET("/", func(c echo.Context) error { return nil })
	subDistrictGroup.POST("/", func(c echo.Context) error { return nil })
	subDistrictGroup.PATCH("/", func(c echo.Context) error { return nil })
	subDistrictGroup.DELETE("/", func(c echo.Context) error { return nil })
	// village
	villageGroup := group.Group("villages")
	villageGroup.GET("/", func(c echo.Context) error { return nil })
	villageGroup.POST("/", func(c echo.Context) error { return nil })
	villageGroup.PATCH("/", func(c echo.Context) error { return nil })
	villageGroup.DELETE("/", func(c echo.Context) error { return nil })

}
