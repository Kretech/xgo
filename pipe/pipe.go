package pipe

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

type Pipe struct {
	cmd *exec.Cmd

	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser

	bufRW *bufio.ReadWriter
}

func NewPipe(cmd *exec.Cmd) *Pipe {
	return &Pipe{cmd: cmd}
}

func (this *Pipe) Start() (err error) {
	this.stdinPipe, err = this.cmd.StdinPipe()
	if err != nil {
		return errors.Wrap(err, ``)
	}

	this.stdoutPipe, err = this.cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, ``)
	}

	err = this.cmd.Start()
	if err != nil {
		return errors.Wrap(err, ``)
	}

	this.bufRW = bufio.NewReadWriter(bufio.NewReader(this.stdoutPipe), bufio.NewWriter(this.stdinPipe))

	return
}

func (this *Pipe) Stop() {
	this.stdinPipe.Close()
	this.stdoutPipe.Close()
	this.cmd.Wait()
}

func (this *Pipe) WriteAndRead(b []byte) (resp []byte, err error) {
	_, err = this.bufRW.Write(b)
	if err != nil {
		err = errors.Wrap(err, ``)
		return
	}

	err = this.bufRW.Flush()
	if err != nil {
		err = errors.Wrap(err, ``)
		return
	}

	resp, _, err = this.bufRW.ReadLine()
	return
}
