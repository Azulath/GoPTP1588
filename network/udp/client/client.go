package main 

import (
	"net"
	"os"
	"fmt"
	"encoding/binary"
)

func main() {
	service := "127.0.0.1:1201"
	
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	//test1 := "yo ppl "
	//test2 := "rage here"
	x := new(Test)
	x.X = 6
	x.Y = 10
	x.S = "hiho"
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, x.X)
	
	written, err := conn.Write([]byte(x.X))
	checkError(err)
	fmt.Println("Bytes written:", written)
	
	var buff [512]byte
	read, err := conn.Read(buff[0:])
	checkError(err)
	fmt.Println("Bytes read:", read)
	
	fmt.Println(string(buff[0:read]))
	fmt.Println(buff[0:read])
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error")
		os.Exit(1)
	}
}

type Test struct {
	X uint16
	Y uint16
	S string
}

