package commands

import (
	datastore "golang/dataStore"
	"net"
	"strconv"
)

type SetMessage struct {
	Key          string
	Value        string
	ExpireTimeMS int64
}

func NewSetMessage(tokens []string) *SetMessage {
	key := tokens[0]
	value := tokens[1]
	if (len(tokens) < 3) || tokens[2] != "PX" {
		return &SetMessage{Key: key, Value: value, ExpireTimeMS: -1}
	}

	expireTimeMS, err := strconv.ParseInt(tokens[3], 10, 64)
	if err != nil {
		panic(err)
	}
	return &SetMessage{Key: key, Value: value, ExpireTimeMS: expireTimeMS}

}
func (m *SetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n" + "$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}

func (m *SetMessage) Handle(conn net.Conn, dataStore *datastore.DataStore) error {
	dataStore.Set(m.Key, m.Value, m.ExpireTimeMS)
	_, err := conn.Write(NewOkMessage().ToBytes())
	if err != nil {
		return err
	}
	return nil
}
