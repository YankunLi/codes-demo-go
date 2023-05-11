package main

import (
	"fmt"
	"runtime"
)

func main() {
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	fmt.Println(*m)
	fmt.Printf("heap %d\n", m.HeapSys)

}
