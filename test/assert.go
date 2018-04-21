package test

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

//	assert a equals b, or show code where error
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if fmt.Sprint(a) != fmt.Sprint(b) {
		file, line := calledBy()
		t.Errorf(
			"Failure in %s:%d\nexpect:\t%v\n   but:\t%v\n----\n%s\n",
			file, line,
			a, b,
			showFile(file, line),
		)
	}
}

func calledBy() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
	file = strings.TrimPrefix(file, os.Getenv(`GOPATH`))
	return `$GOPATH` + file, line
}
