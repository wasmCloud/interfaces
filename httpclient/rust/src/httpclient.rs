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
    context::Context, deserialize, serialize, Message, MessageDispatch, RpcError, RpcResult,
    SendOpts, Transport,
};

pub const SMITHY_VERSION: &str = "1.0";

/// map data structure for holding http headers
///
pub type HeaderMap = std::collections::HashMap<String, HeaderValues>;

pub type HeaderValues = Vec<String>;

/// http request to be sent through the provider
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
pub struct HttpRequest {
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
    pub headers: HeaderMap,
    #[serde(default)]
    pub method: String,
    #[serde(default)]
    pub url: String,
}

/// response from the http request
#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
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
    /// Issue outgoing http request
    async fn request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse>;
}

/// HttpClientReceiver receives messages defined in the HttpClient service trait
/// HttpClient - issue outgoing http requests via an external provider
/// To use this capability, the actor must be linked
/// with "wasmcloud:httpclient"
#[async_trait]
pub trait HttpClientReceiver: MessageDispatch + HttpClient {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Request" => {
                let value: HttpRequest = deserialize(message.arg.as_ref())?;
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
#[derive(Debug)]
pub struct HttpClientSender<'send, T> {
    transport: &'send T,
}

impl<'send, T: Transport> HttpClientSender<'send, T> {
    pub fn new(transport: &'send T) -> Self {
        HttpClientSender { transport }
    }
}

#[async_trait]
impl<'send, T: Transport + std::marker::Sync + std::marker::Send> HttpClient
    for HttpClientSender<'send, T>
{
    #[allow(unused)]
    /// Issue outgoing http request
    async fn request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse> {
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
