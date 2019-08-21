package dynamic

import (
	"log"
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
)

var (
	TString = reflect.TypeOf(``)
)

func tEmptyFunc()                 {}
func tOnlyInFunc(s string)        {}
func tFunc(s1 string) (s2 string) { return }

// Person ...
type Person struct{}

// comment
func (this Person) Name() string {
	return `noname`
}

// Say can say something
func (this Person) Say(c string) string {
	return this.Name() + ` : ` + c
}

func TestGetFuncHeader(t *testing.T) {
	type args struct {
		fn interface{}
	}
	tests := []struct {
		args    args
		wantF   FuncHeader
		wantErr bool
	}{
		// TODO: Add test cases.
		{args{tEmptyFunc}, FuncHeader{Name: `tEmptyFunc`}, false},
		{args{tOnlyInFunc}, FuncHeader{Name: `tOnlyInFunc`, In: []*Parameter{{`s`, TString}}}, false},
		{args{tFunc}, FuncHeader{Name: `tFunc`, In: []*Parameter{{`s1`, TString}}, Out: []*Parameter{{`s2`, TString}}}, false},
		{args{Person{}.Name}, FuncHeader{Name: `Name-fm`, Doc: `// comment`, Out: []*Parameter{{``, TString}}}, false},
		{args{Person{}.Say}, FuncHeader{Name: `Say-fm`, Doc: `// Say can say something`, In: []*Parameter{{`c`, TString}}, Out: []*Parameter{{``, TString}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.wantF.Name, func(t *testing.T) {
			gotF, err := GetFuncHeader(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("FunctionSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !gotF.Equals(&tt.wantF) {
				t.Errorf("FunctionSign() gotF = %v, want %v", gotF.Encode(), tt.wantF.Encode())
			}
		})
	}
}

func BenchmarkGetFuncHeader_i0_o0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFuncHeader(tEmptyFunc)
	}
}

func BenchmarkGetFuncHeader_i1_o1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFuncHeader(tFunc)
	}
}

func BenchmarkGetFuncHeaderNoCache_i1_o1(b *testing.B) {
	debug := func() bool { return false }
	if debug() {
		f, _ := os.Create("dynamic.pprof")
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < b.N; i++ {
		GetFuncHeader(tFunc)
		//GetFuncHeaderNoCache(tFunc)
	}
}

func TestGetFuncHeaderExample(t *testing.T) {
	h, _ := GetFuncHeader(Person{}.Say)
	//h, _ = GetFuncHeader(Person{}.Say)
	t.Log(h.Name)
	t.Log(h.Doc)
	for _, param := range append(h.In, h.Out...) {
		t.Log(param.Name, param.RType)
	}
}

func ExampleGetFuncHeader() {

	// // Person ...
	// type Person struct{}
	//
	// // comment
	// func (this Person) Name() string {
	// 	return `noname`
	// }
	//
	// // Say can say something
	// func (this Person) Say(c string) string {
	// 	return this.Name() + ` : ` + c
	// }

	h, _ := GetFuncHeader(Person{}.Say)
	log.Println(h.Name)
	//: Say-fm

	log.Println(h.Doc)
	//: // Say can say something

	for _, param := range append(h.In, h.Out...) {
		log.Println(param.Name, param.RType)
		//: c string
		//:  string
	}
}
