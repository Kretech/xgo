package encoding

import (
	"testing"

	"github.com/Kretech/common/test"
)

func TestBase64EncodeString(t *testing.T) {
	test.AssertEqual(t, Base64EncodeString(`hello`), `aGVsbG8=`)
}

func TestBase64Decode(t *testing.T) {
	test.AssertEqual(t, Base64Decode(`aGVsbG8=`), `hello`)
}

func TestJsonEncode(t *testing.T) {
	m := make(map[string]interface{})
	m[`a`] = 4
	m[`tt`] = `hello`
	test.AssertEqual(t, JsonEncode(m), `{"a":4,"tt":"hello"}`)
}

func TestJsonDecodeMap(t *testing.T) {
	m := make(map[string]interface{})
	m[`a`] = 4
	m[`tt`] = `hello`

	m2 := make(map[string]interface{})
	JsonDecode(`{"a":4,"tt":"hello"}`, &m2)

	test.AssertEqual(t, JsonEncode(m), JsonEncode(m2))
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

	test.AssertEqual(t, JsonEncode(m1), JsonEncode(m2))
}
