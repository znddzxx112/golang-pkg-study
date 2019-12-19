package http

import (
	"bytes"
	"crypto/tls"
	"github.com/golang/protobuf/proto"
	"github.com/znddzxx112/golang-pkg-study/net/http/http2proto"
	"golang.org/x/net/http2"
	"io"
	"net/http"
	"testing"
)

// openssl genrsa -out default.key 2048

// openssl req -new -x509 -key default.key -out default.pem -days 3650

func TestServer(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("VH", "0_62")
		io.WriteString(w, "http2")
	})
	server := &http.Server{
		Addr:    ":8899",
		Handler: mux,
	}
	if ListenAndServeTLSErr := server.ListenAndServeTLS("./default.pem", "./default.key"); ListenAndServeTLSErr != nil {
		t.Fatal(ListenAndServeTLSErr)
	}
}

func TestClient(t *testing.T) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:    10,
		MaxConnsPerHost: 1,
	}
	if ConfigureTransportErr := http2.ConfigureTransport(transport); ConfigureTransportErr != nil {
		t.Fatal(ConfigureTransportErr)
	}
	client := &http.Client{
		Transport: transport,
	}
	req, NewRequestErr := http.NewRequest("GET", "https://127.0.0.1:8899/", nil)
	if NewRequestErr != nil {
		t.Fatal(NewRequestErr)
	}
	resp, DoErr := client.Do(req)
	if DoErr != nil {
		t.Fatal(DoErr)
	}
	defer resp.Body.Close()
	var cl int64 = 1024
	if resp.ContentLength > 0 {
		cl = resp.ContentLength
	}
	buf := make([]byte, cl)
	io.ReadFull(resp.Body, buf)
	t.Log(string(buf))
}

func TestClientProto(t *testing.T) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:    10,
		MaxConnsPerHost: 1,
	}
	if ConfigureTransportErr := http2.ConfigureTransport(transport); ConfigureTransportErr != nil {
		t.Fatal(ConfigureTransportErr)
	}
	client := &http.Client{
		Transport: transport,
	}

	cmd := &http2proto.Cmd{
		Name:    "hello",
		ArgInfo: []byte("proto"),
	}
	reqCmdStr, MarshalErr := proto.Marshal(cmd)
	if MarshalErr != nil {
		t.Fatal(MarshalErr)
	}

	req, NewRequestErr := http.NewRequest("POST", "https://127.0.0.1:8899/", bytes.NewReader(reqCmdStr))
	if NewRequestErr != nil {
		t.Fatal(NewRequestErr)
	}
	resp, DoErr := client.Do(req)
	if DoErr != nil {
		t.Fatal(DoErr)
	}
	defer resp.Body.Close()
	var cl int64 = 1024
	if resp.ContentLength > 0 {
		cl = resp.ContentLength
	}
	buf := make([]byte, cl)
	io.ReadFull(resp.Body, buf)

	respCmd := &http2proto.Cmd{}
	UnmarshalErr := proto.Unmarshal(buf, respCmd)
	if UnmarshalErr != nil {
		t.Fatal(UnmarshalErr)
	}
	t.Log(string(respCmd.ResInfo))
}

func TestProtoServer(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		buf := make([]byte, req.ContentLength)
		io.ReadFull(req.Body, buf)
		defer req.Body.Close()

		reqCmd := &http2proto.Cmd{}
		proto.Unmarshal(buf, reqCmd)
		reqCmd.ResInfo = reqCmd.GetArgInfo()

		respCmd := &http2proto.Cmd{}
		respCmd.ResInfo = reqCmd.ArgInfo
		respBytes, _ := proto.Marshal(respCmd)
		w.Write(respBytes)

	})
	server := &http.Server{
		Addr:    ":8899",
		Handler: mux,
	}
	if ListenAndServeTLSErr := server.ListenAndServeTLS("./default.pem", "./default.key"); ListenAndServeTLSErr != nil {
		t.Fatal(ListenAndServeTLSErr)
	}
}
