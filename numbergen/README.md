[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-numbergen.svg)](https://crates.io/crates/wasmcloud-interface-numbergen)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=numbergen%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/numbergen/tinygo)

# wasmCloud Number Generator Interface
This is the interface definition for the wasmCloud built-in interface that is guaranteed to be supported by all runtime hosts, `wasmcloud:builtin:numbergen`. The number generator interface provides for the creation of things like random numbers, random numbers within a given range, and globally unique identifiers (GUIDs).

## Capability Provider Implementations
There are no external implementations of this provider as all implementations of the `wasmcloud:builtin:numbergen` contract are built directly into the host runtime(s).

This interface defines the wasmCloud built-in logging interface that comes with each of our supported host runtimes. Actors that use this interface must have the capability contract `wasmcloud:builtin:numbergen` in their claims list (`wash claims sign --cap wasmcloud:builtin:numbergen`).

## Example Usage 
### 🦀 Rust
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_logging::info;
use wasmcloud_interface_numbergen::{generate_guid, random_32, random_in_range};

async fn generate_random() -> Result<(), RpcError> {
    // Generate a Globally Unique ID (GUID)
    let guid: String = generate_guid().await?;
    info!("Generated GUID: {}", guid);
    // Generate a random u32
    let random_num: u32 = random_32().await?;
    info!("Generated number: {}", random_num);
    // Generate a random u32 within an inclusive range
    let random_range: u32 = random_in_range(0, 55).await?;
    info!("Generated number between 0 and 55: {}", random_range);
    Ok(())
}
```
### 🐭 Golang
```go
import numbergen "github.com/wasmcloud/interfaces/numbergen/tinygo"

func GenerateThings(ctx *actor.Context) (string, error) {
  client := numbergen.NewProviderNumberGen()

  // Generate a random U32
  randNum, _ := client.Random32(ctx)
 
  // Generate a random U32 within an inclusive range
  randNum010, _ := client.RandomInRange(ctx, numbergen.RangeLimit{Min: 0, Max: 10})
 
  // Generate a Globally Unique ID (GUID)
  guid, _ := client.GenerateGuid(ctx)
  
  return "", nil
}
```
