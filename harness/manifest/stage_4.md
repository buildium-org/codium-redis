# Handle Concurrent Clients

A production Redis server handles hundreds of clients at once. If your server processes connections one at a time, every client after the first will be blocked until the previous one disconnects. The fix is simple in Go: handle each accepted connection in its own goroutine so clients run independently.

## Requirements

- Accept multiple simultaneous TCP connections
- Handle each connection in its own goroutine so one client does not block another

## What will be tested

- Two clients connect and send `+PING\r\n` at the same time; both receive `+PONG\r\n` without either timing out waiting for the other
