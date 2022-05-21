// This file is generated automatically using wasmcloud/weld-codegen 0.4.3

#[allow(unused_imports)]
use async_trait::async_trait;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use std::{borrow::Borrow, borrow::Cow, io::Write, string::ToString};
#[allow(unused_imports)]
use wasmbus_rpc::{
    cbor::*,
    common::{
        deserialize, message_format, serialize, Context, Message, MessageDispatch, MessageFormat,
        SendOpts, Transport,
    },
    error::{RpcError, RpcResult},
    Timestamp,
};

#[allow(dead_code)]
pub const SMITHY_VERSION: &str = "1.0";

/// Input range for RandomInRange, inclusive. Result will be >= min and <= max
/// Example:
/// random_in_range(RangeLimit{0,4}) returns one the values, 0, 1, 2, 3, or 4.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RangeLimit {
    #[serde(default)]
    pub min: u32,
    #[serde(default)]
    pub max: u32,
}

// Encode RangeLimit as CBOR and append to output stream
#[doc(hidden)]
#[allow(unused_mut)]
pub fn encode_range_limit<W: wasmbus_rpc::cbor::Write>(
    mut e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &RangeLimit,
) -> RpcResult<()> {
    e.array(2)?;
    e.u32(val.min)?;
    e.u32(val.max)?;
    Ok(())
}

// Decode RangeLimit from cbor input stream
#[doc(hidden)]
pub fn decode_range_limit(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<RangeLimit, RpcError> {
    let __result = {
        let mut min: Option<u32> = None;
        let mut max: Option<u32> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct RangeLimit, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.fixed_array()?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => min = Some(d.u32()?),
                    1 => max = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.fixed_map()?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "min" => min = Some(d.u32()?),
                    "max" => max = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        }
        RangeLimit {
            min: if let Some(__x) = min {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RangeLimit.min (#0)".to_string(),
                ));
            },

            max: if let Some(__x) = max {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RangeLimit.max (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
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
    async fn dispatch<'disp__, 'ctx__, 'msg__>(
        &'disp__ self,
        ctx: &'ctx__ Context,
        message: &Message<'msg__>,
    ) -> Result<Message<'msg__>, RpcError> {
        match message.method {
            "GenerateGuid" => {
                let resp = NumberGen::generate_guid(self, ctx).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "NumberGen.GenerateGuid",
                    arg: Cow::Owned(buf),
                })
            }
            "RandomInRange" => {
                let value: RangeLimit = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'RangeLimit': {}", e)))?;
                let resp = NumberGen::random_in_range(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "NumberGen.RandomInRange",
                    arg: Cow::Owned(buf),
                })
            }
            "Random32" => {
                let resp = NumberGen::random_32(self, ctx).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "NumberGen.Random32",
                    arg: Cow::Owned(buf),
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

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
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
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::error::RpcResult<Self> {
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
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.GenerateGuid",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: String = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': String", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a random integer within a range
    /// The result will will be in the range [min,max), i.e., >= min and < max.
    async fn random_in_range(&self, ctx: &Context, arg: &RangeLimit) -> RpcResult<u32> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.RandomInRange",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: u32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': U32", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a 32-bit random number
    async fn random_32(&self, ctx: &Context) -> RpcResult<u32> {
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "NumberGen.Random32",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: u32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': U32", e)))?;
        Ok(value)
    }
}
