package test

import (
	"testing"
)

func TestAssertEqual(t *testing.T) {
	AssertEqual(t, true, true)
	AssertEqual(t, true, 1 == 1)

	AssertEqual(t, 2, 3-1)
	AssertEqual(t, 0, 0)
	AssertEqual(t, int(0), int64(0))

	AssertEqual(t, "hello", "h"+"ello")
}
