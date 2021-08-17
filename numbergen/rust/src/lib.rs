//! wasmcloud:builtin:numbergen

mod numbergen;
pub use numbergen::*;

#[allow(unused_imports)]
use wasmbus_rpc::RpcResult;

// functions below are convenience wrappers around the interface invocations.

#[cfg(target_arch = "wasm32")]
/// Generate a v4-format guid in the form "nnnnnnnn-nnnn-nnnn-nnnn-nnnnnnnnnnnn"
/// where n is a lowercase hex digit and ann bits are random.
pub async fn generate_guid() -> RpcResult<String> {
    use wasmbus_rpc::{actor::prelude::WasmHost, Context};

    let tx = WasmHost::to_provider("wasmcloud:builtin:numbergen", "default").unwrap();
    let ctx = Context::default();
    let ng = NumberGenSender::new(&tx);
    ng.generate_guid(&ctx).await
}

#[cfg(target_arch = "wasm32")]
/// Generate a random integer within an inclusive range. ( min <= n <= max )
pub async fn random_in_range(min: u32, max: u32) -> RpcResult<u32> {
    use wasmbus_rpc::{actor::prelude::WasmHost, Context};

    let tx = WasmHost::to_provider("wasmcloud:builtin:numbergen", "default").unwrap();
    let ctx = Context::default();
    let ng = NumberGenSender::new(&tx);
    ng.random_in_range(&ctx, &RangeLimit { min, max }).await
}

#[cfg(target_arch = "wasm32")]
/// Generate a 32-bit random number
pub async fn random_32() -> RpcResult<u32> {
    use wasmbus_rpc::{actor::prelude::WasmHost, Context};

    let tx = WasmHost::to_provider("wasmcloud:builtin:numbergen", "default").unwrap();
    let ctx = Context::default();
    let ng = NumberGenSender::new(&tx);
    ng.random_32(&ctx).await
}
