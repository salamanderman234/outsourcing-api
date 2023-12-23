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
	return nil
}

func (UserModel) TableName() string {
	return "users"
}

type EmployeeModel struct {
	gorm.Model
	UserID             *uint      `json:"user_id" gorm:"not null"`
	User               *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address            *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname           *string    `json:"fullname" gorm:"not null;type:varchar(255)"`
	IdentityCardNumber *string    `json:"identity_card_number" gorm:"not null;type:varchar(255);unique"`
	Phone              *string    `json:"phone" gorm:"not null;type:varchar(13)"`
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
	UserID             *uint      `json:"user_id" gorm:"not null"`
	User               *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address            *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname           *string    `json:"Fullname" gorm:"not null;type:varchar(255)"`
	IdentityCardNumber *string    `json:"identity_card_number" gorm:"not null;type:varchar(255);unique"`
	Phone              *string    `json:"phone" gorm:"not null;type:varchar(13)"`
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
	return &CategoryPolicy{}
}

func (AdminModel) TableName() string {
	return "admins"
}

// ----- END OF AUTH MODEL -----

// ----- MASTER DATA -----
type CategoryModel struct {
	gorm.Model
	CategoryName *string `json:"category_name" gorm:"not null;type:varchar(255)"`
	Icon         string  `json:"icon" gorm:"type:varchar(255)"`
	Description  string  `json:"description" gorm:"type:text"`
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
	ServiceName  *string                `json:"service_name" gorm:"not null;type:varchar(255)"`
	Description  string                 `json:"description" gorm:"type:varchar(255)"`
	Image        string                 `json:"image" gorm:"type:varchar(255)"`
	Icon         string                 `json:"icon" gorm:"type:varchar(255)"`
	BasePrice    *uint64                `json:"base_price" gorm:"not null;default:0"`
	CategoryID   *uint                  `json:"category_id" gorm:"not null"`
	Category     *CategoryModel         `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ServiceItems []ServiceItemModel     `json:"service_items" gorm:"foreignKey:PartialServiceID"`
	Packages     []*ServicePackageModel `json:"packages" gorm:"many2many:service_package_services"`
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
	Description string                       `json:"description" gorm:"type:varchar(255)"`
	Image       string                       `json:"image" gorm:"type:varchar(255)"`
	Icon        string                       `json:"icon" gorm:"type:varchar(255)"`
	BasePrice   *uint64                      `json:"base_price" gorm:"not null;default:0"`
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
