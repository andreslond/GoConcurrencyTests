package main

import (
	"fmt"
	"math"
	"time"
)

var (
	elevator1 = Elevator{name: "elevator1", location: 1, callStack: make(chan int64, 5)}
)

type Elevator struct {
	name      string
	location  int64
	callStack chan int64
}

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	//wg.Done()
	fmt.Println("Elevator works...")

	go CallElevator(5)
	//go CallElevator(7)

	ElevatorController()
	//wg.Wait()
}

func CallElevator(destinationFloor int64) {
	elevator1.callStack <- destinationFloor
	//currentFloor := elevator1.location
	//flyTime(currentFloor, destinationFloor)
}

func ElevatorController() {
	if len(elevator1.callStack) > 0 {
		fmt.Println(">0")
	}

	for {
		select {
		case destinationFloor := <-elevator1.callStack:
			currentFloor := elevator1.location
			flyTime(currentFloor, destinationFloor)
		}
	}
}

func flyTime(currentFloor int64, destinationFloor int64) {
	floorsToMove := time.Duration(math.Abs(float64(destinationFloor - currentFloor)))
	elevator1.location = 0
	fmt.Println(fmt.Sprintf("going from: %d to %d", currentFloor, destinationFloor))
	time.Sleep(floorsToMove * time.Second)
	fmt.Println(fmt.Sprintf("new location: %d", currentFloor, destinationFloor))
	elevator1.location = destinationFloor
}
