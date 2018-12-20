package version

import (
	"testing"

	"github.com/Kretech/xgo/test"
)

func TestVersionCompare(t *testing.T) {

	cases := [][]interface{}{
		{"1.20.3", "2.99.99", resultLess, nil},
		{"1.20.3", "2.99.99", resultLess, nil},
		{"2.99.99", "1.20.3", resultGreater, nil},
		{"1.20.3", "1.20.3", resultEqual, nil},
		{"1.20.3", "2.99.99", resultLess, nil},
		{"1.20.3", "1.20.3333", resultLess, nil},
		{"1.1.1.1", "9999.9999.9999.9999", resultLess, nil},
	}

	for _, cas := range cases {
		r, err := Compare(cas[0].(string), cas[1].(string))
		must := r == cas[2].(T) && err == cas[3]
		// fmt.Println(must, r, err)
		test.AssertEqual(t, must, true)
		test.AssertEqual(t, r, cas[2].(T))
		test.AssertEqual(t, err, cas[3])
	}
}
