# Key Expiry with PX

Extend SET to support the `PX` option, which sets a key with a millisecond time-to-live. Once the key expires, GET should return a null bulk string instead of the value.

## Requirements

- `SET key value PX milliseconds` stores the key with an expiry
- `GET key` returns the value before the expiry
- `GET key` returns `$-1\r\n` (null bulk string) after the key has expired

## What will be tested

- `SET key1 value1 PX 1000` returns `+OK\r\n`
- `GET key1` immediately returns `$6\r\nvalue1\r\n`
- `GET key1` after 1000ms returns `$-1\r\n`
