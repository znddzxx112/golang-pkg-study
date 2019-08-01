package http

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "hellWorld server")

}

func TestServe(t *testing.T) {
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

func GetHelloServer(w http.ResponseWriter, req *http.Request) {

	log.Println(req.ContentLength)
	log.Println(req.URL)
	log.Println(req.RequestURI)
	if req.ParseForm() == nil {
		log.Println(req.FormValue("foo"))
	}
	log.Println(req.Header.Get("VH"))
	log.Println(req.Header.Get("VH1"))
	log.Println(req.Cookie("cfoo"))
	cok := &http.Cookie{
		Name:     "wfoo",
		Value:    "wbar",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cok)
	io.WriteString(w, req.Host)
	io.WriteString(w, req.Header.Get(http.CanonicalHeaderKey("user-agent")))

}

// get cookie
func TestServerGet(t *testing.T) {
	http.HandleFunc("/gethello", GetHelloServer)

	ln, lerr := net.Listen("tcp", ":8889")
	if lerr != nil {
		log.Fatal("listen:", lerr)
	}

	serr := http.Serve(ln, nil)
	if lerr != nil {
		log.Fatal("serve: ", serr)
	}
}


func PostHelloServer(w http.ResponseWriter, req *http.Request) {

	if req.ParseForm() == nil {
		log.Println(req.PostForm)
	}
	io.WriteString(w, req.PostFormValue("foo"))
}

// post
func TestServerPost(t *testing.T)  {
	http.HandleFunc("/posthello", PostHelloServer)

	ln, lerr := net.Listen("tcp", ":8890")
	if lerr != nil {
		log.Fatal("listen:", lerr)
	}

	serr := http.Serve(ln, nil)
	if lerr != nil {
		log.Fatal("serve: ", serr)
	}
}

func PostJsonHelloServer(w http.ResponseWriter, req *http.Request) {

	var buflen int64 = 4096
	if req.ContentLength > 0 {
		buflen = req.ContentLength
	}
	buf := make([]byte, buflen)
	io.ReadFull(req.Body, buf)
	hellojson := &Hellojson{}
	var UnmarshalErr error
	if UnmarshalErr = json.Unmarshal(buf, hellojson);UnmarshalErr != nil {
		io.WriteString(w, "error")
		return
	}
	io.WriteString(w, hellojson.Name)



}

func TestServerPostJson(t *testing.T) {
	http.HandleFunc("/posthello", PostJsonHelloServer)

	ln, lerr := net.Listen("tcp", ":8890")
	if lerr != nil {
		log.Fatal("listen:", lerr)
	}

	serr := http.Serve(ln, nil)
	if lerr != nil {
		log.Fatal("serve: ", serr)
	}
}

func TestNewServer(t *testing.T)  {
	handle := http.NewServeMux()
	handle.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err == nil {
			writer.Header().Set("Content-Type", "application/json")
			io.WriteString(writer, "hellofoo")
		}

	})

	server := &http.Server{
		Addr:         ":8088",
		Handler:      handle,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
	}
	sErr := server.ListenAndServe()
	defer server.Close()
	if sErr != nil {
		t.Errorf("%q", sErr)
	}
}
