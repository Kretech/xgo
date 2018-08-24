package version

import (
	"testing"

	"github.com/Kretech/xgo/test"
)

func TestVersionCompare(t *testing.T) {

	cases := [][]interface{}{
		{"1.20.3", "2.99.99", -1, nil},
		{"1.20.3", "2.99.99", -1, nil},
		{"2.99.99", "1.20.3", 1, nil},
		{"1.20.3", "1.20.3", 0, nil},
		{"1.20.3", "2.99.99", -1, nil},
		{"1.20.3", "1.20.3333", -1, nil},
		{"1.1.1.1", "9999.9999.9999.9999", -1, nil},
	}

	for _, cas := range cases {
		r, err := Compare(cas[0].(string), cas[1].(string))
		must := r == cas[2].(int) && err == cas[3]
		//fmt.Println(must, r, err)
		test.AssertEqual(t, must, true)
		test.AssertEqual(t, r, cas[2].(int))
		test.AssertEqual(t, err, cas[3])
	}

}
