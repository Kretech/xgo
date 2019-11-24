package pipe

import (
	"log"
	"os/exec"
	"testing"
)

func TestPipes_WriteAndRead(t *testing.T) {
	p := NewPipes(8, exec.Command("./a.out"))
	err := p.Start()
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := p.WriteAndRead([]byte("hello_pipe1\n"))
	log.Println(string(resp), err)

	resp, err = p.WriteAndRead([]byte("hello_pipe2\n"))
	log.Println(string(resp), err)

	resp, err = p.WriteAndRead([]byte("hello_pipe3\n"))
	log.Println(string(resp), err)
}

func BenchmarkPipes_WriteAndReadParallel(b *testing.B) {
	p := NewPipes(4, exec.Command("./a.out"))
	err := p.Start()
	if err != nil {
		b.Error(err)
		return
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err = p.WriteAndRead([]byte("hello_pipe3\n"))
			if err != nil {
				b.Error(err)
			}
		}
	})
}
