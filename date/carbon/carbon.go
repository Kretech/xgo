package carbon

import (
	"time"

	"github.com/Kretech/xgo/date"
)

// Carbon
type Carbon struct {
	t time.Time
}

// 假定多数情况下，一个进程给一个location就够了
var defaultLoc = time.Local

// 设置全局时区
func In(loc *time.Location) {
	defaultLoc = loc
}

func Now() *Carbon {
	return &Carbon{time.Now()}
}

func UnixOf(sec int64, nsec int64) *Carbon {
	t := time.Unix(sec, nsec).In(defaultLoc)
	return &Carbon{t}
}

func TimeOf(t time.Time) *Carbon {
	return &Carbon{t}
}

// 解析时间文本
func StrOf(str string) *Carbon {
	return Parse(str)
}

// todo
// alias to StrOf
func Parse(str string) *Carbon {
	t := time.Now().In(defaultLoc)
	return &Carbon{t}
}

// wrap of time.Parse but with the carbon layout
// 使用 Go time.Time 解析时间文本
func TParse(layout, value string) *Carbon {
	t, _ := time.Parse(date.ToGoFormat(layout), value)
	return TimeOf(t)
}

func carbonOf(c *Carbon) *Carbon {
	return TimeOf(c.t)
}

func (c *Carbon) Clone() *Carbon {
	return carbonOf(c)
}

func (c *Carbon) Time() time.Time {
	return c.t
}

// 设置时区
func (c Carbon) In(loc *time.Location) *Carbon {
	c.t = c.t.In(loc)
	return &c
}

func (c Carbon) Add(d time.Duration) *Carbon {
	c.t = c.t.Add(d)
	return &c
}

// 减去另一个时间：计算两个时间的差
func (c *Carbon) SubTime(t time.Time) time.Duration {
	return c.Time().Sub(t)
}

// 减去一段时间
func (c Carbon) Sub(d time.Duration) *Carbon {
	c.t = c.t.Add(-d)
	return &c
}

// Format 格式化代码
// c.Format("Y-m-d H:i:s")
// @see http://php.net/manual/zh/function.date.php#refsect1-function.date-parameters
func (c *Carbon) Format(format string) string {
	format = date.ToGoFormat(format)
	return c.t.Format(format)
}
