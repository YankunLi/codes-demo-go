package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%b\n", -3>>1)
	fmt.Printf("%b\n", -1<<3)
	fmt.Printf("%b\n", -1)
	fmt.Printf("%b\n", -1^(-1<<3))

}
