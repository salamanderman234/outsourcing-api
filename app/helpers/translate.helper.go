package helpers

import (
	"encoding/json"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var errorMap = map[error]types.GeneralError{
	gorm.ErrRecordNotFound:              types.ErrRecordNotFound,
	bcrypt.ErrMismatchedHashAndPassword: types.ErrNotMatched,
	echo.ErrBadRequest:                  types.ErrBadRequest,
	jwt.ErrTokenExpired:                 types.ErrTokenExpired,
	jwt.ErrTokenSignatureInvalid:        types.ErrInvalidToken,
	gorm.ErrDuplicatedKey:               types.ErrDuplicateEntries,
}

type translateHelper struct{}

func (t translateHelper) TranslateError(err error) types.GeneralError {
	httpErr, ok := err.(*echo.HTTPError)
	if ok {
		msg := httpErr.Message
		returnErr := types.ErrEchoRequest
		returnErr.Status = httpErr.Code
		returnErr.Msg, _ = msg.(string)
		return returnErr
	}
	if conErrs, ok := err.(govalidator.Errors); ok {
		return t.translateGovalidatorError(conErrs)
	}
	if conErr, ok := err.(types.GeneralError); ok {
		return conErr
	}
	result, ok := errorMap[err]
	if !ok {
		return types.ErrInternalServer
	}
	return result
}

func (t translateHelper) loopErrors(errs govalidator.Errors, pasts ...[]types.FieldError) []types.FieldError {
	var results []types.FieldError
	for _, err := range errs {
		con, ok := err.(govalidator.Error)
		if ok {
			field := strings.ToLower(con.Name)
			rule := strings.ToLower(con.Validator)
			detail := con.Err.Error()
			new := types.FieldError{
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

func (t translateHelper) translateGovalidatorError(errs govalidator.Errors) types.GeneralError {
	final := types.ErrValidate
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
