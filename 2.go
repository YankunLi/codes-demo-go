package main

import "fmt"

func main() {
	fmt.Printf("%b\n", ^(-1 << 34))
	fmt.Printf("%b\n", -1^(-1<<34))
	fmt.Printf("%b\n", -1^(-1<<3))
	fmt.Printf("%b\n", ^(1 << 34))
	fmt.Printf("%d\n", ^(int8(3)))
	fmt.Printf("%d\n", ^3)
	fmt.Println("vim-go")
}
