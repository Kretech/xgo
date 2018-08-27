package encoding

import (
	"bytes"
	"encoding/json"
	"log"
)

const (
	OptEscapeHtml = 1 << 1
)

func JsonEncode(s interface{}, opts ...int) string {

	opt := 0
	if len(opts) > 0 {
		opt = opts[0]
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	encoder.SetEscapeHTML(opt&OptEscapeHtml > 0)

	err := encoder.Encode(s)
	if err != nil {
		log.Printf("xgo.encoding.JsonEncode error: %v", err)
	}

	b := buffer.Bytes()

	if len(b) > 0 && b[len(b)-1] == '\n' {
		// 去掉 go std 给加的 \n
		// 正常的 json 末尾是不会有 \n 的
		// @see https://github.com/golang/go/issues/7767
		b = b[:len(b)-1]
	}

	return string(b)
}

func JsonDecode(str interface{}, ele interface{}) {
	if ss, ok := str.(string); ok {
		str = []byte(ss)
	}
	json.Unmarshal(str.([]byte), &ele)
}
