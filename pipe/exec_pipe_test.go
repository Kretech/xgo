package pipe

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"sync"
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
		t.Log(string(line), prefix, err)
	}
}

func TestPipe_Call(t *testing.T) {
	//p := NewPipe(exec.Command("awk", "'{print $0}'"))
	p := NewExecPipe(exec.Command("./a.out"))
	p.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Wait()
		err := p.Stop()
		if err != nil {
			t.Error(err)
		}
	}()

	resp, err := p.WriteAndRead([]byte("hello_pipe\n"))
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(resp), "hello_pipe") {
		t.Error(string(resp))
	}

	wg.Done()
}

func BenchmarkPipe_WriteAndRead(b *testing.B) {
	p := NewExecPipe(exec.Command("./a.out"))
	p.Start()
	defer p.Stop()

	for i := 0; i < b.N; i++ {
		_, _ = p.WriteAndRead([]byte("hello_pipe\n"))
	}
}

func BenchmarkPipe_WriteAndReadParallel(b *testing.B) {
	p := NewExecPipe(exec.Command("./a.out"))
	p.Start()
	defer p.Stop()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = p.WriteAndRead([]byte("hello_pipe\n"))
		}
	})
}
