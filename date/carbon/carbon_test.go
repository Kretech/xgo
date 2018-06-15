package carbon

import (
	"testing"
	"time"

	"github.com/Kretech/common/test"
)

func TestUnixOf(t *testing.T) {
	In(Shanghai)
	test.AssertEqual(t, UnixOf(0, 0).Format("Y-m-d H:i:s"), "1970-01-01 08:00:00")

	In(time.UTC)
	test.AssertEqual(t, UnixOf(0, 0).In(Shanghai).Format("Y-m-d H:i:s"), "1970-01-01 08:00:00")
	test.AssertEqual(t, UnixOf(0, 0).Format("Y-m-d H:i:s"), "1970-01-01 00:00:00")

	In(time.UTC)
	test.AssertEqual(t, UnixOf(0, 0).Time().String(), "1970-01-01 00:00:00 +0000 UTC")
	test.AssertEqual(t, UnixOf(0, 0).Time().Unix(), 0)
	test.AssertEqual(t, UnixOf(0, 0).Format("Y-m-d H:i:s"), "1970-01-01 00:00:00")
	test.AssertEqual(t, UnixOf(0, 0).In(Shanghai).Time(), "1970-01-01 08:00:00 +0800 CST")
	test.AssertEqual(t, UnixOf(0, 0).In(Shanghai).Time().Unix(), 0)
	test.AssertEqual(t, UnixOf(0, 0).In(Shanghai).Format("Y-m-d H:i:s"), "1970-01-01 08:00:00")
}

func TestParse(t *testing.T) {
	Parse("2012-1-1 01:02:03")
	Parse("+1 day")
}

func TestCarbon_Format(t *testing.T) {
	In(Shanghai)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-m-d`), `2018-04-22`)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-n-j`), `2018-4-22`)
	test.AssertEqual(t, UnixOf(1523031400, 0).Format(`Y-n-j`), `2018-4-7`)
	test.AssertEqual(t, UnixOf(1524379525, 0).Format(`Y-m-d H:i:s`), `2018-04-22 14:45:25`)
}

func TestCarbon_In(t *testing.T) {
	t1 := Now()
	test.AssertEqual(t, t1.Time().Location(), time.Local)

	t2 := t1.In(Shanghai)
	// t1未修改
	test.AssertEqual(t, t1.Time().Location(), time.Local)
	// t2修改成功
	test.AssertEqual(t, t2.Time().Location(), Shanghai)
}

func TestCarbon_Sub(t *testing.T) {
	t1 := TParse("Y-m-d H:i:s", "2018-01-02 09:00:00")
	t2 := t1.Sub(time.Hour)
	test.AssertEqual(t, t2.Format("Y-m-d H:i:s"), "2018-01-02 08:00:00")
}
