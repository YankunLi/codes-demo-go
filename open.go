package main

import (
	"fmt"
	"io"
	"os"

	"log"
)

func main() {
	fd, err := os.OpenFile("./homepic.png", os.O_RDWR, 0664)
	if err != nil {
		log.Panicf("open %v", err)
		return
	}

	info, err := fd.Stat()
	if err != nil {
		log.Panicf("stat %v", err)
		return
	}

	count := 0
	buf := make([]byte, 100)
	for {
		//	time.Sleep(1 * time.Second)
		fd.Seek(0, 0)
		n, err := fd.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		if n == 0 {
			//		break
			continue
		}
		count += n
		fmt.Printf("read :%d size: %d \n", count, info.Size())
	}
	fmt.Printf("content: %s\n", string(buf))

	//	fd1, err := os.OpenFile("./homepic.png", os.O_RDWR, 0664)
	//	mmap, err := syscall.Mmap(int(fd1.Fd()), 0, int(100), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(string(mmap))

}
