package main

import (
	"fmt"
	"time"
)

/*
channel 支持 for range 的方式进行遍历，需要注意两个细节。
1.在遍历时，如果 channel 没有关闭，则回出现 deadlock 的错误。
2.在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
3.对于nil channel，无论收发都会被阻塞。
4.写完 chan 之后一定要关闭close chan，否则主协程读的时候，会发生被阻塞。
5.已关闭的Channel(有缓冲的),如果继续读数据，得到的是默认值(对于int或任意类型的指针，就是0, 自定义类型是其默认值), 如果没有关闭，读不到数据了。

select语句中除default外，每个case操作一个channel，要么读要么写。
select语句中除default外，各case执行顺序是随机的。
select语句中如果没有default语句, 则会阻塞等待任一case。
select语句中读操作要判断是否成功读取，关闭的channel也可以读取.
*/

type AA struct {
	ss int
	jj int
}

func main() {
	ch := make(chan *AA, 10)
	fmt.Printf("the length of channel is %d\n", len(ch))
	ch <- &AA{1, 2}
	fmt.Printf("the length of channel is %d\n", len(ch))
	//	close(ch)
	go func() {
		for {
			time.Sleep(time.Second)
			ch <- &AA{3, 4}
			break
		}
	}()
	time.Sleep(time.Second)
	close(ch)
	for it := range ch {
		fmt.Println(*it)
	}
	fmt.Println("close ch")

	for {
		select {
		case v, ok := <-ch:
			time.Sleep(time.Second)
			if ok {
				fmt.Printf("v: %d ok: %v\n", v, ok)
			} else {
				fmt.Printf("v: %d ok: %v\n", v, ok)
				break
			}
		}
	}
}
