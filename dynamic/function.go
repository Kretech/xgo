package dynamic

import (
	"encoding/json"
	"go/ast"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/Kretech/xgo/astutil"
	"github.com/pkg/errors"
)

// Parameter contains input and output info of function
type Parameter struct {
	Name  string
	RType reflect.Type
}

//FuncHeader contains function info
type FuncHeader struct {
	Doc  string // docs above func
	Name string
	In   []*Parameter
	Out  []*Parameter
}

// Equals return two header is equivalent for testing
func (fh *FuncHeader) Equals(other *FuncHeader) bool {
	if !(fh.Name == other.Name && fh.Doc == other.Doc) {
		return false
	}

	if !(len(fh.In) == len(other.In) && len(fh.Out) == len(other.Out)) {
		return false
	}

	a := append(fh.In, fh.Out...)
	b := append(other.In, other.Out...)
	for i := range a {
		if !(a[i].Name == a[i].Name && b[i].RType == b[i].RType) {
			return false
		}
	}

	return true
}

// Encode is convenient for json marshal
func (fh *FuncHeader) Encode() string {
	bytes, _ := json.Marshal(fh)
	return string(bytes)
}

var fhCache sync.Map

//GetFuncHeader return function header in runtime
func GetFuncHeader(originFunc interface{}) (fh FuncHeader, err error) {
	pc := funcPC(originFunc)
	cacheKey := uint(pc)
	value, ok := fhCache.Load(cacheKey)
	if ok {
		fh = value.(FuncHeader)
		return
	}

	fh, err = GetFuncHeaderNoCache(originFunc)
	fhCache.Store(cacheKey, fh)

	return
}

// GetFuncHeaderNoCache is optional way when the cache is incorrect
func GetFuncHeaderNoCache(originFunc interface{}) (fh FuncHeader, err error) { //abc
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

	astFunc := getAstFunc(astPkg.Files[fileLong], funcNameFull)
	if astFunc == nil {
		err = errors.Wrap(err, `unsupport function`)
		return
	}

	fh.addDoc(astFunc)
	fh.addParams(astFunc)

	T := reflect.TypeOf(originFunc)
	for i, p := range append(fh.In) {
		p.RType = T.In(i)
	}
	for i, p := range append(fh.Out) {
		p.RType = T.Out(i)
	}

	return
}

func getAstFunc(file *ast.File, funcNameFull string) *ast.FuncDecl {
	base := path.Base(funcNameFull)
	base = strings.TrimPrefix(base, file.Name.Name+".")
	if !strings.Contains(base, `*`) && strings.HasPrefix(runtime.Version(), `go1.10`) {
		base = strings.Replace(base, "(", "", 1)
		base = strings.Replace(base, ")", "", 1)
	}

	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			fnName := fn.Name.Name
			if fn.Recv != nil {
				recv := fn.Recv.List[0].Type
				recvName := ``
				switch e := recv.(type) {
				case *ast.Ident:
					recvName = e.Name
				case *ast.StarExpr:
					recvName = `(` + `*` + e.X.(*ast.Ident).Name + `)`
				default:
					recvName = astutil.ExprString(recv)
				}
				fnName = recvName + `.` + fnName + `-fm`
			}
			if fnName == base {
				return fn
			}
		}
	}

	return nil
}

func (fh *FuncHeader) addDoc(astFunc *ast.FuncDecl) {
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

func (fh *FuncHeader) addParams(astFunc *ast.FuncDecl) {

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
