package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestOkMessageToBytes(t *testing.T) {
	msg := NewOkMessage()
	if got := string(msg.ToBytes()); got != "+OK\r\n" {
		t.Fatalf("expected +OK\\r\\n, got %q", got)
	}
}

func TestOkMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewOkMessage()

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "+OK\r\n" {
		t.Fatalf("expected +OK\\r\\n, got %q", got)
	}
}
