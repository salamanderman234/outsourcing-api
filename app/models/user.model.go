package models

import "time"

type User struct {
	Model
	ProfilePic         *string      `json:"profile_pic,omitempty"`
	Email              *string      `json:"email,omitempty" visible:"true" gorm:"unique"`
	Password           *string      `json:"password,omitempty" visible:"true"`
	Role               *string      `json:"role,omitempty"`
	VerifiedAt         *time.Time   `json:"verified_at,omitempty"`
	AdminProfile       *Admin       `json:"admin_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	SupervisorProfile  *Supervisor  `json:"supervisor_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	EmployeeProfile    *Employee    `json:"employee_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	ServiceUserProfile *ServiceUser `json:"service_user_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
}

type Admin struct {
	Model
	UserID      *uint      `json:"user_id,omitempty" visible:"true"`
	User        *User      `json:"user" gorm:"foreignKey:UserID" visible:"true"`
	Fullname    *string    `json:"fullname,omitempty" visible:"true"`
	RegencyID   *uint      `json:"regency_id,omitempty" visible:"true"`
	Regency     *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress *string    `json:"full_address,omitempty" visible:"true"`
	AdminRole   *string    `json:"admin_role,omitempty" gorm:"default:operational" visible:"true"`
	BirthPlace  *string    `json:"birth_place,omitempty" visible:"true"`
	BirthDate   *time.Time `json:"birth_date,omitempty" visible:"true"`
	Gender      *string    `json:"gender,omitempty" visible:"true"`
	Phone       *string    `json:"phone,omitempty" visible:"true"`
}

type Supervisor struct {
	Model
	UserID         *uint      `json:"user_id,omitempty" visible:"true"`
	User           *User      `json:"user" gorm:"foreignKey:UserID" visible:"true"`
	Fullname       *string    `json:"fullname,omitempty" visible:"true"`
	RegencyID      *uint      `json:"regency_id,omitempty" visible:"true"`
	Regency        *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress    *string    `json:"full_address,omitempty" visible:"true"`
	BirthPlace     *string    `json:"birth_place,omitempty" visible:"true"`
	BirthDate      *time.Time `json:"birth_date,omitempty" visible:"true"`
	NIK            *string    `json:"nik,omitempty" visible:"true"`
	NPWP           *string    `json:"npwp,omitempty" visible:"true"`
	Gender         *string    `json:"gender,omitempty" visible:"true"`
	MarriageStatus *bool      `json:"marriage_status,omitempty" visible:"true"`
	Phone          *string    `json:"phone,omitempty" visible:"true"`
}

type Employee struct {
	Model
	UserID             *uint      `json:"user_id,omitempty" visible:"true"`
	User               *User      `json:"user" gorm:"foreignKey:UserID" visible:"true"`
	Fullname           *string    `json:"fullname,omitempty" visible:"true"`
	CategoryID         *uint      `json:"category_id,omitempty"`
	Category           *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID" visible:"true"`
	RegencyID          *uint      `json:"regency_id,omitempty" visible:"true"`
	Regency            *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress        *string    `json:"full_address,omitempty" visible:"true"`
	BirthPlace         *string    `json:"birth_place,omitempty" visible:"true"`
	BirthDate          *time.Time `json:"birth_date,omitempty" visible:"true"`
	NIK                *string    `json:"nik,omitempty" visible:"true"`
	NPWP               *string    `json:"npwp,omitempty" visible:"true"`
	Gender             *string    `json:"gender,omitempty" visible:"true"`
	MarriageStatus     *bool      `json:"marriage_status,omitempty" visible:"true"`
	Phone              *string    `json:"phone,omitempty" visible:"true"`
	LastEducation      *string    `json:"last_education,omitempty" visible:"true"`
	Ijazah             *string    `json:"ijazah,omitempty" visible:"true"`
	FieldCertification *string    `json:"field_certification,omitempty" visible:"true"`
}

type ServiceUser struct {
	Model
	UserID      *uint      `json:"user_id,omitempty"`
	User        *User      `json:"user" gorm:"foreignKey:UserID" visible:"true"`
	Fullname    *string    `json:"fullname,omitempty" visible:"true"`
	RegencyID   *uint      `json:"regency_id,omitempty" visible:"true"`
	Regency     *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID" visible:"true"`
	FullAddress *string    `json:"full_address,omitempty" visible:"true"`
	BirthPlace  *string    `json:"birth_place,omitempty" visible:"true"`
	BirthDate   *time.Time `json:"birth_date,omitempty" visible:"true"`
	NIK         *string    `json:"nik,omitempty" visible:"true"`
	Gender      *string    `json:"gender,omitempty" visible:"true"`
	Phone       *string    `json:"phone,omitempty" visible:"true"`
}
