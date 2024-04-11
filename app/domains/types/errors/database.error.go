package custom_errors

import "net/http"

var (
	ErrRecordNotFound = GeneralError{
		Msg:            "requested resource is not found",
		Status:         http.StatusNotFound,
		GeneralMessage: "Not Found Error",
	}
	ErrDuplicateEntries = GeneralError{
		Msg:            "duplicate entries are not allowed",
		Status:         http.StatusConflict,
		GeneralMessage: "Resource Conflict Error",
	}
)
