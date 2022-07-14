[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-numbergen.svg)](https://crates.io/crates/wasmcloud-interface-numbergen)&nbsp;

# wasmCloud Number Generator Interface
This is the interface definition for the wasmCloud built-in interface that is guaranteed to be supported by all runtime hosts, `wasmcloud:builtin:numbergen`. The number generator interface provides for the creation of things like random numbers, random numbers within a given range, and globally unique identifiers (GUIDs).

## Capability Provider Implementations
There are no external implementations of this provider as all implementations of the `wasmcloud:builtin:numbergen` contract are built directly into the host runtime(s).

This interface defines the wasmCloud built-in logging interface that comes with each of our supported host runtimes. Actors that use this interface must have the capability contract `wasmcloud:builtin:numbergen` in their claims list (`wash claims sign --cap wasmcloud:builtin:numbergen`).

## Example Usage (🦀 Rust)

