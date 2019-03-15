package dump_test

import (
	"testing"

	"github.com/Kretech/xgo/dump"
	"github.com/Kretech/xgo/encoding"
)

func TestIsScalar(t *testing.T) {

	mustbe := map[interface{}]bool{
		3:                   true,
		true:                true,
		0.43:                true,
		complex(2, 3):       true,
		"hi":                true,
		struct{}{}:          false,
		new(map[int]string): false,
		new(chan int):       false,
		new([]int):          false,
	}

	for v, b := range mustbe {
		if dump.IsScalar(v) != b {
			t.Error("isScalar failed ", encoding.JsonEncode(v), b)
		}
	}
}
