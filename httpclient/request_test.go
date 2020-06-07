package httpclient_test

import (
	"context"
	"testing"

	"github.com/Kretech/xgo/httpclient"
)

func TestRequest(t *testing.T) {
	ctx := context.Background()
	resp, err := httpclient.GetRequest("https://www.baidu.com/sugrec").Do(ctx)
	if err != nil {
		t.Errorf("%+v", err)
	}

	s, err := resp.String()
	if err != nil {
		t.Errorf("%+v", err)
	}
	if len(s) < 1 {
		t.Error(`toString error`)
	}
}
