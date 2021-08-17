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

#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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
    ///
    /// WriteLog - log a text message
    ///
    async fn write_log(&self, ctx: &Context, arg: &LogEntry) -> RpcResult<()>;
}

/// LoggingReceiver receives messages defined in the Logging service trait
#[async_trait]
pub trait LoggingReceiver: MessageDispatch + Logging {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "WriteLog" => {
                let value: LogEntry = deserialize(message.arg.as_ref()).map_err(|e| {
                    RpcError::Deser(format!(
                        "deserialization for message '{}': {}",
                        message.method, e
                    ))
                })?;
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
#[derive(Debug)]
pub struct LoggingSender<'send, T> {
    transport: &'send T,
}

impl<'send, T: Transport> LoggingSender<'send, T> {
    pub fn new(transport: &'send T) -> Self {
        LoggingSender { transport }
    }
}

#[async_trait]
impl<'send, T: Transport + std::marker::Sync + std::marker::Send> Logging
    for LoggingSender<'send, T>
{
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
                    method: "WriteLog",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        Ok(())
    }
}
