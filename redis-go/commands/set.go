package commands

import "strconv"

type SetMessage struct {
	Key   string
	Value string
}

func NewSetMessage(key string, value string) *SetMessage {
	return &SetMessage{Key: key, Value: value}
}
func (m *SetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n" + "$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}
