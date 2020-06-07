package httpclient

import (
	"net/http"
	"strings"
	"testing"
)

func TestCEDWrapper(t *testing.T) {
	resp := &Response{}
	resp.Response = &http.Response{}
	resp.bytes = []byte(`{"code":-1,"err_msg":"someerror","data":{"id":1}}`)

	v, err := resp.UnwrapToInterface(new(CEDResponseWrapper))
	if v.(map[string]interface{})[`id`] != float64(1) {
		t.Errorf(`failed %v %T`, v, v)
	}

	if !strings.Contains(err.Error(), `someerror`) {
		t.Errorf("%+v", err)
	}
}
