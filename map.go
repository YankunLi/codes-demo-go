package main

import "fmt"

func main() {
	ll := make(map[int][]int, 5)
	s := ll[1]
	fmt.Printf("map len:%d\n", len(ll))
	s = append(s, 1)
	ll[1] = s
	fmt.Println(ll)
}
