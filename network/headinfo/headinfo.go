package main 

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	nameAndService := "www.google.com:80"
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", nameAndService)
	checkError(err)
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	
	num, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	fmt.Println("Num:", num)
	
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	
	fmt.Println(string(result))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

