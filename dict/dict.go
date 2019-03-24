package dict

type any = interface{}

type Dict interface {
	IsEmpty() bool

	String() string
	Len() int

	Exist(k any)
	Get(k any) any
	Set(k any, v any)
	Del(k any)
	Forget(k any)
	Flush()

	ParseMap(data any)
	ParseJson(data []byte) (err error)

	Keys() (keys []string)
	Values() (values []any)

	Filter(fn func(any, string) bool) Dict
	Each(fn func(any, string))

	// Map 把数据导成 Go map
	Map() map[string]any
	Json() string
}
