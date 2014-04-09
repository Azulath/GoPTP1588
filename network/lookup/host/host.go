package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := "www.google.com"
	c, err:= net.LookupCNAME("www.heise.de")
	
	if err != nil {
		fmt.Println("Cname Error:", err.Error())
	} else {
		fmt.Println("Cname:", c)
	}
	
	addrs, err := net.LookupHost(name)
	
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	
	for _, s := range addrs {
		fmt.Println(s)
	}
}

