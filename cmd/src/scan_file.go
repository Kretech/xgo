package src

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ScanFile(filename string, fn func(node ast.Node) bool) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, filename, nil, 0)

	ast.Inspect(f, fn)
}

func ScanFileDecls(filename string, fn func(decl ast.Decl) bool) {
	ScanFile(filename, func(node ast.Node) bool {
		decls := node.(*ast.File).Decls
		for _, decl := range decls {
			fn(decl)
		}
		return false
	})
}
