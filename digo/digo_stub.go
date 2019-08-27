package digo

var _ DIGo = &stub{}

type stub struct {
}

func (di *stub) FillClass(classPtr interface{}) (err error) { return }

func (di *stub) MapSingleton(T Type, V Value) (err error) { return err }

func (di *stub) Listen(hook Hook, fn interface{}) {}

func (di *stub) Singleton(instancePtr interface{}) (err error) { return err }

func (di *stub) SingletonFunc(fn interface{}) (err error) { return err }

func (di *stub) BindFunc(fn interface{}) {}

func (di *stub) InvokePtr(ptr interface{}) error { return nil }

func (di *stub) InvokeFunc(fn interface{}) (fnOut []interface{}, err error) { return fnOut, err }
