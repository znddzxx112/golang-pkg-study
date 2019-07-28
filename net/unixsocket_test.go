package net

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestDialUnix(t *testing.T) {
	runixaddr, ResolveUnixAddrErr := net.ResolveUnixAddr("unix", "./minids.sock")
	if ResolveUnixAddrErr != nil {
		t.Fatal(ResolveUnixAddrErr)
	}
	lunixaddr, ResolveUnixAddrErr := net.ResolveUnixAddr("unix", "./minidscli.sock")
	if ResolveUnixAddrErr != nil {
		t.Fatal(ResolveUnixAddrErr)
	}
	conn ,DialUnixErr :=  net.DialUnix("unix", lunixaddr, runixaddr)
	if DialUnixErr != nil {
		t.Fatal(DialUnixErr)
	}
	defer conn.Close()
	transCoder := NewSTransCoder(conn)
	foo := []byte("foo")
	transCoder.Send(&foo)
	bar,_ := transCoder.Receive()
	t.Log(string(bar))

	os.Remove("./minidscli.sock")


}

func TestListenUnix(t *testing.T) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func(c chan os.Signal) {
		for s := range c{
			switch s  {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				os.Remove("minids.sock")
				os.Exit(0)
			case syscall.SIGUSR1, syscall.SIGUSR2:
				fmt.Println("restart ...")
				fmt.Println("done.")
			default:
				fmt.Println("not catch signal.")
			}
		}
	}(c)

	lunixaddr, ResolveUnixAddrErr := net.ResolveUnixAddr("unix", "./minids.sock")
	if ResolveUnixAddrErr != nil {
		t.Fatal(ResolveUnixAddrErr)
	}
	ln, ListenUnixErr := net.ListenUnix("unix", lunixaddr)
	if ListenUnixErr != nil {
		t.Fatal(ListenUnixErr)
	}
	defer ln.Close()

	for {
		conn, AcceptErr := ln.Accept()
		if AcceptErr != nil {
			t.Error(AcceptErr)
			break
		}
		go func(c net.Conn) {
			defer func() {
				c.Close()
				t.Log("minids close")
			}()
			// echo
			transCoder := NewSTransCoder(c)
			body, ReceiveErr := transCoder.Receive()
			if ReceiveErr != nil {
				if ReceiveErr == io.EOF {
					return
				}
			}
			transCoder.Send(&body)
			cmd := bytes.TrimRight(body, "\r\n")
			if bytes.Equal(cmd , []byte("exit")) {
				return
			}
		}(conn)
	}


}
