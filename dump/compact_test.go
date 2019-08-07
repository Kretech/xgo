package dump_test

import (
	"fmt"
	"testing"

	"github.com/Kretech/xgo/dump"
	"github.com/Kretech/xgo/test"
)

func TestDepthCompact(t *testing.T) {
	a := 3
	keys, kvs := dump.DepthCompact(0, a)
	fmt.Println(keys, kvs)

	test.AssertEqual(t, kvs[`a`], 3)
}
