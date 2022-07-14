[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-keyvalue.svg)](https://crates.io/crates/wasmcloud-interface-keyvalue)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=keyvalue%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/keyvalue/tinygo)
# wasmCloud Key Value Interface
This is the key-value interface with the contract ID of `wasmcloud:keyvalue`. This interface defines a set of common operations for interacting with key-value stores. 

Note that things like consistency guarantees, backup, failover support, replications, and more are all concerns specific to individual providers and not the interface itself.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:keyvalue` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [Redis](https://github.com/wasmCloud/capability-providers/tree/main/kvredis) | wasmCloud | wasmCloud key-value provider for the **Redis** database
| [Vault](https://github.com/wasmCloud/capability-providers/tree/main/kv-vault) | wasmCloud | wasmCloud key-value provider for the Hashicorp [Vault](https://www.vaultproject.io/docs/secrets/kv/kv-v2) secrets engine.

## Example Usage (🦀 Rust)

