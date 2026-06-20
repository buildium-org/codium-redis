# Key Expiry with PX

Redis is commonly used as a cache, and caches need keys that expire automatically. In this stage you'll extend SET to accept the `PX` option, which sets a time-to-live in milliseconds. Once the TTL has elapsed, the key should disappear — GET should return a null bulk string as if the key was never set.

The simplest approach is to record the expiry timestamp when the key is stored, then check it on every GET and treat the key as absent if the current time is past the deadline.

## Requirements

- `SET key value PX milliseconds` stores the key with an expiry time
- `GET key` returns the value before the expiry
- `GET key` returns `$-1\r\n` (null bulk string) after the key has expired

## What will be tested

- `SET key1 value1 PX 1000` returns `+OK\r\n`
- `GET key1` immediately after returns `$6\r\nvalue1\r\n`
- `GET key1` after 1000ms returns `$-1\r\n`
