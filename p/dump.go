package p

import (
	"fmt"
	"reflect"

	"github.com/Kretech/xgo/encoding"
	"github.com/fatih/color"
)

func Dump(args ...interface{}) {
	r := DepthCompact(1, args...)

	for k, originValue := range r {

		txt := color.New(color.Italic, color.FgYellow).Sprint(k) + " => "

		vi := originValue

		if !IsScala(originValue) {
			// txt += fmt.Sprint(reflect.TypeOf(originValue))
			txt += color.New(color.FgCyan).Sprint(reflect.TypeOf(originValue))
			vi = encoding.JsonEncode(originValue, encoding.OptIndentTab)
		}

		txt += fmt.Sprint(vi)

		fmt.Println(txt)
	}

	// fmt.Println(encoding.JsonEncode(r, encoding.OptIndentTab))
}

func IsScala(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {

	case reflect.Map,
		reflect.Struct,
		reflect.Slice, reflect.Array:
		return false

	default:
		return true
	}
}
