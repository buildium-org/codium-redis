package commands

import "strconv"

type GetMessage struct {
	Key string
}

func NewGetMessage(tokens []string) *GetMessage {
	return &GetMessage{Key: tokens[0]}
}
func (m *GetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n")
}
