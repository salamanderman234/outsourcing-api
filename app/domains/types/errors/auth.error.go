package custom_errors

import "net/http"

var (
	ErrForbiden = GeneralError{
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
)
