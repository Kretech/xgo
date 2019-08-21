package astutil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/pkg/errors"
)

//
// use pkgName with last segment in pkgPath
func ReadPackage(pkgPath string) (pkg *ast.Package, err error) {
	short := ``
	for i := len(pkgPath) - 1; i >= 0; i-- {
		if pkgPath[i] == '/' {
			short = pkgPath[:i]
			break
		}
	}
	return ReadPackageWithName(pkgPath, short, func(info os.FileInfo) bool {
		return true
	})
}

//ReadPackageWithName
func ReadPackageWithName(pkgPath string, pkgName string, filter func(os.FileInfo) bool) (pkg *ast.Package, err error) {
	pkgs, err := parser.ParseDir(token.NewFileSet(), pkgPath, filter, parser.ParseComments)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	pkg, ok := pkgs[pkgName]
	if !ok {
		keys := []string{}
		for name := range pkgs {
			keys = append(keys, name)
		}
		err = errors.Errorf("no packge %s in [%s]", pkgName, strings.Join(keys, `,`))
	}

	return
}
