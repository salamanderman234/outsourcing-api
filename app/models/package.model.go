package models

type PackageModel struct {
	Model
	PackageName   *string               `json:"package_name,omitempty" visible:"true"`
	MainImage     *string               `json:"main_image,omitempty" visible:"true"`
	Description   *string               `json:"description,omitempty" visible:"true"`
	Includes      *string               `json:"includes,omitempty" visible:"true"`
	MinContract   *uint                 `json:"min_contract,omitempty" visible:"true"`
	TotalPrice    *uint                 `json:"total_price,omitempty" visible:"true"`
	EmployeePrice *uint                 `json:"employee_price,omitempty" visible:"true"`
	ServicePrice  *uint                 `json:"service_price,omitempty" visible:"true"`
	EtcPrice      *uint                 `json:"etc_price,omitempty" visible:"true"`
	Services      []PackageServiceModel `json:"services" visible:"true"`
}

func (PackageModel) GetTableName() string {
	return "packages"
}

type PackageServiceModel struct {
	Model
	PackageID                     *uint                               `json:"package_id,omitempty"`
	Package                       *PackageModel                       `json:"package,omitempty" visible:"true"`
	ServiceID                     *uint                               `json:"service_id,omitempty"`
	Service                       *ServiceModel                       `json:"service,omitempty" visible:"true"`
	SubTotalPrice                 *uint                               `json:"sub_total_price,omitempty" visible:"true"`
	TotalEmployee                 *uint                               `json:"total_employee,omitempty" visible:"true"`
	AdditionalPackageServiceItems []PackageServiceAdditionalItemModel `json:"additional_package_service_items" visible:"true"`
}

func (PackageServiceModel) GetTableName() string {
	return "package_services"
}

type PackageServiceAdditionalItemModel struct {
	Model
	PackageServiceID        *uint                      `json:"package_service_id,omitempty"`
	PackageService          *PackageServiceModel       `json:"package_service" visible:"true"`
	AdditionalItemServiceID *uint                      `json:"additional_item_service_id,omitempty"`
	AdditionalItemService   AdditionalItemServiceModel `json:"additional_item_service" visible:"true"`
	Quantity                *uint                      `json:"quantity,omitempty" visible:"true"`
	SubTotalPrice           *uint                      `json:"sub_total_price,omitempty" visible:"true"`
}

func (PackageServiceAdditionalItemModel) GetTableName() string {
	return "package_service_additional_items"
}
