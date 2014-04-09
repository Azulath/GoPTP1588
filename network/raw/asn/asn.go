package main 

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	mdata, err := asn1.Marshal(13)
	checkError(err)
	
	fmt.Println("ASN.1 Marshal:", mdata)
	
	var n int
	x, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)
	
	fmt.Println("Value x", x)
	fmt.Println("After marshal/unmarshal:", n)
}

func checkError (err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}