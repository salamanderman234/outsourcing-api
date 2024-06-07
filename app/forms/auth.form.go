package forms

import "time"

type LoginForm struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type ChangePasswordForm struct {
	Email string `json:"email"`
}
type ResetPasswordForm struct {
	Email       string `json:"email" valid:"required"`
	NewPassword string `json:"new_password" valid:"required,stringlength(8|32)"`
	ResetToken  string `json:"reset_token" valid:"required"`
}
type VerifyUserForm struct {
	UserID uint `json:"user_id" form:"user_id" valid:"required"`
}

type UserRegisterForm struct {
	Email              string                          `json:"email" valid:"required,email"`
	Password           string                          `json:"password" valid:"required,stringlength(8|32)"`
	AdminProfile       *AdminProfileRegisterForm       `json:"admin_profile"`
	EmployeeProfile    *EmployeeProfileRegisterForm    `json:"employee_profile"`
	SupervisorProfile  *SupervisorProfileRegisterForm  `json:"supervisor_profile"`
	ServiceUserProfile *ServiceUserProfileRegisterForm `json:"service_user_profile"`
}

type AdminProfileRegisterForm struct {
	Fullname    string    `json:"fullname" valid:"required,stringlength(1|255)"`
	RegencyID   uint      `json:"regency_id" valid:"required,int"`
	FullAddress string    `json:"full_address" valid:"required,stringlength(1|255)"`
	BirthPlace  string    `json:"birth_place" valid:"required,stringlength(1|255)"`
	BirthDate   time.Time `json:"birth_date" valid:"required"`
	Phone       string    `json:"phone" valid:"required,stringlength(12|13)"`
}

// gender gunakan l dan p
type EmployeeProfileRegisterForm struct {
	Fullname       string    `json:"fullname" valid:"required,stringlength(1|255)"`
	RegencyID      uint      `json:"regency_id" valid:"required,int"`
	CategoryID     uint      `json:"category_id,omitempty" valid:"required,int"`
	FullAddress    string    `json:"full_address" valid:"required,stringlength(1|255)"`
	BirthPlace     string    `json:"birth_place" valid:"required,stringlength(1|255)"`
	BirthDate      time.Time `json:"birth_date" valid:"required"`
	Phone          string    `json:"phone" valid:"required,stringlength(12|13)"`
	NIK            string    `json:"nik,omitempty" valid:"required,stringlength(1|255)"`
	NPWP           string    `json:"npwp,omitempty" valid:"required,stringlength(1|255)"`
	Gender         string    `json:"gender,omitempty" valid:"required,in(l|p)"`
	MarriageStatus bool      `json:"marriage_status,omitempty"`
	LastEducation  string    `json:"last_education,omitempty" valid:"required,in(sd|smp|sma|s1)"`
}
type SupervisorProfileRegisterForm struct {
	Fullname       string    `json:"fullname" valid:"required,stringlength(1|255)"`
	RegencyID      uint      `json:"regency_id" valid:"required,int"`
	FullAddress    string    `json:"full_address" valid:"required,stringlength(1|255)"`
	BirthPlace     string    `json:"birth_place" valid:"required,stringlength(1|255)"`
	BirthDate      time.Time `json:"birth_date" valid:"required"`
	Phone          string    `json:"phone" valid:"required,stringlength(12|13)"`
	NIK            string    `json:"nik,omitempty" valid:"required,stringlength(1|255)"`
	NPWP           string    `json:"npwp,omitempty" valid:"required,stringlength(1|255)"`
	Gender         string    `json:"gender,omitempty" valid:"required,in(l|p)"`
	MarriageStatus bool      `json:"marriage_status,omitempty"`
}
type ServiceUserProfileRegisterForm struct {
	Fullname    string    `json:"fullname" valid:"required,stringlength(1|255)"`
	RegencyID   uint      `json:"regency_id" valid:"required,int"`
	FullAddress string    `json:"full_address" valid:"required,stringlength(1|255)"`
	BirthPlace  string    `json:"birth_place" valid:"required,stringlength(1|255)"`
	BirthDate   time.Time `json:"birth_date" valid:"required"`
	Phone       string    `json:"phone" valid:"required,stringlength(12|13)"`
	NIK         string    `json:"nik,omitempty" valid:"required,stringlength(1|255)"`
	Gender      string    `json:"gender,omitempty" valid:"required,in(l|p)"`
}

