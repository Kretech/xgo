package p_test

import (
	"testing"

	"github.com/Kretech/xgo/encoding"
	"github.com/Kretech/xgo/p"
)

type _S struct {
}

func (this *_S) a() (string) {
	return `_s.a`
}

func (this *_S) b(t string) (string) {
	return `_s.b(` + t + `)`
}

func TestDump(t *testing.T) {

	a := 1
	b := `sf`
	c := map[string]interface{}{"name": "z", "age": 14}
	d := []interface{}{&c}

	p.Dump(a, b, c, d)

	userId := func() int { return 4 }
	p.Dump(userId())

	p.Dump(userId2())

	_s := _S{}
	p.Dump(_s.a())
	p.Dump(_s.b(`t`))

	p.Dump(encoding.JsonEncode(`abc`))
	p.Dump(encoding.JsonEncode(map[string]interface{}{"a": a}))
}

func userId2() int {
	return 8
}
