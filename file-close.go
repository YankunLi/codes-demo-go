package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("./2.go")
	if err != nil {
		fmt.Errorf("open fail, err: %v", err)
		return
	}
	fd.Close()
	fd.Close()
}
