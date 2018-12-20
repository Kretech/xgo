package p_test

import (
	"testing"

	"github.com/Kretech/xgo/p"
)

func TestDump(t *testing.T) {

	a := 1
	b := `sf`
	c := map[string]interface{}{"name": "z", "age": 14}
	d := []interface{}{&c}
	p.Dump(a, b, c, d)

}
