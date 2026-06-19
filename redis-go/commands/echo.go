package commands

import "strconv"

type EchoMessage struct {
	Message string
}

func NewEchoMessage(message string) *EchoMessage {
	return &EchoMessage{Message: message}
}
func (m *EchoMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Message)) + "\r\n" + m.Message + "\r\n")
}
