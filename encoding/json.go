package encoding

import "encoding/json"

func JsonEncode(s interface{}) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func JsonDecode(s interface{}, v interface{}) () {
	if ss, ok := s.(string); ok {
		s = []byte(ss)
	}
	json.Unmarshal(s.([]byte), &v)
}
