package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestNewSetMessageWithoutExpiry(t *testing.T) {
	msg := NewSetMessage([]string{"key1", "value1"})
	if msg.Key != "key1" {
		t.Fatalf("expected key key1, got %q", msg.Key)
	}
	if msg.Value != "value1" {
		t.Fatalf("expected value value1, got %q", msg.Value)
	}
	if msg.ExpireTimeMS != -1 {
		t.Fatalf("expected ExpireTimeMS -1, got %d", msg.ExpireTimeMS)
	}
}

func TestNewSetMessageWithExpiry(t *testing.T) {
	msg := NewSetMessage([]string{"key1", "value1", "PX", "1000"})
	if msg.ExpireTimeMS != 1000 {
		t.Fatalf("expected ExpireTimeMS 1000, got %d", msg.ExpireTimeMS)
	}
}

func TestSetMessageToBytes(t *testing.T) {
	msg := NewSetMessage([]string{"key1", "value1"})
	if got := string(msg.ToBytes()); got != "$4\r\nkey1\r\n$6\r\nvalue1\r\n" {
		t.Fatalf("unexpected ToBytes output: %q", got)
	}
}

func TestSetMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewSetMessage([]string{"key1", "value1"})

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "+OK\r\n" {
		t.Fatalf("expected +OK\\r\\n, got %q", got)
	}

	entry, err := store.Get("key1")
	if err != nil {
		t.Fatalf("failed to get key from store: %v", err)
	}
	if entry.Value != "value1" {
		t.Fatalf("expected stored value value1, got %q", entry.Value)
	}
}
