package commands

import (
	datastore "golang/dataStore"
	"testing"
)

func TestEchoMessageToBytes(t *testing.T) {
	msg := NewEchoMessage([]string{"hello"})
	if got := string(msg.ToBytes()); got != "$5\r\nhello\r\n" {
		t.Fatalf("expected $5\\r\\nhello\\r\\n, got %q", got)
	}
}

func TestEchoMessageHandle(t *testing.T) {
	conn := newMockConn()
	store := datastore.NewDataStore()
	msg := NewEchoMessage([]string{"hello"})

	if err := msg.Handle(conn, store); err != nil {
		t.Fatalf("Handle returned error: %v", err)
	}
	if got := conn.written(); got != "$5\r\nhello\r\n" {
		t.Fatalf("expected $5\\r\\nhello\\r\\n, got %q", got)
	}
}
