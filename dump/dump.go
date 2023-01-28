package dump

import (
	"github.com/Kretech/xgo/dynamic"
)

// disable dump in global scope
// use it in production
var Disable = false

func Dump(args ...interface{}) {
	d := &CliDumper{
		name: dynamic.Name{X: `dump`, Y: dynamic.CallerName(true)},
	}
	d.DepthDump(1, args...)
}
