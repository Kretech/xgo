package p

import (
	"fmt"
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

	pc, _, _, _ := runtime.Caller(skip)
	fc := runtime.FuncForPC(pc)
	ss := strings.Split(fc.Name(), "/")

	argNameFunName := ss[len(ss)-1]

	_, file, line, _ := runtime.Caller(skip + 1)

	// todo 一行多次调用时，还需根据 runtime 找到 column 一起定位
	cacheKey := fmt.Sprintf("%s:%d@%s", file, line, argNameFunName)
	return cacheGet(cacheKey, func() interface{} {

		r := []string{}
		found := false

		fset := token.NewFileSet()
		f, _ := parser.ParseFile(fset, file, nil, 0)
		// q.Q(e)
		// q.Q(f.Decls[1].(*ast.FuncDecl).Body.List[1].(*ast.ExprStmt).X.(*ast.CallExpr).Args[0].(*ast.CallExpr).Args[0].(*ast.Ident).Obj)
		ast.Inspect(f, func(node ast.Node) bool {
			if found {
				return false
			}

			if node == nil {
				return false
			}

			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			fn, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			currentName := fn.X.(*ast.Ident).Name + "." + fn.Sel.Name

			if argNameFunName != currentName {
				return true
			}

			// q.Q(argNameFunName, currentName)

			if fset.Position(call.End()).Line != line {
				return true
			}

			for _, arg := range call.Args {
				r = append(r, arg.(*ast.Ident).Name)
			}

			found = true
			return false
		})

		return r
	}).([]string)
}

func Compact(args ...interface{}) (r map[string]interface{}) {
	ps := varNameDepth(1, args...)

	r = make(map[string]interface{}, len(ps))
	for idx, param := range ps {
		r[param] = args[idx]
	}

	return
}

var m = newRWMap()

func cacheGet(key string, backup func() interface{}) interface{} {

	v := m.Get(key)

	if v == nil {
		v = backup()
		m.Set(key, v)
	}

	return v
}
