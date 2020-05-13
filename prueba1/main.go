package main

import (
	"fmt"
	"time"
)

func main() {
	i:= 0
	for ; i < 10000; i++ {
		go func(number int) {
			fmt.Println(fmt.Sprintf("numero: %v",number))
		}(i)
	}
	time.Sleep(5 * time.Second)

}
