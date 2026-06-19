package commands

import "strconv"

type GetMessage struct {
	Key string
}

func NewGetMessage(key string) *GetMessage {
	return &GetMessage{Key: key}
}
func (m *GetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n")
}
