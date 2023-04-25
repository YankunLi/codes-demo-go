package main

import "fmt"

const (
	l    = 1 << iota
	sync = 1 << iota
	ack
	ack1
)

func main() {
	fmt.Printf("sync: %d ack: %d l: %d ack1: %d \n", sync, ack, l, ack1)
}
