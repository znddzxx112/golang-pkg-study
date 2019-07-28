package http

import (
	"io"
	"net/http"
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

