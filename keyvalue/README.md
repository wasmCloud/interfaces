# Key Value Interface
This is the key-value interface with the contract ID of `wasmcloud:keyvalue`. This interface defines a set of common operations for interacting with key-value stores. 

Note that things like consistency guarantees, backup, failover support, replications, and more are all concerns specific to individual providers and not the interface itself.

## Implementations
The following is a list of known implementations of the `wasmcloud:keyvalue` interface. Feel free to submit a PR if you know of more.

| Name | Vendor | Description |
| :---: | :---: | :--- |
| [Redis](https://github.com/wasmCloud/capability-providers/tree/main/kvredis) | wasmCloud | wasmCloud key-value provider for the **Redis** database