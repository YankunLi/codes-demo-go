package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

func main() {
	num32 := 0xffffffff

	var t time.Timer
	fmt.Printf("%x %x %d\n", num32+1, math.MaxUint32+1, unsafe.Sizeof(t))
}
