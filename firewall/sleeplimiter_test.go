package firewall

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"golang.org/x/time/rate"
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

func TestNewSleepLimiterDemo(t *testing.T) {
	l := NewSleepLimiter(1*time.Second, 2)
	time.Sleep(2 * time.Second)
	for {
		l.Acquire()
		log.Println(`hi`)
	}
}

func BenchmarkSleepLimiter_Acquire(b *testing.B) {
	s := NewSleepLimiter(time.Second, 1<<30)
	for i := 0; i < b.N; i++ {
		s.Acquire()
	}
}

func BenchmarkRateLimiter(b *testing.B) {
	var limiter = rate.NewLimiter(1<<30, 1)
	for i := 0; i < b.N; i++ {
		limiter.Wait(context.TODO())
	}
}
