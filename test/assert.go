package test

import "testing"

type Assert struct {
	T testing.T
}

func A(t *testing.T) *Assert {
	return &Assert{*t}
}

func (a *Assert) Equal(actualValue interface{}, expectValue interface{}) {
	AssertEqual(&a.T, actualValue, expectValue)
}

func (a *Assert) True(actualValue interface{}) {
	a.Equal(actualValue, true)
}
