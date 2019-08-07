package dump

import (
	_ "unsafe"

	_ "github.com/Kretech/xgo/p"
)

func DepthCompact(depth int, args ...interface{}) (paramNames []string, paramAndValues map[string]interface{})
