# *XGo*/dump



## Quick Look

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
```

## Option

自定义选项

```go
// 全局禁用
// 可以在生产环境使用该选项避免意外
dump.Disable = true

// 显示代码位置
dump.ShowFileLine1 = true
dump.MarginLine1   = 36     // 显示行号时，前面可以加空格

// 自定义输出
dump.DefaultWriter = os.Stdout

// 数组最多显示多少个，其余的会省略为“...”
dump.MaxSliceLen = 32
dump.MaxMapLen   = 32
```

## More TestCases

https://github.com/Kretech/xgo/blob/master/dump/cli_dumper_test.go
