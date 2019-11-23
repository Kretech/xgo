package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Opt uint8

const (
	OptEscapeHtml Opt = 1 << 1

	OptIndentTab Opt = 1 << 2
)

var (
	logger = log.New(os.Stderr, `xgo.encoding`, log.Ldate|log.Ltime|log.Lshortfile)
)

func (opt Opt) escapeHtml() bool {
	return opt&OptEscapeHtml > 0
}

func (opt Opt) indentTab() bool {
	return opt&OptIndentTab > 0
}

func (opt Opt) Apply(enc *json.Encoder) {

	enc.SetEscapeHTML(opt.escapeHtml())

	if opt.indentTab() {
		enc.SetIndent("", "\t")
	}

}

func JsonEncode(s interface{}, opts ...Opt) string {
	var opt Opt
	if len(opts) > 0 {
		opt = opts[0]
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	opt.Apply(encoder)

	err := encoder.Encode(s)
	if err != nil {
		logger.Output(2, fmt.Sprintf("JsonEncode error: %v", err))
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

func JsonDecode(str interface{}, ele interface{}) error {
	if ss, ok := str.(string); ok {
		str = []byte(ss)
	}
	return json.Unmarshal(str.([]byte), &ele)
}
