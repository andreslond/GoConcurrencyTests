package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("")


	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 15; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
