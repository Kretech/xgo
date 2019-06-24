package astutil

import (
	"fmt"
	"go/ast"
	"strings"
)

//ExprString 获取一个表达式的名字
func ExprString(expr ast.Expr) (name string) {
	switch exp := expr.(type) {

	// 字面值 literal
	case *ast.BasicLit:
		name = exp.Value

		// a.b
	case *ast.SelectorExpr:
		name = ExprString(exp.X) + "." + exp.Sel.Name

	case *ast.CompositeLit:
		name = ExprString(exp.Type) + ExprString(exp.Elts[0])

		if len(exp.Elts) > 0 {
			elts := make([]string, 0, len(exp.Elts))
			for _, elt := range exp.Elts {
				elts = append(elts, ExprString(elt))
			}
			name = `{` + strings.Join(elts, `,`) + `}`
		}

	case *ast.MapType:
		name = fmt.Sprintf("map[%s]%s", ExprString(exp.Key), ExprString(exp.Value))

	//	@todo interface 先都显示 interface{}
	case *ast.InterfaceType:
		name = `interface{}`

	case *ast.KeyValueExpr:
		name = ExprString(exp.Key) + ":" + ExprString(exp.Value)

	case *ast.ArrayType:
		name = "[" + ExprString(exp.Len) + "]" + ExprString(exp.Elt)

	case *ast.StarExpr:
		name = "*" + ExprString(exp.X)

	//	a
	case *ast.Ident:
		name = exp.Name

	case *ast.CallExpr:
		name = ExprString(exp.Fun)

		name += `(`

		if len(exp.Args) > 0 {
			args := make([]string, 0, len(exp.Args))
			for _, arg := range exp.Args {
				args = append(args, ExprString(arg))
			}
			name += strings.Join(args, `,`)
		}

		name += `)`

	//	&a
	case *ast.UnaryExpr:
		name = "&" + ExprString(exp.X)

	//	a["3"]
	case *ast.IndexExpr:
		name = ExprString(exp.X) + "[" + ExprString(exp.Index) + "]"

	case *ast.BinaryExpr:
		name = ExprString(exp.X) + exp.Op.String() + ExprString(exp.Y)

	case nil:
		return ""

	default:
		name = fmt.Sprintf("Unknown(%T)", expr)
	}

	return
}
