// This file is generated automatically using wasmcloud/weld-codegen 0.3.0

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

/// A message to be published
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PubMessage {
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
}

// Encode PubMessage as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_pub_message<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PubMessage,
) -> RpcResult<()> {
    e.array(3)?;
    e.str(&val.subject)?;
    if let Some(val) = val.reply_to.as_ref() {
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.bytes(&val.body)?;
    Ok(())
}

// Decode PubMessage from cbor input stream
#[doc(hidden)]
pub fn decode_pub_message(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<PubMessage, RpcError> {
    let __result = {
        let mut subject: Option<String> = None;
        let mut reply_to: Option<Option<String>> = Some(None);
        let mut body: Option<Vec<u8>> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct PubMessage, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PubMessage: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => subject = Some(d.str()?.to_string()),
                    1 => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    2 => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PubMessage: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "subject" => subject = Some(d.str()?.to_string()),
                    "replyTo" => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "body" => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        }
        PubMessage {
            subject: if let Some(__x) = subject {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field PubMessage.subject (#0)".to_string(),
                ));
            },
            reply_to: reply_to.unwrap(),

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field PubMessage.body (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Reply received from a Request operation
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ReplyMessage {
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
}

// Encode ReplyMessage as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_reply_message<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ReplyMessage,
) -> RpcResult<()> {
    e.array(3)?;
    e.str(&val.subject)?;
    if let Some(val) = val.reply_to.as_ref() {
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.bytes(&val.body)?;
    Ok(())
}

// Decode ReplyMessage from cbor input stream
#[doc(hidden)]
pub fn decode_reply_message(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ReplyMessage, RpcError> {
    let __result = {
        let mut subject: Option<String> = None;
        let mut reply_to: Option<Option<String>> = Some(None);
        let mut body: Option<Vec<u8>> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ReplyMessage, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ReplyMessage: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => subject = Some(d.str()?.to_string()),
                    1 => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    2 => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ReplyMessage: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "subject" => subject = Some(d.str()?.to_string()),
                    "replyTo" => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "body" => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        }
        ReplyMessage {
            subject: if let Some(__x) = subject {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ReplyMessage.subject (#0)".to_string(),
                ));
            },
            reply_to: reply_to.unwrap(),

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ReplyMessage.body (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Message sent as part of a request, with timeout
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RequestMessage {
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// A timeout, in milliseconds
    #[serde(rename = "timeoutMs")]
    #[serde(default)]
    pub timeout_ms: u32,
}

// Encode RequestMessage as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_request_message<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &RequestMessage,
) -> RpcResult<()> {
    e.array(3)?;
    e.str(&val.subject)?;
    e.bytes(&val.body)?;
    e.u32(val.timeout_ms)?;
    Ok(())
}

// Decode RequestMessage from cbor input stream
#[doc(hidden)]
pub fn decode_request_message(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<RequestMessage, RpcError> {
    let __result = {
        let mut subject: Option<String> = None;
        let mut body: Option<Vec<u8>> = None;
        let mut timeout_ms: Option<u32> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct RequestMessage, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct RequestMessage: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => subject = Some(d.str()?.to_string()),
                    1 => body = Some(d.bytes()?.to_vec()),
                    2 => timeout_ms = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct RequestMessage: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "subject" => subject = Some(d.str()?.to_string()),
                    "body" => body = Some(d.bytes()?.to_vec()),
                    "timeoutMs" => timeout_ms = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        }
        RequestMessage {
            subject: if let Some(__x) = subject {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RequestMessage.subject (#0)".to_string(),
                ));
            },

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RequestMessage.body (#1)".to_string(),
                ));
            },

            timeout_ms: if let Some(__x) = timeout_ms {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RequestMessage.timeout_ms (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Message received as part of a subscription
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SubMessage {
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
}

// Encode SubMessage as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_sub_message<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &SubMessage,
) -> RpcResult<()> {
    e.array(3)?;
    e.str(&val.subject)?;
    if let Some(val) = val.reply_to.as_ref() {
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.bytes(&val.body)?;
    Ok(())
}

// Decode SubMessage from cbor input stream
#[doc(hidden)]
pub fn decode_sub_message(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<SubMessage, RpcError> {
    let __result = {
        let mut subject: Option<String> = None;
        let mut reply_to: Option<Option<String>> = Some(None);
        let mut body: Option<Vec<u8>> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct SubMessage, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SubMessage: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => subject = Some(d.str()?.to_string()),
                    1 => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    2 => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SubMessage: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "subject" => subject = Some(d.str()?.to_string()),
                    "replyTo" => {
                        reply_to = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "body" => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        }
        SubMessage {
            subject: if let Some(__x) = subject {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SubMessage.subject (#0)".to_string(),
                ));
            },
            reply_to: reply_to.unwrap(),

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SubMessage.body (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
/// wasmbus.contractId: wasmcloud:messaging
/// wasmbus.actorReceive
#[async_trait]
pub trait MessageSubscriber {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:messaging"
    }
    /// subscription handler
    async fn handle_message(&self, ctx: &Context, arg: &SubMessage) -> RpcResult<()>;
}

/// MessageSubscriberReceiver receives messages defined in the MessageSubscriber service trait
/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
#[doc(hidden)]
#[async_trait]
pub trait MessageSubscriberReceiver: MessageDispatch + MessageSubscriber {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "HandleMessage" => {
                let value: SubMessage = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'SubMessage': {}", e)))?;
                let _resp = MessageSubscriber::handle_message(self, ctx, &value).await?;
                let buf = Vec::new();
                Ok(Message {
                    method: "MessageSubscriber.HandleMessage",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "MessageSubscriber::{}",
                message.method
            ))),
        }
    }
}

/// MessageSubscriberSender sends messages to a MessageSubscriber service
/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
/// client for sending MessageSubscriber messages
#[derive(Debug)]
pub struct MessageSubscriberSender<T: Transport> {
    transport: T,
}

impl<T: Transport> MessageSubscriberSender<T> {
    /// Constructs a MessageSubscriberSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(not(target_arch = "wasm32"))]
impl<'send> MessageSubscriberSender<wasmbus_rpc::provider::ProviderTransport<'send>> {
    /// Constructs a Sender using an actor's LinkDefinition,
    /// Uses the provider's HostBridge for rpc
    pub fn for_actor(ld: &'send wasmbus_rpc::core::LinkDefinition) -> Self {
        Self {
            transport: wasmbus_rpc::provider::ProviderTransport::new(ld, None),
        }
    }
}
#[cfg(target_arch = "wasm32")]
impl MessageSubscriberSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for actor-to-actor messaging
    /// using the recipient actor's public key
    pub fn to_actor(actor_id: &str) -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_actor(actor_id.to_string()).unwrap();
        Self { transport }
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> MessageSubscriber
    for MessageSubscriberSender<T>
{
    #[allow(unused)]
    /// subscription handler
    async fn handle_message(&self, ctx: &Context, arg: &SubMessage) -> RpcResult<()> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "MessageSubscriber.HandleMessage",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        Ok(())
    }
}

/// The Messaging interface describes a service
/// that can deliver messages
/// wasmbus.contractId: wasmcloud:messaging
/// wasmbus.providerReceive
#[async_trait]
pub trait Messaging {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:messaging"
    }
    /// Publish - send a message
    /// The function returns after the message has been sent.
    /// If the sender expects to receive an asynchronous reply,
    /// the replyTo field should be filled with the
    /// subject for the response.
    async fn publish(&self, ctx: &Context, arg: &PubMessage) -> RpcResult<()>;
    /// Request - send a message in a request/reply pattern,
    /// waiting for a response.
    async fn request(&self, ctx: &Context, arg: &RequestMessage) -> RpcResult<ReplyMessage>;
}

/// MessagingReceiver receives messages defined in the Messaging service trait
/// The Messaging interface describes a service
/// that can deliver messages
#[doc(hidden)]
#[async_trait]
pub trait MessagingReceiver: MessageDispatch + Messaging {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Publish" => {
                let value: PubMessage = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'PubMessage': {}", e)))?;
                let _resp = Messaging::publish(self, ctx, &value).await?;
                let buf = Vec::new();
                Ok(Message {
                    method: "Messaging.Publish",
                    arg: Cow::Owned(buf),
                })
            }
            "Request" => {
                let value: RequestMessage = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'RequestMessage': {}", e)))?;
                let resp = Messaging::request(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "Messaging.Request",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "Messaging::{}",
                message.method
            ))),
        }
    }
}

/// MessagingSender sends messages to a Messaging service
/// The Messaging interface describes a service
/// that can deliver messages
/// client for sending Messaging messages
#[derive(Debug)]
pub struct MessagingSender<T: Transport> {
    transport: T,
}

impl<T: Transport> MessagingSender<T> {
    /// Constructs a MessagingSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(target_arch = "wasm32")]
impl MessagingSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a Messaging provider
    /// implementing the 'wasmcloud:messaging' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:messaging", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a Messaging provider
    /// implementing the 'wasmcloud:messaging' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:messaging", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> Messaging for MessagingSender<T> {
    #[allow(unused)]
    /// Publish - send a message
    /// The function returns after the message has been sent.
    /// If the sender expects to receive an asynchronous reply,
    /// the replyTo field should be filled with the
    /// subject for the response.
    async fn publish(&self, ctx: &Context, arg: &PubMessage) -> RpcResult<()> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Messaging.Publish",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        Ok(())
    }
    #[allow(unused)]
    /// Request - send a message in a request/reply pattern,
    /// waiting for a response.
    async fn request(&self, ctx: &Context, arg: &RequestMessage) -> RpcResult<ReplyMessage> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Messaging.Request",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ReplyMessage = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': ReplyMessage", e)))?;
        Ok(value)
    }
}
