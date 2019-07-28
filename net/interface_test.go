package net

import (
	"net"
	"testing"
)

func TestJoinHostPort(t *testing.T) {
	host := net.JoinHostPort("127.0.0.1", "8080")
	t.Log(host)
}

func TestSpiltHostPort(t *testing.T) {
	host, port, _ := net.SplitHostPort("www.baidu.com:8081")
	t.Log(host, port)
}

func TestInterfaceAddrs(t *testing.T) {
	addrs, InterfaceAddrsErr := net.InterfaceAddrs()
	if InterfaceAddrsErr != nil {
		t.Error(InterfaceAddrsErr.Error())
	}
	for _, addr := range addrs {
		t.Logf("%s %s\n", addr.Network(), addr.String())
	}
}

func TestInterfaceByName(t *testing.T) {
	interfaces, InterfaceByNameErr := net.InterfaceByName("wlp4s0")
	if InterfaceByNameErr != nil {
		t.Fatal(InterfaceByNameErr)
	}
	t.Log(interfaces.Name, interfaces.Flags, interfaces.HardwareAddr.String(), interfaces.MTU, interfaces.Index)
	addrs,_ := interfaces.Addrs()
	for _, addr := range addrs {
		t.Logf("%s %s\n", addr.Network(), addr.String())
	}
}

func TestInterfaceByIndex(t *testing.T)  {
	infaces,_ := net.InterfaceByIndex(3)
	t.Log(infaces.Name)
	addrs,_ := infaces.Addrs()
	for _, addr := range addrs {
		t.Logf("%s %s\n", addr.Network(), addr.String())
	}
}

func TestParseCIDR(t *testing.T)  {
	ip, ipnet, _ := net.ParseCIDR("192.168.43.148/24")
	t.Log(ip.String())
	t.Log(ipnet.String())
}