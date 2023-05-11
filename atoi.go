package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s = byte('0')
	f1, _ := os.OpenFile("test.file", os.O_CREATE|os.O_RDWR, 0777)
	f2, _ := os.OpenFile("test.file", os.O_RDWR, 0777)

	var buf = make([]byte, 1)
	var buf1 = make([]byte, 1)
	buf[0] = s
	f1.Write(buf)
	f2.Read(buf1)

	var s1 = buf1[0]
	var s2 byte
	fmt.Printf("sd %v\n", byte(int8(1)))

	var str = string(s1)
	fmt.Printf("%v  %v\n", s1, str)
	fmt.Printf("%v\n", string(s2))
	if v, err := strconv.Atoi(string(s1)); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(v)
	}
}
