package main 

import (
	"fmt"
	"net"
	"os"
	_"io/ioutil"
)

func main() {
	//to ask daytime server
	//nameAndService := "www.google.com:80"
	nameAndService := "127.0.0.1:1200"
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", nameAndService)
	checkError(err)
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	
	written, err := conn.Write([]byte("yo ppl rage here!"))
	checkError(err)
	
	//result, err := ioutil.ReadAll(conn)
	//checkError(err)
	
	var buff [512]byte
	read, err := conn.Read(buff[0:])
	checkError(err)
	
	fmt.Println("Bytes written:", written)
	//fmt.Println(string(result))
	
	fmt.Println("Bytes read:", read)
	fmt.Println(string(buff[0:]))
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

