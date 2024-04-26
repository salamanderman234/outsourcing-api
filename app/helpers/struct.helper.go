package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

type structHelper struct{}

func (s structHelper) GetVisibleAttributes(data any) any {
	if data == nil {
		return nil
	}
	reflectVal := reflect.ValueOf(data)
	mappedData := map[string]any{}
	if reflectVal.Type().Kind() != reflect.Struct {
		return data
	}
	for i := 0; i < reflectVal.NumField(); i++ {
		field := reflectVal.Type().Field(i)
		jsonTag := field.Tag.Get("json")
		name := strings.ReplaceAll(jsonTag, ",omitempty", "")
		isOmitempty := strings.Contains(jsonTag, "omitempty")
		value := reflectVal.Field(i).Interface()
		isVisible := field.Tag.Get("visible")

		fmt.Println(isOmitempty, value, value == nil, name)
		if field.Type.Kind() == reflect.Pointer {
			if isOmitempty && reflect.ValueOf(value).IsNil() {
				continue
			}
		}
		if name == "model" {
			result, _ := s.GetVisibleAttributes(value).(map[string]any)
			for key, value := range result {
				mappedData[key] = value
			}
		} else if field.Type.Kind() == reflect.Struct {
			result := s.GetVisibleAttributes(value)
			if isVisible == "true" {
				mappedData[name] = result
			}
		} else {
			if isVisible == "true" {
				mappedData[name] = value
			}
		}
	}

	return mappedData
}

func (s structHelper) LoopGetVisibleStruct(datas any) any {
	var results []any
	if datas == nil {
		return nil
	}
	rt := reflect.ValueOf(datas)
	desireType := reflect.SliceOf(rt.Type().Elem())
	if rt.Type().ConvertibleTo(desireType) {
		converted := rt.Convert(desireType)
		for i := 0; i < converted.Len(); i++ {
			data := converted.Index(i)
			filtered := s.GetVisibleAttributes(data.Interface())
			results = append(results, filtered)
		}

	}
	return results
}

var Struct = structHelper{}
