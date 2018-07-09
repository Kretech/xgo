package char

func IsUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func IsLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func IsAlpha(c rune) bool {
	return IsLower(c) || IsUpper(c)
}

func IsNumber(c rune) bool {
	return c >= 0 && c <= 9
}
