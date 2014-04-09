package helper

import "fmt"

func ErrHandling (err error) {
	if err != nil {
		fmt.Println(err)
	}
}

