[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-logging.svg)](https://crates.io/crates/wasmcloud-interface-logging)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=logging%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/logging/tinygo)
# wasmCloud Builtin Logging Interface
This interface defines the wasmCloud built-in logging interface that comes with each of our supported host runtimes. Actors that use this interface must have the capability contract `wasmcloud:builtin:logging` in their claims list (`wash claims sign --logging`).

## Capability Provider Implementations

There are no external implementations for this provider as they are built directly into the host runtime.

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)