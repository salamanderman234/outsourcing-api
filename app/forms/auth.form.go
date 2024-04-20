package forms

import "time"

type LoginForm struct {
	Email    string `json:"email" form:"email" valid:"required"`
	Password string `json:"password" form:"password" valid:"required"`
}

type ChangePasswordForm struct {
	Email string `json:"email" form:"email"`
}
type ResetPasswordForm struct {
	Email       string `json:"email" form:"email"`
	NewPassword string `json:"new_password" form:"new_password"`
	ResetToken  string `json:"reset_token" form:"reset_token"`
}
type VerifyUserForm struct {
	UserID uint `json:"user_id" form:"user_id"`
}

type UserRegisterForm struct {
	Email              string                          `json:"email" form:"email" valid:"required"`
	Password           string                          `json:"password" form:"password" valid:"required"`
	AdminProfile       *AdminProfileRegisterForm       `json:"admin_profile" form:"admin_profile"`
	EmployeeProfile    *EmployeeProfileRegisterForm    `json:"employee_profile" form:"employee_profile"`
	SupervisorProfile  *SupervisorProfileRegisterForm  `json:"supervisor_profile" form:"supervisor_profile"`
	ServiceUserProfile *ServiceUserProfileRegisterForm `json:"service_user_profile" form:"service_user_profile"`
}

type AdminProfileRegisterForm struct {
	Fullname    string    `json:"fullname" form:"fullname" valid:"required,stringlength(0|255)"`
	RegencyID   uint      `json:"regency_id" form:"regency_id" valid:"required,int"`
	FullAddress string    `json:"full_address" form:"full_address" valid:"required,stringlength(0|255)"`
	BirthPlace  string    `json:"birth_place" form:"birth_place" valid:"required,stringlength(0|255)"`
	BirthDate   time.Time `json:"birth_date" form:"birth_date" valid:"required"`
	Phone       string    `json:"phone" form:"phone" valid:"required,stringlength(12|13)"`
}

// gender gunakan l dan p
type EmployeeProfileRegisterForm struct{}
type SupervisorProfileRegisterForm struct{}
type ServiceUserProfileRegisterForm struct{}
