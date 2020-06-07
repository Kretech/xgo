package pipe

import (
	"bufio"
	"io"
	"os/exec"
	"sync"

	"github.com/pkg/errors"
)

type ExecPipe struct {
	cmd *exec.Cmd

	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser

	bufRW *bufio.ReadWriter

	lock sync.Mutex
}

func NewExecPipe(cmd *exec.Cmd) *ExecPipe {
	return &ExecPipe{cmd: cmd}
}

func (this *ExecPipe) Start() (err error) {
	this.stdinPipe, err = this.cmd.StdinPipe()
	if err != nil {
		return errors.Wrap(err, `pipe start`)
	}

	this.stdoutPipe, err = this.cmd.StdoutPipe()
	if err != nil {
		return errors.Wrap(err, `pipe start`)
	}

	err = this.cmd.Start()
	if err != nil {
		return errors.Wrap(err, `pipe start`)
	}

	this.bufRW = bufio.NewReadWriter(bufio.NewReader(this.stdoutPipe), bufio.NewWriter(this.stdinPipe))

	return
}

func (this *ExecPipe) Stop() error {
	return this.cmd.Process.Kill()
}

func (this *ExecPipe) WriteAndRead(b []byte) (resp []byte, err error) {
	this.lock.Lock()
	defer this.lock.Unlock()

	_, err = this.bufRW.Write(b)
	if err != nil {
		err = errors.Wrap(err, `pipe write`)
		return
	}

	err = this.bufRW.Flush()
	if err != nil {
		err = errors.Wrap(err, `pipe flush`)
		return
	}

	resp, _, err = this.bufRW.ReadLine()
	if err != nil {
		err = errors.Wrap(err, `pipe readLine`)
	}
	return
}
