package test

import "testing"

type Assert struct {
	T *testing.T
}

func A(t *testing.T) *Assert {
	return &Assert{t}
}

func (a *Assert) Equal(actualValue interface{}, expectValue interface{}) {
	assertEqualSkip(a.T, 1, actualValue, expectValue)
}

func (a *Assert) True(actualValue interface{}) {
	a.Equal(actualValue, true)
}

func (a *Assert) Must(condition interface{}) {
	a.Equal(condition, true)
}
