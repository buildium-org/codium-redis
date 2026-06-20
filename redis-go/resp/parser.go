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
	// "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n" -> ECHO hey
	// "*5\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n$2\r\nPX\r\n:1000\r\n" -> SET key1 value1 PX 1000
	numElements, err := strconv.Atoi(messageParts[0][1:])
	if err != nil {
		return nil, fmt.Errorf("invalid array: %s", messageParts[0])
	}

	elements := []string{}
	for i := 1; i < len(messageParts); i++ {
		if messageParts[i] == "" {
			continue
		}
		op := messageParts[i][0]
		switch op {
		case '$':
			length, err := strconv.Atoi(messageParts[i][1:])
			if err != nil {
				return nil, fmt.Errorf("invalid bulk string: %s", messageParts[i])
			}
			value := messageParts[i+1]
			if len(value) != length {
				return nil, fmt.Errorf("invalid bulk string: %s", value)
			}
			elements = append(elements, value)
			i++
		case ':':
			value, err := strconv.ParseInt(messageParts[i][1:], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer: %s", messageParts[i])
			}
			elements = append(elements, strconv.FormatInt(value, 10))
		default:
			return nil, fmt.Errorf("invalid operation: %s", messageParts[i])
		}
	}

	if len(elements) != numElements {
		return nil, fmt.Errorf("invalid array: %s", messageParts[0])
	}

	return elements, nil
}

func (p *RESPParser) parseBulkString(lengthPart string, valuePart string) (string, error) {
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
