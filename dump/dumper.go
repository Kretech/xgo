package dump

var (
	// disable dump in global scope
	// use it in production
	Disable = false
)

type Dumper interface {
	Dump(args ...interface{})
	DepthDump(depth int, args ...interface{})
}

var defaultDumper Dumper = NewCliDumper()

func Default() Dumper {
	return defaultDumper
}

func Dump(args ...interface{}) {
	defaultDumper.DepthDump(1, args...)
}
