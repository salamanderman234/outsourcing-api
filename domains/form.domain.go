package domains

// ----- AUTH FORM -----
type BasicLoginForm struct {
	Username string `json:"username" form:"username" valid:"required~this field is required"`
	Password string `json:"password" form:"password" valid:"required~this field is required"`
}

// ----- END OF AUTH FORM -----
