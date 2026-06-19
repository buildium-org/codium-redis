package commands

type PingMessage struct{}

func NewPingMessage() *PingMessage {
	return &PingMessage{}
}
func (m *PingMessage) ToBytes() []byte {
	return []byte("+PING\r\n")
}
