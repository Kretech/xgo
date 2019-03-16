# digo

A Denpdency-Injection Container for Golang

## Quick Start

### Singleton

```go
di := digo.NewDIGo()

_ = di.SingletonFunc(func() int { return 316 })
_ = di.SingletonFunc(func(i int) string { return fmt.Sprint(i) })

p := new(struct {Name string;Age  int})
_ = di.FillClass(p)
dump.Dump(p)		
// p => *struct{ Name string; Age int } ({
//	"Name": "316",
//	"Age": 7
// } 

_, _ = di.InvokeFunc(func(age int, name string) {
	println(age, name)
	// 316 316
})
```

## Demo

https://github.com/Kretech/xgo/blob/master/digo/digo_test.go

