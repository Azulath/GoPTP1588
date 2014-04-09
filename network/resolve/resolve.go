package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := "www.google.com"
	
	addr, err := net.ResolveIPAddr("ip6", name)
	// addr, err := net.ResolveIPAddr("ip4", name)
	// addr, err := net.ResolveIPAddr("ip", name)
	
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Resolved address is", addr.String())
}

