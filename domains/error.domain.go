package domains

import "errors"

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
	ErrBadRequest = errors.New("bad request")
)
