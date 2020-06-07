package dynamic_test

import (
	"testing"

	"github.com/Kretech/xgo/dynamic"
)

func TestVarName4Debug(t *testing.T) {
	a := 1
	t.Run(`pkg.VarName`, func(t *testing.T) {
		v := dynamic.VarName(a)
		if len(v) == 0 || v[0] != `a` {
			t.Error(v)
		}
	})

	t.Run(`newName.VarName`, func(t *testing.T) {
		name := dynamic.Name{X: `name`, Y: `VarName`}
		v := name.VarName(a)
		if len(v) == 0 || v[0] != `a` {
			t.Error(v)
		}
	})

}
