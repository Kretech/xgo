package version

import (
	"strings"
)

func escape(s string) string {
	s = strings.TrimLeft(s, "Vv")
	s = strings.Replace(s, "-", ".", -1)

	return s
}
