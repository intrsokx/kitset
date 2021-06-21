package httputil_v2

import (
	"net/http"
	"net/url"
	"time"
)

type params struct {
	//NewHttpUtil param
	timeout time.Duration
	proxy   *url.URL

	//Get Post param
	header map[string]string
	mock   MockFunc
}
type MockFunc func(req *http.Request) (resp *http.Response, err error)

type Option func(opt *params)

func WithTimeout(timeout time.Duration) Option {
	return func(opt *params) {
		opt.timeout = timeout
	}
}

func WithProxy(proxy *url.URL) Option {
	return func(opt *params) {
		opt.proxy = proxy
	}
}

func WithHeader(header map[string]string) Option {
	return func(opt *params) {
		opt.header = header
	}
}

func WithMock(mockFunc MockFunc) Option {
	return func(opt *params) {
		opt.mock = mockFunc
	}
}
