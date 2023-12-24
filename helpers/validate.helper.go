package helpers

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/salamanderman234/outsourcing-api/domains"
)

func Validate(data any) (bool, error) {
	ok, errs := govalidator.ValidateStruct(data)
	if !ok {
		valErr := domains.ErrValidation
		valErr.ValidationErrors = errs
		return false, valErr
	}
	return true, nil
}

func GenerateErrorResponse(errs error) []domains.ErrorDetailResponse {
	var results []domains.ErrorDetailResponse
	errList, ok := errs.(govalidator.Errors)
	if !ok {
		return nil
	}
	for _, err := range errList {
		con := err.(govalidator.Error)
		field := strings.ToLower(con.Name)
		rule := strings.ToLower(con.Validator)
		detail := con.Error()
		new := domains.ErrorDetailResponse{
			Field:  &field,
			Rule:   &rule,
			Detail: &detail,
		}
		results = append(results, new)
	}
	return results
}
