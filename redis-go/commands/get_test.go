package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestGetMessageToBytes(t *testing.T) {
	msg := NewGetMessage([]string{"key1"})
	if got := string(msg.ToBytes()); got != "$4\r\nkey1\r\n" {
		t.Fatalf("expected $4\\r\\nkey1\\r\\n, got %q", got)
	}
}

func TestGetMessageHandleExistingKey(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	store.Set("key1", "value1", -1)
	msg := NewGetMessage([]string{"key1"})

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "$6\r\nvalue1\r\n" {
		t.Fatalf("expected $6\\r\\nvalue1\\r\\n, got %q", got)
	}
}

func TestGetMessageHandleMissingKey(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewGetMessage([]string{"missing"})

	err := msg.Handle(conn, store)
	if err == nil {
		t.Fatal("expected error for missing key")
	}
	if got := conn.written(); got != "" {
		t.Fatalf("expected no response written, got %q", got)
	}
}
