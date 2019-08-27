package test

import "testing"

type TestRunner struct {
	*testing.T
}

func TR(t *testing.T) *TestRunner {
	return &TestRunner{
		t,
	}
}

func (this *TestRunner) Add(fn func(t *Assert)) {
	fn(A(this.T))
}
