package models

type Service struct {
	Model
	CategoryID      *uint                   `json:"category_id,omitempty"`
	Category        *Category               `json:"category,omitempty" visible:"true"`
	RequiredItems   []RequiredItemService   `json:"required_items,omitempty" visible:"true"`
	AdditionalItems []AdditionalItemService `json:"additional_items,omitempty" visible:"true"`
	ServiceName     *string                 `json:"service_name,omitempty" visible:"true"`
	MainImage       *string                 `json:"main_image,omitempty" visible:"true"`
	Description     *string                 `json:"description,omitempty" visible:"true"`
	Includes        *string                 `json:"includes,omitempty" visible:"true"`
	EmployeePrice   *uint                   `json:"employee_price,omitempty" visible:"true"`
	ServicePrice    *uint                   `json:"service_price,omitempty" visible:"true"`
	EtcPrice        *uint                   `json:"etc_price,omitempty" visible:"true"`
	Discount        *uint                   `json:"discount,omitempty" visible:"true"`
	TotalPrice      *uint                   `json:"total_price,omitempty" visible:"true"`
	HourWork        *uint                   `json:"hour_work,omitempty" visible:"true"`
}

type RequiredItemService struct {
	Model
	ServiceID    *uint    `json:"service_id,omitempty"`
	Service      *Service `json:"service,omitempty" visible:"true"`
	ItemName     *string  `json:"item_name,omitempty" visible:"true"`
	Quantity     *uint    `json:"quantity,omitempty" visible:"true"`
	PricePerItem *uint    `json:"price_per_item,omitempty" visible:"true" gorm:"default:0"`
	Subtotal     *uint    `json:"subtotal,omitempty" visible:"true" gorm:"default:0"`
	Description  *string  `json:"description,omitempty" visible:"true"`
}

type AdditionalItemService struct {
	Model
	ServiceID    *uint    `json:"service_id,omitempty"`
	Service      *Service `json:"service,omitempty" visible:"true"`
	ItemName     *string  `json:"item_name,omitempty" visible:"true"`
	MinQuantity  *uint    `json:"min_quantity,omitempty" visible:"true" gorm:"default:0"`
	MaxQuantity  *uint    `json:"max_quantity,omitempty" visible:"true" gorm:"default:0"`
	PricePerItem *uint    `json:"price_per_item,omitempty" visible:"true" gorm:"default:0"`
	Description  *string  `json:"description,omitempty" visible:"true"`
}
