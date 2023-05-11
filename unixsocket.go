package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type UnixSocket struct {
	filename string
	bufsize  int
	handler  func(string) string
}

func NewUnixSocket(filename string, size ...int) *UnixSocket {
	size1 := 10480
	if size != nil {
		size1 = size[0]
	}
	us := UnixSocket{filename: filename, bufsize: size1}
	return &us
}

func (this *UnixSocket) createServer() {
	os.Remove(this.filename)
	addr, err := net.ResolveUnixAddr("unix", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	listener, err := net.ListenUnix("unix", addr)
	defer listener.Close()
	if err != nil {
		panic("Cannot listen to unix domain socket: " + err.Error())
	}
	fmt.Println("Listening on", listener.Addr())
	for {
		c, err := listener.Accept()
		if err != nil {
			panic("Accept: " + err.Error())
		}
		go this.HandleServerConn(c)
	}
}

//接收连接并处理
func (this *UnixSocket) HandleServerConn(c net.Conn) {
	defer c.Close()
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	// 这里，你需要 parse buf 里的数据来决定返回什么给客户端
	// 假设 respnoseData 是你想返回的文件内容
	result := this.HandleServerContext(string(buf[0:nr]))
	_, err = c.Write([]byte(result))
	if err != nil {
		panic("Writes failed.")
	}
}

func (this *UnixSocket) SetContextHandler(f func(string) string) {
	this.handler = f
}

//接收内容并返回结果
func (this *UnixSocket) HandleServerContext(context string) string {
	if this.handler != nil {
		return this.handler(context)
	}
	now := time.Now().String()
	return now
}

func (this *UnixSocket) StartServer() {
	this.createServer()
}

func showContext(content string) string {
	fmt.Println(content)
	return fmt.Sprintf("Server: time: %v", time.Now().UTC())
}

//客户端
func (this *UnixSocket) ClientSendContext(context string) string {
	addr, err := net.ResolveUnixAddr("unix", this.filename)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	//拔号
	c, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}
	//写出
	_, err = c.Write([]byte(context))
	if err != nil {
		panic(fmt.Errorf("Writes failed. err: %v", err))
	}
	//读结果
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	return string(buf[0:nr])
}

func main() {
	var wg sync.WaitGroup
	var us = NewUnixSocket("./admin.sock", 1024)
	us.SetContextHandler(showContext)
	go func() {
		wg.Add(1)
		defer wg.Done()
		us.StartServer()
	}()
	time.Sleep(time.Second)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			fmt.Println(us.ClientSendContext(fmt.Sprintf("Client: time: %v", time.Now().UTC())))
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

}
