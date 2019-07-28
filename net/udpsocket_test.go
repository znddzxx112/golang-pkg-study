package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestListenUdp(t *testing.T) {
	ludpaddr, ResolveUDPAddrErr := net.ResolveUDPAddr("udp", ":8081")
	if ResolveUDPAddrErr != nil {
		t.Fatal(ResolveUDPAddrErr)
		return
	}
	udpconn, ListenUDPErr := net.ListenUDP("udp", ludpaddr)
	if ListenUDPErr != nil {
		t.Fatal(ListenUDPErr)
		return
	}
	defer udpconn.Close()

	rbuf := make([]byte, 1500)
	readAll := 0

	rn, ReadErr := udpconn.Read(rbuf)
	if ReadErr != nil {
		if ReadErr == io.EOF {
			return
		}
		return
	}
	readAll += rn

	fmt.Println(string(rbuf[:readAll]))

	udpconn.Write(rbuf[:readAll])

}

func TestDialUdp(t *testing.T) {
	conn, _ := net.Dial("udp", "127.0.0.1:8081")
	conn.Write([]byte("hello foo bar"))
	conn.Close()
}
