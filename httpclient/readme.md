# *XGo*/httpClient

## Usage

### Basic

```go
resp, _ := httpclient.Default.Get(context.Background(), "https://www.baidu.com/sugrec")

// 打印返回值
log.Println(resp.String()) 

// 反序列化到对象
resp.To(&obj)

// 转为 interface{}，复合对象为 map/array
resp.ToInterface()

// POST
httpclient.Default.PostJSON(ctx, "http://foo.com/bar", struct{})
httpclient.Default.PostForm(ctx, "http://foo.com/bar", map[string][]string{})
```

### Custom Header

```go
resp, _ := httpclient.PostRequest("http://foo.com/bar").
	Header("key", "value").
	Form(map[string][]string{}). // set form-urlencode body
	JSON(struct{}). // set json body
	Body("content-type", io.Reader). // set other body
	Do(ctx)
```

### Custom Response Wrapper

```go
resp, _ := ...
// {
//     "code": -1,
//     "err_msg": "some error",
//     "data": {
//         "id": 1
//     }
// }

resp, err := resp.Unwrap(FooWrapper{})
println(err)	// "some error"
println(resp)	// {"id":1}
```

### Dynamic Request Host

```go
c := httpclient.NewClient().SetHost(`http://foo.com`)
resp, _ := c.Get("/abc")

c := httpclient.NewClient().SetHostGetter(func() string { 
	hosts := []string{"http://192.168.0.1", "http://192.168.0.2"}
	return hosts[rand.Int()]
}})
resp, _ := c.Get("/abc")
```

### Retries

todo 

