package resp

import (
	"golang/commands"
	"testing"
)

func TestParseSimpleString(t *testing.T) {
	parser := &RESPParser{}
	message := "+PING\r\n"
	msg, err := parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.PingMessage); !ok {
		t.Fatalf("expected PingMessage, got %v", msg)
	}
}

func TestParseArray(t *testing.T) {
	parser := &RESPParser{}
	message := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	msg, err := parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.EchoMessage); !ok {
		t.Fatalf("expected EchoMessage, got %v", msg)
	}
	if msg.(*commands.EchoMessage).Message != "hey" {
		t.Fatalf("expected Message to be hey, got %v", msg.(*commands.EchoMessage).Message)
	}

	message = "*1\r\n$4\r\nPING\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.PingMessage); !ok {
		t.Fatalf("expected PingMessage, got %v", msg)
	}

	message = "*3\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.SetMessage); !ok {
		t.Fatalf("expected SetMessage, got %v", msg)
	}
	if msg.(*commands.SetMessage).Key != "key1" {
		t.Fatalf("expected Key to be key1, got %v", msg.(*commands.SetMessage).Key)
	}
	if msg.(*commands.SetMessage).Value != "value1" {
		t.Fatalf("expected Value to be value1, got %v", msg.(*commands.SetMessage).Value)
	}

	message = "*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.GetMessage); !ok {
		t.Fatalf("expected GetMessage, got %v", msg)
	}
	if msg.(*commands.GetMessage).Key != "key1" {
		t.Fatalf("expected Key to be key1, got %v", msg.(*commands.GetMessage).Key)
	}

	message = "*5\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n$2\r\nPX\r\n:1000\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(*commands.SetMessage); !ok {
		t.Fatalf("expected SetMessage, got %v", msg)
	}
	if msg.(*commands.SetMessage).ExpireTimeMS != 1000 {
		t.Fatalf("expected ExpireTimeMS to be 1000, got %v", msg.(*commands.SetMessage).ExpireTimeMS)
	}
}
