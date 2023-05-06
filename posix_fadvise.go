// +build linux

package main

/*
#define _FILE_OFFSET_BITS 64
#define _XOPEN_SOURCE 600
#include <unistd.h>
#include <fcntl.h>
*/

import "C"

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"log"
	"syscall"
)

const (
	POSIX_FADV_NORMAL     = int(C.POSIX_FADV_NORMAL)
	POSIX_FADV_SEQUENTIAL = int(C.POSIX_FADV_SEQUENTIAL)
	POSIX_FADV_RANDOM     = int(C.POSIX_FADV_RANDOM)
	POSIX_FADV_NOREUSE    = int(C.POSIX_FADV_NOREUSE)
	POSIX_FADV_WILLNEED   = int(C.POSIX_FADV_WILLNEED)
	POSIX_FADV_DONTNEED   = int(C.POSIX_FADV_DONTNEED)
)

func Fadvise(fd uintptr, off int64, size int64, advise int) (err error) {
	var errno int
	if errno = int(C.posix_fadvise(C.int(fd), C.__off64_t(off), C.__off64_t(size), C.int(advise))); errno != 0 {
		err = syscall.Errno(errno)
	}
	return
}

func main() {
	fd, err := os.OpenFile("./3gfile", os.O_RDONLY|os.O_CREATE|0, 0664)
	if err != nil {
		log.Panicf("open %v", err)
		return
	}
	_, err = os.OpenFile("./3gfile", os.O_WRONLY|0, 0664)
	if err != nil {
		log.Panicf("open %v", err)
		return
	}

	info, err := fd.Stat()
	if err != nil {
		log.Panicf("stat %v", err)
		return
	}
	err = Fadvise(fd.Fd(), 0, info.Size(), POSIX_FADV_SEQUENTIAL)
	if err != nil {
		log.Panicf("Fadvise %v", err)
		return
	}
	var counts int64
	go func() {
		upc := int64(0)
		tick := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-tick.C:
				log.Printf("IOPS: %d", counts-upc)
			}
			upc = counts
		}
	}()
	size := info.Size() - 4096
	for {
		buffer := make([]byte, 4096*1024)
		counts++
		off := rand.Int63n(size)
		_, err := fd.ReadAt(buffer, off)
		if err != nil {
			log.Panicf("read %v", err)
			return
		}
	}
	fmt.Println("vim-go")
}
