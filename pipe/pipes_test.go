package pipe

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestPipes_WriteAndRead(t *testing.T) {
	pipes := []Pipe{}
	for i := 0; i < 8; i++ {
		pipes = append(pipes, NewExecPipe(exec.Command("./a.out")))
	}
	p := NewPipes(pipes)

	err := p.Start()
	if err != nil {
		t.Error(err)
	}
	defer p.Stop()

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("hello_pipe%d\n", i)
		resp, err := p.WriteAndRead([]byte(s))
		if strings.Contains(string(resp), s) {
			t.Error("expect", s, "got", string(resp), err)
		}
	}
}

func BenchmarkPipes_WriteAndReadParallel(b *testing.B) {
	pipes := []Pipe{}
	for i := 0; i < 8; i++ {
		pipes = append(pipes, NewExecPipe(exec.Command("./a.out")))
	}
	p := NewPipes(pipes)

	err := p.Start()
	if err != nil {
		b.Error(err)
	}
	defer p.Stop()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err = p.WriteAndRead([]byte("hello_pipe3\n"))
			if err != nil {
				b.Error(err)
			}
		}
	})
}
