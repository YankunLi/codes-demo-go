package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("./ll")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Dir is not exist, make it\n")
			os.MkdirAll("./ll/ll", 0775)
		} else {
			fmt.Printf("Stat directory err: %s", err)
		}
	}
}
