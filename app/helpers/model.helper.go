package helpers

import (
	"sync"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"gorm.io/gorm/schema"
)

type modelHelper struct{}

func (modelHelper) GetNamingStrategy(model domains.ModelInterface) string {
	s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return ""
	}
	return s.Table
}

var Model = modelHelper{}
