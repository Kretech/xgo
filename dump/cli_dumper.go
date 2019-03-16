package dump

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Kretech/xgo/p"
	"github.com/fatih/color"
)

type CliDumper struct {
	out io.Writer
}

var (
	DefaultWriter io.Writer = os.Stdout
)

var _ Dumper = NewCliDumper()

// NewCliDumper
func NewCliDumper(opts ...Opt) *CliDumper {
	obj := &CliDumper{
		out: DefaultWriter,
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}

type Opt func(*CliDumper)

func OptOut(w io.Writer) Opt {
	return func(c *CliDumper) {
		c.out = w
	}
}

func (c *CliDumper) Dump(args ...interface{}) {
	c.DepthDump(1, args...)
}

func (c *CliDumper) DepthDump(depth int, args ...interface{}) {
	names, compacted := p.DepthCompact(depth+1, args...)

	for _, name := range names {
		txt := ""

		if strings.HasPrefix(name, "&") {
			txt += color.New(color.Italic, color.FgMagenta).Sprint("&")
			name = name[1:]
		}

		txt += color.New(color.Italic, color.FgCyan).Sprint(name) + SepKv

		txt += Serialize(compacted[name])

		_, _ = fmt.Fprintln(c.out, txt)
	}
}
