package digo

import (
	"reflect"

	"github.com/pkg/errors"
)

type (
	Type = reflect.Type
	Value = reflect.Value
	Kind = reflect.Kind
)

var (
	TypeOf  = reflect.TypeOf
	ValueOf = reflect.ValueOf
)

type DImpl struct {
	//stub

	// 被注入的原始值
	binding map[Type]Value

	// 解析后的值
	store map[Type]Value

	// 标记该类型是否是 singleton
	share map[Type]bool

	// resolved 表示该类型是否已被解析
	// 如果该类型被注入的是一个复杂类型，比如 func，resolved 表示该对象已被 invoked
	resolved map[Type]bool

	startBuilding map[Type]struct{}
	finishBuilt   map[Type]struct{}
}

var _ DIGo = &DImpl{}

func NewDiGoImpl() *DImpl {
	di := &DImpl{
		binding:  make(map[Type]Value),
		store:    make(map[Type]reflect.Value),
		share:    make(map[Type]bool),
		resolved: make(map[Type]bool),

		startBuilding: make(map[Type]struct{}),
		finishBuilt:   make(map[Type]struct{}),
	}
	return di
}

func (di *DImpl) InvokePtr(ptr interface{}) (err error) {
	return di.build(reflect.ValueOf(ptr))
}

func (di *DImpl) build(valuePtr Value) (err error) {
	v := valuePtr.Elem()
	T := v.Type()

	if di.share[T] {
		err = di.resolve(T)
		if err != nil {
			return err
		}
		v.Set(di.store[T])
	}

	return
}

func (di *DImpl) SingletonFunc(fn interface{}) (err error) {
	fnT := TypeOf(fn)
	entryT := fnT.Out(0)

	fnV := ValueOf(fn)

	di.share[entryT] = true

	return di.mapTo(entryT, fnV)
}

func (di *DImpl) Singleton(instancePtr interface{}) (err error) {
	V := ValueOf(instancePtr).Elem()
	T := V.Type()

	di.share[T] = true

	return di.mapTo(T, V)
}

func (di *DImpl) FillClass(classPtr interface{}) (err error) {
	V := ValueOf(classPtr).Elem()
	for i := 0; i < V.NumField(); i++ {
		field := V.Field(i)
		if !field.CanSet() {
			continue
		}

		newValue := reflect.New(field.Type())
		_ = di.build(newValue)
		field.Set(newValue.Elem())
	}
	return
}

func (di *DImpl) MapSingleton(T Type, V Value) (err error) {
	T = T.Elem()

	di.share[T] = true

	return di.mapTo(T, V)
}

func (di *DImpl) mapTo(T Type, V Value) (err error) {
	if _, exist := di.binding[T]; exist {
		// 如果已存在，重置 resolved
		// 下次 build 时就重新 resolve 了
		di.resolved[T] = false
	}

	di.binding[T] = V

	return err
}

func (this *DImpl) BindFunc(fn interface{}) () {

}

func (di *DImpl) InvokeFunc(fn interface{}) (fnOut []interface{}, err error) {
	fnV := ValueOf(fn)
	out, err := di.invokeFunc(fnV)
	for _, value := range out {
		fnOut = append(fnOut, value.Interface())
	}
	return fnOut, err
}

func (di *DImpl) invokeFunc(fnV Value) (out []Value, err error) {
	fnT := fnV.Type()
	in := make([]Value, fnT.NumIn())
	for i := 0; i < len(in); i++ {
		argT := fnT.In(i)
		argV := reflect.New(argT)
		_ = di.build(argV)
		in[i] = argV.Elem()
	}

	out = fnV.Call(in)
	return
}

func (di *DImpl) resolve(t Type) (err error) {
	if di.resolved[t] {
		return err
	}

	bound, ok := di.binding[t]
	if !ok {
		return errors.Wrap(ErrNotBound, t.Name())
	}

	switch bound.Kind() {
	case reflect.Func:
		out, _ := di.invokeFunc(bound)
		di.store[t] = out[0]

	default:
		// 普通类型
		di.store[t] = bound
	}

	di.resolved[t] = true
	return err
}
