package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	httputil "github.com/intrsokx/kitset/upakit/pkg/httputil_v2"
)

var u *httputil.HttpUtil

func main() {
	u = httputil.NewHttpUtil(httputil.WithTimeout(time.Second * 10))
	get()
	post()

	getMock()
}

func post() {
	fmt.Printf("\n\nPOST\n")
	header := map[string]string{
		"Content-Type": "application/json",
	}
	js := map[string]string{}
	js["name"] = "kangxi"
	js["age"] = "23"
	js["sex"] = "男"
	b, _ := json.Marshal(js)
	fmt.Println("post body:", string(b))

	resp, err := u.Post("http://localhost:8080/post", b, httputil.WithHeader(header))

	if err != nil {
		panic(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	fmt.Println(resp.StatusCode)
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("html:", string(html))
}

func get() {
	fmt.Printf("\n\nGET\n")
	resp, err := u.Get("http://localhost:8080/ping")
	if err != nil {
		panic(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	fmt.Println(resp.StatusCode)
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("html:", string(html))
}

func mock(req *http.Request) (resp *http.Response, err error) {
	fmt.Println("mocked url:", req.URL)
	resp = &http.Response{Request: req}
	resp.StatusCode = http.StatusOK

	readCloser := ioutil.NopCloser(bytes.NewBufferString("OK"))
	resp.Body = readCloser

	return resp, nil
}

func getMock() {
	fmt.Printf("\n\nGET MOCK\n")
	resp, err := u.Get("http://localhost:8080/ping", httputil.WithMock(mock))
	if err != nil {
		panic(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	fmt.Println(resp.StatusCode)
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("html:", string(html))
}
