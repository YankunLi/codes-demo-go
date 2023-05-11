package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 1
	c <- 4
	c <- 3
	i, isClose := <-c
	if isClose {
		fmt.Printf("Get Number: %d from channel, channel is not closed\n", i)
	}
	var b bool
	fmt.Println(b)
	close(c)
	select {
	case c <- 1:
		fmt.Println("Put into channel")
	default:
		fmt.Println("default")
	}
	//	if isClose {
	//		fmt.Printf("Get Number: %d from channel, channel has been closed\n", i)
	//	}
}
