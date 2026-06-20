# Respond to PING

Implement PING, the simplest Redis command and the standard way clients check if a server is alive. The server should reply with the inline simple string `+PONG`.

## Requirements

- Read an inline `+PING\r\n` from a client
- Respond with `+PONG\r\n`

## What will be tested

- Sending `+PING\r\n` returns `+PONG\r\n`
