package dump_test

import (
	"testing"

	"github.com/Kretech/xgo/dump"
	"github.com/Kretech/xgo/encoding"
)

func TestDump_Example(t *testing.T) {
	t.Run(`base`, func(t *testing.T) {
		dump.OptShowUint8sAsString = true
		dump.OptShowUint8AsChar = true
		aInt := 1
		bStr := `sf`
		cMap := map[string]interface{}{"name": "z", "age": 14, "bytes": []byte(string(`abc`)), `aByte`: byte('c')}
		dArray := []interface{}{&cMap, aInt, bStr}
		dump.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
	})

	t.Run(`operator`, func(t *testing.T) {
		a := 0.1
		b := 0.2
		cc := 0.3
		dump.Dump(a + b)
		dump.Dump(a+b == cc)
		dump.Dump(a+b > cc)
		dump.Dump(a+b < cc)
	})

	t.Run(`interface`, func(t *testing.T) {
		var err error
		var emptyInterface interface{}
		var emptyMap map[string]interface{}
		var emptySlice []interface{}
		dump.Dump(err, emptyInterface, emptyMap, emptySlice, nil)
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

		p1 := Person{
			Name: "lisan", age: 19, Interests: []string{"a", "b"}, Cars: []*car{{Speed: 120}},
			action: []func() string{
				func() string { return "sayHi" },
				func() string { return "sayBay" },
			},
		}
		p2 := &p1

		dump.Dump(p2)

		type Person2 Person
		p3 := Person2(p1)
		dump.Dump(p3)

		type Person3 = Person2
		p4 := Person3(p1)
		dump.Dump(p4)

		type Person4 struct {
			Person Person
			Person2
		}
		p5 := Person4{p1, p4}
		dump.Dump(p5)
	})

	userId := func() int { return 4 }
	dump.Dump(userId())

	dump.Dump(userId2())

	_s := _S{}
	dump.Dump(_s.a())
	dump.Dump(_s.b(`t`))

	num3 := 3
	dump.Dump(`abc`)
	dump.Dump(encoding.JsonEncode(`abc`))
	dump.Dump(encoding.JsonEncode(map[string]interface{}{"a": num3}))

	ch := make(chan bool)
	dump.Dump(ch)
}

type _S struct{}

func (this *_S) a() string {
	return `_s.a`
}

func (this *_S) b(t string) string {
	return `_s.b(` + t + `)`
}

func userId2() int {
	return 8
}
