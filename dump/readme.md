# xgo/dump



## Quick Start

```go
aInt := 1
bStr := `sf`
cMap := map[string]interface{}{"name": "z", "age": 14}
dArray := []interface{}{&cMap, aInt, bStr}

dump.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
```

![](https://ws1.sinaimg.cn/mw690/8f9ce571ly1g13yuxm4boj20tk0zuncl.jpg)

## Usage

```go
a := 1

// use default instance
dump.Dump(a)

// new instance
c := dump.NewCliDumper()
c.Dump(a)

// 自定义 out
buf := &bytes.Buffer{}
dump.NewCliDumper(dump.OptOut(buf))

// 可以在生成环境全局禁用该功能
dump.Disable = true
```

## More TestCases

https://github.com/Kretech/xgo/blob/master/dump/cli_dumper_test.go
