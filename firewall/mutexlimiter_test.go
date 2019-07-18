package firewall

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMutexLimiter_Acquire(t *testing.T) {
	ttl := time.Second / 123
	limit := 300
	total := 1000

	s := NewMutexLimiter(ttl, limit)
	start := time.Now()
	for i := 0; i < total; i++ {
		s.Acquire()
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println(int(end.Sub(start)/ttl)*limit <= total)
}

func TestNewMutexLimiterDemo(t *testing.T) {
	log.Println(`hi`, 0, 0, time.Now().Unix())
	l := NewMutexLimiter(1*time.Second, 2)
	time.Sleep(2 * time.Second)
	for {
		l.Acquire()
		log.Println(`hi`, time.Now().UnixNano())
	}
}

func BenchmarkMutexLimiter_Acquire(b *testing.B) {
	s := NewMutexLimiter(time.Second, 1<<30)
	for i := 0; i < b.N; i++ {
		s.Acquire()
	}
}

func TestTick(t *testing.T) {
	for range time.Tick(1 * time.Second) {
		log.Println(`hi`)
		//time.Sleep(800 * time.Millisecond)
		time.Sleep(1 * time.Second)
	}
}
