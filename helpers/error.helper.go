package helpers

import (
	"fmt"
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
	fmt.Println(err)
	resp := domains.ErrorBodyResponse{}
	appErr, ok := err.(domains.GeneralError)
	if !ok {
		msg := "there's something wrong"
		resp.Error = &msg
		return http.StatusInternalServerError, "internal server error", &resp
	}
	resp.Error = &appErr.Msg
	if appErr.ValidationErrors != nil {
		errs, _ := GenerateValidationError(appErr.ValidationErrors)
		resp.Errors = errs.Errors
		appErr.ValidationErrors = nil
	} else if appErr.DatabaseError != nil {
		errs := GenerateDatabaseKeyError(appErr.DatabaseError)
		resp.Errors = errs.Errors
		appErr.DatabaseError = nil
	}
	return appErr.Status, appErr.GeneralMessage, &resp
}
