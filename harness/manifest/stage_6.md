# Implement SET and GET

SET and GET are the bread and butter of Redis. Together they turn your server into an actual key-value store. You'll need an in-memory map to hold key-value pairs, and because multiple goroutines can read and write it simultaneously (from concurrent clients), you'll want a mutex or `sync.Map` to avoid data races.

This stage does not involve expiry yet — just plain storage and retrieval.

## Requirements

- `SET key value` stores the key-value pair and returns `+OK\r\n`
- `GET key` returns the stored value as a RESP bulk string
- Keys that have not been set should return a null bulk string (`$-1\r\n`)

## What will be tested

- `SET key1 value1` returns `+OK\r\n`
- `GET key1` returns `$6\r\nvalue1\r\n`
