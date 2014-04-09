package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {
	service := "0.0.0.0:1210"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	
	shorts := readShorts(conn)
	fmt.Println(shorts)
	
	ints := utf16.Decode(shorts)
	fmt.Println(ints)
	
	str := string(ints)
	fmt.Println(str)
}

func readShorts(conn net.Conn) []uint16 {
	var buffer [512]byte
	
	// read everything into the buffer
	n, err := conn.Read(buffer[0:2])
	for true {
		m, err := conn.Read(buffer[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
	}
	
	checkError(err)
	var shorts []uint16
	shorts = make([]uint16, n/2)
	
	if buffer[0] == 0xff && buffer[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buffer[i]) << 8 + uint16(buffer[i+1])
		}
	} else if buffer[1] == 0xff && buffer[0] == 0xfe {
		// little endian
		for i := 2; i < n; i += 2 {
			shorts[i/2] = uint16(buffer[i+1]) << 8 + uint16(buffer[i])
		}
	} else {
		// unknown byte order
		fmt.Println("Unknown order")
	}
	return shorts
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
