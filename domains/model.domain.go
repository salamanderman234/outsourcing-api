package domains

import (
	"time"

	"gorm.io/gorm"
)

// --> BASIC
type Model interface {
	GetPolicy() Policy
}

// ----- AUTH MODEL -----
type UserWithProfileModel struct {
	User    UserModel
	Profile any
}

type UserModel struct {
	gorm.Model
	Profile     string            `json:"profile" gorm:"type:varchar(255)"`
	Email       *string           `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password    *string           `json:"password" gorm:"not null;type:varchar(255)"`
	Role        string            `json:"role" gorm:"not null;default('service_user')"`
	JoinedDate  time.Time         `json:"joined_date"`
	ServiceUser *ServiceUserModel `json:"service_user_profile,omitempty" gorm:"foreignKey:UserID"`
	Supervisor  *SupervisorModel  `json:"supervisor_profile,omitempty" gorm:"foreignKey:UserID"`
	Admin       *AdminModel       `json:"admin_profile,omitempty" gorm:"foreignKey:UserID"`
	Employee    *EmployeeModel    `json:"employee_profile,omitempty" gorm:"foreignKey:UserID"`
}

func (UserModel) GetPolicy() Policy {
	return &UserAuthPolicy{}
}

func (UserModel) TableName() string {
	return "users"
}

type EmployeeModel struct {
	gorm.Model
	UserID             *uint              `json:"user_id" gorm:"not null"`
	User               *UserModel         `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address            *string            `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname           *string            `json:"fullname" gorm:"not null;type:varchar(255)"`
	IdentityCardNumber *string            `json:"identity_card_number" gorm:"not null;type:varchar(255);unique"`
	Phone              *string            `json:"phone" gorm:"not null;type:varchar(13)"`
	CategoryID         uint               `json:"category_id"`
	Category           *CategoryModel     `json:"category" gorm:"foreignKey:CategoryID"`
	Status             EmployeeStatusEnum `json:"employee_status"`
}

func (EmployeeModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}
func (EmployeeModel) TableName() string {
	return "employees"
}

type ServiceUserModel struct {
	gorm.Model
	UserID             *uint      `json:"user_id" gorm:"not null"`
	User               *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address            *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname           *string    `json:"Fullname" gorm:"not null;type:varchar(255)"`
	IdentityCardNumber *string    `json:"identity_card_number" gorm:"not null;type:varchar(255)"`
	Phone              *string    `json:"phone" gorm:"not null;type:varchar(13)"`
}

func (ServiceUserModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}
func (ServiceUserModel) TableName() string {
	return "service_users"
}

type SupervisorModel struct {
	gorm.Model
	UserID             *uint              `json:"user_id" gorm:"not null"`
	User               *UserModel         `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address            *string            `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname           *string            `json:"Fullname" gorm:"not null;type:varchar(255)"`
	IdentityCardNumber *string            `json:"identity_card_number" gorm:"not null;type:varchar(255);unique"`
	Phone              *string            `json:"phone" gorm:"not null;type:varchar(13)"`
	Status             EmployeeStatusEnum `json:"employee_status"`
}

func (SupervisorModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}
func (SupervisorModel) TableName() string {
	return "supervisors"
}

type AdminModel struct {
	gorm.Model
	UserID   *uint      `json:"user_id" gorm:"not null"`
	User     *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address  *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname *string    `json:"Fullname" gorm:"not null;type:varchar(255)"`
	Phone    *string    `json:"phone" gorm:"not null;type:varchar(13)"`
}

func (AdminModel) GetPolicy() Policy {
	return &UserAuthPolicy{}
}

func (AdminModel) TableName() string {
	return "admins"
}

// ----- END OF AUTH MODEL -----

// ----- MASTER DATA -----
type CategoryModel struct {
	gorm.Model
	CategoryName    *string         `json:"category_name" gorm:"not null;type:varchar(255)"`
	Icon            string          `json:"icon" gorm:"type:varchar(255)"`
	Description     string          `json:"description" gorm:"type:text"`
	Employees       []EmployeeModel `json:"employees,omitempty" gorm:"foreignKey:CategoryID"`
	PartialServices []ServiceModel  `json:"partial_services,omitempty" gorm:"foreignKey:CategoryID"`
}

func (CategoryModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}

func (CategoryModel) TableName() string {
	return "categories"
}

type DistrictModel struct {
	gorm.Model
	DisctrictName *string            `json:"district_name" gorm:"not null;type:varchar(255)"`
	Description   string             `json:"description" gorm:"type:varchar(255)"`
	SubDistricts  []SubDistrictModel `json:"sub_districts" gorm:"foreignKey:DistrictID"`
}

func (DistrictModel) GetPolicy() Policy {
	return &DistrictPolicy{}
}

func (DistrictModel) TableName() string {
	return "districts"
}

type SubDistrictModel struct {
	SubDisctrictName *string        `json:"subdistrict_name" gorm:"not null;type:varchar(255)"`
	Description      string         `json:"description" gorm:"type:varchar(255)"`
	DistrictID       *uint          `json:"district_id" gorm:"not null"`
	District         *DistrictModel `json:"district" gorm:"foreignKey:DistrictID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubDistricts     []VillageModel `json:"villages" gorm:"foreignKey:SubDistrictID"`
}

