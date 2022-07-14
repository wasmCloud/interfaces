[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-factorial.svg)](https://crates.io/crates/wasmcloud-interface-factorial)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=factorial%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/factorial/tinygo)
# wasmCloud Factorial Interface
This is the definition for the interface used for examples and illustrations with the contract ID of `wasmcloud:example:factorial`.

This is an interface for a simple service that calculates the factorial of a whole number. 

**NOTE** that this is just an example, and we would not recommend a real-world production scenario where you use an interface and accompanying capability provider for factorial calculations.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:example:factorial` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [Factorial](https://github.com/wasmCloud/examples/tree/main/provider/factorial) | wasmCloud | wasmCloud example implementation of the Factorial interface

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)
