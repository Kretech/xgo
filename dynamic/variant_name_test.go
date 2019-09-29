package dynamic_test

import (
	"testing"

	"github.com/Kretech/xgo/dynamic"
)

func TestVarName4Debug(t *testing.T) {
	a := 1
	t.Run(`pkg.VarName`, func(t *testing.T) {
		if v := dynamic.VarName(a); len(v) == 0 || v[0] != `a` {
			t.Error(v)
		}
	})

	t.Run(`pkg.VarName`, func(t *testing.T) {
		if v := dynamic.VarName(a); len(v) == 0 || v[0] != `a` {
			t.Error(v)
		}
	})

	t.Run(`newName.VarName`, func(t *testing.T) {
		name := dynamic.Name{X: `name`}
		if v := name.VarName(a); len(v) == 0 || v[0] != `a` {
			t.Error(v)
		}
	})

}
