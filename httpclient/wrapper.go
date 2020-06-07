package httpclient

import "errors"

// see wrapper_test.go
type ResponseWrapper interface {
	Error() error
	SetData(interface{})
	GetData() interface{}
}

type CEDResponseWrapper struct {
	Code   int    `json:"code,omitempty"`
	ErrMsg string `json:"err_msg,omitempty"`
	HasData
}

func (ced *CEDResponseWrapper) Error() error {
	if ced.Code == 0 {
		return nil
	}

	return errors.New(ced.ErrMsg)
}

type HasData struct {
	Data interface{} `json:"data,omitempty"`
}

func (w *HasData) SetData(data interface{}) {
	w.Data = data
}

func (w *HasData) GetData() interface{} {
	return w.Data
}
