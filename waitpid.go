package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	pid = flag.Int64("p", -1, "process id")
)

func main() {
	flag.Parse()
	p, err := os.FindProcess(int(*pid))
	if err != nil {
		fmt.Printf("os findProcess id(%d) fail, err: %s", *pid, err)
		return
	}
	if _, err = p.Wait(); err != nil {
		fmt.Printf("Wait process Pid(%d) fail, err: %s", *pid, err)
	}

	return
}
