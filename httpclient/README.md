[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-httpclient.svg)](https://crates.io/crates/wasmcloud-interface-httpclient)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=httpclient%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/httpclient/tinygo)
# wasmCloud HTTP Client Interface
This is the interface definition for the interface with the contract ID `wasmcloud:httpclient`.

Actors utilizing this interface can make HTTP requests and receive HTTP responses for processing. Since this is just an interface, and not an actual provider, you will need to check the documentation for individual provider implementations for a list of link definition values supported by that provider.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:httpclient` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [HTTPClient](https://github.com/wasmCloud/capability-providers/tree/main/httpclient) | wasmCloud | wasmCloud implementation of the HTTP Client Provider

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)
