package astutil

import (
	"go/ast"
	"reflect"
)

var (
	OptReturnOnError = true
)

// Find returns a slice of sub-nodes like jQuery
func Find(parent ast.Node, queries []interface{}) (children []ast.Node, err error) {
	if parent == nil || len(queries) == 0 {
		return
	}

	query := queries[0]
	if query == nil {
		return
	}

	queryT := reflect.TypeOf(query)
	if queryT.Kind() == reflect.Ptr && queryT.Elem().Kind() == reflect.Interface {
		queryT = queryT.Elem()
	}
	vlog.Printf("finding <%v> in %v\n", queryT, Name(parent))

	ast.Inspect(parent, func(node ast.Node) bool {
		if node == nil {
			return false
		}
		node.End()
		nodeT := reflect.TypeOf(node)

		vlog.Println(` > with`, nodeT, nodeT == queryT, nodeT.AssignableTo(queryT), len(queries))
		if nodeT == queryT || nodeT.AssignableTo(queryT) {
			if len(queries) == 1 {
				children = append(children, node)

				// if both parent and child is match, only parent-node can return
				return false
			}

			grandChildren, subErr := Find(node, queries[1:])
			if subErr != nil && OptReturnOnError {
				err = subErr
				return false
			}

			children = append(children, grandChildren...)
		}

		return true
	})

	return
}
