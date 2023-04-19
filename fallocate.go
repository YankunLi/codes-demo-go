package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	nos "bj58.com/nebula2/store/os"
)

var (
	size = flag.Int64("s", 4096, "filesize")
	file = flag.String("n", "test", "filename")
)

func main() {
	flag.Parse()

	fd, err := os.OpenFile(*file, os.O_RDWR|syscall.O_NOATIME|syscall.O_DIRECT, 0664)
	err = nos.Fallocate(fd, nos.FALLOC_FL_KEEP_SIZE, 0, *size)
	if err != nil {
		fmt.Println(err)
	}
}
