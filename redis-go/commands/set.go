package commands

import "strconv"

type SetMessage struct {
	Key   string
	Value string
}

func NewSetMessage(tokens []string) *SetMessage {
	key := tokens[0]
	value := tokens[1]
	return &SetMessage{Key: key, Value: value}
}
func (m *SetMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Key)) + "\r\n" + m.Key + "\r\n" + "$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}
