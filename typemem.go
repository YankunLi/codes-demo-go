package main

import (
	"fmt"
	"unsafe"
)

type SS struct {
	a       string
	b       string
	c       string
	d       string
	e       string
	f       string
	g       string
	content [1024]byte
}

type AA struct {
	items *SS
}

func main() {
	var s []byte
	var p *int
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(p))

}
