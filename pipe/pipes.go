package pipe

import (
	"github.com/pkg/errors"
)

type Pipes struct {
	pipes     []Pipe
	pipesChan chan Pipe
}

//
func NewPipes(pipes []Pipe) *Pipes {
	p := &Pipes{
		pipes:     pipes,
		pipesChan: make(chan Pipe, len(pipes)),
	}
	for _, pipe := range pipes {
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

func (this *Pipes) Stop() (err error) {
	for i := 0; i < len(this.pipes); i++ {
		stop := this.pipes[i].Stop()
		if stop != nil {
			err = errors.Wrapf(stop, "pipes[%d]", i)
		}
	}
	return
}

func (this *Pipes) AcquirePipe() Pipe {
	return <-this.pipesChan
}

func (this *Pipes) ReleasePipe(p Pipe) {
	this.pipesChan <- p
}

func (this *Pipes) WriteAndRead(b []byte) (resp []byte, err error) {
	p := this.AcquirePipe()
	defer this.ReleasePipe(p)

	return p.WriteAndRead(b)
}
