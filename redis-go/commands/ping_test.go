package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestPingMessageToBytes(t *testing.T) {
	msg := NewPingMessage()
	if got := string(msg.ToBytes()); got != "+PING\r\n" {
		t.Fatalf("expected +PING\\r\\n, got %q", got)
	}
}

func TestPingMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewPingMessage()

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "+PONG\r\n" {
		t.Fatalf("expected +PONG\\r\\n, got %q", got)
	}
}
