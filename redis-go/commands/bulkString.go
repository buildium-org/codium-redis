package commands

import (
	datastore "golang/dataStore"
	"net"
	"strconv"
)

type BulkStringMessage struct {
	Value string
}

func NewBulkStringMessage(value string) *BulkStringMessage {
	return &BulkStringMessage{Value: value}
}
func (m *BulkStringMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}

func (m *BulkStringMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	_, err := conn.Write(m.ToBytes())
	if err != nil {
		return err
	}
	return nil
}
