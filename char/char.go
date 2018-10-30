package char

import "regexp"

func IsUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsAlpha(c byte) bool {
	return IsLower(c) || IsUpper(c)
}

func IsNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsHan(b rune) bool {
	return IsHanString(string(b))
}

func IsHanString(s string) bool {
	re := regexp.MustCompile(`[\p{Han}]+`)
	return re.MatchString(s)
}
