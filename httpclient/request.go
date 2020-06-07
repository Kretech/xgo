package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Request is builder for http.Request
type Request struct {
	method string
	url    string
	header http.Header
	body   io.Reader
}

func NewRequestBuilder(method string, url string) *Request {
	return &Request{
		method: method,
		url:    url,
	}
}

func HeadRequest(url string) *Request { return NewRequestBuilder("HEAD", url) }
func GetRequest(url string) *Request  { return NewRequestBuilder("GET", url) }
func PostRequest(url string) *Request { return NewRequestBuilder("POST", url) }
func PutRequest(url string) *Request  { return NewRequestBuilder("PUT", url) }

func (req *Request) Header(key, value string) *Request {
	req.header.Set(key, value)
	return req
}

func (req *Request) Body(contentType string, body io.Reader) *Request {
	req.Header("Content-Type", contentType)
	req.body = body
	return req
}

func (req *Request) Form(form url.Values) *Request {
	return req.Body("application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
}

func (req *Request) JSON(body interface{}) *Request {
	req, err := req.JSONOrError(body)
	if err != nil {
		panic(err)
	}

	return req
}

func (req *Request) JSONOrError(body interface{}) (*Request, error) {
	req.Header("Content-Type", "application/json")
	b, err := json.Marshal(body)
	if err != nil {
		return req, errors.Wrapf(err, `json.Marshal error`)
	}

	return req.Body("application/json", bytes.NewReader(b)), nil
}

func (r *Request) BuildRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(r.method), r.url, r.body)
	if err != nil {
		return nil, errors.Wrapf(err, `NewRequestWithContext`)
	}

	for k, v := range r.header {
		req.Header[k] = v
	}

	return req, nil
}

var (
	DefaultHttpClient = &http.Client{Timeout: 15 * time.Second}
)

func (r *Request) Do(ctx context.Context, clients ...*http.Client) (resp *Response, err error) {
	var c *http.Client
	if len(clients) > 0 {
		c = clients[0]
	} else {
		c = DefaultHttpClient
	}

	req, err := r.BuildRequest(ctx)
	if err != nil {
		return
	}

	stdresp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, `Do`)
	}

	resp = &Response{}
	resp.Response = stdresp
	return
}
