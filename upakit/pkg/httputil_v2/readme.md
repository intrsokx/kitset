# httputil v2

## example

### init httputil
```
u := httputil_v2.NewHttpUtil()

proxy, _ := url.Parse("proxy_str")
u := httputil_v2.NewHttpUtil(httputil_v2.WithTimeout(time.Second*10), httputil_v2.WithProxy(proxy))
```
### get and post
```
u := httputil_v2.NewHttpUtil()

resp, err := u.Get(url)
resp, err := u.Get(url, httputil_v2.WithMock(mock))
resp, err := u.Get(target, httputil_v2.WithHeader(header), httputil_v2.WithMock(mock))

```