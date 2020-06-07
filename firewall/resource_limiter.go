package firewall

type ResourceLimiter interface {
	Acquire() interface{}
	Release(resource interface{})
}

type ChanResourceLimiter struct {
	pool chan interface{}
}

// NewChanResourcePool ...
func NewChanResourcePool(cap int) *ChanResourceLimiter {
	obj := &ChanResourceLimiter{pool: make(chan interface{}, cap)}
	return obj
}

func (this *ChanResourceLimiter) Acquire() interface{} {
	return <-this.pool
}

func (this *ChanResourceLimiter) Release(resource interface{}) {
	this.pool <- resource
}
