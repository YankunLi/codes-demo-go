package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	rootPath := os.Args[1]
	sizeStr := os.Args[2]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return
	}

	//	cntStr := os.Args[3]
	//	cnt, err := strconv.Atoi(cntStr)
	//	if err != nil {
	//		return
	//	}

	wBuf := make([]byte, size)
	for j := 0; j < len(wBuf); j++ {
		wBuf[j] = uint8(rand.Int31())
	}
	//	rBuf := make([]byte, size)

	path := fmt.Sprintf("%s", rootPath)
	w, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	offset := 0
	for i := 0; i < size; i = i + 10 {
		n, err := w.Write(wBuf[offset:i])
		if err != nil {
			return
		}
		offset += n
		w.Sync()
	}

	w.Close()
	//w.Sync()
	fmt.Println("write down######")
	//	time.Sleep(30 * time.Second)
	//w.Close()
	//
	//	fmt.Println("close down")
	//	time.Sleep(30 * time.Second)
	//
	//	for i := 0; i < cnt; i++ {
	//		open, read := int64(0), int64(0)
	//		path := fmt.Sprintf("%s", rootPath)
	//
	//		t1 := time.Now().UnixNano()
	//		r, err := os.Open(path)
	//		if err != nil {
	//			fmt.Println(err)
	//			return
	//		}
	//		t2 := time.Now().UnixNano()
	//		open += t2 - t1
	//
	//		n, err := r.Read(rBuf)
	//		if err != nil || n != len(rBuf) {
	//			return
	//		}
	//
	//		t3 := time.Now().UnixNano()
	//
	//		read += t3 - t2
	//		r.Close()
	//		fmt.Printf("open=%d,read=%d \n", open, read)
	//		time.Sleep(10 * time.Second)
	//	}

}
