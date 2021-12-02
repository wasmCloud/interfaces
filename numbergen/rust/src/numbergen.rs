// This file is generated automatically using wasmcloud-weld and smithy model definitions
//

#![allow(clippy::ptr_arg)]
#[allow(unused_imports)]
use async_trait::async_trait;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use std::{borrow::Cow, string::ToString};
#[allow(unused_imports)]
use wasmbus_rpc::{
    deserialize, serialize, Context, Message, MessageDispatch, RpcError, RpcResult, SendOpts,
    Timestamp, Transport,
};

pub const SMITHY_VERSION: &str = "1.0";

/// Input range for RandomInRange, inclusive. Result will be >= min and <= max
/// Example:
/// random_in_range(RangeLimit{0,4}) returns one the values, 0, 1, 2, 3, or 4.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RangeLimit {
    pub max: u32,
    pub min: u32,
}

/// wasmbus.contractId: wasmcloud:builtin:numbergen
/// wasmbus.providerReceive
#[async_trait]
pub trait NumberGen {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:builtin:numbergen"
    }
    ///
    /// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
    /// These guids are known as "version 4", meaning all bits are random or pseudo-random.
    ///
    async fn generate_guid(&self, ctx: &Context) -> RpcResult<String>;
    /// Request a random integer within a range
    /// The result will will be in the range [min,max), i.e., >= min and < max.
    async fn random_in_range(&self, ctx: &Context, arg: &RangeLimit) -> RpcResult<u32>;
    /// Request a 32-bit random number
    async fn random_32(&self, ctx: &Context) -> RpcResult<u32>;
}

/// NumberGenReceiver receives messages defined in the NumberGen service trait
#[doc(hidden)]
#[async_trait]
pub trait NumberGenReceiver: MessageDispatch + NumberGen {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "GenerateGuid" => {
                let resp = NumberGen::generate_guid(self, ctx).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "NumberGen.GenerateGuid",
                    arg: buf,
                })
            }
            "RandomInRange" => {
                let value: RangeLimit = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = NumberGen::random_in_range(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "NumberGen.RandomInRange",
                    arg: buf,
                })
            }
            "Random32" => {
                let resp = NumberGen::random_32(self, ctx).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "NumberGen.Random32",
                    arg: buf,
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "NumberGen::{}",
                message.method
            ))),
        }
    }
}

/// NumberGenSender sends messages to a NumberGen service
/// client for sending NumberGen messages
#[derive(Debug)]
pub struct NumberGenSender<T: Transport> {
    transport: T,
}

impl<T: Transport> NumberGenSender<T> {
    /// Constructs a NumberGenSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }
}

#[cfg(target_arch = "wasm32")]
impl NumberGenSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a NumberGen provider
    /// implementing the 'wasmcloud:builtin:numbergen' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:builtin:numbergen",
            "default",
        )
        .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a NumberGen provider
    /// implementing the 'wasmcloud:builtin:numbergen' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:builtin:numbergen",
            link_name,
        )?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> NumberGen for NumberGenSender<T> {
    #[allow(unused)]
    ///
    /// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
    /// These guids are known as "version 4", meaning all bits are random or pseudo-random.
    ///
    async fn generate_guid(&self, ctx: &Context) -> RpcResult<String> {
        let arg = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.GenerateGuid",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "GenerateGuid", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a random integer within a range
    /// The result will will be in the range [min,max), i.e., >= min and < max.
    async fn random_in_range(&self, ctx: &Context, arg: &RangeLimit) -> RpcResult<u32> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.RandomInRange",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "RandomInRange", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a 32-bit random number
    async fn random_32(&self, ctx: &Context) -> RpcResult<u32> {
        let arg = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.Random32",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "Random32", e)))?;
        Ok(value)
    }
}
