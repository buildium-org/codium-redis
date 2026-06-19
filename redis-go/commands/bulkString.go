package commands

import "strconv"

type BulkStringMessage struct {
	Value string
}

func NewBulkStringMessage(value string) *BulkStringMessage {
	return &BulkStringMessage{Value: value}
}
func (m *BulkStringMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}
