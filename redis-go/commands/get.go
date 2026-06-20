package commands

import (
	datastore "golang/dataStore"
	"net"
	"strconv"
)

type GetMessage struct {
	Key string
}

func NewGetMessage(tokens []string) *GetMessage {
	return &GetMessage{Key: tokens[0]}
}
func (m *GetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n")
}

func (m *GetMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	value, err := dataStore.Get(m.Key)
	if err != nil {
		return err
	}
	if value == nil {
		_, err = conn.Write(NewNullBulkStringMessage().ToBytes())
	} else {
		_, err = conn.Write(NewBulkStringMessage(value.Value).ToBytes())
	}

	if err != nil {
		return err
	}
	return nil
}
