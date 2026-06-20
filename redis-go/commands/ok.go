package commands

import (
	datastore "golang/dataStore"
	"net"
)

type OkMessage struct{}

func NewOkMessage() *OkMessage {
	return &OkMessage{}
}
func (m *OkMessage) ToBytes() []byte {
	return []byte("+OK\r\n")
}
func (m *OkMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write(m.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