func (SubDistrictModel) GetPolicy() Policy {
	return &SubDistrictPolicy{}
}

func (SubDistrictModel) TableName() string {
	return "sub_districts"
}

type VillageModel struct {
	VillageName   *string           `json:"village_name" gorm:"not null;type:varchar(255)"`
	Description   string            `json:"description" gorm:"type:varchar(255)"`
	SubDistrictID *uint             `json:"subdistrict_id" gorm:"not null"`
	SubDistrict   *SubDistrictModel `json:"subdistrict" gorm:"foreignKey:SubDistrictID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (VillageModel) GetPolicy() Policy {
	return &VillagePolicy{}
}

func (VillageModel) TableName() string {
	return "villages"
}

// ----- END OF MASTER DATA -----
// ----- SERVICE -----
type ServiceItemModel struct {
	gorm.Model
	ItemName         *string       `json:"item_name" gorm:"type:varchar(255);not null;"`
	Description      string        `json:"description" gorm:"type:varchar(255)"`
	MinValue         *uint         `json:"min_value" gorm:"type:not null;default:0;"`
	MaxValue         *uint         `json:"max_value" gorm:"type:not null;default:1;"`
	PricePerItem     *uint64       `json:"price_per_item" gorm:"not null;default:0"`
	IsOptionalChoice *bool         `json:"is_optional_choice" gorm:"not null;default:0"`
	Unit             *string       `json:"unit" gorm:"default:unit;type:varchar(255)"`
	PartialServiceID *uint         `json:"partial_service_id" gorm:"not null"`
	Service          *ServiceModel `json:"service" gorm:"foreignKey:PartialServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ServiceItemModel) GetPolicy() Policy {
	return &ServiceItemPolicy{}
}
func (ServiceItemModel) TableName() string {
	return "partial_service_items"
}

type ServiceModel struct {
	gorm.Model
	ServiceName          *string                `json:"service_name" gorm:"not null;type:varchar(255)"`
	Description          string                 `json:"description" gorm:"type:varchar(255)"`
	Image                string                 `json:"image" gorm:"type:varchar(255)"`
	Icon                 string                 `json:"icon" gorm:"type:varchar(255)"`
	BasePrice            *uint64                `json:"base_price" gorm:"not null;default:0"`
	BaseNumberOfEmployee *uint64                `json:"base_number_of_employee" gorm:"not null;default:0"`
	CategoryID           *uint                  `json:"category_id" gorm:"not null"`
	Category             *CategoryModel         `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceItems         []ServiceItemModel     `json:"service_items" gorm:"foreignKey:PartialServiceID"`
	Packages             []*ServicePackageModel `json:"packages" gorm:"many2many:service_package_services"`
}

func (ServiceModel) GetPolicy() Policy {
	return &ServicePolicy{}
}

func (ServiceModel) TableName() string {
	return "partial_services"
}

type ServicePackageModel struct {
	gorm.Model
	PackageName *string                      `json:"package_name" gorm:"not null;type:varchar(255)"`
	Description string                       `json:"description" gorm:"type:text"`
	Image       string                       `json:"image" gorm:"type:varchar(255)"`
	Icon        string                       `json:"icon" gorm:"type:varchar(255)"`
	TotalPrice  *uint64                      `json:"total_price" gorm:"not null;default:0"`
	FinalPrice  *uint64                      `json:"final_price" gorm:"not null;default:0"`
	Services    []ServicePackageServiceModel `json:"services" gorm:"foreignKey:ServicePackageID"`
}

func (ServicePackageModel) TableName() string {
	return "service_packages"
}

type ServicePackageServiceModel struct {
	gorm.Model
	ServicePackageID *uint                            `json:"service_package_id" gorm:"not null"`
	ServicePackage   *ServicePackageModel             `json:"service_package" gorm:"foreignKey:ServicePackageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceID        *uint                            `json:"service_id" gorm:"not null"`
	Service          *ServiceModel                    `json:"service" gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Items            []ServicePackageServiceItemModel `json:"items" gorm:"foreignKey:ServicePackageServiceID"`
}

func (ServicePackageServiceModel) TableName() string {
	return "service_package_services"
}

type ServicePackageServiceItemModel struct {
	gorm.Model
	ServicePackageServiceID *uint                       `json:"service_package_service_id" gorm:"not null"`
	ServicePackageService   *ServicePackageServiceModel `json:"service_package_service"`
	ServiceItemID           *uint                       `json:"service_item_id" gorm:"not null"`
	ServiceItem             *ServiceItemModel           `json:"service_item" gorm:"foreignKey:ServiceItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Value                   uint                        `json:"value" gorm:"default:0"`
}

