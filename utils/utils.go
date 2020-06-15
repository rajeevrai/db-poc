package utils

import (
	"reflect"
)

func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	reflectVal := reflect.ValueOf(val)

	switch reflectVal.Kind() {
	case reflect.Int:
		return val.(int) == 0

	case reflect.Int64:
		return val.(int64) == 0

	case reflect.String:
		return val.(string) == ""

	case reflect.Map:
		fallthrough

	case reflect.Slice:
		return reflectVal.IsNil() || reflectVal.Len() == 0
	}

	return false
}
