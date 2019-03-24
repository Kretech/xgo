package dict

type Drawer interface {
	Get(string) Drawer

	Int() int

	String() string
}

type MapDrawer struct {
	data interface{}
}

func (d *MapDrawer) Get(key string) Drawer {

	return Drawer(&MapDrawer{
		d.data.(map[string]interface{})[key],
	})
}

func (*MapDrawer) Int() int {
	panic("implement me")
}

func (*MapDrawer) String() string {
	panic("implement me")
}
