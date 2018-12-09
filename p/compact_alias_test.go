package p_test

import (
	"testing"

	pp "github.com/Kretech/xgo/p"
	"github.com/Kretech/xgo/test"
)

func TestArgsNameWithAlias(t *testing.T) {
	as := test.A(t)

	a := 3
	b := 4
	a1 := pp.VarName(a, b)
	as.Equal(a1, []string{`a`, `b`})
}
