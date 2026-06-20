# Respond to PING

PING is the simplest Redis command and the standard way clients check that a server is alive. In this stage you'll read a raw inline command from the socket and write back a response — your first real interaction with the Redis wire format.

The inline format (`+PING\r\n`) is a shorthand used by simple clients and the Redis CLI. It isn't a full RESP array, so you don't need a protocol parser yet — just read the line and write `+PONG\r\n` back.

## Requirements

- Read an inline `+PING\r\n` from the connection
- Respond with the simple string `+PONG\r\n`

## What will be tested

- Sending `+PING\r\n` returns `+PONG\r\n`
