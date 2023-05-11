package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	go func(cancel context.CancelFunc) {
		time.Sleep(10 * time.Second)
		cancel()
		fmt.Printf("Cancel had been call\n")
	}(cancel)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Printf("Parent goroutine has been cancel, err:%s\n", ctx.Err())
		}
		fmt.Printf("child goroutin exit\n")
	}(ctx)
	wg.Wait()
}
