[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-lattice-control.svg)](https://crates.io/crates/wasmcloud-interface-lattice-control)&nbsp;
<!-- [![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=lattice-control%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/lattice-control/tinygo) -->
# wasmCloud Lattice Control Interface
The lattice control interface is a smithy-defined interface contract that is expected to be consumed in one of two different ways:

* Directly - A [NATS client](https://github.com/wasmcloud/control-interface-client) library may use the data structures from this interface to communicate over the lattice control interface topic
* Indirectly - Either side of the `wasmcloud:latticecontrol` contract
    * Capability Providers - Capability providers can simply provide a wrapper around the NATS client, exposing lattice control functionality to actors
    * Actors - Actors can make use of this crate as they would any other wasmCloud interface crate

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:latticecontrol` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [Lattice Controller](https://github.com/wasmCloud/capability-providers/tree/main/lattice-controller) | wasmCloud | Implementation using the [wasmcloud-control-interface](https://crates.io/crates/wasmcloud-control-interface) crate.

## Example Usage (🦀 Rust)

<!-- ## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo) -->