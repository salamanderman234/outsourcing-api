package domains

import (
	"errors"
	"net/http"
)

type GeneralError struct {
	Msg              string
	Status           int
	GeneralMessage   string
	ValidationErrors error
	DatabaseError    error
}

func (a GeneralError) Error() string {
	return a.Msg
}

var (
	ErrRepository = GeneralError{
		Msg:            "repository error",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Internal Server Error",
	}
	ErrRepositoryInterfaceConversion = GeneralError{
		Msg:            "failed to perform type conversion on repository level",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Conversion Error",
	}
	// service
	ErrHashingPassword = errors.New("hashing password error")
	ErrInvalidCreds    = GeneralError{
		Msg:            "invalid credentials",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Request Error",
	}
	ErrConversionType = GeneralError{
		Msg:            "failed to perform type conversion",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Conversion Error",
	}
	ErrValidation = GeneralError{
		Msg:            "request does not match validation",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	// token
	ErrGenerateToken = errors.New("generate token error")
	ErrExpiredToken  = GeneralError{
		Msg:            "token is expired, generate new one",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Token Error",
	}
	ErrInvalidToken = GeneralError{
		Msg:            "token is invalid",
		Status:         http.StatusUnauthorized,
		GeneralMessage: "Token Error",
	}
	// convert
	ErrJsonConvert = GeneralError{
		Msg:            "failed to perform json conversion",
		Status:         http.StatusNotFound,
		GeneralMessage: "Conversion Error",
	}
	// bind
	ErrGenerateBindingErrs = GeneralError{
		Msg:            "failed to generate binding error",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "Binding Error",
	}
	// http
	ErrBadRequest = GeneralError{
		Msg:            "bad request",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	// gorm
	ErrRecordNotFound = GeneralError{
		Msg:            "requested resource is not found",
		Status:         http.StatusNotFound,
		GeneralMessage: "Not Found Error",
	}
	ErrDuplicateEntries = GeneralError{
		Msg:            "duplicate entries not allowed",
		Status:         http.StatusConflict,
		GeneralMessage: "Resource Conflict Error",
	}

	ErrForeignKeyViolated = GeneralError{
		Msg:            "invalid foreign key",
		Status:         http.StatusUnprocessableEntity,
		GeneralMessage: "Foreign Key Violated  Error",
	}
	// access
	ErrInvalidAccess = GeneralError{
		Msg:            "access denied",
		Status:         http.StatusForbidden,
		GeneralMessage: "Invalid Access Error",
	}
	// role
	ErrInvalidRole = GeneralError{
		Msg:            "invalid role",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Invalid Access Error",
	}
	// file
	ErrGetMultipartFormData = GeneralError{
		Msg:            "wrong content type, must be multipart form data",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
	}
	ErrFileOpen = GeneralError{
		Msg:            "failed to open file",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "File Error",
	}
	ErrFileCreate = GeneralError{
		Msg:            "failed to create file",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "File Error",
	}
	ErrFileCopy = GeneralError{
		Msg:            "failed to copy file",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "File Error",
	}
	ErrDeleteFile = GeneralError{
		Msg:            "failed to delete file",
		Status:         http.StatusInternalServerError,
		GeneralMessage: "File Error",
	}
)

type DatabaseKeyError struct {
	Msg   string
	Field string
}

func (e DatabaseKeyError) Error() string {
	return e.Msg
}

type DatabaseKeyErrors struct {
	Errors []DatabaseKeyError
	Msg    string
}

func (e DatabaseKeyErrors) Error() string {
	return e.Msg
}
