package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type HostGetter func() string

type Client struct {
	host HostGetter

	*http.Client
}

func NewClient() *Client {
	c := &Client{
		Client: http.DefaultClient,
	}
	return c
}

func (c *Client) SetHost(host string) *Client {
	c.host = func() string { return host }
	return c
}

func (c *Client) SetHostGetter(hg HostGetter) *Client {
	c.host = hg
	return c
}

func (c *Client) Get(ctx context.Context, uri string, query ...map[string]string) (resp *Response, err error) {
	uri = c.mergeHost(uri)
	if len(query) > 0 {
		u, _ := url.Parse(uri)
		q := u.Query()
		for k, v := range query[0] {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
		uri = u.String()
	}

	req := NewRequestBuilder(`GET`, uri)
	return c.DoRequest(ctx, req)
}

func (c *Client) PostForm(ctx context.Context, url string, form url.Values) (*Response, error) {
	return c.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
}

func (c *Client) PostJSON(ctx context.Context, url string, body interface{}) (*Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrapf(err, `json.Marshal error`)
	}

	return c.Post(ctx, url, "application/json", bytes.NewReader(b))
}

func (c *Client) Post(ctx context.Context, url, contentType string, body io.Reader) (resp *Response, err error) {
	url = c.mergeHost(url)

	req := NewRequestBuilder(`POST`, url).Body(contentType, body)

	return c.DoRequest(ctx, req)
}

func (c *Client) DoRequest(ctx context.Context, req *Request) (resp *Response, err error) {
	return req.Do(ctx, c.Client)
}

func (c *Client) mergeHost(rawurl string) string {
	hasSchema := strings.HasPrefix(rawurl, `http://`) || strings.HasPrefix(rawurl, "https://")
	if !hasSchema && c.host != nil {
		host := c.host()
		rawurl = fmt.Sprintf("%s/%s", strings.TrimSuffix(host, `/`), strings.TrimPrefix(rawurl, `/`))
	}
	return rawurl
}

var Default = NewClient()
