package db

import (
	"reflect"
	"strings"
)

// Extracts a comma-separated list of DB fields from a struct (including aliases like "as ...")
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

		// Embedded structs (recursive)
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
