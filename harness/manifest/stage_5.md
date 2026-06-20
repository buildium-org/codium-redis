# Implement ECHO

Add the ECHO command, which introduces RESP array parsing. Commands sent by real Redis clients are encoded as RESP arrays, so implementing ECHO lays the groundwork for all future commands.

## Requirements

- Parse RESP array-formatted commands (e.g. `*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n`)
- Respond with a RESP bulk string containing the echoed argument

## What will be tested

- `ECHO hey` returns `$3\r\nhey\r\n`
