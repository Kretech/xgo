package dump

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/Kretech/xgo/encoding"
	"github.com/fatih/color"
)

var (
	MaxSliceLen = 32
	MaxMapLen   = 32

	SepKv = " => "

	StringQuota = `"`
)

func Serialize(originValue interface{}) (txt string) {
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
		result = fmt.Sprintf(`%s%v%s`, StringQuota, rV.Interface(), StringQuota)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		result = fmt.Sprint(rV.Interface())
	}

	if isPtr {
		txt += color.New(color.FgMagenta).Sprint("*")
	}

	if !IsScalar(originValue) {
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
				buf.WriteString(Serialize(v))

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
				buf.WriteString(Serialize(key.Interface()))
				buf.WriteString(SepKv)
				buf.WriteString(Serialize(v))

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
