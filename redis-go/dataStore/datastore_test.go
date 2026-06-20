package datastore

import (
	"testing"
	"time"
)

func TestDataStore(t *testing.T) {
	dataStore := NewDataStore()
	dataStore.Set("key1", "value1", 100)
	value, err := dataStore.Get("key1")
	if err != nil {
		t.Fatalf("failed to get value: %v", err)
	}
	if value.Value != "value1" {
		t.Fatalf("expected value to be value1, got %v", value.Value)
	}
	time.Sleep(100 * time.Millisecond)
	value, err = dataStore.Get("key1")
	if err != nil {
		t.Fatalf("failed to get value: %v", err)
	}
	if value != nil {
		t.Fatalf("expected value to be nil, got %v", value)
	}
}
