package main

import (
	"flag"
	"fmt"
	"os"

	"log"
)

var (
	dir = flag.String("d", "./", "directory")
)

func main() {
	for i := 0; i < 203; i++ {
		fd, err := os.OpenFile(fmt.Sprintf("%s/data_%d", *dir, i), os.O_RDONLY|0, 0664)
		if err != nil {
			log.Panicf("open %v", err)
			return
		}

		info, err := fd.Stat()
		if err != nil {
			log.Panicf("stat %v", err)
			return
		}

		size := info.Size() - 4096
		if size < 40960000 {
			continue
		}
		buffer := make([]byte, 4096*4096)
		for {
			_, err := fd.Read(buffer)
			if err != nil {
				break
			}
		}
		fd.Close()
	}
	fmt.Println("vim-go")
}
