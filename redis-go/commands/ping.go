package commands

import (
	datastore "golang/dataStore"
	"net"
)

type PingMessage struct{}

func NewPingMessage() *PingMessage {
	return &PingMessage{}
}
func (m *PingMessage) ToBytes() []byte {
	return []byte("+PING\r\n")
}
func (m *PingMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		return err
	}
	return nil
}
