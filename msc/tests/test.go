package main

import (
	"fmt"
)

func main() {
	i := 4
	s := "bla " + fmt.Sprintf("%v", i) + " end"
	fmt.Println(s)
	str := "8080"
	str += ":"
	fmt.Println(str)
	var iTest uint32 = 3231748232
	fmt.Println("test")
	fmt.Println(iTest)
	first := uint8((iTest >> 24) & 0xc0ff)
	second := uint8(iTest >> 16)
	third := uint8(iTest >> 8)
	fourth := uint8(iTest)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(0xff)
}


