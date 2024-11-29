package dump

import (
	"strings"
	"testing"
	"time"
)

type (
	Integer int

	String      string
	StringPtr   *string
	StringAlias = string

	car struct {
		Speed int
		Owner interface{}
	}

	Person struct {
		Name      String
		age       int
		Interests []string

		friends [4]*Person

		Cars []*car

		action []func() string
	}
)

func TestSerialize(t *testing.T) {
	time.Local = time.UTC

	type args struct {
		originValue interface{}
	}
	tests := []struct {
		name           string
		originValue    interface{}
		wantSerialized string
	}{
		// TODO: Add test cases.
		{"byte", byte('a'), "a"},
		{"uint8", uint8('a'), "a"},
		{"int", 3, "3"},
		// {"*int", new(int), "*0"},
		{"Integer", Integer(3), "3"},
		{"float", 0.3, "0.3"},
		{"string", "abc", `"abc"`},
		{"*string", ptrString("abc"), `"abc"`},
		{"String", String("abc"), `"abc"`},
		{"[]byte", []byte("abc"), `[]uint8 (len=3) "abc"`},
		{"map", map[string]int{"a": 1, "b": 2, "c": 3}, `map[string]int(len=3){"a"=>1"b"=>2"c"=>3}`},
		{"slice", []int{1, 3, 2}, "[]int(len=3)[0=>1\n1=>3\n2=>2]"},
		{"time", time.Unix(1500000000, 0), `time.Time "2017-07-14 02:40:00 +0000 UTC"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSerialized := Serialize(tt.originValue); !equalNoSpace(gotSerialized, tt.wantSerialized) {
				t.Errorf("Serialize() = %v want %v", gotSerialized, tt.wantSerialized)
				// t.Errorf("Serialize() = %v escape = %v want %v", gotSerialized, escapeSpace(gotSerialized, " \t\n\r"), tt.wantSerialized)
			}
		})
	}
}

func ptrString(s string) *string {
	return &s
}

// 忽略空格比较
func equalNoSpace(a, b string) bool {
	charset := " \t\n\r"
	return escapeSpace(a, charset) == escapeSpace(b, charset)
}

func escapeSpace(a string, charset string) string {
	for _, char := range charset {
		a = strings.Replace(a, string(char), "", -1)
	}
	return a
}
