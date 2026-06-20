# Handle Concurrent Clients

Make your server handle multiple clients connecting at the same time. This requires handling each connection in its own goroutine so one client doesn't block another.

## Requirements

- Accept multiple simultaneous TCP connections
- Handle each connection independently (e.g. in a goroutine)

## What will be tested

- Two clients connect and send PING concurrently; both receive `+PONG\r\n`
