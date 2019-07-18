# Firewall

### Intro

| files | how to work |      |
| ------------ | ------------------------------------- | ---- |
| Limiter      | limiter interface |      |
| SleepLimiter | sleep process when tokens are run out |      |
| MutexLimiter | use mutex counter and ticker to rate  |      |
|              |                                       |      |
| Semaphore    | Semaphore                             |      |
|              |                                       |      |
|              |                                       |      |
|              |                                       |      |

### How to use

```go
limiter := New{XXX}Limiter(1*time.Second, 300)	// 300 tokens in one second
limiter.Acquire()	// return immediately or waiting useable token
```

