package httpclient

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Response struct {
	*http.Response

	bytes []byte
}

func (r *Response) Bytes() (bytes []byte, err error) {
	if len(r.bytes) > 0 {
		return r.bytes, nil
	}

	r.bytes, err = ioutil.ReadAll(r.Response.Body)
	if err != nil {
		err = errors.Wrapf(err, "Response.ReadAll")
		return []byte{}, err
	}
	defer r.Body.Close()

	return r.bytes, nil
}

func (r *Response) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return ``, err
	}
	return string(bytes), nil
}

var (
	Unmarshalers = map[string]func([]byte, interface{}) error{
		`application/json`: JsonUnmarshaler,
		`application/xml`:  XmlUnmarshaler,
	}

	DefaultUnmarshaler = JsonUnmarshaler

	JsonUnmarshaler = func(data []byte, v interface{}) error { return json.Unmarshal(data, v) }
	XmlUnmarshaler  = func(data []byte, v interface{}) error { return xml.Unmarshal(data, v) }
)

func (r *Response) To(v interface{}) error {
	bytes, err := r.Bytes()
	if err != nil {
		return err
	}

	var found bool
	contentType := r.Response.Header.Get("Content-Type")
	for key, unmarshaler := range Unmarshalers {
		if strings.Contains(contentType, key) {
			err = unmarshaler(bytes, v)
			if err == nil {
				found = true
				break
			}

			err = errors.Wrapf(err, `Unmarshal error with: %s`, bytes)
			return err
		}
	}
	if !found {
		if DefaultUnmarshaler == nil {
			return errors.Errorf(`nonsupport Content-Type: ` + contentType)
		}

		err = DefaultUnmarshaler(bytes, v)
		if err != nil {
			return errors.Wrapf(err, `Unmarshal error with: %s`, bytes)
		}
	}

	return nil
}

func (r *Response) ToInterface() (interface{}, error) {
	var i interface{}
	err := r.To(&i)
	return i, err
}

func (r *Response) UnwrapTo(w ResponseWrapper, v interface{}) error {
	w.SetData(v)
	err := r.To(w)
	if err != nil {
		return err
	}

	err = w.Error()
	if err != nil {
		return errors.Wrapf(err, "response.user.error")
	}

	return nil
}

func (r *Response) UnwrapToInterface(w ResponseWrapper) (interface{}, error) {
	var i interface{}
	err := r.UnwrapTo(w, &i)
	return i, err
}
