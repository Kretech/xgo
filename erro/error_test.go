package erro

import (
	std "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestDemo(t *testing.T) {

	base := func(m string) error { return std.New(m) }

	t.Run(`std/errors`, func(t *testing.T) {
		err := std.New(`base error`)
		fmt.Println(err)
	})

	t.Run(``, func(t *testing.T) {
		err := fmt.Errorf("%w", std.New(`base`))
		t.Log(err)
	})

	t.Run(`pkg/errors`, func(t *testing.T) {
		func() {
			fmt.Println(base(`pkg/errors.New`))
			fmt.Println(errors.Cause(base(`pkg/errors.Cause`)))
			fmt.Println(errors.WithStack(base(`pkg/errors.WithStack`)))
			fmt.Println(errors.WithMessage(base(`with`), `message`))
			fmt.Println(errors.WithMessagef(base(`withMessageF`), `%v`, `args`))
			fmt.Println(errors.Wrap(base(`wrap`), `message`))
		}()
	})

	a()
}

func a() { b() }
func b() { c() }
func c() {
	fmt.Println(errors.WithStack(errors.New(`a/b/c`)))
}
