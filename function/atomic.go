package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

func main() {
	var right bool
	var s = true
	var num = uint32(3)
	fmt.Printf("%d\n", num)
	atomic.AddUint32(&num, ^uint32(0))
	fmt.Printf("%d\n", num)
	fmt.Println(^uint32(0))
	pp := unsafe.Pointer(&right)
	atomic.StorePointer(&pp, unsafe.Pointer(&s))

	if right {
		fmt.Println("we right")
	}
}
