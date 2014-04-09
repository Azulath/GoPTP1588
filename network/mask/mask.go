package main 

import (
	"net"
	"fmt"
	"os"
)

func main() {
	name := "192.168.0.1"
	
	addr := net.ParseIP(name)
	
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	
	fmt.Println("Adress is", addr.String(),
				"\nDefault mask length is", bits,
				"\nLeading ones count is", ones,
				"\nMask is (hex)", mask.String(),
				"\nNetwork is", network.String())
}

