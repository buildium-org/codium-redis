package datastore

import (
	"fmt"
	"sync"
)

type DataStore struct {
	Data  map[string]string
	Mutex sync.Mutex
}

func NewDataStore() *DataStore {
	return &DataStore{Data: make(map[string]string), Mutex: sync.Mutex{}}
}

func (d *DataStore) Set(key string, value string) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	d.Data[key] = value
}

func (d *DataStore) Get(key string) (string, error) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	value, ok := d.Data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return value, nil
}
