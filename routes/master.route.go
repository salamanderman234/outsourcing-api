package routes

import "github.com/labstack/echo/v4"

func registerMasterRoutes(group *echo.Group) {
	// service category
	serviceCategoryGroup := group.Group("service-categories")
	serviceCategoryGroup.GET("/", func(c echo.Context) error { return nil })
	serviceCategoryGroup.POST("/", func(c echo.Context) error { return nil })
	serviceCategoryGroup.PATCH("/", func(c echo.Context) error { return nil })
	serviceCategoryGroup.DELETE("/", func(c echo.Context) error { return nil })
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
