package main

import (
	"log"
	"net"
)

func main() {
	addr := "0.0.0.0:8080"

	tcpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		log.Fatalf("net.ResovleTCPAddr fail:%s", addr)
	}

	listener, err := net.ListenUDP("udp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	} else {
		log.Println("listening", addr)
	}

	var buffer []byte = []byte("You are welcome. I'm server.")
	for {
		n, _, err := listener.ReadFromUDP(buffer)
		if err != nil {
			log.Println("listener.Accept error:", err)
			continue
		}
		log.Println("count:", n, "receve Msg: ", string(buffer))

	}

}
