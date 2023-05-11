package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
	"unsafe"

	"log"
)

const (
	AlignSize = 512
)

// 在 block 这个字节数组首地址，往后找，找到符合 AlignSize 对齐的地址，并返回
// 这里用到位操作，速度很快；
func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

// 分配 BlockSize 大小的内存块
// 地址按照 512 对齐
func AlignedBlock(BlockSize int) []byte {
	// 分配一个，分配大小比实际需要的稍大
	block := make([]byte, BlockSize+AlignSize)

	// 计算这个 block 内存块往后多少偏移，地址才能对齐到 512
	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	// 偏移指定位置，生成一个新的 block，这个 block 将满足地址对齐 512；
	block = block[offset : offset+BlockSize]
	if BlockSize != 0 {
		// 最后做一次校验
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}

	return block
}

var (
	file = flag.String("f", "test", "file name")
)

func main() {
	flag.Parse()
	fd, err := os.OpenFile(*file, os.O_RDWR|syscall.O_NOATIME|syscall.O_DIRECT, 0664)
	if err != nil {
		log.Panicf("open %v %s", err, *file)
		return
	}

	buf := AlignedBlock(4096) //make([]byte, 12)
	start := time.Now()
	_, err = fd.Read(buf[:])
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Printf("costTime: %d\n", time.Since(start).Nanoseconds())

	//	fd1, err := os.OpenFile("./homepic.png", os.O_RDWR, 0664)
	//	mmap, err := syscall.Mmap(int(fd1.Fd()), 0, int(100), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(string(mmap))

}
