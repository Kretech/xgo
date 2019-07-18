package firewall

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChanSemaphore(t *testing.T) {
	cost := time.Second / 130
	limit := 6
	total := 20

	wg := sync.WaitGroup{}
	wg.Add(total)

	s := NewChanSemaphore(limit)
	start := time.Now()
	for i := 0; i < total; i++ {
		go func(i int) {
			s.Acquire()
			time.Sleep(cost)
			//fmt.Println(i)
			wg.Done()
			s.Release()
		}(i)
	}

	wg.Wait()

	end := time.Now()

	fmt.Println(end.Sub(start))
	// 单位时间实际执行了 总任务/总时间 total/T
	// 单位时间理论最多执行 limit/cost
	// total/T <= limit/cost
	fmt.Println(total*int(cost) <= limit*int(end.Sub(start)))
}

func BenchmarkNewChanSemaphore(b *testing.B) {
	s := NewChanSemaphore(1<<30 - 1)

	for i := 0; i < b.N; i++ {
		s.Acquire()
		s.Release()
	}
}
