package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("failed to listen on :6379: %v", err)
	}
	defer listener.Close()

	fmt.Println("Listening on :6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("connection from %s", conn.RemoteAddr())
}
