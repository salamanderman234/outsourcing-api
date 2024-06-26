package models

import "time"

type User struct {
	Model
	ProfilePic             *string          `json:"profile_pic"`
	Email                  *string          `json:"email" visible:"true" gorm:"unique"`
	Password               *string          `json:"password" visible:"true"`
	Role                   *string          `json:"role"`
	VerifiedAt             *time.Time       `json:"verified_at"`
	AdminProfile           *Admin           `json:"admin_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	SuperAdminProfile      *SuperAdmin      `json:"super_admin_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	SupervisorProfile      *Supervisor      `json:"supervisor_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	EmployeeProfile        *Employee        `json:"employee_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	ServiceUserProfile     *ServiceUser     `json:"service_user_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	ApplicationUserProfile *ApplicationUser `json:"application_user_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
}

type Admin struct {
	Model
	UserID      *uint      `json:"user_id" visible:"true"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	Fullname    *string    `json:"fullname" visible:"true"`
	RegencyID   *uint      `json:"regency_id" visible:"true"`
	Regency     *Regency   `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress *string    `json:"full_address" visible:"true"`
	AdminRole   *string    `json:"admin_role" gorm:"default:operational" visible:"true"`
	BirthPlace  *string    `json:"birth_place" visible:"true"`
	BirthDate   *time.Time `json:"birth_date" visible:"true"`
	Gender      *string    `json:"gender" visible:"true"`
	Phone       *string    `json:"phone" visible:"true"`
}
type SuperAdmin struct {
	Model
	UserID      *uint      `json:"user_id" visible:"true"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	Fullname    *string    `json:"fullname" visible:"true"`
	RegencyID   *uint      `json:"regency_id" visible:"true"`
	Regency     *Regency   `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress *string    `json:"full_address" visible:"true"`
	AdminRole   *string    `json:"admin_role" gorm:"default:operational" visible:"true"`
	BirthPlace  *string    `json:"birth_place" visible:"true"`
	BirthDate   *time.Time `json:"birth_date" visible:"true"`
	Gender      *string    `json:"gender" visible:"true"`
	Phone       *string    `json:"phone" visible:"true"`
}

type Supervisor struct {
	Model
	UserID         *uint      `json:"user_id" visible:"true"`
	User           *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	Fullname       *string    `json:"fullname" visible:"true"`
	RegencyID      *uint      `json:"regency_id" visible:"true"`
	Regency        *Regency   `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress    *string    `json:"full_address" visible:"true"`
	BirthPlace     *string    `json:"birth_place" visible:"true"`
	BirthDate      *time.Time `json:"birth_date" visible:"true"`
	NIK            *string    `json:"nik" visible:"true"`
	NPWP           *string    `json:"npwp" visible:"true"`
	Gender         *string    `json:"gender" visible:"true"`
	MarriageStatus *bool      `json:"marriage_status" visible:"true"`
	Phone          *string    `json:"phone" visible:"true"`
	Status         *string    `json:"status"`
}

type Employee struct {
	Model
	UserID             *uint      `json:"user_id" visible:"true"`
	User               *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	Fullname           *string    `json:"fullname" visible:"true"`
	CategoryID         *uint      `json:"category_id"`
	Category           *Category  `json:"category" gorm:"foreignKey:CategoryID" visible:"true"`
	RegencyID          *uint      `json:"regency_id" visible:"true"`
	Regency            *Regency   `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress        *string    `json:"full_address" visible:"true"`
	BirthPlace         *string    `json:"birth_place" visible:"true"`
	BirthDate          *time.Time `json:"birth_date" visible:"true"`
	NIK                *string    `json:"nik" visible:"true"`
	NPWP               *string    `json:"npwp" visible:"true"`
	Gender             *string    `json:"gender" visible:"true"`
	MarriageStatus     *bool      `json:"marriage_status" visible:"true"`
	Phone              *string    `json:"phone" visible:"true"`
	LastEducation      *string    `json:"last_education" visible:"true"`
	Ijazah             *string    `json:"ijazah" visible:"true"`
	FieldCertification *string    `json:"field_certification" visible:"true"`
	Status             *string    `json:"status"`
}

type ServiceUser struct {
	Model
	UserID      *uint      `json:"user_id"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	Fullname    *string    `json:"fullname" visible:"true"`
	RegencyID   *uint      `json:"regency_id" visible:"true"`
	Regency     *Regency   `json:"regency" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress *string    `json:"full_address" visible:"true"`
	BirthPlace  *string    `json:"birth_place" visible:"true"`
	BirthDate   *time.Time `json:"birth_date" visible:"true"`
	NIK         *string    `json:"nik" visible:"true"`
	Gender      *string    `json:"gender" visible:"true"`
	Phone       *string    `json:"phone" visible:"true"`
}

type ApplicationUser struct {
	Model
	UserID          *uint      `json:"user_id"`
	User            *User      `json:"user,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	ApplicationName *string    `json:"application_name"`
	ValidUntil      *time.Time `json:"valid_until"`
	AccessID        *uint      `json:"access_id"`
	Access          *Access    `json:"access" gorm:"foreignKey:AccessID"`
}
