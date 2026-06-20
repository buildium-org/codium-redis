package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestPongMessageToBytes(t *testing.T) {
	msg := NewPongMessage()
	if got := string(msg.ToBytes()); got != "+PONG\r\n" {
		t.Fatalf("expected +PONG\\r\\n, got %q", got)
	}
}

func TestPongMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewPongMessage()

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "+PONG\r\n" {
		t.Fatalf("expected +PONG\\r\\n, got %q", got)
	}
}
