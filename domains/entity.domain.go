package domains

import (
	"mime/multipart"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/salamanderman234/outsourcing-api/configs"
)

// basic
type Entity interface {
	IsEntity() bool
}
type BasicEntity struct {
	ID uint `json:"id,omitempty"`
}

type FileWrapper struct {
	File   *multipart.FileHeader
	Config configs.FileConfig
	Dest   string
	Field  string
}

func (BasicEntity) IsEntity() bool {
	return true
}

// response entity
type BasicResponse struct {
	Message string `json:"message"`
	Body    any    `json:"payload"`
}

type DataBodyResponse struct {
	Data       any         `json:"data,omitempty"`
	Datas      any         `json:"datas,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Updated    any         `json:"updated,omitempty"`
	ID         *uint       `json:"id,omitempty"`
	Affected   *int        `json:"affected_row,omitempty"`
}

type ErrorBodyResponse struct {
	Error  *string               `json:"error,omitempty"`
	Errors []ErrorDetailResponse `json:"errors,omitempty"`
}

type ErrorDetailResponse struct {
	Field  *string `json:"field"`
	Rule   *string `json:"rule"`
	Detail *string `json:"detail"`
}

// file
type EntityFileMap struct {
	Field string
	File  *multipart.FileHeader
}

// pagination entity
type Pagination struct {
	Next     uint           `json:"next"`
	Current  uint           `json:"current"`
	Previous uint           `json:"previous"`
	MaxPage  uint           `json:"max_page"`
	Queries  map[string]any `json:"queries"`
}

type PaginationQuery struct {
	Query string `json:"query,omitempty"`
	Value any    `json:"value,omitempty"`
}

// authentication entity
type JWTClaims struct {
	jwt.RegisteredClaims
	JWTPayload
}
type JWTPayload struct {
	ID         uint    `json:"id"`
	Username   *string `json:"username"`
	Role       *string `json:"role"`
	ProfilePic *string `json:"profile_pic"`
}
type TokenPair struct {
	Refresh string `json:"refresh_token"`
	Access  string `json:"access_token"`
}

// ----- AUTH ------
type UserEntity struct {
	BasicEntity
	Email       string             `json:"email,omitempty"`
	Role        string             `json:"role,omitempty"`
	ServiceUser *ServiceUserEntity `json:"service_user_profile,omitempty"`
	Supervisor  *SupervisorEntity  `json:"supervisor_profile,omitempty"`
	Admin       *AdminEntity       `json:"admin_profile,omitempty"`
	Employee    *EmployeeEntity    `json:"employee_profile,omitempty"`
}
type ServiceUserEntity struct {
	BasicEntity
	User               *UserEntity `json:"user,omitempty"`
	Address            string      `json:"address,omitempty"`
	Fullname           string      `json:"fullname,omitempty"`
	IdentityCardNumber string      `json:"identity_card_number,omitempty"`
	Phone              string      `json:"phone,omitempty"`
}
type SupervisorEntity struct {
	BasicEntity
	User               *UserEntity         `json:"user,omitempty"`
	Address            string              `json:"address"`
	Fullname           string              `json:"fullname"`
	IdentityCardNumber string              `json:"identity_card_number"`
	Phone              string              `json:"phone"`
	Status             EmployeeStatusEnum  `json:"employee_status"`
	PlacementStatus    PlacementStatusEnum `json:"placement_status"`
}
type AdminEntity struct {
	BasicEntity
	User     *UserEntity `json:"user,omitempty"`
	Address  string      `json:"address"`
	Fullname string      `json:"fullname"`
	Phone    string      `json:"phone"`
}
type EmployeeEntity struct {
	BasicEntity
	User               *UserEntity         `json:"user,omitempty"`
	Address            string              `json:"address"`
	Fullname           string              `json:"fullname"`
	IdentityCardNumber string              `json:"identity_card_number"`
	Phone              string              `json:"phone"`
	Status             EmployeeStatusEnum  `json:"employee_status"`
	PlacementStatus    PlacementStatusEnum `json:"placement_status"`
}
type UserWithProfileEntity struct {
	User    UserEntity `json:"created_user"`
	Profile any        `json:"created_profile"`
}

// ----- END OF AUTH ----
// ----- MASTER DATA -----
type CategoryEntity struct {
	BasicEntity
	CategoryName *string `json:"category_name"`
	Icon         string  `json:"icon"`
	Description  string  `json:"description"`
}
type DistrictEntity struct {
	BasicEntity
	DisctrictName string              `json:"district_name"`
	Description   string              `json:"description"`
	SubDistricts  []SubDistrictEntity `json:"sub_districts"`
}
type SubDistrictEntity struct {
	BasicEntity
	SubDisctrictName string         `json:"subdistrict_name"`
	Description      string         `json:"description"`
	District         DistrictEntity `json:"district"`
	SubDistricts     []VillageModel `json:"villages"`
}
type VillageEntity struct {
	BasicEntity
	VillageName   *string           `json:"village_name"`
	Description   string            `json:"description"`
	SubDistrictID *uint             `json:"subdistrict_id"`
	SubDistrict   *SubDistrictModel `json:"subdistrict"`
}

// ----- END OF MASTER DATA -----
// ----- SERVICE -----
type ServiceItemEntity struct {
	BasicEntity
	ItemName         string         `json:"item_name"`
	Description      string         `json:"description"`
	MinValue         uint           `json:"min_value"`
	MaxValue         uint           `json:"max_value"`
	ServiceID        uint           `json:"partial_service_id"`
	Service          *ServiceEntity `json:"service,omitempty"`
	PricePerItem     uint64         `json:"price_per_item"`
	IsOptionalChoice bool           `json:"is_optional_choice"`
	Unit             string         `json:"unit"`
}
type ServiceEntity struct {
	BasicEntity
	ServiceName  string              `json:"service_name"`
	Description  string              `json:"description"`
	Image        string              `json:"image"`
	Icon         string              `json:"icon"`
	BasePrice    uint64              `json:"base_price"`
	CategoryID   uint                `json:"category_id"`
	Category     *CategoryEntity     `json:"category,omitempty"`
	ServiceItems []ServiceItemEntity `json:"service_items"`
}
type ServicePackageEntity struct {
	BasicEntity
	PackageName string                        `json:"package_name"`
	Description string                        `json:"description"`
	Image       string                        `json:"image"`
	Icon        string                        `json:"icon"`
	BasePrice   uint64                        `json:"base_price"`
	Services    []ServicePackageServiceEntity `json:"services"`
}
type ServicePackageServiceEntity struct {
	BasicEntity
	ServicePackageID uint                             `json:"service_package_id"`
	ServicePackage   *ServicePackageEntity            `json:"service_package,omitempty"`
	ServiceID        uint                             `json:"service_id"`
	Service          ServiceEntity                    `json:"service"`
	Items            []ServicePackageServiceItemModel `json:"items"`
}

// ----- END OF SERVICE -----
// ----- ORDER ENTITY -----
type ServiceOrderEntity struct {
	BasicEntity
	// PaymentStatus       string                     `json:"payment_status"`
	ServiceUserID       uint                       `json:"service_user_id"`
	PurchasePrice       *uint64                    `json:"purchase_price"`
	TotalDiscount       uint64                     `json:"total_discount"`
	TotalPrice          uint64                     `json:"total_price"`
	TotalItem           uint                       `json:"total_item"`
	Date                time.Time                  `json:"date"`
	ContractDuration    uint                       `json:"contract_duration"`
	StartDate           time.Time                  `json:"start_date"`
	Address             string                     `json:"address"`
	Note                string                     `json:"buyer_note"`
	PaymentType         PaymentTypeEnum            `json:"payment_type"`
	Status              OrderStatusEnum            `json:"status"`
	MOU                 string                     `json:"mou"`
	ServicePackageID    uint                       `json:"service_package_id"`
	ServicePackage      *ServicePackageEntity      `json:"service_package,omitempty"`
	ServiceUser         *ServiceUserEntity         `json:"service_user"`
	ServiceOrderDetails []ServiceOrderDetailEntity `json:"order_details"`
}
type ServiceOrderDetailEntity struct {
	BasicEntity
	ServiceOrderID      uint                           `json:"service_order_id"`
	PartialServiceID    uint                           `json:"partial_service_id"`
	ServicePrice        uint64                         `json:"service_price"`
	AdditionalPrice     uint64                         `json:"additional_price"`
	TotalPrice          uint64                         `json:"total_price"`
	PartialServiceItems []ServiceOrderDetailItemEntity `json:"order_detail_items"`
	PartialService      *ServiceEntity                 `json:"partial_service"`
}
type ServiceOrderDetailItemEntity struct {
	BasicEntity
	ServiceOrderDetailID uint               `json:"service_order_detail_id"`
	PartialServiceItemID uint               `json:"partial_service_item_id"`
	Value                uint               `json:"value"`
	ItemPrice            uint64             `json:"item_price"`
	TotalPrice           uint64             `json:"total_price"`
	ServiceItem          *ServiceItemEntity `json:"partial_service_item"`
}

// ----- END OF ORDER ENTITY -----
// ----- PLACEMENT ENTITY -----
type ServiceOrderPlacementEntity struct{}
type ServiceOrderPlacementEmployeeEntity struct{}

// ----- END OF PLACEMENT ENTITY -----
