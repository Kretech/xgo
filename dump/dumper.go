package dump

type Dumper interface {
	Dump(args ...interface{})
	DepthDump(depth int, args ...interface{})
}

var defaultDumper Dumper = NewCliDumper()

func Dump(args ...interface{}) {
	defaultDumper.DepthDump(1, args...)
}
