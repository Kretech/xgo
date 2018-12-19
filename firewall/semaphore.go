package firewall

type Semaphore interface {
	// P wait()
	Acquire()

	// V signal()
	Release()
}

// NewSemaphore
func NewSemaphore(cap int) Semaphore {
	return NewChanSemaphore(cap)
}

// Semaphore 限制同一时刻，最多可以发生的次数
type ChanSemaphore struct {
	pool chan struct{}
}

// NewChanSemaphore
func NewChanSemaphore(cap int) *ChanSemaphore {
	obj := &ChanSemaphore{pool: make(chan struct{}, cap)}
	return obj
}

func (this *ChanSemaphore) Release() {
	<-this.pool
}

func (this *ChanSemaphore) Acquire() {
	this.pool <- struct{}{}
}
