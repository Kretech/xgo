package p

import (
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
	"strings"
	"sync"
)

var compactInitOnce sync.Once

func VarName(args ...interface{}) []string {
	return varNameDepth(1, args...)
}

func varNameDepth(skip int, args ...interface{}) (c []string) {
	var argNameFunName = ``
	pc, _, _, _ := runtime.Caller(skip)
	fc := runtime.FuncForPC(pc)
	ss := strings.Split(fc.Name(), "/")

	argNameFunName = ss[len(ss)-1]

	_, file, line, _ := runtime.Caller(skip + 1)

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, file, nil, 0)
	// q.Q(e)
	// q.Q(f.Decls[1].(*ast.FuncDecl).Body.List[1].(*ast.ExprStmt).X.(*ast.CallExpr).Args[0].(*ast.CallExpr).Args[0].(*ast.Ident).Obj)
	ast.Inspect(f, func(node ast.Node) bool {
		if node == nil {
			return false
		}

		call, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		// q.Q(call)
		fn := call.Fun.(*ast.SelectorExpr)
		currentName := fn.X.(*ast.Ident).Name + "." + fn.Sel.Name

		if argNameFunName != currentName {
			return true
		}

		// q.Q(argNameFunName, currentName)

		if fset.Position(call.End()).Line != line {
			return true
		}

		for _, arg := range call.Args {
			c = append(c, arg.(*ast.Ident).Name)
		}

		return false
	})

	return
}

func Compact(args ...interface{}) (r map[string]interface{}) {
	ps := varNameDepth(1, args...)

	r = make(map[string]interface{}, len(ps))
	for idx, param := range ps {
		r[param] = args[idx]
	}

	return
}
