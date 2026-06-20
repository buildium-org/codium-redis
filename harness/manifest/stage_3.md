# Handle Multiple PINGs

Verify that your server can handle more than one command on the same connection without closing or resetting. Redis clients reuse connections for many commands.

## Requirements

- Keep the connection open after responding to the first PING
- Respond correctly to each subsequent PING on the same connection

## What will be tested

- Two sequential `+PING\r\n` messages on one connection each return `+PONG\r\n`
