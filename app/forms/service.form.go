package forms

// create form
type RequiredItemAddForm struct {
	ServiceID    uint   `json:"service_id" valid:"required,int"`
	ItemName     string `json:"item_name" valid:"required,stringlength(1|255)"`
	Quantity     uint   `json:"quantity,omitempty" valid:"required,int"`
	PricePerItem uint   `json:"price_per_item,omitempty" valid:"optional,int"`
	Description  string `json:"description" valid:"optional,stringlength(1|3000)"`
}
type AdditionalItemServiceAddForm struct {
	ServiceID    uint   `json:"service_id" valid:"required,int"`
	ItemName     string `json:"item_name" valid:"required,stringlength(1|255)"`
	MinQuantity  uint   `json:"min_quantity,omitempty" valid:"optional,int"`
	MaxQuantity  uint   `json:"max_quantity,omitempty" valid:"optional,int"`
	PricePerItem uint   `json:"price_per_item,omitempty" valid:"required,int"`
	Description  string `json:"description" valid:"optional,stringlength(1|3000)"`
}
type RequiredItemServiceCreateForm struct {
	ItemName     string `json:"item_name" valid:"required,stringlength(1|255)"`
	Quantity     uint   `json:"quantity,omitempty" valid:"optional,int"`
	PricePerItem uint   `json:"price_per_item,omitempty" valid:"optional,int"`
	Description  string `json:"description" valid:"optional,stringlength(1|3000)"`
}

type AdditionalItemServiceCreateForm struct {
	ItemName     string `json:"item_name" valid:"required,stringlength(1|255)"`
	MinQuantity  uint   `json:"min_quantity,omitempty" valid:"optional,int"`
	MaxQuantity  uint   `json:"max_quantity,omitempty" valid:"optional,int"`
	PricePerItem uint   `json:"price_per_item,omitempty" valid:"required,int"`
	Description  string `json:"description" valid:"optional,stringlength(1|3000)"`
}

type ServiceCreateForm struct {
	CategoryID      uint                              `json:"category_id" valid:"required,int"`
	RequiredItems   []RequiredItemServiceCreateForm   `json:"required_items" valid:"optional"`
	AdditionalItems []AdditionalItemServiceCreateForm `json:"additional_items" valid:"optional"`
	ServiceName     string                            `json:"service_name" valid:"required,stringlength(1|255)"`
	MainImage       string                            `json:"main_image"`
	Icon            string                            `json:"icon"`
	Description     string                            `json:"description" valid:"optional,stringlength(1|30000)"`
	Includes        string                            `json:"includes" valid:"optional,stringlength(1|30000)"`
	EmployeePrice   uint                              `json:"employee_price" valid:"required,int"`
	ServicePrice    uint                              `json:"service_price" valid:"required,int"`
	Discount        uint                              `json:"discount" valid:"optional,int"`
	HourWork        uint                              `json:"hour_work" valid:"optional,int"`
}

// update form
type RequiredItemServiceUpdateForm struct {
	RequiredItemServiceID uint    `json:"required_item_service_id" valid:"required,int"`
	ItemName              *string `json:"item_name" valid:"optional,stringlength(1|255)"`
	Quantity              uint    `json:"quantity,omitempty" valid:"required,int"`
	PricePerItem          *uint   `json:"price_per_item,omitempty" valid:"optional,int"`
	Description           *string `json:"description" valid:"optional,stringlength(1|3000)"`
}

type AdditionalItemServiceUpdateForm struct {
	AdditionalItemServiceID uint    `json:"required_item_service_id" valid:"required,int"`
	ItemName                *string `json:"item_name" valid:"optional,stringlength(1|255)"`
	MinQuantity             *uint   `json:"min_quantity,omitempty" valid:"optional,int"`
	MaxQuantity             *uint   `json:"max_quantity,omitempty" valid:"optional,int"`
	PricePerItem            *uint   `json:"price_per_item,omitempty" valid:"optional,int"`
	Description             *string `json:"description" valid:"optional,stringlength(1|3000)"`
}

type ServiceUpdateForm struct {
	CategoryID *uint `json:"category_id" valid:"optional,int"`
	// RequiredItems   *[]RequiredItemServiceUpdateForm   `json:"required_items" valid:"optional"`
	// AdditionalItems *[]AdditionalItemServiceUpdateForm `json:"additional_items" valid:"optional"`
	ServiceName   *string `json:"service_name" valid:"optional,stringlength(1|255)"`
	MainImage     *string `json:"main_image" valid:"optional"`
	Icon          *string `json:"icon" valid:"optional"`
	Description   *string `json:"description" valid:"optional,stringlength(1|30000)"`
	Includes      *string `json:"includes" valid:"optional,stringlength(1|30000)"`
	EmployeePrice *uint   `json:"employee_price" valid:"optional,int"`
	ServicePrice  *uint   `json:"service_price" valid:"optional,int"`
	Discount      *uint   `json:"discount" valid:"optional,int"`
	HourWork      *uint   `json:"hour_work" valid:"optional,int"`
}
