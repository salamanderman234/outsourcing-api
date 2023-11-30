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
}

type ServiceUserRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"Fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"id_card_number" form:"id_card_number" valid:"required~this field is required,stringLength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type AdminRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"Fullname" form:"fullname" valid:"required~this field is required"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type EmployeeRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"Fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"id_card_number" form:"id_card_number" valid:"required~this field is required,stringLength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}
type SupervisorRegisterForm struct {
	BasicRegisterForm
	Address  *string `json:"address" valid:"required~this field is required,stringlength(0|255)~maximum 255 character"`
	Fullname *string `json:"Fullname" form:"fullname" valid:"required~this field is required"`
	CardID   *string `json:"id_card_number" form:"id_card_number" valid:"required~this field is required,stringLength(16|16)~must contain 16 character,numeric~only numeric character"`
	Phone    *string `json:"phone" form:"phone" valid:"required~this field is required,stringlength(12|13)~must contain minimum 12 character and maximum 13 character,numeric~only numeric character"`
}

// ----- END OF AUTH FORM -----
