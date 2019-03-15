package word

import (
	"bytes"
	"strings"
)

// CamelCase convert `a_b_c` to `aBC`
func CamelCase(v string) string {
	buf := bytes.NewBuffer([]byte{})
	length := len(v)
	for i := 0; i < length; i++ {
		if v[i] != '_' {
			buf.WriteByte(v[i])
		} else {
			i++
			if i > length {
				continue
			}
			if v[i] >= 'a' && v[i] <= 'z' {
				buf.WriteByte(v[i] + 'A' - 'a')
			} else {
				buf.WriteByte(v[i])
			}
		}
	}

	return buf.String()
}

// UnderlineCase convert `ABC` to `a_b_c`
func UnderlineCase(v string) string {
	buf := bytes.NewBuffer([]byte{})
	length := len(v)
	for i := 0; i < length; i++ {
		if i > 0 && v[i] >= 'A' && v[i] <= 'Z' {
			buf.WriteByte('_')
			buf.WriteByte(v[i] + 'a' - 'A')
		} else {
			buf.WriteByte(v[i])
		}
	}
	return buf.String()
}

func UpperFirst(v string) string {
	return strings.ToUpper(string(v[0])) + v[1:]
}

func LowerFirst(v string) string {
	return strings.ToLower(string(v[0])) + v[1:]
}
