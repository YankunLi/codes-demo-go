package main

import "fmt"

func main() {
	var ss = []int{0, 1, 2, 3, 4}
	for i, v := range ss {
		if v == 4 {
			fmt.Println("index: ", i, "val: ", v)
			ss = append(ss[:i], ss[i+1:]...)
		}
	}
	fmt.Println(ss)
	fmt.Println(len(ss), " ", ss[4:])
}
