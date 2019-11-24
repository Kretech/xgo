package pipe

import (
	"os/exec"

	"github.com/pkg/errors"
)

type Pipes struct {
	pipes     []*Pipe
	pipesChan chan *Pipe
}

//
func NewPipes(size int, cmdPtr *exec.Cmd) *Pipes {
	p := &Pipes{
		pipes:     make([]*Pipe, size),
		pipesChan: make(chan *Pipe, size),
	}
	for i := 0; i < size; i++ {
		cmd := *cmdPtr
		pipe := NewPipe(&cmd)
		p.pipes[i] = pipe
		p.pipesChan <- pipe
	}
	return p
}

func (this *Pipes) Start() (err error) {
	for i := 0; i < len(this.pipes); i++ {
		err = this.pipes[i].Start()
		if err != nil {
			return errors.Wrapf(err, "pipes[%d]", i)
		}
	}
	return
}

func (this *Pipes) AcquirePipe() *Pipe {
	return <-this.pipesChan
}

func (this *Pipes) ReleasePipe(p *Pipe) {
	this.pipesChan <- p
}

func (this *Pipes) WriteAndRead(b []byte) (resp []byte, err error) {
	p := this.AcquirePipe()
	defer this.ReleasePipe(p)

	return p.WriteAndRead(b)
}
