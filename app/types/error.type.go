package types

import (
	"net/http"
)

type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
	Error string `json:"error"`
}

type GeneralError struct {
	Name             string
	Msg              string
	Status           int
	GeneralMessage   string
	ValidationErrors []FieldError
	DatabaseError    error
}

func (a GeneralError) Error() string {
	return a.Msg
}

var (
	ErrForbiden = GeneralError{
		Name:           "Forbidden Error",
		Msg:            "you dont have permission to this resources",
		Status:         http.StatusForbidden,
		GeneralMessage: "Forbidden Error",
	}
	ErrUnauthorized = GeneralError{
		Msg:            "invalid authentication",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Unauthorized Error",
	}
	ErrNotMatched = GeneralError{
		Msg:            "email or password is wrong",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Unauthorized Error",
	}
	ErrNotVerifiedUser = GeneralError{
		Msg:            "verify your account first",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Unauthorized Error",
	}
	ErrTokenExpired = GeneralError{
		Msg:            "access token is expired",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Token Error",
	}
	ErrInvalidToken = GeneralError{
		Msg:            "access token is invalid",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Token Error",
	}
	ErrInternalServer = GeneralError{
		Name:           "Server Error",
		Msg:            "there is something wrong",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Internal Server Error",
	}
	ErrTimeOut = GeneralError{
		Name:           "Timeout Error",
		Msg:            "server taking too long to respond",
		Status:         http.StatusRequestTimeout,
		GeneralMessage: "Request Error",
	}
	ErrBadRequest = GeneralError{
		Name:           "Request Error",
		Msg:            "invalid request",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	ErrRouteNotFound = GeneralError{
		Name:           "Not Found Error",
		Msg:            "requested route not found",
		Status:         http.StatusNotFound,
		GeneralMessage: "Request Error",
	}
	ErrEchoRequest = GeneralError{
		Name:           "Echo Error",
		Msg:            "invalid request body",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	ErrRecordNotFound = GeneralError{
		Name:           "Record Not Found Error",
		Msg:            "requested resource is not found",
		Status:         http.StatusNotFound,
		GeneralMessage: "Not Found Error",
	}
	ErrDuplicateEntries = GeneralError{
		Name:           "Duplicate Entry Error",
		Msg:            "duplicate entries are not allowed",
		Status:         http.StatusConflict,
		GeneralMessage: "Resource Conflict Error",
	}
	ErrValidate = GeneralError{
		Name:           "Validation Error",
		Msg:            "validation does not meet the requirements",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Validation Error",
	}
	ErrUnprocessableEntity = GeneralError{
		Name:           "Unprocessable Entity",
		Msg:            "server cant process your request",
		Status:         http.StatusUnprocessableEntity,
		GeneralMessage: "Request Error",
	}
)
