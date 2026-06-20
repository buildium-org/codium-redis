# Build Your Own Redis

In this tutorial you'll build a Redis-compatible server in Go from scratch. Starting from a bare TCP socket, you'll work up through the RESP wire protocol and implement the core commands that make Redis useful.

## What you'll build

By the end you'll have a server that:

- Accepts TCP connections on port 6379
- Handles concurrent clients
- Speaks the RESP (REdis Serialization Protocol) wire format
- Supports `PING`, `ECHO`, `SET`, and `GET`
- Expires keys set with the `PX` option

## Stages

1. **Bind to a TCP Port** — get the server listening on port 6379
2. **Respond to PING** — reply with `+PONG` to the simplest Redis command
3. **Handle Multiple PINGs** — keep connections alive and handle sequential commands
4. **Handle Concurrent Clients** — serve multiple clients at the same time with goroutines
5. **Implement ECHO** — parse RESP array commands and reply with bulk strings
6. **Implement SET and GET** — store and retrieve key-value pairs in memory
7. **Key Expiry with PX** — expire keys after a millisecond timeout
