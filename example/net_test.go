package main

import (
	"fmt"
	"net"
	"testing"
)

func TestCIDRFunc(t *testing.T) {
	ip, net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		fmt.Println("error is", err)
	}
	fmt.Println(ip)
	fmt.Println(net)
}
