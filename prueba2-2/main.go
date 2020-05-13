package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	now := time.Now()

	//1st goroutine
	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(100 * time.Millisecond)
		}
		numGR := runtime.NumGoroutine()
		fmt.Println(numGR)
	}()

	//2nd goroutine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
		}
		numGR := runtime.NumGoroutine()
		fmt.Println(numGR)
	}()
	//3 goroutine
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
		}
		numGR := runtime.NumGoroutine()
		fmt.Println(numGR)
	}()

	go longAndHeavyTask()

	elapsed := time.Since(now)
	fmt.Println("tiempo: " + elapsed.String())
	time.Sleep(3 * time.Second)
}

func longAndHeavyTask() {
	time.Sleep(5 * time.Second)
}
