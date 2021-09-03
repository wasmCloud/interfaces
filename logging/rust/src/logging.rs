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
    Transport,
};

pub const SMITHY_VERSION: &str = "1.0";

#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct LogEntry {
    /// severity level: debug,info,warn,error
    #[serde(default)]
    pub level: String,
    /// message to log
    #[serde(default)]
    pub text: String,
}

/// wasmbus.contractId: wasmcloud:builtin:logging
/// wasmbus.providerReceive
#[async_trait]
pub trait Logging {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:builtin:logging"
    }
    ///
    /// WriteLog - log a text message
    ///
    async fn write_log(&self, ctx: &Context, arg: &LogEntry) -> RpcResult<()>;
}

/// LoggingReceiver receives messages defined in the Logging service trait
#[doc(hidden)]
#[async_trait]
pub trait LoggingReceiver: MessageDispatch + Logging {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "WriteLog" => {
                let value: LogEntry = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = Logging::write_log(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Logging.WriteLog",
                    arg: buf,
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "Logging::{}",
                message.method
            ))),
        }
    }
}

/// LoggingSender sends messages to a Logging service
/// client for sending Logging messages
#[derive(Debug)]
pub struct LoggingSender<T: Transport> {
    transport: T,
}

impl<T: Transport> LoggingSender<T> {
    /// Constructs a LoggingSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }
}

#[cfg(target_arch = "wasm32")]
impl LoggingSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a Logging provider
    /// implementing the 'wasmcloud:builtin:logging' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:builtin:logging",
            "default",
        )
        .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a Logging provider
    /// implementing the 'wasmcloud:builtin:logging' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport = wasmbus_rpc::actor::prelude::WasmHost::to_provider(
            "wasmcloud:builtin:logging",
            link_name,
        )?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> Logging for LoggingSender<T> {
    #[allow(unused)]
    ///
    /// WriteLog - log a text message
    ///
    async fn write_log(&self, ctx: &Context, arg: &LogEntry) -> RpcResult<()> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Logging.WriteLog",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        Ok(())
    }
}
