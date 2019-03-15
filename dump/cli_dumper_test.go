package dump_test

import (
	"bytes"
	"testing"

	"github.com/Kretech/xgo/dump"
	"github.com/Kretech/xgo/encoding"
)

type _S struct {
}

func (this *_S) a() string {
	return `_s.a`
}

func (this *_S) b(t string) string {
	return `_s.b(` + t + `)`
}

// go test -v -vet off -test.run ^TestCliDumper_Dump ./dump/...
func TestCliDumper_Dump(t *testing.T) {
	c := dump.NewCliDumper()

	aInt := 1
	bStr := `sf`
	cMap := map[string]interface{}{"name": "z", "age": 14}
	dArray := []interface{}{&cMap, aInt, bStr}
	//dump.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
	c.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])

	userId := func() int { return 4 }
	c.Dump(userId())

	c.Dump(userId2())

	_s := _S{}
	c.Dump(_s.a())
	c.Dump(_s.b(`t`))

	c.Dump(encoding.JsonEncode(`abc`))
	c.Dump(encoding.JsonEncode(map[string]interface{}{"a": aInt}))

	buf := &bytes.Buffer{}
	c = dump.NewCliDumper(dump.OptOut(buf))
	c.Dump(1)

	if buf.Len() < 1 {
		t.Error("dump failed", buf.Len(), buf.String())
	}
}

func userId2() int {
	return 8
}
