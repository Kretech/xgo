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

	t.Run(`base`, func(t *testing.T) {
		dump.OptShowUint8sAsString = true
		dump.OptShowUint8AsChar = true
		aInt := 1
		bStr := `sf`
		cMap := map[string]interface{}{"name": "z", "age": 14, "bytes": []byte(string(`abc`)), `aByte`: byte('c')}
		dArray := []interface{}{&cMap, aInt, bStr}
		//dump.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
		c.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
	})

	t.Run(`operator`, func(t *testing.T) {
		a := 0.1
		b := 0.2
		cc := 0.3
		c.Dump(a + b)
		c.Dump(a+b == cc)
		c.Dump(a+b > cc)
		c.Dump(a+b < cc)
	})

	t.Run(`interface`, func(t *testing.T) {
		var err error
		var emptyInterface interface{}
		var emptyMap map[string]interface{}
		var emptySlice []interface{}
		c.Dump(err, emptyInterface, emptyMap, emptySlice, nil)
	})

	t.Run(`struct`, func(t *testing.T) {
		dump.ShowFileLine1 = true
		defer func() {
			dump.ShowFileLine1 = false
		}()

		type String string

		type car struct {
			Speed int
			Owner interface{}
		}

		type Person struct {
			Name      String
			age       int
			Interests []string

			friends [4]*Person

			Cars []*car

			action []func() string
		}

		p1 := Person{Name: "lisan", age: 19, Interests: []string{"a", "b"}, Cars: []*car{{Speed: 120}},
			action: []func() string{
				func() string { return "sayHi" },
				func() string { return "sayBay" },
			}}
		p2 := &p1

		c.Dump(p2)

		type Person2 Person
		p3 := Person2(p1)
		c.Dump(p3)

		type Person3 = Person2
		p4 := Person3(p1)
		c.Dump(p4)

		type Person4 struct {
			Person Person
			Person2
		}
		p5 := Person4{p1, p4}
		c.Dump(p5)
	})

	userId := func() int { return 4 }
	c.Dump(userId())

	c.Dump(userId2())

	_s := _S{}
	c.Dump(_s.a())
	c.Dump(_s.b(`t`))

	num3 := 3
	c.Dump(`abc`)
	c.Dump(encoding.JsonEncode(`abc`))
	c.Dump(encoding.JsonEncode(map[string]interface{}{"a": num3}))

	ch := make(chan bool)
	c.Dump(ch)

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
