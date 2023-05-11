package main

import "fmt"

type A struct {
	b int
}

func main() {
	var a *A
	a = nil
	if a.b != 1 && a != nil {
		fmt.Println(a)
	}

	fmt.Println("vim-go")
}
