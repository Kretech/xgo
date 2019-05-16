package dump

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/Kretech/xgo/encoding"
	"github.com/fatih/color"
)

var (
	MaxSliceLen = 32
	MaxMapLen   = 32

	SepKv = " => "

	StringQuota = `"`
)

func Serialize(originValue interface{}) (serialized string) {
	if originValue == nil {
		return "<nil>"
	}

	result := originValue

	T := reflect.TypeOf(originValue)
	V := reflect.ValueOf(originValue)
	isPtr := false

	if T.Kind() == reflect.Ptr {
		isPtr = true
		T = T.Elem()
		V = V.Elem()
	}

	if !V.IsValid() {
		return "<zeroValue>"
	}

	// 基础类型
	switch T.Kind() {
	case reflect.String:
		quota := StringQuota
		s := V.Interface().(string)
		if strings.Contains(s, `"`) {
			quota = "`"
		}
		result = fmt.Sprintf(`%s%v%s`, quota, s, quota)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		result = fmt.Sprint(V.Interface())
	}

	if IsScalar(originValue) {
		serialized = fmt.Sprint(result)
		return
	}

	if isPtr {
		serialized += color.New(color.FgMagenta).Sprint("*")
	}

	rTName := strings.Replace(T.String(), " ", "", 1)
	head := color.New(color.FgGreen).Sprint(rTName) + " "

	func() {
		defer func() {
			recover()
		}()

		if hasLen(T.Kind()) {
			head += "("
			head += fmt.Sprintf("len=%v", color.New(color.FgYellow).Sprint(V.Len()))
			//txt += fmt.Sprintf("cap=%v ", color.New(color.FgGreen).Sprint(reflect.ValueOf(originValue).Cap()))
			head += ") "
		}
	}()

	// 恶心。。
	serialized += head

	// ...

	switch T.Kind() {
	case reflect.Array, reflect.Slice:

		buf := bytes.Buffer{}
		buf.WriteString("[")
		for i := 0; i < V.Len(); i++ {
			v := V.Index(i).Interface()
			buf.WriteByte('\n')
			buf.WriteString(fmt.Sprintf("%d%s", i, SepKv))
			buf.WriteString(Serialize(v))

			if i+1 >= MaxSliceLen {
				buf.WriteString(fmt.Sprintf("\n...\nother %d items...\n", V.Len()-MaxSliceLen))
				break
			}
		}

		body := withTab(buf.String())

		body += "\n]"

		result = body

	case reflect.Map:

		buf := bytes.Buffer{}
		buf.WriteString("{")
		for i, key := range V.MapKeys() {
			v := V.MapIndex(key).Interface()
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

	case reflect.Struct:
		buf := bytes.Buffer{}
		buf.WriteString("{")
		for i := 0; i < V.NumField(); i++ {
			field := V.Field(i)
			fieldT := V.Type().Field(i)
			buf.WriteByte('\n')
			buf.WriteString(fieldT.Name)
			buf.WriteString(": ")
			if field.CanInterface() {
				buf.WriteString(Serialize(field.Interface()))
			} else {
				newValue := reflect.NewAt(fieldT.Type, unsafe.Pointer(field.UnsafeAddr())).Elem()
				buf.WriteString(Serialize(newValue.Interface()))
			}

			if i+1 >= MaxMapLen {
				break
			}
		}

		body := withTab(buf.String())

		body += "\n}"

		result = body

	case reflect.Func:
		result = fmt.Sprintf("{ &%v }", originValue)

	case reflect.Chan:
		result = fmt.Sprintf("{...}")

	default:
		result = fmt.Sprintf("(%T)", originValue) + encoding.JsonEncode(originValue, encoding.OptIndentTab)
	}

	serialized += fmt.Sprint(result)

	return
}
