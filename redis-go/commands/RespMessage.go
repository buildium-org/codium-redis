package commands

import (
	datastore "golang/dataStore"
	"net"
)

type RespMessage interface {
	ToBytes() []byte
	Handle(conn net.Conn, dataStore *datastore.DataStore) error
}
