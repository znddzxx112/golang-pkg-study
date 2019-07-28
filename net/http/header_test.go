package http

import (
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCanonicalHeaderKey(t *testing.T)  {
	key := http.CanonicalHeaderKey("content-length")
	t.Log(key)
}

func TestDetectContentType(t *testing.T)  {
	im,OpenErr := os.Open("./contentType.jpg")
	if OpenErr != nil {
		t.Fatal(OpenErr)
	}
	defer im.Close()
	header := make([]byte, 512)
	_, ReadErr := im.Read(header)
	if ReadErr != nil {
		t.Fatal(ReadErr)
	}
	t.Log(http.DetectContentType(header))
}

// http2
// http://www.imooc.com/article/details/id/78928

func TestParseHttpVersion(t *testing.T) {
	var major,minor int
	var ok bool
	if major, minor, ok = http.ParseHTTPVersion("HTTP/1.1");!ok {
		t.Fail()
	}
	t.Log(major, minor)
}

func TestParseTime(t *testing.T) {
	var clock time.Time
	var ok error
	if clock,ok = http.ParseTime("Sun, 28 Jul 2019 05:35:32 GMT");ok != nil {
		t.Fail()
	}
	t.Log(clock.Format("2006-01-02 15:04:05"))
}

func TestStatusText(t *testing.T)  {
	var text string
	text = http.StatusText(200)
	t.Log(text)
	text = http.StatusText(503)
	t.Log(text)
}