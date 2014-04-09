package main 

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("test")
	fmt.Println("Starting UDP-Dayime-Server")
	service := ":1200"
	
	conn, err := net.ListenPacket("udp", service)
	checkError(err)
	
	for {
		handleClient(conn)
	}
}

func handleClient(conn net.PacketConn) {
	var buff [512]byte
	
	_, addr, err := conn.ReadFrom(buff[0:])
	
	if err != nil {
		return
	}
	
	daytime := time.Now().String()
	conn.WriteTo([]byte(daytime), addr)
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error", err.Error())
		os.Exit(1)
	}
}

