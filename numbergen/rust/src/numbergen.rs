// This file is generated automatically using wasmcloud-weld and smithy model definitions
//

#![allow(clippy::ptr_arg)]
#[allow(unused_imports)]
use async_trait::async_trait;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use std::borrow::Cow;
#[allow(unused_imports)]
use wasmbus_rpc::{
    client, context, deserialize, serialize, Message, MessageDispatch, RpcError, Transport,
};

pub const SMITHY_VERSION: &str = "1.0";

/// Input range for RandomInRange. Result will be >= min and < max
/// Example:
/// random_in_range(RangeLimit{0,5}) returns one the values, 0, 1, 2, 3, or 4.
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
pub struct RangeLimit {
    pub max: u32,
    pub min: u32,
}

/// wasmbus.contractId: wasmcloud:builtin:numbergen
/// wasmbus.providerReceive
#[async_trait]
pub trait NumberGen {
    ///
    /// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
    /// These guids are known as "version 4", meaning all bits are random or pseudo-random.
    ///
    async fn generate_guid(&self, ctx: &context::Context<'_>) -> Result<String, RpcError>;
    /// Request a random integer within a range
    /// The result will will be in the range [min,max), i.e., >= min and < max.
    async fn random_in_range(
        &self,
        ctx: &context::Context<'_>,
        arg: &RangeLimit,
    ) -> Result<u32, RpcError>;
    /// Request a 32-bit random number
    async fn random_32(&self, ctx: &context::Context<'_>) -> Result<u32, RpcError>;
}

/// NumberGenReceiver receives messages defined in the NumberGen service trait
#[async_trait]
pub trait NumberGenReceiver: MessageDispatch + NumberGen {
    async fn dispatch(
        &self,
        ctx: &context::Context<'_>,
        message: &Message<'_>,
    ) -> Result<Message<'_>, RpcError> {
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
                let value: RangeLimit = deserialize(message.arg.as_ref())?;
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
#[derive(Debug)]
pub struct NumberGenSender<T> {
    transport: T,
    config: client::SendConfig,
}

impl<T: Transport> NumberGenSender<T> {
    pub fn new(config: client::SendConfig, transport: T) -> Self {
        NumberGenSender { transport, config }
    }
}

#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> NumberGen for NumberGenSender<T> {
    #[allow(unused)]
    ///
    /// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
    /// These guids are known as "version 4", meaning all bits are random or pseudo-random.
    ///
    async fn generate_guid(&self, ctx: &context::Context<'_>) -> Result<String, RpcError> {
        let arg = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                &self.config,
                Message {
                    method: "GenerateGuid",
                    arg: Cow::Borrowed(&arg),
                },
            )
            .await?;
        let value = deserialize(resp.arg.as_ref())?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a random integer within a range
    /// The result will will be in the range [min,max), i.e., >= min and < max.
    async fn random_in_range(
        &self,
        ctx: &context::Context<'_>,
        arg: &RangeLimit,
    ) -> Result<u32, RpcError> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                &self.config,
                Message {
                    method: "RandomInRange",
                    arg: Cow::Borrowed(&arg),
                },
            )
            .await?;
        let value = deserialize(resp.arg.as_ref())?;
        Ok(value)
    }
    #[allow(unused)]
    /// Request a 32-bit random number
    async fn random_32(&self, ctx: &context::Context<'_>) -> Result<u32, RpcError> {
        let arg = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                &self.config,
                Message {
                    method: "Random32",
                    arg: Cow::Borrowed(&arg),
                },
            )
            .await?;
        let value = deserialize(resp.arg.as_ref())?;
        Ok(value)
    }
}
