package domains

// ----- AUTH FORM -----
type BasicLoginForm struct {
	Email    string `json:"email" form:"email" valid:"required~this field is required"`
	Password string `json:"password" form:"password" valid:"required~this field is required"`
	Remember bool   `json:"remember" form:"remember"`
}

type BasicRegisterForm struct {
	Remember bool   `json:"remember" form:"remember"`
	Email    string `json:"email" form:"email" valid:"required~this field is required"`
	Password string `json:"password" form:"password" valid:"required~this field is required"`
	Role     string `json:"role" form:"role" valid:"required~this field is required,in(admin|service_user|supervisor|employee)~not a valid value"`
}

type ServiceUserRegisterForm struct {
	Email    string `json:"email" form:"email" valid:"required~this field is required"`
	Password string `json:"password" form:"password" valid:"required~this field is required"`
}

// ----- END OF AUTH FORM -----
