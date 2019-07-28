package net

import (
	"net"
	"testing"
)

func TestLookupAddr(t *testing.T)  {
	names,_:=net.LookupAddr("173.194.127.81")
	t.Log(names)
}

func TestLookupHost(t *testing.T)  {
	addrs , _ := net.LookupHost("t.10jqka.com.cn")
	t.Log(addrs)
}

func TestLookupCname(t *testing.T)  {
	cname, _ :=net.LookupCNAME("www.10jqka.com.cn")
	t.Log(cname)
}

func TestLookupPort(t *testing.T)  {
	port, _ := net.LookupPort("tcp", "http")
	t.Log(port) // 80
	
}

func TestLookupTxt(t *testing.T)  {
	txt, _ := net.LookupTXT("192.168.43.1")
	t.Log(txt)
}