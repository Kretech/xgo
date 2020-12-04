package httpclient_test

import (
	"context"
	"testing"
	"time"

	"github.com/Kretech/xgo/httpclient"
)

func TestRequest(t *testing.T) {
	c := httpclient.Default
	c.Timeout = time.Second * 2
	resp, err := c.Get(context.Background(), "https://www.baidu.com/sugrec")
	if err != nil {
		t.Errorf("%+v", err)
	}

	tos, err := resp.String()
	if err != nil {
		t.Errorf("%+v", err)
	}
	if len(tos) < 1 {
		t.Error("toString failed")
	}

	m := map[string]interface{}{}
	err = resp.To(&m)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if len(m) < 1 {
		t.Error("toMap failed")
	}

	iface, err := resp.ToInterface()
	if err != nil {
		t.Errorf("%+v", err)
	}
	if iface == nil {
		t.Error("ToInteface failed")
	}
}
