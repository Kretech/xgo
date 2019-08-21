package dynamic

import (
	"reflect"
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

func (this Person) Say(c string) string {
	return this.Name() + ` : ` + c
}

func TestFunctionSign(t *testing.T) {
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
		{args{Person{}.Say}, FuncHeader{Name: `Say-fm`, In: []*Parameter{{`c`, TString}}, Out: []*Parameter{{``, TString}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.wantF.Name, func(t *testing.T) {
			gotF, err := GetFuncHeader(tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("FunctionSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("FunctionSign() gotF = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func TestGetFuncHeaderExample(t *testing.T) {
	h, _ := GetFuncHeader(Person{}.Say)
	t.Log(h.Name)
	t.Log(h.Doc)
	for _, param := range append(h.In, h.Out...) {
		t.Log(param.Name, param.RType)
	}
}
