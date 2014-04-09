package main 

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting UDP-Dayime-Server")
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buff [512]byte
	
	_, addr, err := conn.ReadFromUDP(buff[0:])
	
	if err != nil {
		return
	}
	
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error", err.Error())
		os.Exit(1)
	}
}

