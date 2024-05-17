package models

type Package struct {
	Model
	IsCustom      *bool            `json:"is_custom" gorm:"default:0"`
	PackageName   *string          `json:"package_name,omitempty" visible:"true"`
	MainImage     *string          `json:"main_image,omitempty" visible:"true"`
	Description   *string          `json:"description,omitempty" visible:"true"`
	Includes      *string          `json:"includes,omitempty" visible:"true"`
	MinContract   *uint            `json:"min_contract,omitempty" visible:"true"`
	EmployeePrice *uint            `json:"employee_price,omitempty" visible:"true"`
	ServicePrice  *uint            `json:"service_price,omitempty" visible:"true"`
	EtcPrice      *uint            `json:"etc_price,omitempty" visible:"true"`
	Discount      *uint            `json:"discount,omitempty" visible:"true"`
	TotalPrice    *uint            `json:"total_price,omitempty" visible:"true"`
	Services      []PackageService `json:"services" visible:"true"`
}

type PackageService struct {
	Model
	PackageID                     *uint                          `json:"package_id,omitempty"`
	Package                       *Package                       `json:"package,omitempty" visible:"true"`
	ServiceID                     *uint                          `json:"service_id,omitempty"`
	Service                       *Service                       `json:"service,omitempty" visible:"true"`
	TotalEmployee                 *uint                          `json:"total_employee,omitempty" visible:"true"`
	EmployeePrice                 *uint                          `json:"employee_price,omitempty" visible:"true"`
	ServicePrice                  *uint                          `json:"service_price,omitempty" visible:"true"`
	SubTotalPrice                 *uint                          `json:"sub_total_price,omitempty" visible:"true"`
	AdditionalPackageServiceItems []PackageServiceAdditionalItem `json:"additional_package_service_items" visible:"true"`
}

type PackageServiceAdditionalItem struct {
	Model
	PackageServiceID        *uint                 `json:"package_service_id,omitempty"`
	PackageService          *PackageService       `json:"package_service" visible:"true"`
	AdditionalItemServiceID *uint                 `json:"additional_item_service_id,omitempty"`
	AdditionalItemService   AdditionalItemService `json:"additional_item_service" visible:"true"`
	Quantity                *uint                 `json:"quantity,omitempty" visible:"true"`
	Price                   *uint                 `json:"price,omitempty" visible:"true"`
	SubTotalPrice           *uint                 `json:"sub_total_price,omitempty" visible:"true"`
}
