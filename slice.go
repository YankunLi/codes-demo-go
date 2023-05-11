package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println(len(s))
	s = s[:5]
	for i, _ := range s {
		fmt.Println(i)
	}
	fmt.Println(len(s))
	var ll []byte = nil
	for _, v := range ll {
		fmt.Println(v)
	}
}
