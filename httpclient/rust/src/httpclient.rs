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

/// http request to be sent through the provider
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpRequest {
    /// request body, defaults to empty
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// optional headers. defaults to empty
    pub headers: HeaderMap,
    /// http method, defaults to "GET"
    #[serde(default)]
    pub method: String,
    #[serde(default)]
    pub url: String,
}

/// response from the http request
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpResponse {
    /// response body
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    /// Case is not guaranteed to be normalized, so
    /// actors checking response headers need to do their own
    /// case conversion.
    /// Example (rust):
    /// // check for 'Content-Type' header
    /// let content_type:Option<&Vec<String>> = header.iter()
    /// .map(|(k,_)| k.to_ascii_lowercase())
    /// .find(|(k,_)| k == "content-type")
    /// .map(|(_,v)| v);
    pub header: HeaderMap,
    /// response status code
    #[serde(rename = "statusCode")]
    pub status_code: u16,
}

/// HttpClient - issue outgoing http requests via an external provider
/// To use this capability, the actor must be linked
/// with "wasmcloud:httpclient"
/// wasmbus.contractId: wasmcloud:httpclient
/// wasmbus.providerReceive
#[async_trait]
pub trait HttpClient {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:httpclient"
    }
    /// Issue outgoing http request
    async fn request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse>;
}

/// HttpClientReceiver receives messages defined in the HttpClient service trait
/// HttpClient - issue outgoing http requests via an external provider
/// To use this capability, the actor must be linked
/// with "wasmcloud:httpclient"
#[doc(hidden)]
#[async_trait]
pub trait HttpClientReceiver: MessageDispatch + HttpClient {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Request" => {
                let value: HttpRequest = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = HttpClient::request(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "HttpClient.Request",
                    arg: buf,
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "HttpClient::{}",
                message.method
            ))),
        }
    }
}

/// HttpClientSender sends messages to a HttpClient service
/// HttpClient - issue outgoing http requests via an external provider
/// To use this capability, the actor must be linked
/// with "wasmcloud:httpclient"
/// client for sending HttpClient messages
#[derive(Debug)]
pub struct HttpClientSender<T: Transport> {
    transport: T,
}

impl<T: Transport> HttpClientSender<T> {
    /// Constructs a HttpClientSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }
}

#[cfg(target_arch = "wasm32")]
impl HttpClientSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a HttpClient provider
    /// implementing the 'wasmcloud:httpclient' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:httpclient", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a HttpClient provider
    /// implementing the 'wasmcloud:httpclient' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:httpclient", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> HttpClient for HttpClientSender<T> {
    #[allow(unused)]
    /// Issue outgoing http request
    async fn request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "HttpClient.Request",
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
