# Implement SET and GET

Add in-memory key-value storage with SET and GET. These are the core commands that make Redis useful as a cache or data store.

## Requirements

- `SET key value` stores a key-value pair and returns `+OK\r\n`
- `GET key` returns the stored value as a RESP bulk string

## What will be tested

- `SET key1 value1` returns `+OK\r\n`
- `GET key1` returns `$6\r\nvalue1\r\n`
