package dump_test

import (
	"testing"

	"github.com/Kretech/xgo/dump"
	"github.com/Kretech/xgo/test"
)

func TestDepthCompact(t *testing.T) {
	a := 3
	b := `B`
	t.Run(`DepthCompact`, func(t *testing.T) {
		keys, kvs := dump.DepthCompact(0, a, b)
		t.Log(keys, kvs)
		test.AssertEqual(t, kvs[`a`], 3)
	})

	t.Run(`Compact`, func(t *testing.T) {
		keys, kvs := dump.Compact(a, b)
		t.Log(keys, kvs)
		test.AssertEqual(t, kvs[`a`], 3)
	})
}

func TestCompact(t *testing.T) {
	a := 3
	b := `B`
	keys, kvs := dump.Compact(a, b)
	dump.Dump(keys, kvs)
}
