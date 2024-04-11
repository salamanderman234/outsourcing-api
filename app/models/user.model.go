package models

import "time"

type User struct {
	Model
	ProfilePic          *string     `json:"profile_pic,omitempty"`
	Email               *string     `json:"email,omitempty" visible:"true" gorm:"unique"`
	Password            *string     `json:"password,omitempty" visible:"true"`
	Role                *string     `json:"role,omitempty"`
	VerifiedAt          *time.Time  `json:"verified_at,omitempty"`
	ChangePasswordToken *string     `json:"change_password_token,omitempty"`
	AdminProfile        *Admin      `json:"admin_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	SupervisorProfile   *Supervisor `json:"supervisor_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	EmployeeProfile     *Employee   `json:"employee_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
	ServiceUserProfile  *Admin      `json:"service_user_profile,omitempty" gorm:"foreignKey:UserID" visible:"true"`
}

type Admin struct {
	Model
	UserID      *uint      `json:"user_id,omitempty"`
	User        *User      `json:"user" gorm:"foreignKey:UserID"`
	Fullname    *string    `json:"fullname,omitempty"`
	RegencyID   *uint      `json:"regency_id,omitempty"`
	Regency     *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID"`
	FullAddress *string    `json:"full_address,omitempty"`
	AdminRole   *string    `json:"admin_role,omitempty"`
	BirthPlace  *string    `json:"birth_place,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
	Phone       *string    `json:"phone,omitempty"`
}

type Supervisor struct {
	Model
	UserID         *uint      `json:"user_id,omitempty"`
	User           *User      `json:"user" gorm:"foreignKey:UserID"`
	Fullname       *string    `json:"fullname,omitempty"`
	RegencyID      *uint      `json:"regency_id,omitempty"`
	Regency        *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID"`
	FullAddress    *string    `json:"full_address,omitempty"`
	AdminRole      *string    `json:"admin_role,omitempty"`
	BirthPlace     *string    `json:"birth_place,omitempty"`
	BirthDate      *time.Time `json:"birth_date,omitempty"`
	NIK            *string    `json:"nik,omitempty"`
	NPWP           *string    `json:"npwp,omitempty"`
	Gender         *string    `json:"gender,omitempty"`
	MarriageStatus *bool      `json:"marriage_status,omitempty"`
	Phone          *string    `json:"phone,omitempty"`
}

type Employee struct {
	Model
	UserID             *uint      `json:"user_id,omitempty"`
	User               *User      `json:"user" gorm:"foreignKey:UserID"`
	Fullname           *string    `json:"fullname,omitempty"`
	CategoryID         *uint      `json:"category_id,omitempty"`
	Category           *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	RegencyID          *uint      `json:"regency_id,omitempty"`
	Regency            *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID"`
	FullAddress        *string    `json:"full_address,omitempty"`
	BirthPlace         *string    `json:"birth_place,omitempty"`
	BirthDate          *time.Time `json:"birth_date,omitempty"`
	NIK                *string    `json:"nik,omitempty"`
	NPWP               *string    `json:"npwp,omitempty"`
	Gender             *string    `json:"gender,omitempty"`
	MarriageStatus     *bool      `json:"marriage_status,omitempty"`
	Phone              *string    `json:"phone,omitempty"`
	LastEducation      *string    `json:"last_education,omitempty"`
	Ijazah             *string    `json:"ijazah,omitempty"`
	FieldCertification *string    `json:"field_certification,omitempty"`
}

type ServiceUser struct {
	Model
	UserID      *uint      `json:"user_id,omitempty"`
	User        *User      `json:"user" gorm:"foreignKey:UserID"`
	Fullname    *string    `json:"fullname,omitempty"`
	RegencyID   *uint      `json:"regency_id,omitempty"`
	Regency     *Regency   `json:"regency,omitempty" gorm:"foreignKey:RegencyID"`
	FullAddress *string    `json:"full_address,omitempty"`
	BirthPlace  *string    `json:"birth_place,omitempty"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
	NIK         *string    `json:"nik,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	Phone       *string    `json:"phone,omitempty"`
}
