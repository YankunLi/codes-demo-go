package main

import (
	"fmt"
	"unsafe"
)

type inode struct {
	mode uint32
	id   uint64
	uid  uint32
	size uint64
	gid  uint32
}

func main() {
	var in inode
	fmt.Printf("algin: %d\n", unsafe.Sizeof(in))
}
