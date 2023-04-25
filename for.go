package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		var s = 1
		s += 2
		fmt.Print(s)
	}
}
