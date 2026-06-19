package commands

type OkMessage struct{}

func NewOkMessage() *OkMessage {
	return &OkMessage{}
}
func (m *OkMessage) ToBytes() []byte {
	return []byte("+OK\r\n")
}
