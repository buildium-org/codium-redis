package main

import (
	"fmt"
	datastore "golang/dataStore"
	"golang/resp"
	"io"
	"log"
	"net"
)

func main() {
	dataStore := datastore.NewDataStore()
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
		go handleConn(conn, dataStore)
	}
}

func handleConn(conn net.Conn, dataStore *datastore.DataStore) {
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
		if err != nil {
			log.Printf("parse error from %s: %v", conn.RemoteAddr(), err)
			return
		}
		err = msg.Handle(conn, dataStore)
		if err != nil {
			log.Printf("handle error from %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}
