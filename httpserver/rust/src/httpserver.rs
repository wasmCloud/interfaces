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

/// map data structure for holding http headers
///
pub type HeaderMap = std::collections::HashMap<String, HeaderValues>;

pub type HeaderValues = Vec<String>;

/// HttpRequest contains data sent to actor about the http request
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpRequest {
    /// Request body as a byte array. May be empty.
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// map of request headers (string key, string value)
    pub header: HeaderMap,
    /// HTTP method. One of: GET,POST,PUT,DELETE,HEAD,OPTIONS,CONNECT,PATCH,TRACE
    #[serde(default)]
    pub method: String,
    /// full request path
    #[serde(default)]
    pub path: String,
    /// query string. May be an empty string if there were no query parameters.
    #[serde(rename = "queryString")]
    #[serde(default)]
    pub query_string: String,
}

/// HttpResponse contains the actor's response to return to the http client
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpResponse {
    /// Body of response as a byte array. May be an empty array.
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// Map of headers (string keys, list of values)
    pub header: HeaderMap,
    /// statusCode is a three-digit number, usually in the range 100-599,
    /// A value of 200 indicates success.
    #[serde(rename = "statusCode")]
    pub status_code: u16,
}

/// HttpServer is the contract to be implemented by actor
/// wasmbus.contractId: wasmcloud:httpserver
/// wasmbus.actorReceive
#[async_trait]
pub trait HttpServer {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:httpserver"
    }
    async fn handle_request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse>;
}

/// HttpServerReceiver receives messages defined in the HttpServer service trait
/// HttpServer is the contract to be implemented by actor
#[doc(hidden)]
#[async_trait]
pub trait HttpServerReceiver: MessageDispatch + HttpServer {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "HandleRequest" => {
                let value: HttpRequest = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = HttpServer::handle_request(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "HttpServer.HandleRequest",
                    arg: buf,
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "HttpServer::{}",
                message.method
            ))),
        }
    }
}

/// HttpServerSender sends messages to a HttpServer service
/// HttpServer is the contract to be implemented by actor
/// client for sending HttpServer messages
#[derive(Debug)]
pub struct HttpServerSender<T: Transport> {
    transport: T,
}

impl<T: Transport> HttpServerSender<T> {
    /// Constructs a HttpServerSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }
}
#[cfg(not(target_arch = "wasm32"))]
impl<'send> HttpServerSender<wasmbus_rpc::provider::ProviderTransport<'send>> {
    /// Constructs a Sender using an actor's LinkDefinition,
    /// Uses the provider's HostBridge for rpc
    pub fn for_actor(ld: &'send wasmbus_rpc::core::LinkDefinition) -> Self {
        Self {
            transport: wasmbus_rpc::provider::ProviderTransport::new(ld, None),
        }
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> HttpServer for HttpServerSender<T> {
    #[allow(unused)]
    async fn handle_request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "HttpServer.HandleRequest",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "HandleRequest", e)))?;
        Ok(value)
    }
}
