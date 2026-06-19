package commands

type PongMessage struct{}

func NewPongMessage() *PongMessage {
	return &PongMessage{}
}
func (m *PongMessage) ToBytes() []byte {
	return []byte("+PONG\r\n")
}
