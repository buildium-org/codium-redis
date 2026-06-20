package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestNullBulkStringMessageToBytes(t *testing.T) {
	msg := NewNullBulkStringMessage()
	if got := string(msg.ToBytes()); got != "$-1\r\n" {
		t.Fatalf("expected $-1\\r\\n, got %q", got)
	}
}

func TestNullBulkStringMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewNullBulkStringMessage()

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "$-1\r\n" {
		t.Fatalf("expected $-1\\r\\n, got %q", got)
	}
}
