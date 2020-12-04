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

// NewRequest
func NewRequest(method string, url string) *Request {
	return &Request{
		method: method,
		url:    url,
	}
}

// Header set header
func (r *Request) Header(key, value string) *Request {
	r.header.Set(key, value)
	return r
}

// Body set body with content-type
func (r *Request) Body(contentType string, body io.Reader) *Request {
	r.Header("Content-Type", contentType)
	r.body = body
	return r
}

// Form set body with application/x-www-form-urlencoded
func (r *Request) Form(form url.Values) *Request {
	return r.Body("application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
}

// JSON set body with application/json
func (r *Request) JSON(body interface{}) *Request {
	r, err := r.JSONOrError(body)
	if err != nil {
		panic(err)
	}

	return r
}

// JSONOrError set body with application/json
// return error when marshal failed
func (r *Request) JSONOrError(body interface{}) (*Request, error) {
	r.Header("Content-Type", "application/json")
	b, err := json.Marshal(body)
	if err != nil {
		return r, errors.Wrapf(err, `json.Marshal error`)
	}

	return r.Body("application/json", bytes.NewReader(b)), nil
}

// Build to http.Request
func (r *Request) Build(ctx context.Context) (*http.Request, error) {
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

// Do client.do
func (r *Request) Do(ctx context.Context, clients ...*http.Client) (resp *Response, err error) {
	var c *http.Client
	if len(clients) > 0 {
		c = clients[0]
	} else {
		c = DefaultHttpClient
	}

	req, err := r.Build(ctx)
	if err != nil {
		return
	}

	httpResp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, `Do`)
	}

	resp = &Response{}
	resp.Response = httpResp
	return
}
