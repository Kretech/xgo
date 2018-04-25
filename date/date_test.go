package date

import (
	"testing"

	"github.com/Kretech/common/test"
)

func TestDate(t *testing.T) {
	test.AssertEqual(t, Format(`Y-m-d`, 1524379525), `2018-04-22`)
	test.AssertEqual(t, Format(`Y-n-j`, 1524379525), `2018-4-22`)
	test.AssertEqual(t, Format(`Y-n-j`, 1523031400), `2018-4-7`)
	test.AssertEqual(t, Format(`Y-m-d H:i:s`, 1524379525), `2018-04-22 14:45:25`)
}

func TestStrToTime(t *testing.T) {

}
