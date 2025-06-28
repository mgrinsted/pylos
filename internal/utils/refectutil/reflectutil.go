package refectutil

import (
	"reflect"
	"strings"
)

func GetDBFields(dto interface{}) string {
	return getDBFields(reflect.TypeOf(dto))
}

func getDBFields(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			fields = append(fields, getDBFields(field.Type))
			continue
		}

		dbTag := field.Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			fields = append(fields, dbTag)
		}
	}
	return strings.Join(fields, ", ")
}