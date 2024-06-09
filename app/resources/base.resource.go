package resources

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/salamanderman234/outsourcing-api/app/domains"
	"github.com/salamanderman234/outsourcing-api/app/types"
	"github.com/salamanderman234/outsourcing-api/configs"
	"gorm.io/gorm/schema"
)

type Resource struct {
	path       string
	model      domains.ModelInterface
	fileConfig types.FileConfig
	field      string
	policy     domains.Policy
}

func (r Resource) GetFullPath() string {
	return fmt.Sprintf("%s%s%s", configs.ResourceConfig.BasePath, r.fileConfig.BasePath, r.path)
}

func (r Resource) GetFieldValue() string {
	v := reflect.ValueOf(r.model)
	val := reflect.Indirect(v).FieldByName(r.field).Interface()
	if val == nil {
		return ""
	}
	if _, ok := val.(*string); !ok {
		return ""
	}
	strVal := val.(*string)
	if strVal != nil {
		return *strVal
	}
	return ""
}

func (r Resource) GetName() (string, error) {
	s, err := schema.Parse(r.model, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", s.Table, r.field), nil
}

func (r Resource) GetFileConfig() any {
	return r.fileConfig
}

func (r Resource) GetPolicy() domains.Policy {
	return r.policy
}

func (r Resource) GetData() domains.ModelInterface {
	return r.model
}

func (r Resource) GetFieldName() string {
	return r.field
}

func (r *Resource) SetData(data domains.ModelInterface) {
	dataName := reflect.TypeOf(data).Elem().Name()
	resourceName := reflect.TypeOf(r.model).Elem().Name()

	if dataName == resourceName {
		r.model = data
	}
}
