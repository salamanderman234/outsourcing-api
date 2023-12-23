package domains

import "errors"

type GeneralError struct {
	Msg    string
	Field  *string
	Status int
}

func (a GeneralError) Error() string {
	return a.Msg
}

var (
	ErrRepository                    = errors.New("gorm error")
	ErrRepositoryInterfaceConversion = errors.New("repository interface error")
	// service
	ErrHashingPassword = errors.New("hashing password error")
	ErrInvalidCreds    = errors.New("invalid credentials error")
	ErrConversionType  = errors.New("conversion type error")
	ErrValidation      = errors.New("validation error")
	// token
	ErrGenerateToken = errors.New("generate token error")
	ErrExpiredToken  = errors.New("expired token error")
	ErrInvalidToken  = errors.New("invalid token error")
	// convert
	ErrJsonConvert = errors.New("json convert error")
	// bind
	ErrGenerateBindingErrs = errors.New("generate binding error")
	// http
	ErrBadRequest = errors.New("bad request error")
	// gorm
	ErrRecordNotFound     = errors.New("resource not found error")
	ErrDuplicateEntries   = errors.New("duplicate entries error")
	ErrForeignKeyViolated = errors.New("invalid foreign key error")
	ErrCardIdDuplicate    = errors.New("duplicate entries for identity card number field")
	ErrEmailDuplicate     = errors.New("duplicate entries for email field")
	// access
	ErrInvalidAccess = errors.New("invalid access error")
	// role
	ErrInvalidRole = errors.New("invalid request role")
	// file
	ErrGetMultipartFormData = errors.New("get multipart formdata error")
	ErrFileOpen             = errors.New("file open error")
	ErrFileCreate           = errors.New("file create error")
	ErrFileCopy             = errors.New("file copy error")
	ErrDeleteFile           = errors.New("file delete error")
)

type DatabaseKeyError struct {
	Msg    string
	Field  string
	Status int
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
