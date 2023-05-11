package main

import "fmt"

func main() {
	b := 0
	c := 1
	if c == 0 || b == 0 || c%b == 0 {
		fmt.Println("invalid parameter")
		return
	}
	a := 5 % b
	fmt.Println(a)
}
