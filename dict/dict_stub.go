package dict

type DictStub struct {
}

func (*DictStub) IsEmpty() bool {
	panic("implement me")
}

func (*DictStub) String() string {
	panic("implement me")
}

func (*DictStub) Len() int {
	panic("implement me")
}

func (*DictStub) Exist(k any) {
	panic("implement me")
}

func (*DictStub) Get(k any) any {
	panic("implement me")
}

func (*DictStub) Set(k any, v any) {
	panic("implement me")
}

func (*DictStub) Del(k any) {
	panic("implement me")
}

func (*DictStub) Forget(k any) {
	panic("implement me")
}

func (*DictStub) Flush() {
	panic("implement me")
}

func (*DictStub) ParseMap(data any) {
	panic("implement me")
}

func (*DictStub) ParseJson(data []byte) (err error) {
	panic("implement me")
}

func (*DictStub) Keys() (keys []string) {
	panic("implement me")
}

func (*DictStub) Values() (values []any) {
	panic("implement me")
}

func (*DictStub) Filter(fn func(any, string) bool) Dict {
	panic("implement me")
}

func (*DictStub) Each(fn func(any, string)) {
	panic("implement me")
}

func (*DictStub) Map() map[string]any {
	panic("implement me")
}

func (*DictStub) Json() string {
	panic("implement me")
}
