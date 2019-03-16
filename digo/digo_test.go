package digo_test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/Kretech/xgo/digo"
	"github.com/Kretech/xgo/dump"
)

func TestDemo_Singleton(t *testing.T) {
	di := digo.NewDIGo()

	// 注册两个基本类型
	_ = di.SingletonFunc(func() int { return 316 })
	_ = di.SingletonFunc(func(i int) string { return fmt.Sprint(i) })

	var age int
	var name string

	_ = di.InvokePtr(&age)
	_ = di.InvokePtr(&name)

	if age != 316 {
		t.Error(dump.Serialize(age))
	}

	if name != `316` {
		t.Error(dump.Serialize(name))
	}
	println(age, name)
	// 利用容器提供的默认数据调用函数
	t.Run(`InvokeFunc()`, func(t *testing.T) {
		invoked := ``
		_, _ = di.InvokeFunc(func(a int, s string) {
			invoked = fmt.Sprintf("a=%v s=%v", a, s)
		})
		if invoked != `a=316 s=316` {
			t.Error(dump.Serialize(invoked))
		}
	})

	// 利用容器提供的默认数据构造对象
	type P struct {
		Name string
		Age  int
	}
	_ = di.SingletonFunc(func(name string, age int) P {
		return P{Name: name, Age: age}
	})

	t.Run(`InvokePtr()`, func(t *testing.T) {
		p1 := new(P)
		_ = di.InvokePtr(p1)
		if !(p1.Name == name && p1.Age == age) {
			t.Error(dump.Serialize(p1))
		}

		// 测试单例生成的 struct 是同一个
		// 注意，p2 的类型是 *P，如果 p2 定义为 P，Build 时传 &p，虽然不会影响测试结果，但你自己已经维护了两个 struct 了
		p2 := new(P)
		_ = di.InvokePtr(p2)
		if *p2 != *p1 {
			t.Error(dump.Serialize(p2))
		}

		// 用 Singleton() 注册对象
		_ = di.Singleton(p2)
		p3 := new(P)
		_ = di.InvokePtr(p3)
		if *p2 != *p3 {
			t.Error(dump.Serialize(p3))
		}

		// FillClass
		t.Run(`FillClass()`, func(t *testing.T) {
			p4 := new(P)
			_ = di.FillClass(p4)
			if *p2 != *p4 {
				t.Error(dump.Serialize(p4))
			}
		})
	})

	// 测试 MapSingleton
	t.Run(`MapSingleton()`, func(t *testing.T) {
		_ = di.MapSingleton(reflect.TypeOf(new(int)), reflect.ValueOf(7))
		var num7 int
		_ = di.InvokePtr(&num7)
		if num7 != 7 {
			t.Error(dump.Serialize(num7))
		}
	})

	// 测试 MapSingleton 绑定 interface 和 implement
	t.Run(`MapSingleton()`, func(t *testing.T) {
		buf := bytes.Buffer{}
		_ = di.MapSingleton(reflect.TypeOf(new(io.Writer)), reflect.ValueOf(&buf))
		_, _ = di.InvokeFunc(func(w io.Writer) {
			_, _ = w.Write([]byte(`hi`))
		})
		if buf.Len() != 2 {
			t.Error(dump.Serialize(buf.String()))
		}
	})

	p := new(struct {
		Name string;
		Age  int
	})
	_ = di.FillClass(p)
	dump.Dump(p)

}
