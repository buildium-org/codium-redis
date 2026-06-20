package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestBulkStringMessageToBytes(t *testing.T) {
	msg := NewBulkStringMessage("hello")
	if got := string(msg.ToBytes()); got != "$5\r\nhello\r\n" {
		t.Fatalf("expected $5\\r\\nhello\\r\\n, got %q", got)
	}
}

func TestBulkStringMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewBulkStringMessage("hello")

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "$5\r\nhello\r\n" {
		t.Fatalf("expected $5\\r\\nhello\\r\\n, got %q", got)
	}
}
