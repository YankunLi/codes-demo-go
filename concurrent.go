package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a int
	runtime.GOMAXPROCS(4)
	go func() {
		for {
			a = 1
			runtime.Gosched()
			a = 2
			runtime.Gosched()
		}
	}()
	go func() {
		var b int
		for {
			b = a
			fmt.Println(b)
		}
	}()
	time.Sleep(1000 * time.Second)
}
