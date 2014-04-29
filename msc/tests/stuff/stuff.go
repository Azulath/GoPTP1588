package stuff

import (
	"time"
	"fmt"
)

var (
	tCar Car
)

func TestCar(c Car, counter int, sleep time.Duration) {
	tCar = c
	time.Sleep(sleep * time.Millisecond)
	fmt.Println(tCar, counter)
}

type Car struct {
	Id int
	//name string
}
