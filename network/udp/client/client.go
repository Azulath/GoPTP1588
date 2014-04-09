package main 

import (
	"net"
	"os"
	"fmt"
)

func main() {
	service := "127.0.0.1:1201"
	
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	
	written, err := conn.Write([]byte("yo ppl rage here!"))
	checkError(err)
	fmt.Println("Bytes written:", written)
	
	var buff [512]byte
	read, err := conn.Read(buff[0:])
	checkError(err)
	fmt.Println("Bytes read:", read)
	
	fmt.Println(string(buff[0:read]))
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error")
		os.Exit(1)
	}
}

