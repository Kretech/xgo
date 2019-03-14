package p

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/Kretech/xgo/encoding"
	"github.com/fatih/color"
)

var (
	MaxSliceLen = 32
	MaxMapLen   = 32

	SepKv = " => "

	Out io.Writer = os.Stdout
)

// Dump
func Dump(args ...interface{}) {
	DepthDump(1, args...)
}

// DD means Dump and Die
func DD(args ...interface{}) {
	DepthDump(1, args...)

	panic(`DD`)
}

func DepthDump(depth int, args ...interface{}) {
	names, compacted := DepthCompact(depth+1, args...)

	for _, name := range names {
		txt := ""

		if strings.HasPrefix(name, "&") {
			txt += color.New(color.Italic, color.FgMagenta).Sprint("&")
			name = name[1:]
		}

		txt += color.New(color.Italic, color.FgCyan).Sprint(name) + SepKv

		txt += serialize(compacted[name])

		_, _ = fmt.Fprintln(Out, txt)
	}
}

func serialize(originValue interface{}) (txt string) {
	result := originValue

	rT := reflect.TypeOf(originValue)
	rV := reflect.ValueOf(originValue)
	isPtr := false

	if rT.Kind() == reflect.Ptr {
		isPtr = true
		rT = rT.Elem()
		rV = rV.Elem()
	}

	// 基础类型
	switch rT.Kind() {
	case reflect.String:
		result = fmt.Sprintf(`"%v"`, rV.Interface())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		result = fmt.Sprint(rV.Interface())
	}

	if isPtr {
		txt += color.New(color.FgMagenta).Sprint("*")
	}

	if !IsScala(originValue) {
		// txt += fmt.Sprint(reflect.TypeOf(originValue))

		rTName := strings.Replace(rT.String(), " ", "", 1)
		head := color.New(color.FgGreen).Sprint(rTName) + " "

		func() {
			defer func() {
				recover()
			}()

			head += "("
			head += fmt.Sprintf("len=%v", color.New(color.FgYellow).Sprint(rV.Len()))
			//txt += fmt.Sprintf("cap=%v ", color.New(color.FgGreen).Sprint(reflect.ValueOf(originValue).Cap()))
			head += ") "
		}()

		// 恶心。。
		txt += head

		// ...

		switch rT.Kind() {
		case reflect.Slice:

			buf := bytes.Buffer{}
			buf.WriteString("[")
			for i := 0; i < rV.Len(); i++ {
				v := rV.Index(i).Interface()
				buf.WriteByte('\n')
				buf.WriteString(fmt.Sprintf("%d%s", i, SepKv))
				buf.WriteString(serialize(v))

				if i+1 >= MaxSliceLen {
					buf.WriteString(fmt.Sprintf("\n...\nother %d items...\n", rV.Len()-MaxSliceLen))
					break
				}
			}

			body := withTab(buf.String())

			body += "\n]"

			result = body

		case reflect.Map:

			buf := bytes.Buffer{}
			buf.WriteString("{")
			for i, key := range rV.MapKeys() {
				v := rV.MapIndex(key).Interface()
				buf.WriteByte('\n')
				buf.WriteString(serialize(key.Interface()))
				buf.WriteString(SepKv)
				buf.WriteString(serialize(v))

				if i+1 >= MaxMapLen {
					break
				}
			}

			body := withTab(buf.String())

			body += "\n}"

			result = body

		default:
			result = encoding.JsonEncode(originValue, encoding.OptIndentTab)
		}
	}

	txt += fmt.Sprint(result)

	return
}

// 简单类型
func IsScala(v interface{}) bool {
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

	default:
		return true
	}
}

func withTab(str string) string {
	return strings.Replace(str, "\n", "\n\t", -1)
}
