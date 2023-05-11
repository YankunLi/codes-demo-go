package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t1 := strings.Replace(time.Now().Format("2006-01-02 15:04:05"), " :", "_", 1)
	//	t2 := strings.Replace(t1, ":", "_", -1)

	fmt.Printf("Time: %s, time: %s\n", time.Now().String(), t1)
	t2 := time.Now()
	time.Sleep(1 * time.Second)
	t3 := time.Now()
	fmt.Println(t2.Sub(t3).Milliseconds())
}
