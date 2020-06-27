package dump

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"

	"github.com/fatih/color"

	"github.com/Kretech/xgo/encoding"
)

var (
	// uint8(97) => 'a'
	OptShowUint8AsChar = true

	// []uint8{'a','b'} => "ab"
	OptShowUint8sAsString = true

	// 按字典序显示map.keys
	OptSortMapKeys = true
)

var (
	MaxSliceLen = 32
	MaxMapLen   = 32

	SepKv = " => "

	StringQuota = `"`
)

const (
	Zero = `<zero>`
	Nil  = "<nil>"
)

func serializeScalar(V reflect.Value) (result string) {

	switch V.Type().Kind() {
	case reflect.String:
		quota := StringQuota
		s := fmt.Sprint(V.Interface())

		if strings.Contains(s, StringQuota) {
			quota = "`"
		}
		result = fmt.Sprintf(`%s%v%s`, quota, s, quota)

	case reflect.Uint8:
		if OptShowUint8AsChar {
			result = fmt.Sprintf("%c", V.Uint())
		} else {
			result = fmt.Sprint(V.Interface())
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint /*reflect.Uint8,*/, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:

		result = fmt.Sprint(V.Interface())

	default:
		result = fmt.Sprint(V.Interface())
	}

	return
}

func Serialize(originValue interface{}) (serialized string) {
	if originValue == nil {
		return Nil
	}

	result := originValue

	var V reflect.Value

	switch v := originValue.(type) {
	case reflect.Value:
		V = v
	default:
		V = reflect.ValueOf(originValue)
	}

	T := V.Type()
	isPtr := false

	if T.Kind() == reflect.Ptr {
		isPtr = true
		T = T.Elem()
		V = V.Elem()
	}

	if !V.IsValid() {
		return Zero
	}

	// 基础类型
	if IsScalar(originValue) {
		serialized = serializeScalar(V)
		return
	}

	if isPtr {
		serialized += color.New(color.FgMagenta).Sprint("*")
	}

	//复合类型
	// 1. 先打 head，标明类型
	// 2. reflect

	rTName := strings.Replace(T.String(), " ", "", 1)
	head := color.New(color.FgGreen).Sprint(rTName) + " "

	func() {
		if hasLen(T.Kind()) {
			head += "("
			head += fmt.Sprintf("len=%v", color.New(color.FgYellow).Sprint(V.Len()))
			//txt += fmt.Sprintf("cap=%v ", color.New(color.FgGreen).Sprint(reflect.ValueOf(originValue).Cap()))
			head += ") "
		}
	}()

	// 恶心。。
	serialized += head

	// special complex
	switch v := originValue.(type) {
	case []byte:
		if !OptShowUint8sAsString {
			break
		}

		serialized += Serialize(string(v))
		return
	}

	// ...

	switch T.Kind() {
	case reflect.Array, reflect.Slice:

		buf := bytes.Buffer{}
		buf.WriteString("[")
		notEmpty := false
		for i := 0; i < V.Len(); i++ {
			v := V.Index(i)
			vi := v.Interface()
			vs := Serialize(vi)
			if vs != Zero {
				buf.WriteByte('\n')
				buf.WriteString(fmt.Sprintf("%d%s", i, SepKv))
				buf.WriteString(vs)

				notEmpty = true
			}

			if i+1 >= MaxSliceLen {
				buf.WriteString(fmt.Sprintf("\n...\nother %d items...\n", V.Len()-MaxSliceLen))
				break
			}
		}

		body := buf.String()
		if notEmpty {
			body = withTab(body) + "\n"
		}
		body += "]"

		result = body

	case reflect.Map:

		type item struct {
			key   string
			value string
		}
		items := make([]item, 0, V.Len())

		buf := bytes.Buffer{}
		for i, key := range V.MapKeys() {
			v := V.MapIndex(key).Interface()
			items = append(items, item{Serialize(key.Interface()), Serialize(v)})

			if i+1 >= MaxMapLen {
				break
			}
		}

		if OptSortMapKeys {
			sort.Slice(items, func(i, j int) bool {
				return items[i].key < items[j].key
			})
		}

		buf.WriteString("{")
		for _, item := range items {
			buf.WriteByte('\n')
			buf.WriteString(item.key)
			buf.WriteString(SepKv)
			buf.WriteString(item.value)
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
			} else if field.CanAddr() {
				newValue := reflect.NewAt(fieldT.Type, unsafe.Pointer(field.UnsafeAddr())).Elem()
				buf.WriteString(Serialize(newValue.Interface()))
			} else {
				buf.WriteString("unaddressable: " + field.String())
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
