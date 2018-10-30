package test

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

var (
	// Alias
	BeTrue  = AssertTrue
	BeNil   = AssertNil
	BeEqual = AssertEqual
)

func AssertTrue(t *testing.T, resultValue interface{}) {
	AssertEqual(t, resultValue, true)
}

func AssertNil(t *testing.T, resultValue interface{}) {
	AssertEqual(t, resultValue, nil)
}

// 	assert a equals b, or show code where error
func AssertEqual(t *testing.T, resultValue interface{}, expectValue interface{}) {

	if isEqual(resultValue, expectValue) {
		return
	}

	resultValue = fmt.Sprintf("%+v", resultValue)
	expectValue = fmt.Sprintf("%+v", expectValue)

	file, line := calledBy()
	t.Errorf(
		"Failure in %s:%d\nresult:(%d)\t%+v\nexpect:(%d)\t%+v\n----\n%s\n",
		file, line,
		len(resultValue.(string)), resultValue,
		len(expectValue.(string)), expectValue,
		showFile(file, line),
	)
}

func isEqual(actualValue interface{}, expectValue interface{}) bool {
	if actualValue == nil || expectValue == nil {
		return actualValue == expectValue
	}

	switch reflect.TypeOf(expectValue).Kind() {

	case reflect.Map, reflect.Struct, reflect.Slice, reflect.Array:
		return reflect.DeepEqual(actualValue, expectValue)

	default:
		actual := fmt.Sprintf("%v", actualValue)
		expect := fmt.Sprintf("%v", expectValue)
		return actual == expect
	}
}

func calledBy() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
	file = strings.TrimPrefix(file, os.Getenv(`GOPATH`))
	return `$GOPATH` + file, line
}
