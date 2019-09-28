package astutil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

var (
	// 开了缓存，ops 可以提高三个数量级
	OptPackageCache = true
)

// ReadPackage is simple wrapper for parserDir
// use pkgName with last segment in pkgPath
func ReadPackage(pkgPath string) (pkg *ast.Package, err error) {
	short := ``
	for i := len(pkgPath) - 1; i >= 0; i-- {
		if pkgPath[i] == '/' {
			short = pkgPath[:i]
			break
		}
	}
	return ReadPackageWithName(pkgPath, short, `.`, func(info os.FileInfo) bool {
		return true
	})
}

var pkgCache sync.Map

// ReadPackageWithName read package with specified package name
// fileScope used for cache key
func ReadPackageWithName(pkgPath string, pkgName string, fileScope string, filter func(os.FileInfo) bool) (pkg *ast.Package, err error) {

	if !OptPackageCache {
		return readPackageWithNameNoCache(pkgPath, pkgName, filter)
	}

	cacheKey := pkgPath + `|` + pkgName + `|` + fileScope
	value, ok := pkgCache.Load(cacheKey)
	if ok {
		pkg = value.(*ast.Package)
		return
	}

	pkg, err = readPackageWithNameNoCache(pkgPath, pkgName, filter)

	pkgCache.Store(cacheKey, pkg)

	return
}

func readPackageWithNameNoCache(pkgPath string, pkgName string, filter func(os.FileInfo) bool) (pkg *ast.Package, err error) {
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
		err = errors.Errorf("no package %s in [%s]", pkgName, strings.Join(keys, `,`))
	}

	return
}
