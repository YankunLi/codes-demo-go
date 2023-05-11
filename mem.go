package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: allocBuffer,
}

const (
	DefaultFileCacheSize = 4 * 1024 * 1024
)

func allocBuffer() interface{} {
	b := make([]byte, DefaultFileCacheSize)
	return &b
}

func getBuffer() *[]byte {
	b := bufPool.Get().(*[]byte)
	return b
}

func putBuffer(b *[]byte) {
	bufPool.Put(b)
}

func main() {
	//var mm = make([]uintptr, 0, 1024)
	var mm = make([]*[]byte, 0, 1024)
	var sum int64
	for i := 0; i < 10240; i++ {
		tmp := getBuffer()
		for _, it := range *tmp {
			if it == '1' {
			}
		}
		sum += int64(len(*tmp))
		//mm = append(mm, uintptr(unsafe.Pointer(tmp)))
		mm = append(mm, tmp)
	}

	time.Sleep(30 * time.Second)
	fmt.Println("putbuffer")
	fmt.Println("size of memory allocated is: ", sum)
	for _, tmp := range mm {
		for _, it := range *tmp {
			if it == '1' {
			}
		}
		//putBuffer((*[]byte)(unsafe.Pointer(tmp)))
	}
	runtime.GC()

	mm = nil
	time.Sleep(3 * time.Second)
	runtime.GC()
	fmt.Println("vim-go")
	time.Sleep(300 * time.Second)
}
