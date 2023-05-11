package main

import (
	"fmt"
	"os"
)

func main() {
	mode := os.FileMode(0777)
	fmt.Printf("2: %b\n", mode)
	fmt.Printf("d: %d\n", mode)
	fmt.Printf("o: %o\n", mode)
	fmt.Printf("is dir: %t\n", os.FileMode(os.ModeDir|mode).IsDir())
	fmt.Printf("is dir: %t\n", os.FileMode(mode).IsDir())
	fmt.Println("vim-go\n")
}
