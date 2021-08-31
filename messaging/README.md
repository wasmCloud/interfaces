# Messaging Interface
This is the interface for the `wasmcloud:messaging` contract. This contract is a very simple abstraction over the concept of a message broker. This contract does not have controls or knobs or settings for things like delivery guarantees (e.g. "at most once" vs "at least once"), persistence of messages, etc. 

Any details like that are the responsibility of individual providers.

## Implementations
The following is a list of known implementations of the `wasmcloud:messaging` contract. Feel free to submit a PR if you know of additional providers.

| Name | Vendor | Description |
| :---: | :---: | :--- |
| [NATS](https://github.com/wasmCloud/capability-providers/tree/main/nats) | wasmCloud | wasmCloud Messaging Provider for the [NATS](https://nats.io) broker