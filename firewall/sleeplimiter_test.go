package firewall

import (
	"fmt"
	"testing"
	"time"
)

func TestSleepLimiter_Acquire(t *testing.T) {
	ttl := time.Second / 123
	limit := 300
	total := 1000

	s := NewSleepLimiter(ttl, limit)
	start := time.Now()
	for i := 0; i < total; i++ {
		s.Acquire()
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println(int(end.Sub(start)/ttl)*limit <= total)
}

func BenchmarkSleepLimiter_Acquire(b *testing.B) {
	s := NewSleepLimiter(time.Second, 1<<30)
	for i := 0; i < b.N; i++ {
		s.Acquire()
	}
}
