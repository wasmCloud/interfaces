[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-keyvalue.svg)](https://crates.io/crates/wasmcloud-interface-keyvalue)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=keyvalue%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/keyvalue/tinygo)
# wasmCloud Messaging Interface
This is the interface for the `wasmcloud:messaging` contract. This contract is a very simple abstraction over the concept of a message broker. This contract does not have controls or knobs or settings for things like delivery guarantees (e.g. "at most once" vs "at least once"), persistence of messages, etc. Such details are the responsibility of individual providers.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:messaging` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [NATS](https://github.com/wasmCloud/capability-providers/tree/main/nats) | wasmCloud | wasmCloud Messaging Provider for the [NATS](https://nats.io) broker

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)