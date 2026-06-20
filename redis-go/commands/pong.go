package commands

import (
	datastore "golang/dataStore"
	"net"
)

type PongMessage struct{}

func NewPongMessage() *PongMessage {
	return &PongMessage{}
}
func (m *PongMessage) ToBytes() []byte {
	return []byte("+PONG\r\n")
}
func (m *PongMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write(m.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
