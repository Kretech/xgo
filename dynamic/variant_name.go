package dynamic

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
	"sync"

	"github.com/Kretech/xgo/astutil"
)

type Name struct {

	// X.Y means how user call this function
	// if user called dynamic.VarName(a), X=dynamic Y=VarName
	// if user called dump.Dump(a) X=dump Y=Dump
	// set this variant used to avoid to scanner import alias
	X string
	Y string
}

// VarName return the variable names
// VarName(a, b) => []string{"a", "b"}
func (n Name) VarName(args ...interface{}) []string {
	return n.VarNameDepth(1, args...)
}

// VarNameDepth return the variable names of args...
// args is a stub for call this method
//
// how can we do this ?
// first, get function info by runtime.Caller
// second, get astFile with ast/parser
// last, match function parameters in ast and get the code
func (n Name) VarNameDepth(skip int, args ...interface{}) (names []string) {
	_, calledFile, calledLine, _ := runtime.Caller(skip + 1)

	shouldCalledExpr := n.X + `.` + n.Y

	// todo 一行多次调用时，还需根据 runtime 找到 column 一起定位
	cacheKey := fmt.Sprintf("%s:%d@%s", calledFile, calledLine, shouldCalledExpr)
	return cacheGet(cacheKey, func() interface{} {

		r := []string{}
		found := false

		fset := token.NewFileSet()
		f, _ := parser.ParseFile(fset, calledFile, nil, 0)

		ast.Inspect(f, func(node ast.Node) (goon bool) {
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

			if fset.Position(call.End()).Line != calledLine {
				return true
			}

			exprStr := astutil.SrcOf(call.Fun)
			if exprStr != shouldCalledExpr {
				return true
			}

			// 拼装每个参数的名字
			for _, arg := range call.Args {
				name := astutil.SrcOf(arg)
				r = append(r, name)
			}

			found = true
			return false
		})

		return r
	}).([]string)
}

func VarName(args ...interface{}) []string {
	name := Name{X: `dynamic`, Y: `VarName`}
	return name.VarNameDepth(1, args...)
}

func VarNameDepth(skip int, args ...interface{}) []string {
	name := Name{X: `dynamic`, Y: `VarNameDepth`}
	return name.VarNameDepth(skip+1, args...)
}

var m sync.Map

func cacheGet(key string, backup func() interface{}) interface{} {
	v, found := m.Load(key)
	if !found {
		v = backup()
		m.Store(key, v)
	}

	return v
}
