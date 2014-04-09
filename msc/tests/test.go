package main

import (
	"fmt"
)

func main(){
	i := 4
	s := "bla " + fmt.Sprintf("%v", i) + " end"
	fmt.Println(s)
	str := "8080"
	str += ":"
	fmt.Println(str)
}


