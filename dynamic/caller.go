package dynamic

import (
	"runtime"
	"strings"
)

func CallerName(short bool) string {
	return CallerNameSkip(1, short)
}

func CallerNameSkip(skip int, short bool) string {
	pc, _, _, _ := runtime.Caller(skip + 1)
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	if short {
		idx := strings.LastIndex(name, `.`)
		if idx+1 < len(name) {
			return name[idx+1:]
		}
	}

	return name
}
