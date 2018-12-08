// package p_test is testing for p
// which p is a php adapter

package p_test

import (
	"testing"

	"github.com/Kretech/xgo/p"
	. "github.com/Kretech/xgo/short"
	"github.com/Kretech/xgo/test"
)

func TestAll(t *testing.T) {
	cas := test.TR(t)

	cas.Add(testArgsName)
	cas.Add(testCompact)

	// cas.Add(func(t *test.Assert) {
	// 	q.Q(runtime.NumGoroutine(), p.GoID())
	// })
}

func testArgsName(as *test.Assert) {
	a := 3
	b := 4
	a1 := p.VarName(a, b)
	as.Equal(a1, []string{`a`, `b`})
}

func testCompact(as *test.Assert) {
	age := 3
	name := `zhang`

	result := p.Compact(age, name)
	expect := map[string]Any{`age`: 3, `name`: `zhang`}

	as.Equal(result, expect)
}

func BenchmarkVarName(b *testing.B) {
	a := 3
	for i := 0; i < b.N; i++ {
		p.VarName(a)
	}
}
