package astutil

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"reflect"
	"testing"
)

var code = `
package main
// this package is ...

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a+b)
}

`

var astFile, _ = parser.ParseFile(token.NewFileSet(), "", code, parser.ParseComments)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestChildren(t *testing.T) {

	t.Run(`print`, func(t *testing.T) {
		buf := bytes.NewBufferString(``)
		err := printer.Fprint(buf, token.NewFileSet(), astFile)
		t.Log(buf.String(), err)
	})

	t.Run(`queryField`, func(t *testing.T) {
		EnableLog()
		f, err := Find(astFile, []interface{}{new(ast.FuncDecl)})
		t.Log(Name(f[0]), err) // assert len(f)>0

		children, err := Find(astFile, []interface{}{new(ast.FuncDecl), new(ast.AssignStmt)})
		for _, child := range children {
			t.Log(Name(child), SrcOf(child))
		}

		exps, err := Find(astFile, []interface{}{new(ast.Expr)})
		for _, child := range exps {
			t.Log(Name(child), SrcOf(child))
		}
		//t.Log(Name(exp[0]), err)
	})

}

type _a interface {
	Do()
}

type _b struct{}

func (_ *_b) Do() {}

func TestExampleChildren(t *testing.T) {
	t.Log(reflect.TypeOf(new(_b)).AssignableTo(reflect.TypeOf((*_a)(nil)).Elem()))
}
