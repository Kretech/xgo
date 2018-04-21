package encoding

import "encoding/base64"

func Base64Decode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func Base64EncodeString(s string) string {
	return Base64Encode([]byte(s))
}

func Base64Encode(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}
