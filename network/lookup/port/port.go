package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	networkType := "udp"
	service := "http"
	
	port, err := net.LookupPort(networkType, service)
	
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Service Port", port)
}

