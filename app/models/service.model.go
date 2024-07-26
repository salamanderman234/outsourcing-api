package models

type Service struct {
	Model
	CategoryID      *uint                   `json:"category_id"`
	Category        *Category               `json:"category" visible:"true"`
	RequiredItems   []RequiredItemService   `json:"required_items" visible:"true"`
	AdditionalItems []AdditionalItemService `json:"additional_items" visible:"true"`
	ServiceName     *string                 `json:"service_name" visible:"true"`
	Icon            *string                 `json:"icon" visible:"true"`
	MainImage       *string                 `json:"main_image" visible:"true"`
	Description     *string                 `json:"description" visible:"true"`
	Includes        *string                 `json:"includes" visible:"true"`
	EmployeePrice   *uint                   `json:"employee_price" visible:"true"`
	ServicePrice    *uint                   `json:"service_price" visible:"true"`
	EtcPrice        *uint                   `json:"etc_price" visible:"true"`
	Discount        *uint                   `json:"discount" visible:"true"`
	TotalPrice      *uint                   `json:"total_price" visible:"true"`
	HourWork        *uint                   `json:"hour_work" visible:"true"`
}

type RequiredItemService struct {
	Model
	ServiceID    *uint    `json:"service_id"`
	Service      *Service `json:"service" visible:"true"`
	ItemName     *string  `json:"item_name" visible:"true"`
	Quantity     *uint    `json:"quantity" visible:"true"`
	PricePerItem *uint    `json:"price_per_item" visible:"true" gorm:"default:0"`
	Subtotal     *uint    `json:"subtotal" visible:"true" gorm:"default:0"`
	Description  *string  `json:"description" visible:"true"`
}

type AdditionalItemService struct {
	Model
	ServiceID    *uint    `json:"service_id"`
	Service      *Service `json:"service" visible:"true"`
	ItemName     *string  `json:"item_name" visible:"true"`
	MinQuantity  *uint    `json:"min_quantity" visible:"true" gorm:"default:0"`
	MaxQuantity  *uint    `json:"max_quantity" visible:"true" gorm:"default:0"`
	PricePerItem *uint    `json:"price_per_item" visible:"true" gorm:"default:0"`
	Description  *string  `json:"description" visible:"true"`
}
