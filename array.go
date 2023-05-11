package main

import (
	"fmt"

	"bj58.com/nebula2/util"
)

func main() {
	var ar [8]byte
	ss, _ := util.GenUUID()
	copy(ar[:], ss)
	fmt.Println(string(ar[:]))

}
