package dump

import (
	"reflect"
	"strings"
)

// IsScalar 简单类型
func IsScalar(v interface{}) bool {
	if v == nil {
		return true
	}

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {

	case reflect.Map,
		reflect.Struct,
		reflect.Slice, reflect.Array:
		return false

	case reflect.Chan:
		return false

	default:
		return true
	}
}

func hasLen(k reflect.Kind) bool {
	switch k {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return true
	default:
		return false
	}
}

func withTab(str string) string {
	return strings.Replace(str, "\n", "\n\t", -1)
}
