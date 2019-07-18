package firewall

import (
	"runtime"
	"sync"
	"time"
)

type MutexLimiter struct {
	ttl time.Duration

	limit   int
	current int
	flushed int

	lock sync.Mutex
}

func (l *MutexLimiter) Acquire() {
	l.lock.Lock()

	if l.current < l.limit {
		l.current++
		l.lock.Unlock()
		return
	}

	// if overflow
	// release lock and acquire later
	l.lock.Unlock()
	runtime.Gosched()
	l.Acquire()
}

func NewMutexLimiter(ttl time.Duration, limit int) *MutexLimiter {
	obj := &MutexLimiter{ttl: ttl, limit: limit}

	start := make(chan struct{})
	go func() {
		start <- struct{}{}
		for range time.Tick(ttl) {
			obj.lock.Lock()
			obj.current = 0
			obj.flushed++
			obj.lock.Unlock()
		}
	}()
	<-start

	return obj
}
