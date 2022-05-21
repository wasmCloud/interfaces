// This file is generated automatically using wasmcloud/weld-codegen 0.4.3

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

#[allow(dead_code)]
pub const SMITHY_VERSION: &str = "1.0";

/// map data structure for holding http headers
///
pub type HeaderMap = std::collections::HashMap<String, HeaderValues>;

// Encode HeaderMap as CBOR and append to output stream
#[doc(hidden)]
#[allow(unused_mut)]
pub fn encode_header_map<W: wasmbus_rpc::cbor::Write>(
    mut e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &HeaderMap,
) -> RpcResult<()> {
    e.map(val.len() as u64)?;
    for (k, v) in val {
        e.str(k)?;
        encode_header_values(e, v)?;
    }
    Ok(())
}

// Decode HeaderMap from cbor input stream
#[doc(hidden)]
pub fn decode_header_map(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<HeaderMap, RpcError> {
    let __result = {
        {
            let map_len = d.fixed_map()? as usize;
            let mut m: std::collections::HashMap<String, HeaderValues> =
                std::collections::HashMap::with_capacity(map_len);
            for _ in 0..map_len {
                let k = d.str()?.to_string();
                let v = decode_header_values(d).map_err(|e| {
                    format!(
                        "decoding 'org.wasmcloud.interface.httpserver#HeaderValues': {}",
                        e
                    )
                })?;
                m.insert(k, v);
            }
            m
        }
    };
    Ok(__result)
}
pub type HeaderValues = Vec<String>;

// Encode HeaderValues as CBOR and append to output stream
#[doc(hidden)]
#[allow(unused_mut)]
pub fn encode_header_values<W: wasmbus_rpc::cbor::Write>(
    mut e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &HeaderValues,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        e.str(item)?;
    }
    Ok(())
}

// Decode HeaderValues from cbor input stream
#[doc(hidden)]
pub fn decode_header_values(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<HeaderValues, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<String> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(d.str()?.to_string())
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<String> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(d.str()?.to_string()),
                }
            }
            arr
        }
    };
    Ok(__result)
}
/// HttpRequest contains data sent to actor about the http request
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpRequest {
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
    /// map of request headers (string key, string value)
    pub header: HeaderMap,
    /// Request body as a byte array. May be empty.
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
}

// Encode HttpRequest as CBOR and append to output stream
#[doc(hidden)]
#[allow(unused_mut)]
pub fn encode_http_request<W: wasmbus_rpc::cbor::Write>(
    mut e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &HttpRequest,
) -> RpcResult<()> {
    e.array(5)?;
    e.str(&val.method)?;
    e.str(&val.path)?;
    e.str(&val.query_string)?;
    encode_header_map(e, &val.header)?;
    e.bytes(&val.body)?;
    Ok(())
}

