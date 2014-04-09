package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hello\n")

	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for i := 0; i < hashSize; i += 4 {
		var val uint32
		val = uint32(hashValue[i])<<24 +
			uint32(hashValue[i+1])<<16 +
			uint32(hashValue[i+2])<<8 +
			uint32(hashValue[i+3])
		fmt.Printf("%x", val)
	}
	
	fmt.Println()
}