// update profile form
type UserUpdateForm struct {
	ProfilePic               *string                       `json:"profile_pic"`
	Password                 *string                       `json:"password" valid:"optional,stringlength(8|32)"`
	AdminUpdateProfile       *AdminUpdateProfileForm       `json:"admin_profile"`
	EmployeeUpdateProfile    *EmployeeUpdateProfileForm    `json:"employee_profile"`
	SupervisorUpdateProfile  *SupervisorUpdateProfileForm  `json:"supervisor_profile"`
	ServiceUserUpdateProfile *ServiceUserUpdateProfileForm `json:"service_user_profile"`
}

type AdminUpdateProfileForm struct {
	Fullname    *string    `json:"fullname" valid:"optional,stringlength(1|255)"`
	RegencyID   *uint      `json:"regency_id" valid:"optional,int"`
	FullAddress *string    `json:"full_address" valid:"optional,stringlength(1|255)"`
	BirthPlace  *string    `json:"birth_place" valid:"optional,stringlength(1|255)"`
	BirthDate   *time.Time `json:"birth_date" valid:"optional"`
	Phone       *string    `json:"phone" valid:"optional,stringlength(12|13)"`
}

type EmployeeUpdateProfileForm struct {
	Fullname       *string    `json:"fullname" valid:"optional,stringlength(1|255)"`
	RegencyID      *uint      `json:"regency_id" valid:"optional,int"`
	CategoryID     *uint      `json:"category_id,omitempty" valid:"optional,int"`
	FullAddress    *string    `json:"full_address" valid:"optional,stringlength(1|255)"`
	BirthPlace     *string    `json:"birth_place" valid:"optional,stringlength(1|255)"`
	BirthDate      *time.Time `json:"birth_date" valid:"optional"`
	Phone          *string    `json:"phone" valid:"optional,stringlength(12|13)"`
	NIK            *string    `json:"nik,omitempty" valid:"optional"`
	NPWP           *string    `json:"npwp,omitempty" valid:"optional"`
	Gender         *string    `json:"gender,omitempty" valid:"optional,in(l|p)"`
	MarriageStatus *bool      `json:"marriage_status,omitempty"`
	LastEducation  *string    `json:"last_education,omitempty" valid:"optional,in(sd,smp,sma,s1)"`
}
type SupervisorUpdateProfileForm struct {
	Fullname       *string    `json:"fullname" valid:"optional,stringlength(1|255)"`
	RegencyID      *uint      `json:"regency_id" valid:"optional,int"`
	FullAddress    *string    `json:"full_address" valid:"optional,stringlength(1|255)"`
	BirthPlace     *string    `json:"birth_place" valid:"optional,stringlength(1|255)"`
	BirthDate      *time.Time `json:"birth_date" valid:"optional"`
	Phone          *string    `json:"phone" valid:"optional,stringlength(12|13)"`
	NIK            *string    `json:"nik,omitempty" valid:"optional"`
	NPWP           *string    `json:"npwp,omitempty" valid:"optional"`
	Gender         *string    `json:"gender,omitempty" valid:"optional,in(l|p)"`
	MarriageStatus *bool      `json:"marriage_status,omitempty"`
}
type ServiceUserUpdateProfileForm struct {
	Fullname    *string    `json:"fullname" valid:"optional,stringlength(1|255)"`
	RegencyID   *uint      `json:"regency_id" valid:"optional,int"`
	FullAddress *string    `json:"full_address" valid:"optional,stringlength(1|255)"`
	BirthPlace  *string    `json:"birth_place" valid:"optional,stringlength(1|255)"`
	BirthDate   *time.Time `json:"birth_date" valid:"optional"`
	Phone       *string    `json:"phone" valid:"optional,stringlength(12|13)"`
	NIK         *string    `json:"nik,omitempty" valid:"optional"`
	Gender      *string    `json:"gender,omitempty" valid:"optional,in(l|p)"`
}
