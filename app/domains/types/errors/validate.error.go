package custom_errors

import "net/http"

var (
	ErrValidate = GeneralError{
		Msg:            "validation does not meet the requirements",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Validation Error",
	}
)
