package resp

import (
	"fmt"
	"golang/commands"
	"strconv"
	"strings"
)

type RESPParser struct{}

func (p *RESPParser) Parse(message string) (commands.RespMessage, error) {
	messageParts := strings.Split(message, "\r\n")

	switch messageParts[0][0] {
	case '+':
		return p.parseSimpleString(messageParts[0])
	case '*':
		tokens, err := p.parseArray(messageParts)
		if err != nil {
			return nil, err
		}
		return p.parseTokens(tokens)
	default:
		return nil, fmt.Errorf("unknown message type: %v", messageParts[0][0])
	}

}

func (p *RESPParser) parseTokens(messageParts []string) (commands.RespMessage, error) {
	switch messageParts[0] {
	case "ECHO":
		return commands.NewEchoMessage(messageParts[1]), nil
	case "PING":
		return commands.NewPingMessage(), nil
	case "SET":
		return commands.NewSetMessage(messageParts[1], messageParts[2]), nil
	case "GET":
		return commands.NewGetMessage(messageParts[1]), nil
	default:
		return nil, fmt.Errorf("unknown token: %s", messageParts[0])
	}
}

func (p *RESPParser) parseSimpleString(messagePart string) (commands.RespMessage, error) {
	// +PONG\r\n -> PONG
	if messagePart[0] != '+' {
		return nil, fmt.Errorf("invalid simple string: %s", messagePart)
	}
	if messagePart[1:] == "PING" {
		return commands.NewPingMessage(), nil
	}
	if messagePart[1:] == "OK" {
		return commands.NewOkMessage(), nil
	}
	return nil, fmt.Errorf("unknown simple string: %s", messagePart[1:])
}

func (p *RESPParser) parseArray(messageParts []string) ([]string, error) {
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

	return elements, nil
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
