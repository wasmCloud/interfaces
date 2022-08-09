[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-lattice-control.svg)](https://crates.io/crates/wasmcloud-interface-lattice-control)&nbsp;
# wasmCloud Lattice Control Interface
The lattice control interface is a smithy-defined interface contract that outlines the operations and data structures supported by a capability provider supporting the `wasmcloud:latticecontrol` contract.

* Capability Providers - Capability providers can simply provide a wrapper around a NATS client, exposing lattice control functionality to actors
* Actors - Actors can make use of this crate as they would any other wasmCloud interface crate, thus enabling an actor to contain business logic that manipulates a lattice.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:latticecontrol` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [Lattice Controller](https://github.com/wasmCloud/capability-providers/tree/main/lattice-controller) | wasmCloud | First party implementation of the lattice controller provider

## Example Usage (🦀 Rust)

Start 250 instances of the echo actor actor on a host
```rust
use wasmbus_rpc::actor::prelude::{Context, RpcResult};
use wasmcloud_interface_lattice_control::{
    CtlOperationAck, LatticeController, LatticeControllerSender, StartActorCommand,
};
use wasmcloud_interface_logging::debug;

async fn start_actor(ctx: &Context) -> RpcResult<CtlOperationAck> {
    let lattice = LatticeControllerSender::new();
    let cmd = StartActorCommand {
        lattice_id: "default".to_string(),
        actor_ref: "wasmcloud.azurecr.io/echo:0.3.4".to_string(),
        annotations: None,
        count: 250,
        host_id: "NB67YNOVU5YB3526RUNCKNZBCQDH2L5NZJKQ6FWOVWGSHNHHEO65RP4A".to_string(),
    };

    debug!(
        "Starting {} instance(s) of actor {} on host {}",
        cmd.count, cmd.actor_ref, cmd.host_id
    );

    lattice.start_actor(ctx, &cmd).await
}

```

Get all hosts in a lattice
```rust
use wasmbus_rpc::actor::prelude::{Context, RpcResult};
use wasmcloud_interface_lattice_control::{Host, LatticeController, LatticeControllerSender};
use wasmcloud_interface_logging::info;

async fn get_hosts(ctx: &Context) -> RpcResult<Vec<Host>> {
    let lattice = LatticeControllerSender::new();
    let hosts = lattice.get_hosts(ctx, GetHostsRequest {
        lattice_id: "default".to_string()
    }).await?;

    info!("There are {} hosts in this lattice", hosts.len());
    Ok(hosts)
}
```