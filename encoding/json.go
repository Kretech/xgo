package encoding

import (
	"bytes"
	"encoding/json"

	"github.com/lunny/log"
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
		log.Errorf("xgo.encoding.JsonEncode error: %v", err)
	}

	return string(buffer.Bytes())
}

func JsonDecode(str interface{}, ele interface{}) () {
	if ss, ok := str.(string); ok {
		str = []byte(ss)
	}
	json.Unmarshal(str.([]byte), &ele)
}
