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

/// The Factorial service has a single method, calculate, which
/// calculates the factorial of its whole number parameter.
/// wasmbus.contractId: wasmcloud:example:factorial
/// wasmbus.providerReceive
/// wasmbus.actorReceive
#[async_trait]
pub trait Factorial {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:example:factorial"
    }
    /// Calculates the factorial (n!) of the input parameter
    async fn calculate(&self, ctx: &Context, arg: &u32) -> RpcResult<u64>;
}

/// FactorialReceiver receives messages defined in the Factorial service trait
/// The Factorial service has a single method, calculate, which
/// calculates the factorial of its whole number parameter.
#[doc(hidden)]
#[async_trait]
pub trait FactorialReceiver: MessageDispatch + Factorial {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Calculate" => {
                let value: u32 = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = Factorial::calculate(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Factorial.Calculate",
                    arg: buf,
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "Factorial::{}",
                message.method
            ))),
        }
    }
}

/// FactorialSender sends messages to a Factorial service
/// The Factorial service has a single method, calculate, which
/// calculates the factorial of its whole number parameter.
/// client for sending Factorial messages
#[derive(Debug)]
pub struct FactorialSender<T: Transport> {
    transport: T,
}

impl<T: Transport> FactorialSender<T> {
    /// Constructs a FactorialSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }
}

#[cfg(not(target_arch = "wasm32"))]
impl<'send> FactorialSender<wasmbus_rpc::provider::ProviderTransport<'send>> {
    /// Constructs a Sender using an actor's LinkDefinition,
    /// Uses the provider's HostBridge for rpc
    pub fn for_actor(ld: &'send wasmbus_rpc::core::LinkDefinition) -> Self {
        Self {
            transport: wasmbus_rpc::provider::ProviderTransport::new(ld, None),
        }
    }
}
#[cfg(target_arch = "wasm32")]
impl FactorialSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for actor-to-actor messaging
    /// using the recipient actor's public key
    pub fn to_actor(actor_id: &str) -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_actor(actor_id.to_string()).unwrap();
        Self { transport }
    }
}

#[cfg(target_arch = "wasm32")]
impl FactorialSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a Factorial provider
    /// implementing the 'wasmcloud:example:factorial' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:example:factorial",
            "default",
        )
        .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a Factorial provider
    /// implementing the 'wasmcloud:example:factorial' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:example:factorial",
            link_name,
        )?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> Factorial for FactorialSender<T> {
    #[allow(unused)]
    /// Calculates the factorial (n!) of the input parameter
    async fn calculate(&self, ctx: &Context, arg: &u32) -> RpcResult<u64> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Factorial.Calculate",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "Calculate", e)))?;
        Ok(value)
    }
}
