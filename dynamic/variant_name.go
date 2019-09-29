package dynamic

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/Kretech/xgo/astutil"
)

type Name struct {
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
	pc, calledFile, calledLine, _ := runtime.Caller(skip + 1)
	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	log.Printf("%#v", frame)
	log.Println(calledFile, calledLine, pc)

	// userCalledFunc means how user call this function
	// user can call in these ways:
	// - dynamic.VarName(a)
	// - name.VarName(a)
	// - name.VarName(a)
	userCalledFunc := runtime.FuncForPC(pc)

	// eg: github.com/Kretech/xgo/dynamic.VarName
	userCalledFuncName := userCalledFunc.Name()
	log.Println(userCalledFuncName)

	// shouldCallXXX build statements how users call this
	// users could call this in these ways:
	// - pkg.VarName()
	// - pkgAlias.VarName()
	// - F() // import pkg as .

	// shouldCallSel means the SelectorExpr.X such as pkg, pkgAlias or empty
	shouldCalledSel := userCalledFunc.Name()[:strings.LastIndex(userCalledFuncName, `.`)]

	splitName := strings.Split(userCalledFuncName, "/")
	shouldCalledExpr := splitName[len(splitName)-1]

	// 粗匹配 dump.(*CliDumper).Dump
	// 针对 d:=dumper(); d.Dump() 的情况
	if strings.Contains(shouldCalledExpr, ".(") {
		// 简单的正则来估算是不是套了一层 struct{}
		matched, _ := regexp.MatchString(`\w+\.(.+)\.\w+`, shouldCalledExpr)
		if matched {
			// 暂时不好判断前缀 d 是不是 dumper 类型，先略过
			// 用特殊的 . 前缀表示这个 sel 不处理
			shouldCalledSel = ""
			shouldCalledExpr = shouldCalledExpr[strings.LastIndex(shouldCalledExpr, "."):]
		}
	}

	//fmt.Println("userCalledFunc   =", userCalledFunc.Name())
	//fmt.Println("shouldCalledSel  =", shouldCalledSel)
	//fmt.Println("shouldCalledExpr =", shouldCalledExpr)

	//_, calledFile, calledLine, _ := runtime.Caller(skip + 1)
	//fmt.Printf("%v:%v\n", file, line)

	// todo 一行多次调用时，还需根据 runtime 找到 column 一起定位
	cacheKey := fmt.Sprintf("%s:%d@%s", calledFile, calledLine, shouldCalledExpr)
	return cacheGet(cacheKey, func() interface{} {

		r := []string{}
		found := false

		fset := token.NewFileSet()
		f, _ := parser.ParseFile(fset, calledFile, nil, 0)

		// import alias
		aliasImport := make(map[string]string)
		for _, decl := range f.Decls {
			decl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range decl.Specs {
				is, ok := spec.(*ast.ImportSpec)
				if !ok {
					continue
				}

				if is.Name != nil && strings.Trim(is.Path.Value, `""`) == shouldCalledSel {
					aliasImport[is.Name.Name] = shouldCalledSel
					shouldCalledExpr = is.Name.Name + "." + strings.Split(shouldCalledExpr, ".")[1]

					shouldCalledExpr = strings.TrimLeft(shouldCalledExpr, `.`)
				}
			}
		}

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

			// 检查是不是调用 argsName 的方法
			isArgsNameFunc := func(expr *ast.CallExpr, shouldCallName string) bool {
				exprStr := astutil.ExprString(expr.Fun)
				return exprStr == shouldCallName
			}

			if fset.Position(call.End()).Line != calledLine {
				return true
			}

			if !isArgsNameFunc(call, shouldCalledExpr) {
				return true
			}

			// 拼装每个参数的名字
			for _, arg := range call.Args {
				name := astutil.ExprString(arg)
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
