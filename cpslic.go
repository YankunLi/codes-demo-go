package main

import "fmt"

func main() {
	s := make([]int, 2)
	ss := make([]int, 2)
	s[0] = 1
	s[1] = 2
	copy(ss, s[0:0])
	fmt.Printf("%v\n", s[0:0])
	fmt.Printf("%v\n", ss)
}
