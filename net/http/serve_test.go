package http

import (
	"io"
	"log"
	"net"
	"net/http"
	"testing"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "hellWorld server")

}

func TestServe(t *testing.T)  {
	http.HandleFunc("/hello", HelloServer)

	ln, lerr := net.Listen("tcp", ":8888")
	if lerr != nil {
		log.Fatal("listen:", lerr)
	}

	serr := http.Serve(ln, nil)
	if lerr != nil {
		log.Fatal("serve: ", serr)
	}
}
