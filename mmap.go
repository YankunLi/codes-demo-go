package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	map_file, _ := os.OpenFile("./homepic.png", os.O_RDONLY|syscall.O_DIRECT, 0755)

	mmap, err := syscall.Mmap(int(map_file.Fd()), 0, int(85), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(mmap))
}
