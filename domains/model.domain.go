package domains

import (
	"time"

	"gorm.io/gorm"
)

// Registry
func ModelRegistry() []any {
	return nil
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

func (AdminModel) TableName() string {
	return "admins"
}

// ----- END OF AUTH MODEL -----
