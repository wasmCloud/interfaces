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

/// A message to be published
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PubMessage {
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
}

/// Reply received from a Request operation
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ReplyMessage {
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
}

/// Message sent as part of a request, with timeout
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RequestMessage {
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
    /// A timeout, in milliseconds
    #[serde(rename = "timeoutMs")]
    pub timeout_ms: u32,
}

/// Message received as part of a subscription
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SubMessage {
    /// The message payload
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// An optional topic on which the reply should be sent.
    #[serde(rename = "replyTo")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub reply_to: Option<String>,
    /// The subject, or topic, of the message
    #[serde(default)]
    pub subject: String,
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
                let value: SubMessage = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = MessageSubscriber::handle_message(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "MessageSubscriber.HandleMessage",
                    arg: buf,
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
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "MessageSubscriber.HandleMessage",
                    arg: Cow::Borrowed(&arg),
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
                let value: PubMessage = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = Messaging::publish(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Messaging.Publish",
                    arg: buf,
                })
            }
            "Request" => {
                let value: RequestMessage = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = Messaging::request(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Messaging.Request",
                    arg: buf,
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
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Messaging.Publish",
                    arg: Cow::Borrowed(&arg),
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
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Messaging.Request",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "Request", e)))?;
        Ok(value)
    }
}
