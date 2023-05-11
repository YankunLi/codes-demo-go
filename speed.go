package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"syscall"
	"time"
)

var (
	filename    = flag.String("f", "", "filename")
	traffic     = flag.Int("t", 10, "traffic ")
	enablePause = flag.Bool("p", false, "enable pause")
)

func main() {
	flag.Parse()
	fd, err := os.OpenFile(*filename, syscall.O_DIRECT|os.O_RDONLY, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	speed := *traffic * 1024 * 1024
	timesPerSecond := math.Ceil(float64(speed) / float64(64*1024))

	var sum int

	buf := make([]byte, 64*1024)

	var pauseTime = int64(1000-timesPerSecond*20) / int64(timesPerSecond)

	if *enablePause {
		fmt.Println("enable pause")
		lastTime := time.Now()
		for {
			n, err := fd.Read(buf)
			if err != nil {
				fmt.Println("pause ", err)
				return
			}
			sum += n
			if sum > speed {
				curTime := time.Now()
				toSleep := lastTime.Add(time.Second).Sub(curTime).Milliseconds()
				pauseTime = pauseTime + toSleep/int64(timesPerSecond)

				lastTime = time.Now()
				sum = 0
			}
			time.Sleep(time.Duration(pauseTime) * time.Millisecond)
		}
	} else {
		lastTime := time.Now()
		for {
			n, err := fd.Read(buf)
			if err != nil {
				fmt.Println("narmal ", err)
				return
			}
			sum += n
			if sum > speed {
				curTime := time.Now()
				if curTime.Before(lastTime.Add(time.Second)) {
					time.Sleep(lastTime.Add(time.Second).Sub(curTime))
				}
				lastTime = time.Now()
				sum = 0
			}
		}
	}
}
