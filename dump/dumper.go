package dump

import (
	"os"

	"github.com/Kretech/xgo/dynamic"
)

var (
	// disable dump in global scope
	// use it in production
	Disable = false
)

func Dump(args ...interface{}) {
	d := &CliDumper{
		out:  os.Stdout,
		name: dynamic.Name{X: `dump`, Y: `Dump`},
	}
	d.DepthDump(1, args...)
}
