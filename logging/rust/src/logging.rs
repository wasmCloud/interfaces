// This file is generated automatically using wasmcloud/weld-codegen 0.4.2

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

// Encode LogEntry as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_log_entry<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &LogEntry,
) -> RpcResult<()> {
    e.array(2)?;
    e.str(&val.level)?;
    e.str(&val.text)?;
    Ok(())
}

// Decode LogEntry from cbor input stream
#[doc(hidden)]
pub fn decode_log_entry(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<LogEntry, RpcError> {
    let __result = {
        let mut level: Option<String> = None;
        let mut text: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct LogEntry, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct LogEntry: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => level = Some(d.str()?.to_string()),
                    1 => text = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct LogEntry: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "level" => level = Some(d.str()?.to_string()),
                    "text" => text = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        LogEntry {
            level: if let Some(__x) = level {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field LogEntry.level (#0)".to_string(),
                ));
            },

            text: if let Some(__x) = text {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field LogEntry.text (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
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
    async fn dispatch<'disp__, 'ctx__, 'msg__>(
        &'disp__ self,
        ctx: &'ctx__ Context,
        message: &Message<'msg__>,
    ) -> Result<Message<'msg__>, RpcError> {
        match message.method {
            "WriteLog" => {
                let value: LogEntry = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'LogEntry': {}", e)))?;
                let _resp = Logging::write_log(self, ctx, &value).await?;
                let buf = Vec::new();
                Ok(Message {
                    method: "Logging.WriteLog",
                    arg: Cow::Owned(buf),
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

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
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
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::error::RpcResult<Self> {
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
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Logging.WriteLog",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        Ok(())
    }
}
