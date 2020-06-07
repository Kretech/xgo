package pipe

type Pipe interface {
	Start() error
	Stop() error
	WriteAndRead([]byte) ([]byte, error)
}
