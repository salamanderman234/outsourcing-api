package helpers

import (
	"encoding/json"

	"github.com/salamanderman234/outsourcing-api/domains"
)

func Convert(source any, target any) error {
	jsonEncode, err := json.Marshal(source)
	if err != nil {
		return domains.ErrJsonConvert
	}
	err = json.Unmarshal(jsonEncode, target)
	if err != nil {
		return domains.ErrJsonConvert
	}
	return nil
}
