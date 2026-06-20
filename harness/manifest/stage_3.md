# Handle Multiple PINGs

Real Redis clients send many commands over a single long-lived connection. If your server closes the connection after the first response, or stops reading after one command, it will fail here. This stage checks that your read-respond loop keeps running for the lifetime of the connection.

## Requirements

- Keep the connection open after responding to the first PING
- Continue reading and responding to each subsequent command on the same connection

## What will be tested

- Two sequential `+PING\r\n` messages sent on a single connection both return `+PONG\r\n`
