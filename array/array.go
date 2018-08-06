package array

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"

	"github.com/Kretech/xgo/dict"
)

// Stream 提供对一组数据集的操作，接口 api 参考自 Laravel 的 Collection 和 Java 的 stream
type Array struct {
	*list.List
}

func (a *Array) String() string {
	b := bytes.NewBuffer([]byte{})
	a.each(func(it *list.Element) {
		b.WriteString(fmt.Sprintf("%v", it.Value))
		if it.Next() != nil {
			b.WriteByte(' ')
		}
	})

	return b.String()
}

func newArray() *Array {
	return &Array{
		list.New(),
	}
}

// Of 提供把普通列表转换成 Array 的入口
func Values(elements ...interface{}) *Array {
	a := newArray()
	for idx, _ := range elements {
		a.PushBack(elements[idx])
	}
	return a
}

func Slice(slice interface{}) *Array {
	e := reflect.ValueOf(slice)
	if e.Kind() != reflect.Slice {
		panic("array.Slice() must receive a slice ([]type)")
	}

	l := newArray()
	for i := 0; i < e.Len(); i++ {
		l.PushBack(e.Index(i).Interface())
	}

	return l
}

func (this *Array) KeyBy(field string) *dict.MapDict {
	d := dict.NewMapDict()

	for it := this.Front(); it != nil; it = it.Next() {
		key := getField(it.Value, field)
		d.Set(key, it.Value)
	}
	return d
}

func (this *Array) each(fn func(*list.Element)) {
	for it := this.Front(); it != nil; it = it.Next() {
		fn(it)
	}
}

func getField(v interface{}, field string) interface{} {

	elem := reflect.ValueOf(v).Elem()
	switch elem.Kind() {
	case reflect.Struct:
		return getStructField(elem, field)
	case reflect.Map:
		return getMapField(elem, field)
	}

	panic("")
}

func getMapField(v reflect.Value, field string) interface{} {
	return v.MapIndex(reflect.ValueOf(field)).String()
}

func getStructField(v reflect.Value, field string) interface{} {
	return v.FieldByName(field)
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
