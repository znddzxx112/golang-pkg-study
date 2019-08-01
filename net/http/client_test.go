package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	url2 "net/url"
	"strings"
	"testing"
)

// 0.05s
func TestGoClient(t *testing.T) {

	total := 1000
	c := make(chan string, total)
	resCh := make(chan string)
	for j := 1; j <= 20; j++ {
		go func(ch chan string) {
			client := &http.Client{
				Transport: &http.Transport{
					DisableKeepAlives:   false,
					MaxIdleConns:        2,
					MaxConnsPerHost:     1,
					MaxIdleConnsPerHost: 1,
				},
			}
			defer client.CloseIdleConnections()

			var url string
			var ok bool
			for {
				select {
				case url, ok = <-ch:
					if !ok {
						return
					}
					req, NewRequestErr := http.NewRequest("get", url, nil)
					if NewRequestErr != nil {
						t.Fatal(NewRequestErr)
					}

					resp, DoErr := client.Do(req)
					if DoErr != nil {
						t.Fatal(DoErr)
					}
					buf := make([]byte, resp.ContentLength)
					io.ReadFull(resp.Body, buf)
					resp.Body.Close()
					resCh <- string(buf)
				}
			}

		}(c)
	}

	for i := 1; i <= total; i++ {
		url := "http://127.0.0.1:8888/hello"
		c <- url
	}
	close(c)

	for i := 1; i <= total; i++ {
		select {
		case buf, _ := <-resCh:
			t.Log(buf)
		}
	}
}

// 0.37
func TestShortClient(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			MaxIdleConns:        2,
			MaxConnsPerHost:     1,
			MaxIdleConnsPerHost: -1,
		},
	}
	defer client.CloseIdleConnections()

	url := "http://127.0.0.1:8888/hello"

	total := 1000
	for i := 1; i <= total; i++ {
		req, NewRequestErr := http.NewRequest("get", url, nil)
		if NewRequestErr != nil {
			t.Fatal(NewRequestErr)
		}
		resp, DoErr := client.Do(req)
		if DoErr != nil {
			t.Fatal(DoErr)
		}
		buf := make([]byte, resp.ContentLength)
		io.ReadFull(resp.Body, buf)
		resp.Body.Close()
		t.Log(string(buf))
	}
}

// 0.18s
func TestLongClient(t *testing.T)  {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			MaxIdleConns:        2,
			MaxConnsPerHost:     2,
			MaxIdleConnsPerHost: 1,
		},
	}
	defer client.CloseIdleConnections()

	url := "http://127.0.0.1:8888/hello"

	total := 1000
	for i := 1; i <= total; i++ {
		req, NewRequestErr := http.NewRequest("get", url, nil)
		if NewRequestErr != nil {
			t.Fatal(NewRequestErr)
		}
		resp, DoErr := client.Do(req)
		if DoErr != nil {
			t.Fatal(DoErr)
		}
		buf := make([]byte, resp.ContentLength)
		io.ReadFull(resp.Body, buf)
		resp.Body.Close()
		t.Log(string(buf))
	}
}

func TestClientGet(t *testing.T)  {
	client := &http.Client{
		Transport:&http.Transport{
			MaxConnsPerHost:2,
			MaxIdleConnsPerHost:1,
		},
	}
	defer client.CloseIdleConnections()

	values := url2.Values{}
	values.Add("foo", "bar")
	values.Add("foo2", "bar2")

	url := fmt.Sprintf(
		"%s://%s/gethello?%s",
		"http", net.JoinHostPort("127.0.0.1", "8889"), values.Encode())

	cok := &http.Cookie{
		Name : "cfoo",
		Value:"cbar",
	}

	header := http.Header{}
	header.Add("VH", "0_62")

	req,_ := http.NewRequest("GET",url, nil)
	req.AddCookie(cok)
	req.Header = header
	response, GetErr := client.Do(req)
	if GetErr != nil {
		t.Fail()
	}
	defer response.Body.Close()
	var buflen int64 = 4096
	if response.ContentLength > 0 {
		buflen = response.ContentLength
	}
	buf := make([]byte, buflen)
	io.ReadFull(response.Body, buf)
	t.Log(string(buf))
	t.Log(string(buf))
	t.Log(response.StatusCode)
	t.Log(response.Status)
	t.Log(response.Header)
	t.Log(response.Proto)
	t.Log(response.ProtoMajor)
	t.Log(response.ProtoMinor)
	t.Log(response.ContentLength)
	t.Log(response.Header.Get(http.CanonicalHeaderKey("content-Type")))
	t.Log(response.Header.Get(http.CanonicalHeaderKey("content-Typ")))
}

func TestClientPost(t *testing.T)  {
	client := &http.Client{
		Transport:&http.Transport{
			MaxConnsPerHost:2,
			MaxIdleConnsPerHost:1,
		},
	}
	defer client.CloseIdleConnections()

	values := url2.Values{}
	values.Add("foo", "bar")
	values.Add("foo2", "bar2")

	url := fmt.Sprintf(
		"%s://%s/posthello",
		"http", net.JoinHostPort("127.0.0.1", "8890"))

	req,_ := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	req.Header.Set(http.CanonicalHeaderKey("Content-Type"), "application/x-www-form-urlencoded")

	response, _ := client.Do(req)
	defer response.Body.Close()
	var buflen int64 = 4096
	if response.ContentLength > 0 {
		buflen = response.ContentLength
	}
	buf := make([]byte, buflen)
	io.ReadFull(response.Body, buf)
	t.Log(string(buf))
}

type Hellojson struct {
	Id int64 `json:id`
	Name string `json:name`
}

func TestClientPostJson(t *testing.T)  {
	client := &http.Client{
		Transport:&http.Transport{
			MaxConnsPerHost:2,
			MaxIdleConnsPerHost:1,
		},
	}
	defer client.CloseIdleConnections()


	url := fmt.Sprintf(
		"%s://%s/posthello",
		"http", net.JoinHostPort("127.0.0.1", "8890"))


	hello := Hellojson{Id:12,Name:"baike"}
	MarshalStr, _ := json.Marshal(hello)

	req,_ := http.NewRequest("POST", url, bytes.NewReader(MarshalStr))
	req.Header.Set(http.CanonicalHeaderKey("Content-Type"), "application/json")

	response, _ := client.Do(req)
	defer response.Body.Close()
	var buflen int64 = 4096
	if response.ContentLength > 0 {
		buflen = response.ContentLength
	}
	buf := make([]byte, buflen)
	io.ReadFull(response.Body, buf)
	t.Log(string(buf))
}
