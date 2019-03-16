package random

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Println(Intn(100))
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intn(6)
	}
}

type T struct {
	// Write func(t *T, p []byte) (n int, err error)
	io.Writer
}

var _ io.Writer = &T{}

type lambdaWriter func(p []byte) (n int, err error)

func (f lambdaWriter) Write(p []byte) (n int, err error) {
	return f(p)
}

func TestQiang(t *testing.T) {

	o := T{}

	// m := reflect.TypeOf(&T{}).Method(0)
	fmt.Println(reflect.TypeOf(&T{}).Method(0))
	fmt.Println(reflect.TypeOf(o).Field(0))
	fmt.Println(reflect.TypeOf(o.Writer))

	o.Writer = lambdaWriter(func(p []byte) (n int, err error) {
		fmt.Println(`hi`)
		return
	})
	o.Write(nil)
}
