package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for l := 0; l < 15; l++ {
			c <- l
			time.Sleep(time.Second)
		}
		fmt.Println("envío: ", runtime.NumGoroutine())
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Recibo: ", runtime.NumGoroutine())
}
