package domains

import (
	"errors"
	"net/http"

	"github.com/asaskevich/govalidator"
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
	ErrRequestTimeOut = GeneralError{
		Msg:            "server taking too long to respond",
		Status:         http.StatusRequestTimeout,
		GeneralMessage: "Request Error",
	}
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
	ErrUnprocessAbleEntity = GeneralError{
		Msg:            "rule violation",
		Status:         http.StatusUnprocessableEntity,
		GeneralMessage: "Status Error",
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
	ErrUnmatchedData = GeneralError{
		Msg:            "this data pair cannot be used",
		Status:         http.StatusUnprocessableEntity,
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
	ErrMissingId = GeneralError{
		Msg:            "query params id is required and must be an unsigned integer",
		Status:         http.StatusBadRequest,
		GeneralMessage: "Request Error",
		ValidationErrors: govalidator.Errors{
			govalidator.Error{
				Name:                     "id",
				Validator:                "required and must be an unsigned integer",
				CustomErrorMessageExists: true,
				Err:                      errors.New("this field is required and must be an unsigned integer"),
			},
		},
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
	ErrInvalidFileType = GeneralError{
		Msg:            "invalid file type",
		Status:         http.StatusUnprocessableEntity,
		GeneralMessage: "File Error",
	}
	ErrFileSize = GeneralError{
		Msg:            "file size is too large",
		Status:         http.StatusRequestEntityTooLarge,
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
