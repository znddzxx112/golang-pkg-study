package pprof

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"testing"
)

func TestClientGet(t *testing.T) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost:     2,
			MaxIdleConnsPerHost: 1,
		},
	}
	defer client.CloseIdleConnections()

	values := url.Values{}
	values.Add("foo", "bar")
	values.Add("foo2", "bar2")

	url := fmt.Sprintf(
		"%s://%s/gethello?%s",
		"http", net.JoinHostPort("127.0.0.1", "8889"), values.Encode())

	cok := &http.Cookie{
		Name:  "cfoo",
		Value: "cbar",
	}

	header := http.Header{}
	header.Add("VH", "0_62")

	req, _ := http.NewRequest("GET", url, nil)
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
