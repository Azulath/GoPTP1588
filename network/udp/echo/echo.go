package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("test")
	fmt.Println("Starting UDP-Echo-Server...")
	service := ":1201"
	
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
	
	n, addr, err := conn.ReadFromUDP(buff[0:])
	go PrintStudd()
	
	if err != nil {
		return
	}

	fmt.Printf("Received: %+v", string(buff[0:n]))
	_, err2 := conn.WriteToUDP(([]byte)(buff[0:n]), addr)
	if err2 != nil {
		return
	}
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error")
		os.Exit(1)
	}
}

func PrintStudd() {
	fmt.Println("go routine printening")
}
