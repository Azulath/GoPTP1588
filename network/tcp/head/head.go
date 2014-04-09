package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	service := "www.google.com:80"
	
	conn, err := net.Dial("tcp", service)
	checkError(err)
	
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	
	//resultOld, err1 := readFully(conn)
	//checkError(err1)
	
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	
	//fmt.Println(string(resultOld))
	//fmt.Println("=====================================")
	fmt.Println(string(result))
	
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	
	result := bytes.NewBuffer(nil)
	var buff [512]byte
	
	for {
		n, err := conn.Read(buff[0:])
		result.Write(buff[0:n])
		
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
