package main

import (
	"fmt"
	"golang/commands"
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

		switch msg := msg.(type) {
		case *commands.PingMessage:
			_, err := conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		case *commands.EchoMessage:
			_, err := conn.Write(commands.NewBulkStringMessage(msg.Message).ToBytes())
			if err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		case *commands.SetMessage:
			dataStore.Set(msg.Key, msg.Value, msg.ExpireTimeMS)
			_, err := conn.Write(commands.NewOkMessage().ToBytes())
			if err != nil {
				log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
				return
			}
		case *commands.GetMessage:
			value, err := dataStore.Get(msg.Key)
			if err != nil {
				log.Printf("get error from %s: %v", conn.RemoteAddr(), err)
				return
			}

			if value == nil {
				_, err = conn.Write(commands.NewNullBulkStringMessage().ToBytes())
				if err != nil {
					log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
					return
				}
			} else {
				_, err = conn.Write(commands.NewBulkStringMessage(value.Value).ToBytes())
				if err != nil {
					log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
					return
				}
			}
		default:
			log.Printf("unknown message type: %v", msg)
		}
	}
}