// Decode HttpRequest from cbor input stream
#[doc(hidden)]
pub fn decode_http_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<HttpRequest, RpcError> {
    let __result = {
        let mut method: Option<String> = None;
        let mut path: Option<String> = None;
        let mut query_string: Option<String> = None;
        let mut header: Option<HeaderMap> = None;
        let mut body: Option<Vec<u8>> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct HttpRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.fixed_array()?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => method = Some(d.str()?.to_string()),
                    1 => path = Some(d.str()?.to_string()),
                    2 => query_string = Some(d.str()?.to_string()),
                    3 => {
                        header = Some(decode_header_map(d).map_err(|e| {
                            format!(
                                "decoding 'org.wasmcloud.interface.httpserver#HeaderMap': {}",
                                e
                            )
                        })?)
                    }
                    4 => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.fixed_map()?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "method" => method = Some(d.str()?.to_string()),
                    "path" => path = Some(d.str()?.to_string()),
                    "queryString" => query_string = Some(d.str()?.to_string()),
                    "header" => {
                        header = Some(decode_header_map(d).map_err(|e| {
                            format!(
                                "decoding 'org.wasmcloud.interface.httpserver#HeaderMap': {}",
                                e
                            )
                        })?)
                    }
                    "body" => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        }
        HttpRequest {
            method: if let Some(__x) = method {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpRequest.method (#0)".to_string(),
                ));
            },

            path: if let Some(__x) = path {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpRequest.path (#1)".to_string(),
                ));
            },

            query_string: if let Some(__x) = query_string {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpRequest.query_string (#2)".to_string(),
                ));
            },

            header: if let Some(__x) = header {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpRequest.header (#3)".to_string(),
                ));
            },

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpRequest.body (#4)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// HttpResponse contains the actor's response to return to the http client
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct HttpResponse {
    /// statusCode is a three-digit number, usually in the range 100-599,
    /// A value of 200 indicates success.
    #[serde(rename = "statusCode")]
    #[serde(default)]
    pub status_code: u16,
    /// Map of headers (string keys, list of values)
    pub header: HeaderMap,
    /// Body of response as a byte array. May be an empty array.
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub body: Vec<u8>,
}

// Encode HttpResponse as CBOR and append to output stream
#[doc(hidden)]
#[allow(unused_mut)]
pub fn encode_http_response<W: wasmbus_rpc::cbor::Write>(
    mut e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &HttpResponse,
) -> RpcResult<()> {
    e.array(3)?;
    e.u16(val.status_code)?;
    encode_header_map(e, &val.header)?;
    e.bytes(&val.body)?;
    Ok(())
}

// Decode HttpResponse from cbor input stream
#[doc(hidden)]
pub fn decode_http_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<HttpResponse, RpcError> {
    let __result = {
        let mut status_code: Option<u16> = None;
        let mut header: Option<HeaderMap> = None;
        let mut body: Option<Vec<u8>> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct HttpResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.fixed_array()?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => status_code = Some(d.u16()?),
                    1 => {
                        header = Some(decode_header_map(d).map_err(|e| {
                            format!(
                                "decoding 'org.wasmcloud.interface.httpserver#HeaderMap': {}",
                                e
                            )
                        })?)
                    }
                    2 => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.fixed_map()?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "statusCode" => status_code = Some(d.u16()?),
                    "header" => {
                        header = Some(decode_header_map(d).map_err(|e| {
                            format!(
                                "decoding 'org.wasmcloud.interface.httpserver#HeaderMap': {}",
                                e
                            )
                        })?)
                    }
                    "body" => body = Some(d.bytes()?.to_vec()),
                    _ => d.skip()?,
                }
            }
        }
        HttpResponse {
            status_code: if let Some(__x) = status_code {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpResponse.status_code (#0)".to_string(),
                ));
            },

            header: if let Some(__x) = header {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpResponse.header (#1)".to_string(),
                ));
            },

            body: if let Some(__x) = body {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field HttpResponse.body (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
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
    async fn dispatch<'disp__, 'ctx__, 'msg__>(
        &'disp__ self,
        ctx: &'ctx__ Context,
        message: &Message<'msg__>,
    ) -> Result<Message<'msg__>, RpcError> {
        match message.method {
            "HandleRequest" => {
                let value: HttpRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'HttpRequest': {}", e)))?;
                let resp = HttpServer::handle_request(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "HttpServer.HandleRequest",
                    arg: Cow::Owned(buf),
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

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
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
#[cfg(target_arch = "wasm32")]
impl HttpServerSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for actor-to-actor messaging
    /// using the recipient actor's public key
    pub fn to_actor(actor_id: &str) -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_actor(actor_id.to_string()).unwrap();
        Self { transport }
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> HttpServer for HttpServerSender<T> {
    #[allow(unused)]
    async fn handle_request(&self, ctx: &Context, arg: &HttpRequest) -> RpcResult<HttpResponse> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "HttpServer.HandleRequest",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: HttpResponse = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': HttpResponse", e)))?;
        Ok(value)
    }
}
