package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	j := 4
	for {
		switch i {
		case 1:
			j++
			fmt.Printf("j: %d\n", j)
			time.Sleep(1 * time.Second)
			continue
			if j > 10 {
				fmt.Printf("run break\n")
				break
			}
		}
		fmt.Printf("for end\n")
	}
}
