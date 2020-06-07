package dump

import (
	_ "unsafe"

	"github.com/Kretech/xgo/p"
)

// use this to init function in p package
var _ = p.DepthCompact

func DepthCompact(depth int, args ...interface{}) (paramNames []string, paramAndValues map[string]interface{})

func Compact(args ...interface{}) (paramNames []string, paramAndValues map[string]interface{}) {
	return DepthCompact(1, args...)
}
