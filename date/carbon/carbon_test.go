package carbon

import (
	"testing"
	"time"

	"github.com/Kretech/common/test"
)

var shanghai *time.Location

func init() {
	shanghai, _ = time.LoadLocation("Asia/Shanghai")
}

func TestUnixOf(t *testing.T) {
	In(shanghai)
	test.AssertEqual(t, UnixOf(0, 0).Format("Y-m-d H:i:s"), "1970-01-01 08:00:00")

	In(time.UTC)
	test.AssertEqual(t, UnixOf(0, 0).Format("Y-m-d H:i:s"), "1970-01-01 00:00:00")
}

func TestOf(t *testing.T) {
	Of("2012-1-1 01:02:03")
}

func TestCarbon_Format(t *testing.T) {
	In(shanghai)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-m-d`), `2018-04-22`)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-n-j`), `2018-4-22`)
	test.AssertEqual(t, UnixOf(1523031400, 0).Format(`Y-n-j`), `2018-4-7`)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-m-d H:i:s`), `2018-04-22 14:45:25`)
}
