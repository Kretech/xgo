# *XGo*/p

Tricking implements for php runtime functions


[TOC]

[所有函数列表 => GoWalker](https://gowalker.org/github.com/Kretech/xgo/p)

## Sample

### Dumper

```go
aInt := 1
bStr := `sf`
cMap := map[string]interface{}{"name": "z", "age": 14}
dArray := []interface{}{&cMap, aInt, bStr}
c := cMap

p.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, c, cMap["name"], dArray[2], dArray[aInt])
```

![](https://i.loli.net/2019/03/14/5c8a541bd8497.png)

## 已知问题

- compact 目前只支持一行调一次
	- 目前定位调用代码是通过行号和函数名，所以只支持一行一个

## todo

- [ ] 获取函数签名
- [ ] ClassLoader
- [ ] Functional http API
- [ ] JSON-RPC
