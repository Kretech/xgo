package digo

import (
	"errors"
)

var (
	ErrNotBound = errors.New(`build not-bound types`)
)

type DIGo interface {
	// 	注册单例
	//
	// 	s := new(T)
	// 	di.Singleton(s)
	Singleton(instancePtr interface{}) (err error)

	// 通过闭包方式注册
	// 闭包的第一个返回值的类型作为注册的 T
	// 闭包的参数，会根据类型从容器中取数据
	//
	// 	di.SingletonFunc(func(a *A) (T) {
	// 		// a will build first
	// 		return T{}
	// 	})
	SingletonFunc(fn interface{}) (err error)

	// MapSingleton 表示从接口到实体的映射
	//
	// di.MapSingleton(op.Writer, os.Stdout)
	MapSingleton(T Type, V Value) (err error)

	// 	BindFunc register beans which need build repeat
	//
	// 	di.BindFunc(func(a *A) (*) {
	// 		return T{}
	// 	}
	//BindFunc(fn interface{})

	// 利用容器提供的默认数据赋值
	// 	ptr := new(T)
	// 	di.Build(ptr)
	InvokePtr(ptr interface{}) (err error)

	// 利用容器提供的默认数据调用函数
	// 	di.Invoke(func(a *A, b *B) {
	// 		// a, b built
	// 	}
	InvokeFunc(fn interface{}) (fnOut []interface{}, err error)

	// 利用容器提供的默认数据填充对象的成员
	//
	//
	FillClass(classPtr interface{}) (err error)

	//
	// 	di.Listen(di.BeforeBuilding, func(a *A) {
	// 	})
	//Listen(hook Hook, fn interface{})
}

type Hook uint16

// NewDIGo
func NewDIGo() DIGo {
	return NewDiGoImpl()
}
