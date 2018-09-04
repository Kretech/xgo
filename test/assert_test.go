package test

import "testing"

func TestAssert_Equal(t *testing.T) {
	a := A(t)

	a.Equal(true, true)
	a.Equal(true, 1 == 1)

	a.Equal(2, 3-1)
	a.Equal(0, 0)
	a.Equal(int(0), int64(0))

	a.Equal("hello", "h"+"ello")
}
