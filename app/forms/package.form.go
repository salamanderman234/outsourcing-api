package forms

// create form
type PackageServiceItemCreateForm struct {
	AdditionalItemServiceID uint `json:"additional_item_service_id" valid:"required"`
	Quantity                uint `json:"quantity" valid:"required"`
}

type PackageServiceCreateForm struct {
	ServiceID                     uint                            `json:"service_id" valid:"required"`
	TotalEmployee                 uint                            `json:"total_employee" valid:"required"`
	AdditionalPackageServiceItems *[]PackageServiceItemCreateForm `json:"additional_package_service_items" valid:"optional"`
}

type PackageCreateForm struct {
	PackageName string                     `json:"package_name" valid:"required,stringlength(1|255)"`
	MainImage   string                     `json:"main_image"`
	Description string                     `json:"description" valid:"required,stringlength(1|50000)"`
	Includes    string                     `json:"includes" valid:"required,stringlength(1|50000)"`
	MinContract uint                       `json:"min_contract" valid:"required"`
	Discount    uint                       `json:"discount" valid:"optional,int"`
	Services    []PackageServiceCreateForm `json:"services" valid:"required"`
}

// update form
type PackageServiceItemUpdateForm struct {
	PackageServiceItemID uint  `json:"package_service_item_id" valid:"required"`
	Quantity             *uint `json:"quantity" valid:"optional"`
}
type PackageServiceUpdateForm struct {
	PackageServiceID              uint                            `json:"package_service_id" valid:"required"`
	TotalEmployee                 *uint                           `json:"total_employee" valid:"optional"`
	AdditionalPackageServiceItems *[]PackageServiceItemUpdateForm `json:"additional_package_service_items" valid:"optional"`
}
type PackageUpdateForm struct {
	PackageName *string `json:"package_name" valid:"optional,stringlength(1|255)"`
	MainImage   *string `json:"main_image"`
	Description *string `json:"description" valid:"optional,stringlength(1|50000)"`
	Includes    *string `json:"includes" valid:"optional,stringlength(1|50000)"`
	MinContract *uint   `json:"min_contract" valid:"optional"`
	Discount    *uint   `json:"discount" valid:"optional"`
	// Services    *[]PackageServiceUpdateForm `json:"services" valid:"optional"`
}
