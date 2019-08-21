package dynamic

import (
	"go/ast"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/Kretech/xgo/astutil"
	"github.com/pkg/errors"
)

type Parameter struct {
	Name  string
	RType reflect.Type
}

type FuncHeader struct {
	Doc  string // docs above func
	Name string
	In   []*Parameter
	Out  []*Parameter
}

//GetFuncHeader return function header in runtime
func GetFuncHeader(originFunc interface{}) (fh FuncHeader, err error) { //abc
	pc := funcPC(originFunc)
	runtimeFunc := runtime.FuncForPC(pc)
	funcNameFull := runtimeFunc.Name()
	funcName := funcNameFull[strings.LastIndexByte(funcNameFull, '.')+1:]

	fh.Name = funcName

	fileLong, _ := runtimeFunc.FileLine(pc)

	pkgPath := path.Dir(fileLong)
	pkgShort := path.Base(pkgPath)
	fileShort := path.Base(fileLong)

	astPkg, err := astutil.ReadPackageWithName(pkgPath, pkgShort, func(info os.FileInfo) bool {
		return strings.Contains(info.Name(), fileShort)
	})
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	getAstFunc := func(file *ast.File, funcNameFull string) *ast.FuncDecl {
		base := path.Base(funcNameFull)
		base = strings.TrimPrefix(base, file.Name.Name+".")

		for _, d := range file.Decls {
			if fn, ok := d.(*ast.FuncDecl); ok {
				fnName := fn.Name.Name
				if fn.Recv != nil {
					fnName = fn.Recv.List[0].Type.(*ast.Ident).Name + `.` + fnName + `-fm`
				}
				if fnName == base {
					return fn
				}
			}
		}
		return file.Scope.Lookup(funcName).Decl.(*ast.FuncDecl)
	}

	astFunc := getAstFunc(astPkg.Files[fileLong], funcNameFull)
	addDoc(&fh, astFunc)
	addParams(&fh, astFunc)

	T := reflect.TypeOf(originFunc)
	for i, p := range append(fh.In) {
		p.RType = T.In(i)
	}
	for i, p := range append(fh.Out) {
		p.RType = T.Out(i)
	}

	return
}

func addDoc(fh *FuncHeader, astFunc *ast.FuncDecl) {
	if astFunc.Doc == nil {
		return
	}

	for _, c := range astFunc.Doc.List {
		if fh.Doc != `` {
			fh.Doc += "\n"
		}
		fh.Doc += c.Text
	}
}
func addParams(fh *FuncHeader, astFunc *ast.FuncDecl) {

	for _, field := range astFunc.Type.Params.List {
		pa := Parameter{}
		for _, name := range field.Names {
			if pa.Name != `` {
				pa.Name += `,`
			}
			pa.Name += name.Name
		}

		//typeStr := p.Type.(*ast.Ident).Name // string

		fh.In = append(fh.In, &pa)
	}

	if astFunc.Type.Results != nil {
		for _, field := range astFunc.Type.Results.List {
			pa := Parameter{}
			for _, name := range field.Names {
				if pa.Name != `` {
					pa.Name += `,`
				}
				pa.Name += name.Name
			}

			fh.Out = append(fh.Out, &pa)
		}
	}
}

// sys.PtrSize
const PtrSize = 4 << (^uintptr(0) >> 63) // unsafe.Sizeof(uintptr(0)) but an ideal const

// copy from runtime/funcPC
// copy from syscall/funcPC
func funcPC(f interface{}) uintptr {
	return reflect.ValueOf(f).Pointer()
	//return **(**uintptr)(add(unsafe.Pointer(&f), PtrSize))
	//return **(**uintptr)(unsafe.Pointer(&f))
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}
