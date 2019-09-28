package astutil

import (
	"go/ast"
	"testing"
)

func TestExprString(t *testing.T) {
	exprs, err := Find(astFile, []interface{}{new(ast.Expr)})
	if err != nil {
		t.Error(err)
	}

	for _, exp := range exprs {
		t.Log(SrcOf(exp), ExprString(exp.(ast.Expr)))
	}
}
