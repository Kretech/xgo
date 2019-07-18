package firewall

type Limiter interface {
	Acquire()
}
