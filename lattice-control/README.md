![Crates.io](https://img.shields.io/crates/v/wasmcloud-interface-lattice-control)
[![Documentation](https://img.shields.io/badge/Docs-Documentation-blue)](https://wasmcloud.dev)
[![Rustdocs](https://docs.rs/lattice-control-interface/badge.svg)](https://docs.rs/wasmcloud-interface-lattice-control)

# Lattice Control Interface
The lattice control interface is a smithy-defined interface contract that is expected to be consumed in one of two different ways:

* Directly - A [NATS client](https://github.com/wasmcloud/control-interface-client) library may use the data structures from this interface to communicate over the lattice control interface topic
* Indirectly - Either side of the `wasmcloud:latticecontrol` contract
    * Capability Providers - Capability providers can simply provide a wrapper around the NATS client, exposing lattice control functionality to actors
    * Actors - Actors can make use of this crate as they would any other wasmCloud interface crate

