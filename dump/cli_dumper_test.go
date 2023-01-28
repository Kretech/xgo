package dump

import (
	"os"
)

func ExampleNewCliDumper() {
	a := 1
	b := `2`

	ShowFileLine1 = false
	DefaultWriter = os.Stdout

	g1 := NewCliDumper(`g1`)
	g1.Dump(a, b)

	g2 := MyPackage{MyDump: MyDump}
	g2.MyDump(a, b)
	// Output:
	// a => 1
	// b => "2"
	// a => 1
	// b => "2"
}

// MyPackage mock the package name. so we can called it by `x.y()`
type MyPackage struct {
	MyDump func(...interface{})
}

func MyDump(args ...interface{}) {
	_g := NewCliDumper(`g2`) // means package name
	_g.DepthDump(1, args...)
}