func (ServicePackageServiceItemModel) TableName() string {
	return "service_package_service_items"
}

// ----- END OF SERVICE -----
// ----- ORDER MODEL -----
type ServiceOrderModel struct {
	gorm.Model
	// PaymentStatus       *string                   `json:"payment_status" gorm:"not null"`
	ServiceUserID       *uint                     `json:"service_user_id" gorm:"not null"`
	PurchasePrice       *uint64                   `json:"purchase_price" gorm:"not null"`
	TotalDiscount       uint64                    `json:"total_discount" gorm:"not null;default:0"`
	TotalPrice          *uint64                   `json:"total_price" gorm:"not null"`
	TotalItem           *uint                     `json:"total_item" gorm:"not null"`
	Date                *time.Time                `json:"date" gorm:"not null"`
	ContractDuration    *uint                     `json:"contract_duration" gorm:"not null"`
	StartDate           *time.Time                `json:"start_date" gorm:"not null"`
	Address             *string                   `json:"address" gorm:"not null;type:text"`
	Note                string                    `json:"buyer_note" gorm:"type:text"`
	PaymentType         *PaymentTypeEnum          `json:"payment_type" gorm:"not null"`
	Status              *OrderStatusEnum          `json:"status" gorm:"not null"`
	MOU                 string                    `json:"mou" gorm:"type:varchar(255)"`
	ServicePackageID    *uint                     `json:"service_package_id" gorm:"default:null"`
	ServicePackage      *ServicePackageModel      `json:"service_package" gorm:"foreignKey:ServicePackageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceUser         *ServiceUserModel         `json:"service_user" gorm:"foreignKey:ServiceUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceOrderDetails []ServiceOrderDetailModel `json:"order_details" gorm:"foreignKey:ServiceOrderID"`
}

func (ServiceOrderModel) TableName() string {
	return "service_orders"
}
func (ServiceOrderModel) GetPolicy() Policy {
	return &ServiceOrderPolicy{}
}

type ServiceOrderDetailModel struct {
	gorm.Model
	ServiceOrderID      *uint                         `json:"service_order_id" gorm:"not null"`
	PartialServiceID    *uint                         `json:"partial_service_id" gorm:"not null"`
	ServicePrice        *uint64                       `json:"service_price" gorm:"not null"`
	AdditionalPrice     *uint64                       `json:"additional_price" gorm:"not null"`
	TotalPrice          *uint64                       `json:"total_price" gorm:"not null"`
	PartialServiceItems []ServiceOrderDetailItemModel `json:"order_detail_items" gorm:"foreignKey:ServiceOrderDetailID"`
	PartialService      *ServiceModel                 `json:"partial_service" gorm:"foreignKey:PartialServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ServiceOrderDetailModel) TableName() string {
	return "service_order_details"
}
func (ServiceOrderDetailModel) GetPolicy() Policy {
	return &ServiceOrderPolicy{}
}

type ServiceOrderDetailItemModel struct {
	gorm.Model
	ServiceOrderDetailID *uint             `json:"service_order_detail_id" gorm:"not null"`
	PartialServiceItemID *uint             `json:"partial_service_item_id" gorm:"not null"`
	Value                *uint             `json:"value" gorm:"not null"`
	ItemPrice            *uint64           `json:"item_price" gorm:"not null"`
	TotalPrice           *uint64           `json:"total_price" gorm:"not null"`
	ServiceItem          *ServiceItemModel `json:"partial_service_item" gorm:"foreignKey:PartialServiceItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (ServiceOrderDetailItemModel) TableName() string {
	return "service_order_detail_items"
}
func (ServiceOrderDetailItemModel) GetPolicy() Policy {
	return &ServiceOrderPolicy{}
}

