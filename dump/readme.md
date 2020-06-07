# *XGo*/dump



## Quick Look

```go
aInt := 1
bStr := `sf`
cMap := map[string]interface{}{"name": "z", "age": 14}
dArray := []interface{}{&cMap, aInt, bStr}

dump.Dump(aInt, &aInt, &bStr, bStr, cMap, dArray, cMap["name"], dArray[2], dArray[aInt])
```

![](https://i.loli.net/2020/06/03/ojTUsJAui5WKFMf.png)

## Usage

```go
a := 1

// use default instance
dump.Dump(a)

// new instance
c := dump.NewCliDumper()
c.Dump(a)

// use custom writer
buf := &bytes.Buffer{}
dump.NewCliDumper(dump.OptOut(buf))
```

## Option

custom options

```go
// disable dump in global, default false means enable
// disable it in production to ignore the unnecessary output and risks
dump.Disable = false

// show variant line in head
dump.ShowFileLine1 = true

// change writer in global
dump.DefaultWriter = os.Stdout

// only show the first-n elements
// others will show as '...'
// * these options is design for debug clearly
dump.MaxSliceLen = 32
dump.MaxMapLen   = 32

// serialize options

// render map with sorted keys
dump.OptSortMapKeys = true

// render []uint8(`ab`) as string("ab")
OptShowUint8sAsString = true

// render uint8(97) as char('a')
OptShowUint8AsChar
```

## More TestCases

https://github.com/Kretech/xgo/blob/master/dump/cli_dumper_test.go

https://github.com/Kretech/xgo/blob/master/dump/serialize_test.go
