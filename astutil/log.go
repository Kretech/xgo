package astutil

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Writer struct {
	io.Writer
}

func (w *Writer) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

func (w *Writer) SetOutput(writer io.Writer) {
	w.Writer = writer
}

var logOutput = &Writer{ioutil.Discard}

func SetLogOutput(writer io.Writer) {
	logOutput.Writer = writer
}

func EnableLog() {
	SetLogOutput(os.Stderr)
}

func DisableLog() {
	SetLogOutput(ioutil.Discard)
}

var vlog = log.New(logOutput, `ast_util: `, log.LstdFlags|log.Lshortfile)
