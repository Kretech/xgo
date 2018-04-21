package dict

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Kretech/common/encoding"
)

var (
	ErrNotDict = errors.New(`parsed object is not map[string]interface{}`)
)

type Dict struct {
	data map[string]interface{}
}

func NewDict() *Dict {
	return &Dict{
		data: newMap(),
	}
}

func (d *Dict) IsEmpty() bool {
	return len(d.data) == 0
}

func (d *Dict) Len() int {
	return len(d.data)
}

func (d *Dict) Get(k interface{}) interface{} {
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

func (d *Dict) Set(k interface{}, v interface{}) {
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

func (d *Dict) Forget(k interface{}) {
	d.Set(k, nil)
}

func (d *Dict) ParseJsonString(data []byte) (err error) {
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

func (d *Dict) Keys() (keys []string) {
	for k := range d.data {
		keys = append(keys, k)
	}
	return
}

func (d *Dict) Values() (values []interface{}) {
	for _, v := range d.data {
		values = append(values, v)
	}
	return
}

func (d *Dict) Pluck() {}

func (d *Dict) Filter(fn func(interface{}, string) bool) *Dict {
	instance := NewDict()
	d.Each(func(v interface{}, k string) {
		if fn(v, k) {
			instance.Set(k, v)
		}
	})
	return instance
}

func (d *Dict) Each(fn func(interface{}, string)) {
	for k, v := range d.data {
		fn(v, k)
	}
}

func (d *Dict) Data() map[string]interface{} {
	return d.data
}

func (d *Dict) SetData(data map[string]interface{}) {
	d.data = data
}

func (d *Dict) Json() string {
	return encoding.JsonEncode(d.data)
}

func toString(k interface{}) string {
	switch k.(type) {
	case string:
		return k.(string)
	case int:
		return strconv.FormatInt(int64(k.(int)), 10)
	default:
		return fmt.Sprint(k)
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
