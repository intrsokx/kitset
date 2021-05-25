package httputil

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var defaultReqHeader = map[string]string{
	"Content-Type":   "application/json",
	"Accept-Charset": "UTF-8",
}

type HttpUtil struct {
	client *http.Client
	header map[string]string
}

func NewHttpUtil(timeout time.Duration) *HttpUtil {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:      false,
			DisableCompression:     true,
			MaxIdleConns:           100,
			MaxIdleConnsPerHost:    10,
			MaxConnsPerHost:        0, //0表示不限制
			ResponseHeaderTimeout:  timeout,
			MaxResponseHeaderBytes: 1 << 20,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: timeout,
	}

	util := &HttpUtil{
		client: client,
		header: defaultReqHeader,
	}

	return util
}

func NewHttpUtilWithProxy(timeout time.Duration, proxy *url.URL) *HttpUtil {
	/**client配置
	Idle: 空闲
	Coon: 连接
	Host: 单台服务器，多个不同端口的服务是不同的host
	*/
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:                  http.ProxyURL(proxy),
			DisableKeepAlives:      false,
			DisableCompression:     true,
			MaxIdleConns:           100,
			MaxIdleConnsPerHost:    10,
			MaxConnsPerHost:        0, //0表示不限制
			ResponseHeaderTimeout:  timeout,
			MaxResponseHeaderBytes: 1 << 20,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: timeout,
	}

	util := &HttpUtil{
		client: client,
		header: defaultReqHeader,
	}

	return util
}

func (util *HttpUtil) SetHeader(header map[string]string) {
	util.header = header
}

type Response struct {
	Body       []byte
	StatusCode int
}

func (util *HttpUtil) Post(url string, data []byte) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for k, v := range util.header {
		req.Header.Set(k, v)
	}

	resp, err := util.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()

	ret := &Response{}
	ret.StatusCode = resp.StatusCode
	ret.Body, err = ioutil.ReadAll(resp.Body)

	return ret, err
}
