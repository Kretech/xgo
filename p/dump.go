package p

import (
	"fmt"
	"reflect"

	"github.com/Kretech/xgo/encoding"
)

func Dump(args ...interface{}) {
	r := DepthCompact(1, args...)

	for k, v := range r {
		vi := v

		if !IsScala(v) {
			vi = encoding.JsonEncode(v, encoding.OptIndentTab)
		}

		fmt.Printf("%v => %v\n", k, vi)
	}

	// fmt.Println(encoding.JsonEncode(r, encoding.OptIndentTab))
}

func IsScala(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {

	case reflect.Map, reflect.Slice, reflect.Struct:
		return false

	default:
		return true
	}
}
