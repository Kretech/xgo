package dict

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Kretech/xgo/encoding"
)

var (
	ErrNotDict = errors.New(`parsed object is not map[string]interface{}`)
)

type MapDict struct {
	data map[string]interface{}
}

func (m *MapDict) ToMap() map[string]interface{} {
	panic("implement me")
}

func (d *MapDict) String() string {
	return ""
}

func NewMapDict() *MapDict {
	return &MapDict{
		data: newMap(),
	}
}

func (d *MapDict) IsEmpty() bool {
	return len(d.data) == 0
}

func (d *MapDict) Len() int {
	return len(d.data)
}

func (d *MapDict) Get(k interface{}) interface{} {
	paths := strings.Split(toString(k), `.`)

	var current interface{}
	current = toMap(d.data)

	size := len(paths)
	for i := 0; i < size-1; i++ {
		m := toMap(current)
		current = m[paths[i]]
	}

	return toMap(current)[paths[size-1]]
}

func (d *MapDict) Set(k interface{}, v interface{}) {
	paths := strings.Split(toString(k), `.`)

	parent := d.data

	size := len(paths)
	for idx := 0; idx < size-1; idx++ {
		//fmt.Println(idx, d.provider, parent, &d.provider, &parent)
		seq := paths[idx]

		i := parent[seq]
		if _, ok := i.(map[string]interface{}); !ok {
			parent[seq] = newMap()
			parent = parent[seq].(map[string]interface{})
		} else {
			parent = i.(map[string]interface{})
		}
	}

	parent[paths[size-1]] = v
}

func (d *MapDict) Forget(k interface{}) {
	d.Set(k, nil)
}

func (d *MapDict) ParseJsonString(data []byte) (err error) {
	d.data, err = JsonToMap(data)
	return
}

func JsonToMap(data []byte) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})

	var i interface{}
	err = json.Unmarshal(data, &i)
	if err != nil {
		return
	}

	m, ok := i.(map[string]interface{})
	if !ok {
		return m, ErrNotDict
	}

	return
}

func (d *MapDict) Keys() (keys []string) {
	for k := range d.data {
		keys = append(keys, k)
	}
	return
}

func (d *MapDict) Values() (values []interface{}) {
	for _, v := range d.data {
		values = append(values, v)
	}
	return
}

func (d *MapDict) Filter(fn func(interface{}, string) bool) *MapDict {
	instance := NewMapDict()
	d.Each(func(v interface{}, k string) {
		if fn(v, k) {
			instance.Set(k, v)
		}
	})
	return instance
}

func (d *MapDict) Each(fn func(interface{}, string)) {
	for k, v := range d.data {
		fn(v, k)
	}
}

func (d *MapDict) Data() map[string]interface{} {
	return d.data
}

func (d *MapDict) SetData(data map[string]interface{}) {
	d.data = data
}

func (d *MapDict) Json() string {
	return encoding.JsonEncode(d.data)
}

func toString(k interface{}) string {
	switch k.(type) {
	case fmt.Stringer:
		return k.(fmt.Stringer).String()
	case string:
		return k.(string)
	case int:
		return strconv.FormatInt(int64(k.(int)), 10)
	default:
		return fmt.Sprintf("%v", k)
	}
}

func newMap() map[string]interface{} {
	return make(map[string]interface{})
}

func toMap(i interface{}) map[string]interface{} {
	m, ok := i.(map[string]interface{})
	if !ok {
		//fmt.Printf("%v,%v\n", i, m)
		m = make(map[string]interface{})
		//fmt.Printf("%v,%v\n", i, m)
	}

	return m
}
