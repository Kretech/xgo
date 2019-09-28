package astutil

import (
	"go/ast"
	"go/printer"
	"go/token"
	"reflect"
	"strings"
)

//SrcOf print ast.Node to a string
func SrcOf(node ast.Node) string {
	buf := strings.Builder{}
	err := printer.Fprint(&buf, token.NewFileSet(), node)
	if err != nil {
		vlog.Println(`print error `, err)
	}
	return buf.String()
}

func Name(node ast.Node) string {
	return reflect.TypeOf(node).String()
}

func typeNoPtr(v interface{}) (t reflect.Type) {
	if v == nil {
		return
	}

	t = reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return
}
