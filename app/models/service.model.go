package models

type ServiceModel struct {
	Model
	CategoryID      uint                         `json:"category_id,omitempty"`
	Category        *Category                    `json:"category,omitempty" visible:"true"`
	RequiredItems   []RequiredItemServiceModel   `json:"required_items,omitempty" visible:"true"`
	AdditionalItems []AdditionalItemServiceModel `json:"additional_items,omitempty" visible:"true"`
	ServiceName     string                       `json:"service_name,omitempty" visible:"true"`
	MainImage       string                       `json:"main_image,omitempty" visible:"true"`
	Description     string                       `json:"description,omitempty" visible:"true"`
	Includes        string                       `json:"includes,omitempty" visible:"true"`
	TotalPrice      uint                         `json:"total_price,omitempty" visible:"true"`
	EmployeePrice   uint                         `json:"employee_price,omitempty" visible:"true"`
	ServicePrice    uint                         `json:"service_price,omitempty" visible:"true"`
	EtcPrice        uint                         `json:"etc_price,omitempty" visible:"true"`
	HourWork        uint                         `json:"hour_work,omitempty" visible:"true"`
}

func (ServiceModel) GetTableName() string {
	return "services"
}

type RequiredItemServiceModel struct {
	Model
	ServiceID   *uint         `json:"service_id,omitempty"`
	Service     *ServiceModel `json:"service,omitempty" visible:"true"`
	ItemName    *string       `json:"item_name,omitempty" visible:"true"`
	Quantity    *uint         `json:"quantity,omitempty" visible:"true"`
	Description *string       `json:"description,omitempty" visible:"true"`
}

func (RequiredItemServiceModel) GetTableName() string {
	return "required_service_items"
}

type AdditionalItemServiceModel struct {
	RequiredItemServiceModel
}

func (AdditionalItemServiceModel) GetTableName() string {
	return "additional_service_items"
}
