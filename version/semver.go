package version

import (
	"fmt"
	"strconv"
	"strings"
)

const dot = `.`

type SemVer struct {
	Major int
	Minor int
	Patch int
}

//Parse
func Parse(s string) (v *SemVer) {
	v = new(SemVer)

	s = escape(s)
	vs := strings.Split(s, dot)
	if len(vs) > 0 {
		v.Major, _ = strconv.Atoi(vs[0])
	}
	if len(vs) > 1 {
		v.Minor, _ = strconv.Atoi(vs[1])
	}
	if len(vs) > 2 {
		v.Patch, _ = strconv.Atoi(vs[2])
	}

	return
}

func (v *SemVer) String() string {
	return fmt.Sprintf("v%v.%v.%v", v.Major, v.Minor, v.Patch)
}

func (v *SemVer) NumberString() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}

func (v SemVer) NextMajor() *SemVer {
	return &SemVer{Major: v.Major + 1}
}

func (v SemVer) NextMinor() *SemVer {
	return &SemVer{Major: v.Major, Minor: v.Minor + 1}
}

func (v SemVer) NextPatch() *SemVer {
	v.Patch += 1
	return &v
}
