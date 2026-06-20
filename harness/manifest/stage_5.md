# Implement ECHO

The first four stages were about getting the network layer right — binding a port, staying alive across commands, and handling concurrent connections. Now it's time to actually speak Redis.

Almost every real Redis command is sent as a RESP array (REdis Serialization Protocol), not the inline format you handled for PING. ECHO is the simplest command that uses this encoding, making it a good first target. Before implementing ECHO, this is a great time to invest in a proper RESP parser — you'll need it for every command from here on. A clean parser will pay dividends immediately in the next two stages.

A RESP array looks like this:

```
*2\r\n        ← array of 2 elements
$4\r\n        ← bulk string of length 4
ECHO\r\n      ← the command name
$3\r\n        ← bulk string of length 3
hey\r\n       ← the argument
```

## Requirements

- Parse incoming RESP array commands
- Respond to `ECHO <message>` with a RESP bulk string containing the message

## What will be tested

- `ECHO hey` (encoded as `*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n`) returns `$3\r\nhey\r\n`
