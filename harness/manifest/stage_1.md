# Bind to a TCP Port

Redis listens for connections on TCP port 6379 by default. Before any commands can be processed, your server needs to open that port and accept incoming connections. This stage is purely about getting the network layer in place — no protocol parsing yet.

## Requirements

- Bind and listen on TCP port 6379
- Accept incoming client connections

## What will be tested

- A TCP connection to `localhost:6379` succeeds without being refused or timing out
