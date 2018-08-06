package test

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

//	assert a equals b, or show code where error
func AssertEqual(t *testing.T, resultValue interface{}, expectValue interface{}) {
	if isEqual(resultValue, expectValue) {
		file, line := calledBy()
		t.Errorf(
			"Failure in %s:%d\nresult:\t%+v;\nexpect:\t%+v;\n----\n%s\n",
			file, line,
			resultValue, expectValue,
			showFile(file, line),
		)
	}
}

func isEqual(resultValue interface{}, expectValue interface{}) bool {
	switch reflect.TypeOf(expectValue).Kind() {
	case reflect.Map:
		return reflect.DeepEqual(resultValue, expectValue)
	default:
		value := fmt.Sprint(resultValue)
		expect := fmt.Sprint(expectValue)
		return value == expect
	}
}

func calledBy() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
	file = strings.TrimPrefix(file, os.Getenv(`GOPATH`))
	return `$GOPATH` + file, line
}
