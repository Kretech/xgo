package astutil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/pkg/errors"
)

var (
	// 开了缓存，ops 可以提高三个数量级
	OptPackageCache = true
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

var pkgCache sync.Map

//ReadPackageWithName
func ReadPackageWithName(pkgPath string, pkgName string, filter func(os.FileInfo) bool) (pkg *ast.Package, err error) {

	if !OptPackageCache {
		return readPackageWithNameNoCache(pkgPath, pkgName, filter)
	}

	cacheKey := pkgPath + `|` + pkgName + `|F` + funcCacheKey(filter)
	value, ok := pkgCache.Load(cacheKey)
	if ok {
		pkg = value.(*ast.Package)
		return
	}

	pkg, err = readPackageWithNameNoCache(pkgPath, pkgName, filter)

	pkgCache.Store(cacheKey, pkg)

	return
}

func funcCacheKey(filter func(os.FileInfo) bool) (s string) {
	t := **(**uintptr)(unsafe.Pointer(&filter))
	return strconv.FormatUint(uint64(t), 16)
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
		err = errors.Errorf("no packge %s in [%s]", pkgName, strings.Join(keys, `,`))
	}

	return
}
