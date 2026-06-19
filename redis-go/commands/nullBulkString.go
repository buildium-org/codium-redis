package commands

type NullBulkStringMessage struct{}

func NewNullBulkStringMessage() *NullBulkStringMessage {
	return &NullBulkStringMessage{}
}
func (m *NullBulkStringMessage) ToBytes() []byte {
	return []byte("$-1\r\n")
}
