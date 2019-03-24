package main

import (
	"go/ast"
	"io"

	"github.com/Kretech/xgo/cmd/src"
)

var InterfaceHandler handler = func(in []string, out io.Writer) (err error) {
	filename := in[0]

	//classes :=

	src.ScanFileDecls(filename, func(decl ast.Decl) bool {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok {
			return false
		}

		println(
			`name`,
			fd.Name.Name,
			fd.Name.Obj,
		)

		return false
	})

	return
}
