package p_test

import (
	"testing"

	. "github.com/Kretech/xgo/p"
	"github.com/Kretech/xgo/test"
)

func TestArgsNameWithInternalAlias(t *testing.T) {
	as := test.A(t)

	a := 3
	b := 4
	a1 := VarName(a, b)
	as.Equal(a1, []string{`a`, `b`})
}

func TestInternalCompact(t *testing.T) {
	as := test.A(t)

	age := 3
	name := `zhang`

	result := Compact(age, name)
	expect := map[string]interface{}{`age`: 3, `name`: `zhang`}

	as.Equal(result, expect)
}
