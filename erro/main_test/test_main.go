package main

import (
	std "errors"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	base := func(m string) error { return std.New(m) }

	err := std.New(`base error`)
	fmt.Println(err)

	func() {
		fmt.Println(base(`pkg/errors.New`))
		fmt.Println(errors.Cause(base(`pkg/errors.Cause`)))
		fmt.Println(errors.WithStack(base(`pkg/errors.WithStack`)))
		fmt.Println(errors.WithMessage(base(`with`), `message`))
		fmt.Println(errors.WithMessagef(base(`withMessageF`), `%v`, `args`))
		fmt.Println(errors.Wrap(base(`wrap`), `message`))
	}()

	for _, arg := range os.Environ() {
		fmt.Println(arg)
	}

}
