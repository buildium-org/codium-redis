package resp

import "testing"

func TestParseSimpleString(t *testing.T) {
	parser := &RESPParser{}
	message := "+PING\r\n"
	msg, err := parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(PingMessage); !ok {
		t.Fatalf("expected PingMessage, got %v", msg)
	}

	message = "+OK\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(OkMessage); !ok {
		t.Fatalf("expected OkMessage, got %v", msg)
	}
}

func TestParseArray(t *testing.T) {
	parser := &RESPParser{}
	message := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"
	msg, err := parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(EchoMessage); !ok {
		t.Fatalf("expected EchoMessage, got %v", msg)
	}
	if msg.(EchoMessage).Message != "hey" {
		t.Fatalf("expected Message to be hey, got %v", msg.(EchoMessage).Message)
	}

	message = "*1\r\n$4\r\nPING\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(PingMessage); !ok {
		t.Fatalf("expected PingMessage, got %v", msg)
	}

	message = "*3\r\n$3\r\nSET\r\n$4\r\nkey1\r\n$6\r\nvalue1\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(SetMessage); !ok {
		t.Fatalf("expected SetMessage, got %v", msg)
	}
	if msg.(SetMessage).Key != "key1" {
		t.Fatalf("expected Key to be key1, got %v", msg.(SetMessage).Key)
	}
	if msg.(SetMessage).Value != "value1" {
		t.Fatalf("expected Value to be value1, got %v", msg.(SetMessage).Value)
	}

	message = "*2\r\n$3\r\nGET\r\n$4\r\nkey1\r\n"
	msg, err = parser.Parse(message)
	if err != nil {
		t.Fatalf("failed to parse message: %v", err)
	}
	if _, ok := msg.(GetMessage); !ok {
		t.Fatalf("expected GetMessage, got %v", msg)
	}
	if msg.(GetMessage).Key != "key1" {
		t.Fatalf("expected Key to be key1, got %v", msg.(GetMessage).Key)
	}
}
