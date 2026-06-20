package commands

import (
	datastore "golang/dataStore"
	"net"
)

type NullBulkStringMessage struct{}

func NewNullBulkStringMessage() *NullBulkStringMessage {
	return &NullBulkStringMessage{}
}
func (m *NullBulkStringMessage) ToBytes() []byte {
	return []byte("$-1\r\n")
}
func (m *NullBulkStringMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write(m.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
