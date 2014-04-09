package main

import (
	"fmt"
	"net"
	"os"
)

func main () {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	
	for {
		conn, err := listener.Accept()
		
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}

func handleClient (conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	
	var buff [512]byte
	
	for {
		// read upto 512 bytes
		n, err := conn.Read(buff[0:])
		
		if err != nil {
			return
		}
		
		fmt.Println("Written", string(buff[0:n]))
		
		// write the n bytes read
		_, err2 := conn.Write(buff[0:n])
		
		if err2 != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

