package helpers

import (
	"encoding/json"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	custom_errors "github.com/salamanderman234/outsourcing-api/app/domains/types/errors"
	"github.com/salamanderman234/outsourcing-api/app/domains/types/responses"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var errorMap = map[error]custom_errors.GeneralError{
	gorm.ErrRecordNotFound:              custom_errors.ErrRecordNotFound,
	bcrypt.ErrMismatchedHashAndPassword: custom_errors.ErrNotMatched,
	echo.ErrBadRequest:                  custom_errors.ErrBadRequest,
	jwt.ErrTokenExpired:                 custom_errors.ErrTokenExpired,
	jwt.ErrTokenSignatureInvalid:        custom_errors.ErrInvalidToken,
}

type translateHelper struct{}

func (t translateHelper) TranslateError(err error) custom_errors.GeneralError {
	if conErrs, ok := err.(govalidator.Errors); ok {
		return t.translateGovalidatorError(conErrs)
	}
	if conErr, ok := err.(custom_errors.GeneralError); ok {
		return conErr
	}
	result, ok := errorMap[err]
	if !ok {
		return custom_errors.ErrInternalServer
	}
	return result
}

func (t translateHelper) loopErrors(errs govalidator.Errors, pasts ...[]responses.FieldError) []responses.FieldError {
	var results []responses.FieldError
	for _, err := range errs {
		con, ok := err.(govalidator.Error)
		if ok {
			field := strings.ToLower(con.Name)
			rule := strings.ToLower(con.Validator)
			detail := con.Err.Error()
			new := responses.FieldError{
				Field: field,
				Rule:  rule,
				Error: detail,
			}
			results = append(results, new)
		} else if con2, ok2 := err.(govalidator.Errors); ok2 {
			callResults := t.loopErrors(con2, results)
			results = append(results, callResults...)
		}
	}
	return results
}

func (t translateHelper) translateGovalidatorError(errs govalidator.Errors) custom_errors.GeneralError {
	final := custom_errors.ErrValidate
	final.ValidationErrors = t.loopErrors(errs)
	return final
}

func (translateHelper) TranslateStruct(source any, target any) error {
	jsonEncode, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonEncode, target)
	if err != nil {
		return err
	}
	return nil
}

var Translator = translateHelper{}
