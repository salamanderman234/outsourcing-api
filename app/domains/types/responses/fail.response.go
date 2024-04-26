package responses

type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
	Error string `json:"error"`
}

type FailResponse struct {
	BaseResponse
	Detail    string       `json:"detail"`
	Errors    []FieldError `json:"errors,omitempty"`
	DebugMsg  *string      `json:"debug,omitempty"`
	ErrorType *string      `json:"error_type,omitempty"`
}
