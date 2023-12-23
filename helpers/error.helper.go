package helpers

import (
	"errors"
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

func GenerateDatabaseKeyError(err error) *domains.ErrorBodyResponse {
	databaseKey, ok := err.(domains.DatabaseKeyError)
	if !ok {
		return nil
	}
	rule := "foreign key"
	resp := domains.ErrorBodyResponse{
		Errors: []domains.ErrorDetailResponse{
			{Field: &databaseKey.Field, Rule: &rule, Detail: &databaseKey.Msg},
		},
	}
	return &resp
}

func GenerateValidationError(errs error) (domains.ErrorBodyResponse, error) {
	resp := domains.ErrorBodyResponse{}
	for _, err := range errs.(govalidator.Errors) {
		cErr, ok := err.(govalidator.Error)
		if ok {
			field := strings.ToLower(cErr.Name)
			rule := cErr.Validator
			msg := cErr.Error()
			resp.Errors = append(resp.Errors, domains.ErrorDetailResponse{
				Field:  &field,
				Rule:   &rule,
				Detail: &msg,
			})
		} else {
			nestedErrs, ok := err.(govalidator.Errors)
			if ok {
				for _, nestedErr := range nestedErrs {
					nestedErrConvert := nestedErr.(govalidator.Error)
					field := strings.ToLower(nestedErrConvert.Name)
					rule := nestedErrConvert.Validator
					msg := nestedErrConvert.Error()
					resp.Errors = append(resp.Errors, domains.ErrorDetailResponse{
						Field:  &field,
						Rule:   &rule,
						Detail: &msg,
					})
				}
			}

		}
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
	_, ok = err.(domains.DatabaseKeyError)
	if ok {
		resp := GenerateDatabaseKeyError(err)
		return http.StatusUnprocessableEntity, domains.ErrForeignKeyViolated.Error(), resp
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
		return http.StatusConflict, msg, &resp
	} else if errors.Is(err, domains.ErrEmailDuplicate) {
		rule := "unique"
		field := "email"
		detail := "this email is already exists"
		resp.Errors = []domains.ErrorDetailResponse{
			{
				Field:  &field,
				Rule:   &rule,
				Detail: &detail,
			},
		}
		return http.StatusConflict, msg, &resp
	} else if err == domains.ErrInvalidAccess {
		errString := "don't have access to these resources"
		resp.Error = &errString
		return http.StatusForbidden, msg, &resp
	} else if errors.Is(err, domains.ErrInvalidRole) {
		errString := "invalid role"
		resp.Error = &errString
		return http.StatusBadRequest, msg, &resp
	} else if errors.Is(err, domains.ErrForeignKeyViolated) {
		errString := "foreign key error"
		resp.Error = &errString
		return http.StatusUnprocessableEntity, msg, &resp
	} else if errors.Is(err, domains.ErrExpiredToken) {
		errString := "token is expired"
		resp.Error = &errString
		return http.StatusUnauthorized, msg, &resp
	} else if errors.Is(err, domains.ErrInvalidCreds) {
		errString := "wrong email or password"
		resp.Error = &errString
		return http.StatusUnauthorized, msg, &resp
	} else if errors.Is(err, domains.ErrGetMultipartFormData) {
		errString := "request content type must be multipart/form-data"
		resp.Error = &errString
		return http.StatusBadRequest, msg, &resp
	} else if err != nil {
		errString := "there's something wrong"
		resp.Error = &errString
		return http.StatusInternalServerError, "internal server error", &resp
	}
	return 200, "ok", nil
}
