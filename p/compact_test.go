// package p_test is testing for p
// which p is a php adapter

package p_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Kretech/xgo/p"
	. "github.com/Kretech/xgo/short"
	"github.com/Kretech/xgo/test"
)

func TestAll(t *testing.T) {
	cas := test.TR(t)

	// cas.Add(testArgsName)
	// cas.Add(testConcurrenceArgsName)
	// cas.Add(testCompact)

	// todo 这是一个bug，同一行调用多次 varName 时，无法识别是第几个，现在都认为是第一个
	cas.Add(func(t *test.Assert) {
		a := 3
		b := 4
		c := 4
		fmt.Println(p.VarName(a, b),
			p.VarName(b, a), p.VarName(c))
	})
}

func TestArgsName(t *testing.T) {
	as := test.A(t)

	a := 3
	b := 4
	a1 := p.VarName(a, b)
	as.Equal(a1, []string{`a`, `b`})
}

func TestConcurrenceArgsName(t *testing.T) {
	as := test.A(t)

	a := 3
	c := 4
	wg := sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 1; i < 20; i++ {
				if i%2 == 1 {
					as.Equal(p.VarName(a, c), []string{`a`, `c`})
				} else {
					as.Equal(p.VarName(c), []string{`c`})
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCompact(t *testing.T) {
	as := test.A(t)

	age := 3
	name := `zhang`

	_, result := p.Compact(age, name)
	expect := map[string]Any{`age`: 3, `name`: `zhang`}

	as.Equal(result, expect)
}

func BenchmarkVarName(b *testing.B) {
	a := 3
	for i := 0; i < b.N; i++ {
		p.VarName(a)
	}
}
