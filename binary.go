package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

func main() {
	val := make([]byte, 8)
	val[7] = 1
	val[6] = 1
	start := time.Now()
	for i := 1; i < 10; i++ {
		s := make([]byte, 4096*1024)
		s[0] = 0
	}
	size := binary.BigEndian.Uint64(val)
	elapsed := time.Since(start)
	fmt.Printf("elpased: %d, size: %d", elapsed.Nanoseconds(), size)
}
