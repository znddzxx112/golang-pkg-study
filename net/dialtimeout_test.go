package net

import (
	"net"
	"testing"
	"time"
)

func TestDialTimeout(t *testing.T) {
	conn, DialTimeoutErr := net.DialTimeout("tcp", "www.163.com:81", time.Second*2)
	if DialTimeoutErr != nil {
		t.Fatal(DialTimeoutErr)
	}
	defer conn.Close()
	time.Sleep(time.Second * 10)
}
