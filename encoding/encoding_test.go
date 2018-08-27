package encoding

import (
	"encoding/json"
	"strings"
	"testing"

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
