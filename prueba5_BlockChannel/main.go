package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	//c <- "three"

	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}
