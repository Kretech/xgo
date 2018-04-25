package carbon

import (
	"time"

	"github.com/Kretech/common/date"
)

type Carbon struct {
	t time.Time
}

var defaultLoc = time.Local

func In(loc *time.Location) {
	defaultLoc = loc
}

func Now() *Carbon {
	return &Carbon{time.Now()}
}

func UnixOf(sec int64, nsec int64) *Carbon {
	t := time.Unix(sec, nsec)
	return &Carbon{t}
}

func Of(value string) *Carbon {
	t := time.Now().In(defaultLoc)
	return &Carbon{t}
}

func (c *Carbon) Format(format string) string {
	format = date.ToGoFormat(format)
	return c.t.In(defaultLoc).Format(format)
}