// ----- END OF ORDER MODEL -----
// ----- PLACEMENT MODEL -----
type ServiceOrderPlacementDailyReportModel struct {
	gorm.DB
	ServiceOrderPlacementID uint      `json:"service_order_placement_id"`
	File                    string    `json:"file"`
	Date                    time.Time `json:"date"`
}

func (ServiceOrderPlacementDailyReportModel) TableName() string {
	return "service_order_placement_daily_reports"
}
func (ServiceOrderPlacementDailyReportModel) GetPolicy() Policy {
	return &PlacementPolicy{}
}

type ServiceOrderPlacementModel struct {
	gorm.Model
	SupervisorID      *uint                                   `json:"supervisor_id"`
	Supervisor        *SupervisorModel                        `json:"supervisor" gorm:"foreignKey:SupervisorID;"`
	ServiceOrderID    *uint                                   `json:"service_order_id"`
	ServiceOrder      *ServiceOrderModel                      `json:"service_order" gorm:"foreignKey:ServiceOrderID"`
	StartDate         time.Time                               `json:"start_date"`
	Duration          uint                                    `json:"duration" gorm:"default:0"`
	EndDate           time.Time                               `json:"end_date"`
	Address           string                                  `json:"address"`
	CompanyName       string                                  `json:"company_name"`
	Status            PlacementStatusEnum                     `json:"status"`
	ServicePlacements []ServiceOrderPlacementServiceModel     `json:"service_placements" gorm:"foreignKey:ServiceOrderPlacementID"`
	Reports           []ServiceOrderPlacementDailyReportModel `json:"reports" gorm:"foreignKey:ServiceOrderPlacementID"`
}

func (ServiceOrderPlacementModel) TableName() string {
	return "service_order_placements"
}
func (ServiceOrderPlacementModel) GetPolicy() Policy {
	return &PlacementPolicy{}
}

type ServiceOrderPlacementServiceModel struct {
	gorm.Model
	ServiceOrderPlacementID *uint                                       `json:"service_order_placement_id"`
	ServiceOrderPlacement   *ServiceOrderPlacementModel                 `json:"service_order_placement" gorm:"foreignKey:ServiceOrderPlacementID"`
	PartialServiceID        *uint                                       `json:"partial_service_id"`
	PartialService          *ServiceModel                               `json:"partial_service"`
	Employees               []ServiceOrderPlacementServiceEmployeeModel `json:"employees" gorm:"foreignKey:ServiceOrderPlacementServiceID"`
}

func (ServiceOrderPlacementServiceModel) TableName() string {
	return "service_order_placement_service"
}
func (ServiceOrderPlacementServiceModel) GetPolicy() Policy {
	return &PlacementPolicy{}
}

type ServiceOrderPlacementServiceEmployeeScheduleModel struct {
	gorm.Model
	ServiceOrderPlacementServiceEmployeeID *uint  `json:"service_order_placement_service_employee_id"`
	WeekDay                                string `json:"weekday"`
	StartShift                             uint   `json:"start_shift"`
	EndShift                               uint   `json:"end_shift"`
}

func (ServiceOrderPlacementServiceEmployeeScheduleModel) TableName() string {
	return "service_order_placement_service_employee_schedules"
}
func (ServiceOrderPlacementServiceEmployeeScheduleModel) GetPolicy() Policy {
	return &PlacementPolicy{}
}

type ServiceOrderPlacementServiceEmployeeModel struct {
	gorm.Model
	EmployeeID                     *uint                                               `json:"employee_id"`
	Employee                       *EmployeeModel                                      `json:"employee" gorm:"foreignKey:EmployeeID;"`
	ServiceOrderPlacementServiceID *uint                                               `json:"service_order_placement_service_id"`
	ServiceOrderPlacementService   *ServiceOrderPlacementServiceModel                  `json:"service_order_placement_service" gorm:"foreignKey:ServiceOrderPlacementServiceID"`
	PlacementDate                  time.Time                                           `json:"placement_date"`
	Status                         EmployeePlacementStatusEnum                         `json:"status"`
	Schedules                      []ServiceOrderPlacementServiceEmployeeScheduleModel `json:"schedules" gorm:"foreignKey:ServiceOrderPlacementServiceID"`
}

func (ServiceOrderPlacementServiceEmployeeModel) TableName() string {
	return "service_order_placement_service_employees"
}
func (ServiceOrderPlacementServiceEmployeeModel) GetPolicy() Policy {
	return &PlacementPolicy{}
}

// ----- END OF PLACEMENT MODEL -----
