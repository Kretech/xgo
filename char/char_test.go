package char

import (
	"fmt"
	"testing"
)

func BenchmarkIsAlpha(b *testing.B) {
}

func TestIsAlpha(t *testing.T) {
	fmt.Println(
		IsUpper('A'),
		IsUpper('a'),
		IsUpper('z'),
		IsUpper('Z'),
		IsNumber('0'),
		IsNumber('9'),
		IsNumber(9),
	)
}

func TestIsHan(t *testing.T) {
	fmt.Println(IsHan('汉'))
	fmt.Println(IsHan('a'))

}
