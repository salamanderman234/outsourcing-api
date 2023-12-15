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
	UserID   *uint      `json:"user_id" gorm:"not null"`
	User     *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address  *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname *string    `json:"fullname" gorm:"not null;type:varchar(255)"`
	CardID   *string    `json:"id_card_number" gorm:"not null;type:varchar(255)"`
	Phone    *string    `json:"phone" gorm:"not null;type:varchar(13)"`
}

func (EmployeeModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}
func (EmployeeModel) TableName() string {
	return "employees"
}

type ServiceUserModel struct {
	gorm.Model
	UserID   *uint      `json:"user_id" gorm:"not null"`
	User     *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address  *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname *string    `json:"Fullname" gorm:"not null;type:varchar(255)"`
	CardID   *string    `json:"id_card_number" gorm:"not null;type:varchar(255)"`
	Phone    *string    `json:"phone" gorm:"not null;type:varchar(13)"`
}

func (ServiceUserModel) GetPolicy() Policy {
	return &CategoryPolicy{}
}
func (ServiceUserModel) TableName() string {
	return "service_users"
}

type SupervisorModel struct {
	gorm.Model
	UserID   *uint      `json:"user_id" gorm:"not null"`
	User     *UserModel `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address  *string    `json:"address" gorm:"not null;type:varchar(255)"`
	Fullname *string    `json:"Fullname" gorm:"not null;type:varchar(255)"`
	CardID   *string    `json:"id_card_number" gorm:"not null;type:varchar(255)"`
	Phone    *string    `json:"phone" gorm:"not null;type:varchar(13)"`
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
	Description  string  `json:"description" gorm:"type:varchar(255)"`
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
