package helpers

import "github.com/asaskevich/govalidator"

type validateHelper struct{}

func (validateHelper) Validate(data any) error {
	_, err := govalidator.ValidateStruct(data)
	return err
}

var Validator = validateHelper{}
