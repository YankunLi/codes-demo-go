package main

import "fmt"

func main() {
	var a = false
	var b = true
	c := a || b
	fmt.Println(c)
	c = b || a
	fmt.Println(c)
}
