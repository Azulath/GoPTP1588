package main 

import (
	"fmt"
	"net"
)

func main() {
	name4 := "127.0.0.1"
	name6 := "0:0:0:0:0:0:0:1"
	
	addr4 := net.ParseIP(name4)
	addr6 := net.ParseIP(name6)
	
	if addr4 == nil {
		fmt.Println("Invalid IPv4 address")
	} else {
		fmt.Println("The IPv4 adress is:", addr4)
	}
	if addr6 == nil {
		fmt.Println("Invalid IPv4 address")
	} else {
		fmt.Println("The IPv6 adress is:", addr6)
	}
}

