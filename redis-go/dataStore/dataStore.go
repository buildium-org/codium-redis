package datastore

import (
	"fmt"
	"sync"
	"time"
)

type DataStoreEntry struct {
	Value        string
	ExpireTimeMS int64
}
type DataStore struct {
	Data  map[string]DataStoreEntry
	Mutex sync.Mutex
}

func NewDataStore() *DataStore {
	return &DataStore{Data: make(map[string]DataStoreEntry), Mutex: sync.Mutex{}}
}

func (d *DataStore) Set(key string, value string, expireTimeMS int64) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	if expireTimeMS > 0 {
		expireTimeMS = time.Now().UnixMilli() + expireTimeMS
	} else {
		expireTimeMS = -1
	}
	d.Data[key] = DataStoreEntry{Value: value, ExpireTimeMS: expireTimeMS}
}

func (d *DataStore) Get(key string) (*DataStoreEntry, error) {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	value, ok := d.Data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}

	if value.ExpireTimeMS > 0 && time.Now().UnixMilli() > value.ExpireTimeMS {
		delete(d.Data, key)
		return nil, nil
	}
	return &value, nil
}
