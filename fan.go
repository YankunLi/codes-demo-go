package main

import "fmt"

const (
	mask = iota
)

func main() {
	//var j uint8 = 1
	var i uint8 = 2
	fmt.Printf(" %b\n", i)
	fmt.Printf(" %b\n", ^i)
	fmt.Printf(" %b\n", i&(^i))
	cc := make(chan int, 5)
	cc <- 1
	cc <- 2
	cc <- 3
	close(cc)
	for {
		select {
		case x, ok := <-cc:
			if ok {
				fmt.Println(x)
			} else {
				return
			}
		}
	}

}
