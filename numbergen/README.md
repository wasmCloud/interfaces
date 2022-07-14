[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-logging.svg)](https://crates.io/crates/wasmcloud-interface-logging)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=logging%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/logging/tinygo)
# wasmCloud Number Generator Interface
This is the interface definition for the wasmCloud built-in interface that is guaranteed to be supported by all runtime hosts, `wasmcloud:builtin:numbergen`. The number generator interface provides for the creation of things like random numbers, random numbers within a given range, and globally unique identifiers (GUIDs).

## Capability Provider Implementations
There are no external implementations of this provider as all implementations of the `wasmcloud:builtin:numbergen` contract are built directly into the host runtime(s).

This interface defines the wasmCloud built-in logging interface that comes with each of our supported host runtimes. Actors that use this interface must have the capability contract `wasmcloud:builtin:numbergen` in their claims list (`wash claims sign --cap wasmcloud:builtin:numbergen`).

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)