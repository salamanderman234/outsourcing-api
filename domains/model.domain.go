package domains

import "gorm.io/gorm"

// Registry
func ModelRegistry() []any {
	return nil
}

// ----- AUTH MODEL -----
type EmployeeModel struct {
}

func (EmployeeModel) TableName() string {
	return "employees"
}

type ServiceUserModel struct {
	gorm.Model
	Email    *string `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password *string `json:"password" gorm:"not null;type:varchar(255)"`
	Profile  string  `json:"profile" gorm:"type:varchar(255)"`
}

func (ServiceUserModel) TableName() string {
	return "service_users"
}

type SupervisorModel struct {
}

func (SupervisorModel) TableName() string {
	return "supervisors"
}

type AdminModel struct {
}

func (AdminModel) TableName() string {
	return "admins"
}

// ----- END OF AUTH MODEL -----
