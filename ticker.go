package main

import (
	"fmt"
	"time"
)

func main() {
	tk := time.NewTicker(1 * time.Second)
	defer tk.Stop()
	tk1 := time.NewTicker(2 * time.Second)
	defer tk1.Stop()
	i := <-tk1.C
	fmt.Println("Receive a message2: ", i.UTC(), "Now: ", time.Now().UTC())

	time.Sleep(10 * time.Second)
	for {
		select {
		case i := <-tk.C:
			fmt.Println("Receive a message1: ", i.UTC(), "Now: ", time.Now().UTC())
		}
	}
}
