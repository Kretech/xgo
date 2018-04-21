package word

import (
	"testing"

	"github.com/Kretech/common/test"
)

func TestCamelCase(t *testing.T) {
	test.AssertEqual(t, CamelCase(`a_bc_d`), `aBcD`)
}

func TestUnderlineCase(t *testing.T) {
	test.AssertEqual(t, UnderlineCase(`helloWorld`), `hello_world`)
}
