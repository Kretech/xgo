# Firewall

### Intro

| files | how it works | Benchmark / 2.6 GHz Intel Core i7 |
| ------------ | ------------------------------------- | ---- |
| Limiter      | limiter interface |  |
| SleepLimiter | sleep process when tokens are run out | 20M / qps      81.9 ns/op |
| MutexLimiter | use mutex counter and ticker to rate  | 100M / qps     14.2 ns/op |
|              |                                       |      |
| Semaphore    | wraped chan struct{}  | 30M / qps      40.4 ns/op |
|              |                                       |      |
|              |                                       |      |
|              |                                       |      |

### How to use

```go
// RateLimiter
limiter := New{XXX}Limiter(1*time.Second, 300)	// 300 tokens in one second
limiter.Acquire()	// return immediately or waiting useable token

// Semaphore
sema := NewChanSemaphore(cap)
sema.Acquire()
sema.Release()
```

