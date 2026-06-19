package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("read error from %s: %v", conn.RemoteAddr(), err)
			}
			return
		}

		if strings.TrimSpace(line) == "PING" {
			if _, err := conn.Write([]byte("+PONG\r\n")); err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		}
	}
}
