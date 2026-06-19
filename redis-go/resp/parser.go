package resp

import (
	"fmt"
	"strconv"
	"strings"
)

// *1\r\n$4\r\nPING\r\n -> PING
// +PONG\r\n -> PONG
// *2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n -> ECHO hey

type RespMessage interface{}

type PingMessage struct{}

type ArrayMessage struct {
	Elements []string
}
type EchoMessage struct {
	Message string
}
type SetMessage struct {
	Key   string
	Value string
}
type GetMessage struct {
	Key string
}

type OkMessage struct{}

func NewOkMessage() *OkMessage {
	return &OkMessage{}
}
func (m *OkMessage) ToBytes() []byte {
	return []byte("+OK\r\n")
}

type BulkStringMessage struct {
	Value string
}

func NewBulkStringMessage(value string) *BulkStringMessage {
	return &BulkStringMessage{Value: value}
}
func (m *BulkStringMessage) ToBytes() []byte {
	return []byte("$" + strconv.Itoa(len(m.Value)) + "\r\n" + m.Value + "\r\n")
}

type RESPParser struct{}

func (p *RESPParser) Parse(message string) (RespMessage, error) {
	messageParts := strings.Split(message, "\r\n")

	switch messageParts[0][0] {
	case '+':
		return p.parseSimpleString(messageParts[0])
	case '*':
		tokens, err := p.parseArray(messageParts)
		if err != nil {
			return nil, err
		}
		return p.parseTokens(tokens.Elements)
	default:
		return nil, fmt.Errorf("unknown message type: %v", messageParts[0][0])
	}

}

func (p *RESPParser) parseTokens(messageParts []string) (RespMessage, error) {
	switch messageParts[0] {
	case "ECHO":
		return EchoMessage{Message: messageParts[1]}, nil
	case "PING":
		return PingMessage{}, nil
	case "SET":
		return SetMessage{Key: messageParts[1], Value: messageParts[2]}, nil
	case "GET":
		return GetMessage{Key: messageParts[1]}, nil
	default:
		return nil, fmt.Errorf("unknown token: %s", messageParts[0])
	}
}

func (p *RESPParser) parseSimpleString(messagePart string) (RespMessage, error) {
	// +PONG\r\n -> PONG
	if messagePart[0] != '+' {
		return nil, fmt.Errorf("invalid simple string: %s", messagePart)
	}
	if messagePart[1:] == "OK" {
		return OkMessage{}, nil
	}
	if messagePart[1:] == "PING" {
		return PingMessage{}, nil
	}
	return nil, fmt.Errorf("unknown simple string: %s", messagePart[1:])
}

func (p *RESPParser) parseArray(messageParts []string) (*ArrayMessage, error) {
	// *2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n -> ECHO hey
	if messageParts[0][0] != '*' {
		return nil, fmt.Errorf("invalid array: %s", messageParts[0])
	}

	numElements, err := strconv.Atoi(messageParts[0][1:])
	if err != nil {
		return nil, fmt.Errorf("invalid array: %s", messageParts[0])
	}

	elements := []string{}
	for i := 1; i < numElements*2; i += 2 {
		lengthPart := messageParts[i]
		valuePart := messageParts[i+1]
		value, err := p.parseBulkString(lengthPart, valuePart)
		if err != nil {
			return nil, fmt.Errorf("invalid array: %s", err)
		}
		elements = append(elements, value)
	}

	return &ArrayMessage{Elements: elements}, nil
}

func (p *RESPParser) parseBulkString(lengthPart string, valuePart string) (string, error) {
	// $4\r\nECHO\r\n -> ECHO
	if lengthPart[0] != '$' {
		return "", fmt.Errorf("invalid bulk string: %s", lengthPart)
	}
	length, err := strconv.Atoi(lengthPart[1:])
	if err != nil {
		return "", fmt.Errorf("invalid bulk string: %s", lengthPart)
	}
	if len(valuePart) != length {
		return "", fmt.Errorf("invalid bulk string: %s", valuePart)
	}
	return valuePart, nil
}
