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
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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

/// The Messaging interface describes a service
/// that can deliver messages
/// wasmbus.contractId: wasmcloud:messaging
/// wasmbus.providerReceive
#[async_trait]
pub trait Messaging {
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
#[async_trait]
pub trait MessagingReceiver: MessageDispatch + Messaging {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Publish" => {
                let value: PubMessage = deserialize(message.arg.as_ref())?;
                let resp = Messaging::publish(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Messaging.Publish",
                    arg: buf,
                })
            }
            "Request" => {
                let value: RequestMessage = deserialize(message.arg.as_ref())?;
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
#[derive(Debug)]
pub struct MessagingSender<'send, T> {
    transport: &'send T,
}

impl<'send, T: Transport> MessagingSender<'send, T> {
    pub fn new(transport: &'send T) -> Self {
        MessagingSender { transport }
    }
}

#[async_trait]
impl<'send, T: Transport + std::marker::Sync + std::marker::Send> Messaging
    for MessagingSender<'send, T>
{
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
                    method: "Publish",
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
                    method: "Request",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)?;
        Ok(value)
    }
}

/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
/// wasmbus.contractId: wasmcloud:messaging
/// wasmbus.actorReceive
#[async_trait]
pub trait MessageSubscriber {
    /// subscription handler
    async fn handle_message(&self, ctx: &Context, arg: &SubMessage) -> RpcResult<()>;
}

/// MessageSubscriberReceiver receives messages defined in the MessageSubscriber service trait
/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
#[async_trait]
pub trait MessageSubscriberReceiver: MessageDispatch + MessageSubscriber {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "HandleMessage" => {
                let value: SubMessage = deserialize(message.arg.as_ref())?;
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
#[derive(Debug)]
pub struct MessageSubscriberSender<'send, T> {
    transport: &'send T,
}

impl<'send, T: Transport> MessageSubscriberSender<'send, T> {
    pub fn new(transport: &'send T) -> Self {
        MessageSubscriberSender { transport }
    }
}

#[async_trait]
impl<'send, T: Transport + std::marker::Sync + std::marker::Send> MessageSubscriber
    for MessageSubscriberSender<'send, T>
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
                    method: "HandleMessage",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        Ok(())
    }
}
