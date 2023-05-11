package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Rename("/opt/ll.txt", "/data00/ll.txt")
	if err != nil {
		fmt.Println(err)
	}
}
