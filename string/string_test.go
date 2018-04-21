package string

import (
	"testing"

	"github.com/Kretech/common/test"
)

func TestString_Equal(t *testing.T) {
	test.AssertEqual(t, New(`abc`).Equal([]byte{'a', 'b', 'c'}), true)
}

func TestString_Contains(t *testing.T) {
	test.AssertEqual(t, New(`abc`).Contains(`b`), true)
	test.AssertEqual(t, New(`abc`).Contains(`d`), false)

	test.AssertEqual(t, New([]byte{'a', 'b', 'c'}).Contains(`b`), true)
	test.AssertEqual(t, New([]rune{'李', '二', '狗'}).Contains(`二`), true)
}

func TestString_Trim(t *testing.T) {
	test.AssertEqual(t, New(` abc `).Trim(` `).Equal(`abc`), true)
	test.AssertEqual(t, New(`李二狗`).Trim(`狗`).Equal(`李二`), true)
}

func TestString_Replace(t *testing.T) {
	test.AssertEqual(t, New(`aabc`).Replace(`a`, `b`).Equal(`bbbc`), true)
	test.AssertEqual(t, New(`aabc`).Replace(`a`, `b`, 1).Equal(`babc`), true)
	test.AssertEqual(t, New(`李二狗`).Replace(`二`, `三`).Equal(`李三狗`), true)
}
