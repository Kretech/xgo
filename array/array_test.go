package array

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Kretech/xgo/test"
)

type people struct {
	Id    int
	Name  string
	extra interface{}
}

func TestArray_Demo(t *testing.T) {
	peoples := []*people{{3, "zhang", nil}, {5, "wang", nil}, {4, "li", nil}, {6, "zhao", nil}}
	Values(peoples[0], peoples[1], peoples[2], peoples[3])
}

func TestArray_KeyBy(t *testing.T) {

	peoples := []*people{{3, "zhang", nil}, {5, "wang", nil}, {4, "li", nil}, {6, "zhao", nil}}
	a1 := Values(peoples[0], peoples[1], peoples[2], peoples[3])

	test.AssertEqual(t, a1.String(), `&{3 zhang <nil>} &{5 wang <nil>} &{4 li <nil>} &{6 zhao <nil>}`)

	// d1 := a1.KeyBy("Id")
	v := getStructField(reflect.ValueOf(peoples[0]).Elem(), "Id")
	fmt.Println(v)

	// fmt.Println(d1)
	// fmt.Println(d1.Data())
	// fmt.Println(encoding.JsonEncode(d1.Data()))

	// a2 := Slice(peoples)
	// fmt.Println(a2)

}
