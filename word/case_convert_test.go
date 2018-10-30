package word

import (
	"testing"

	"github.com/Kretech/xgo/test"
)

func TestCamelCase(t *testing.T) {
	test.AssertEqual(t, CamelCase(`a_bc_d`), `aBcD`)
}

func TestUnderlineCase(t *testing.T) {
	test.AssertEqual(t, UnderlineCase(`helloWorld`), `hello_world`)
}

func TestUpperFirst(t *testing.T) {
	test.BeEqual(t, UpperFirst(`a`), `A`)
	test.BeEqual(t, UpperFirst(`ab`), `Ab`)
	test.BeEqual(t, UpperFirst(`_a`), `_a`)

	test.BeEqual(t, LowerFirst(`_a`), `_a`)
	test.BeEqual(t, LowerFirst(`A`), `a`)
	test.BeEqual(t, LowerFirst(`Ac`), `ac`)
	test.BeEqual(t, LowerFirst(`ac`), `ac`)
}
