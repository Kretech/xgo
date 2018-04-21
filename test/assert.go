package test

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
)

//	assert a equals b, or show code where error
func AssertEqual(t *testing.T, resultValue interface{}, expectValue interface{}) {
	value := fmt.Sprint(resultValue)
	expect := fmt.Sprint(expectValue)
	if value != expect {
		file, line := calledBy()
		t.Errorf(
			"Failure in %s:%d\nresult:\t%v;\nexpect:\t%v;\n----\n%s\n",
			file, line,
			value, expect,
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
