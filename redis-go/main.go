package main

import (
	"fmt"
	resp "golang/resp"
	"io"
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

	parser := &resp.RESPParser{}
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("read error from %s: %v", conn.RemoteAddr(), err)
			}
			return
		}

		msg, err := parser.Parse(string(buf[:n]))

		switch msg := msg.(type) {
		case resp.PingMessage:
			_, err := conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		case resp.EchoMessage:
			_, err := conn.Write(resp.NewBulkStringMessage(msg.Message).ToBytes())
			if err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		default:
			log.Printf("unknown message type: %v", msg)
		}
	}
}
