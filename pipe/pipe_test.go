package pipe

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"testing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestPipe(t *testing.T) {
	//cmd := exec.Command("awk '{print $0}'")
	cmd := exec.Command("./a.out")
	//cmd := exec.Command("top")

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		t.Error(err)
		return
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Error(err)
		return
	}

	err = cmd.Start()
	if err != nil {
		t.Error(err)
		return
	}

	rn := bufio.NewReadWriter(bufio.NewReader(stdoutPipe), bufio.NewWriter(stdinPipe))
	for i := 0; i < 10; i++ {
		_, err := rn.WriteString("abc" + fmt.Sprint(rand.Int31()) + "\n")
		_ = rn.Writer.Flush()

		line, prefix, err := rn.ReadLine()
		log.Println(string(line), prefix, err)
	}
}

func TestPipe_Call(t *testing.T) {
	//p := NewPipe(exec.Command("awk", "'{print $0}'"))
	p := NewPipe(exec.Command("./a.out"))
	p.Start()

	resp, err := p.WriteAndRead([]byte("hello_pipe\n"))
	log.Println(string(resp), err)
}

func BenchmarkPipe_WriteAndRead(b *testing.B) {
	p := NewPipe(exec.Command("./a.out"))
	p.Start()

	for i := 0; i < b.N; i++ {
		_, _ = p.WriteAndRead([]byte("hello_pipe\n"))
	}
}

func BenchmarkPipe_WriteAndReadParallel(b *testing.B) {
	p := NewPipe(exec.Command("./a.out"))
	p.Start()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = p.WriteAndRead([]byte("hello_pipe\n"))
		}
	})
}
