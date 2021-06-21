package httputil_v2

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

var (
	DefaultMaxIdleConns        = 100
	DefaultMaxIdleConnsPerHost = 10
	DefaultForceAttemptHTTP2   = true

	DefaultInsecureSkipVerify = true

	defaultHeader = map[string]string{
		"Content-Type":   "application/json",
		"Accept-Charset": "UTF-8",
	}
)

type HttpUtil struct {
	c *http.Client
}

func NewHttpUtil(opts ...Option) *HttpUtil {
	c := &http.Client{}

	params := &params{}
	for _, do := range opts {
		do(params)
	}

	if params.timeout != 0 {
		c.Timeout = params.timeout
	}
	if params.proxy != nil {
		c.Transport = &http.Transport{
			Proxy:               http.ProxyURL(params.proxy),
			MaxIdleConns:        DefaultMaxIdleConns,
			MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
			ForceAttemptHTTP2:   DefaultForceAttemptHTTP2,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: DefaultInsecureSkipVerify,
			},
		}
	}

	return &HttpUtil{c}
}

//tips:方法直接将resp返回，记得关闭resp.body
func (u *HttpUtil) Get(url string, opts ...Option) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	param := &params{}
	for _, do := range opts {
		do(param)
	}

	if param.header != nil {
		for k, v := range param.header {
			req.Header.Set(k, v)
		}
	} else {
		for k, v := range defaultHeader {
			req.Header.Set(k, v)
		}
	}
	if param.mock != nil {
		return param.mock(req)
	}

	return u.c.Do(req)
}

//tips:方法直接将resp返回，记得关闭resp.body
func (u *HttpUtil) Post(url string, body []byte, opts ...Option) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return
	}

	param := &params{}
	for _, do := range opts {
		do(param)
	}

	if param.header != nil {
		for k, v := range param.header {
			req.Header.Set(k, v)
		}
	} else {
		for k, v := range defaultHeader {
			req.Header.Set(k, v)
		}
	}
	if param.mock != nil {
		return param.mock(req)
	}

	return u.c.Do(req)
}
