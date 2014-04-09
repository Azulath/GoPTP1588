package main 

import (
	"net"
	"os"
	"fmt"
)

func main() {
	name := "www.google.com"
	addr, err := net.ResolveIPAddr("ip", name)
	
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Resolved adress is", addr.String())
}

