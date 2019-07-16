package firewall

import (
	"sync"
	"time"
)

// SleepLimiter 限制单位时间内的可执行次数
type SleepLimiter struct {
	ttl   time.Duration
	limit int

	current     int
	currentTime time.Time

	sync.Mutex
}

func NewSleepLimiter(ttl time.Duration, limit int) *SleepLimiter {
	obj := &SleepLimiter{ttl: ttl, limit: limit}
	return obj
}

func (this *SleepLimiter) Acquire() {
	this.Lock()

	expireAt := this.currentTime.Add(this.ttl)
	now := time.Now()
	// expire
	if expireAt.Before(now) {
		this.current = 0
		this.currentTime = now

		expireAt = this.currentTime.Add(this.ttl)
	}

	this.current++
	if this.current <= this.limit {
		this.Unlock()
		return
	}

	time.Sleep(expireAt.Sub(now))
	this.Unlock()

	// 直接尾递归
	this.Acquire()
}
