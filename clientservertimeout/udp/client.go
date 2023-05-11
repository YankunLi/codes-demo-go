package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connTimeout := 3 * time.Second
	conn, err := net.DialTimeout("udp", "127.0.0.1:8080", connTimeout) // 3s timeout
	if err != nil {
		log.Println("dial failed:", err)
		os.Exit(1)
	}
	defer conn.Close()

	readTimeout := 2 * time.Second

	err = conn.SetWriteDeadline(time.Now().Add(readTimeout)) // timeout
	if err != nil {
		log.Println("setReadDeadline failed:", err)
	}

	buffer := []byte("hello world")

	for {

		n, err := conn.Write(buffer)
		if err != nil {
			log.Println("Read failed:", err)
			break
			//break
		}

		log.Println("write count:", n)
	}

}
