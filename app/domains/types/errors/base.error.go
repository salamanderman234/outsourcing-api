package custom_errors

import (
	"net/http"

	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
)

type GeneralError struct {
	Msg              string
	Status           int
	GeneralMessage   string
	ValidationErrors []responses.FieldError
	DatabaseError    error
}

func (a GeneralError) Error() string {
	return a.Msg
}

var (
	ErrInternalServer = GeneralError{
		Msg:            "there is something wrong",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Internal Server Error",
	}
	ErrTimeOut = GeneralError{
		Msg:            "server taking too long to respond",
		Status:         http.StatusRequestTimeout,
		GeneralMessage: "Request Error",
	}
	ErrBadRequest = GeneralError{
		Msg:            "invalid request",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	ErrRouteNotFound = GeneralError{
		Msg:            "requested route not found",
		Status:         http.StatusNotFound,
		GeneralMessage: "Request Error",
	}
	ErrEchoBinding = GeneralError{
		Msg:            "missing request body",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
)
