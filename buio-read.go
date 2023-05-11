package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	filename  = flag.String("f", "", "filename")
	chunksize = flag.Int("c", 1, "chunk size")
)

func main() {
	flag.Parse()
	rfile, err := os.OpenFile(*filename, os.O_RDONLY|os.O_SYNC, 0664)
	if err != nil {
		fmt.Errorf("open error")
		return
	}

	var buf = make([]byte, 200*1024)

	start := time.Now()
	rd := bufio.NewReaderSize(rfile, *chunksize*1024*1024)
	for {
		_, err := rd.Read(buf)
		if err != nil {
			break
		}
	}
	fmt.Printf("cost: %d\n", time.Since(start).Nanoseconds()/1e6)
}
