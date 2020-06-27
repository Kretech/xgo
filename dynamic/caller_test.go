package dynamic_test

import (
	"testing"

	"github.com/Kretech/xgo/dynamic"
)

func a() string {
	return b()
}

func b() string {
	return dynamic.CallerName(true)
}

func TestCallerName(t *testing.T) {
	name := a()
	t.Log(name)

	name = b()
	t.Log(name)
}
