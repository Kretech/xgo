package src

import "go/ast"

type AstQuery struct {
	node ast.Node
}

func (this *AstQuery) FindNodeType() (nodes []ast.Node) {
	ast.Inspect(this.node, func(node ast.Node) bool {




		return true
	})

	return
}
