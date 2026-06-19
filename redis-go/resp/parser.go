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

	var tokens []string
	var err error
	switch messageParts[0][0] {
	case '+':
		tokens, err = p.parseSimpleString(messageParts)
	case '*':
		tokens, err = p.parseArray(messageParts)
	default:
		return nil, fmt.Errorf("unknown message type: %v", messageParts[0][0])
	}
	if err != nil {
		return nil, err
	}
	return p.parseTokens(tokens)
}

func (p *RESPParser) parseTokens(tokens []string) (commands.RespMessage, error) {
	switch tokens[0] {
	case "ECHO":
		return commands.NewEchoMessage(tokens[1:]), nil
	case "PING":
		return commands.NewPingMessage(), nil
	case "SET":
		return commands.NewSetMessage(tokens[1:]), nil
	case "GET":
		return commands.NewGetMessage(tokens[1:]), nil
	default:
		return nil, fmt.Errorf("unknown token: %s", tokens[0])
	}
}

func (p *RESPParser) parseSimpleString(messageParts []string) ([]string, error) {
	token := messageParts[0][1:]
	return []string{token}, nil
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
