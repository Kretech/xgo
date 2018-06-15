package carbon

import (
	"time"

	"github.com/Kretech/common/date"
)

type Carbon struct {
	t time.Time
}

// 假定多数情况下，一个进程给一个location就够了
var defaultLoc = time.Local

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

func (c Carbon) In(loc *time.Location) *Carbon {
	c.t = c.t.In(loc)
	return &c
}

func (c Carbon) Add(d time.Duration) *Carbon {
	c.t = c.t.Add(d)
	return &c
}

func (c *Carbon) SubTime(t time.Time) (time.Duration) {
	return c.Time().Sub(t)
}

func (c Carbon) Sub(d time.Duration) *Carbon {
	c.t = c.t.Add(-d)
	return &c
}

func (c *Carbon) Format(format string) string {
	format = date.ToGoFormat(format)
	return c.t.Format(format)
}
