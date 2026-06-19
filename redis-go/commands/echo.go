package commands

import "strconv"

type EchoMessage struct {
	Message string
}

func NewEchoMessage(tokens []string) *EchoMessage {
	return &EchoMessage{Message: tokens[0]}
}
func (m *EchoMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Message)) + "\r\n" + m.Message + "\r\n")
}
