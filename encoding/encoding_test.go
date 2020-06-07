package encoding

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Kretech/xgo/test"
)

func TestBase64EncodeString(t *testing.T) {
	test.AssertEqual(t, Base64EncodeString(`hello`), `aGVsbG8=`)
}

func TestBase64Decode(t *testing.T) {
	test.AssertEqual(t, Base64Decode(`aGVsbG8=`), `hello`)
}

func TestJsonEncodeMap(t *testing.T) {
	m := make(map[string]interface{})
	m[`a`] = 4.0
	m[`tt`] = `hello`
	b, _ := json.Marshal(m)

	expect := make(map[string]interface{})
	JsonDecode(b, &expect)
	test.AssertEqual(t, m, expect)

	m[`escape`] = `&`
	test.AssertTrue(t, strings.Contains(JsonEncode(m), `"escape":"&"`))
	test.AssertTrue(t, strings.Contains(JsonEncode(m, OptEscapeHtml), `"escape":"\u0026"`))
}

func TestJsonDecodeObject(t *testing.T) {

	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	m1 := User{
		Id:   4,
		Name: "hi",
	}

	m2 := User{}
	JsonDecode(`{"id":4,"name":"hi"}`, &m2)

	test.AssertEqual(t, m1, m2)
}

func TestJsonMarshalTime(t *testing.T) {
	m := map[string]interface{}{
		"t": time.Date(2001, 1, 1, 1, 1, 0, 0, time.Local),
		"X": time.Date(1002, 1, 1, 1, 1, 0, 0, time.Local),
	}
	s := JsonEncode(m)
	fmt.Println(s)

	type A struct {
		X time.Time
	}

	m2 := A{}
	err := JsonDecode(s, &m2)
	fmt.Println(m2, err)
}
