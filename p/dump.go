package p

import (
	"fmt"
	"os"
	"reflect"

	"github.com/Kretech/xgo/encoding"
	"github.com/fatih/color"
)

// Dump
func Dump(args ...interface{}) {
	DepthDump(1, args...)
}

// DD means Dump and Die
func DD(args ...interface{}) {
	DepthDump(1, args...)

	os.Exit(0)
}

func DepthDump(depth int, args ...interface{}) {
	r := DepthCompact(depth+1, args...)

	for k, originValue := range r {

		txt := color.New(color.Italic, color.FgYellow).Sprint(k) + " => "

		txt += serialize(originValue)

		fmt.Println(txt)
	}

	// fmt.Println(encoding.JsonEncode(r, encoding.OptIndentTab))
}

func serialize(originValue interface{}) (txt string) {
	vi := originValue

	if !IsScala(originValue) {
		// txt += fmt.Sprint(reflect.TypeOf(originValue))

		txt += color.New(color.FgCyan).Sprint(reflect.TypeOf(originValue))
		vi = encoding.JsonEncode(originValue, encoding.OptIndentTab)
	}

	txt += fmt.Sprint(vi)

	return
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
