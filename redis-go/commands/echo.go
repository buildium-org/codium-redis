package commands

import (
	datastore "golang/dataStore"
	"net"
	"strconv"
)

type EchoMessage struct {
	Message string
}

func NewEchoMessage(tokens []string) *EchoMessage {
	return &EchoMessage{Message: tokens[0]}
}
func (m *EchoMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Message)) + "\r\n" + m.Message + "\r\n")
}

func (m *EchoMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write(m.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
