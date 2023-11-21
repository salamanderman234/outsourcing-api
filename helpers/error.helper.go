package helpers

import (
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func GenerateBindingError(errs []error) (domains.ErrorBodyResponse, error) {
	resp := domains.ErrorBodyResponse{}
	for _, err := range errs {
		cErr, ok := err.(*echo.BindingError)
		if !ok {
			return resp, domains.ErrConversionType
		}
		field := strings.ToLower(cErr.Field)
		msg := cErr.Message.(string)
		resp.Errors = append(resp.Errors, domains.ErrorDetailResponse{
			Field:  &field,
			Detail: &msg,
		})
	}
	return resp, nil
}

func GenerateValidationError(errs error) (domains.ErrorBodyResponse, error) {
	resp := domains.ErrorBodyResponse{}
	for _, err := range errs.(govalidator.Errors) {
		cErr, ok := err.(govalidator.Error)
		if !ok {
			return resp, domains.ErrConversionType
		}
		field := strings.ToLower(cErr.Name)
		rule := cErr.Validator
		msg := cErr.Error()
		resp.Errors = append(resp.Errors, domains.ErrorDetailResponse{
			Field:  &field,
			Rule:   &rule,
			Detail: &msg,
		})
	}
	return resp, nil
}

func HandleError(err error) (int, string, *domains.ErrorBodyResponse) {
	resp := domains.ErrorBodyResponse{}
	_, ok := err.(govalidator.Errors)
	if ok {
		resp, _ = GenerateValidationError(err)
		return http.StatusBadRequest, domains.ErrValidation.Error(), &resp
	}
	msg := err.Error()
	if err == domains.ErrBadRequest {
		errString := "invalid user request"
		resp.Error = &errString
		return http.StatusBadRequest, msg, &resp
	} else if err == domains.ErrInvalidToken {
		errString := "token is invalid"
		resp.Error = &errString
		return http.StatusUnauthorized, msg, &resp
	} else if err == domains.ErrRecordNotFound {
		errString := "not found"
		resp.Error = &errString
		return http.StatusNotFound, msg, &resp
	} else if err == domains.ErrDuplicateEntries {
		errString := "duplicate data entries"
		resp.Error = &errString
		return http.StatusUnprocessableEntity, msg, &resp
	} else if err == domains.ErrInvalidAccess {
		errString := "don't have access to these resources"
		resp.Error = &errString
		return http.StatusForbidden, msg, &resp
	} else if err == domains.ErrForeignKeyViolated {
		errString := "token is invalid"
		resp.Error = &errString
		return http.StatusUnprocessableEntity, msg, &resp
	} else if err == domains.ErrExpiredToken {
		errString := "token is expired"
		resp.Error = &errString
		return http.StatusUnauthorized, msg, &resp
	} else if err == domains.ErrInvalidCreds {
		errString := "wrong email or password"
		resp.Error = &errString
		return http.StatusUnauthorized, msg, &resp
	} else if err != nil {
		errString := "there's something wrong"
		resp.Error = &errString
		return http.StatusInternalServerError, "internal server error", &resp
	}
	return 200, "ok", nil
}
