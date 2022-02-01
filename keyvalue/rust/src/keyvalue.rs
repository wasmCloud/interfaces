// This file is generated automatically using wasmcloud/weld-codegen 0.2.4

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

/// Response to get request
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct GetResponse {
    /// the value, if it existed
    #[serde(default)]
    pub value: String,
    /// whether or not the value existed
    #[serde(default)]
    pub exists: bool,
}

// Encode GetResponse as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_get_response<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &GetResponse,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.value)?;
    e.bool(val.exists)?;
    Ok(())
}

// Decode GetResponse from cbor input stream
#[doc(hidden)]
pub fn decode_get_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<GetResponse, RpcError> {
    let __result = {
        let mut value: Option<String> = None;
        let mut exists: Option<bool> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct GetResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetResponse: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => value = Some(d.str()?.to_string()),
                    1 => exists = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetResponse: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "value" => value = Some(d.str()?.to_string()),
                    "exists" => exists = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        }
        GetResponse {
            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field GetResponse.value (#0)".to_string(),
                ));
            },

            exists: if let Some(__x) = exists {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field GetResponse.exists (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct IncrementRequest {
    /// name of value to increment
    #[serde(default)]
    pub key: String,
    /// amount to add to value
    #[serde(default)]
    pub value: i32,
}

// Encode IncrementRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_increment_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &IncrementRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.key)?;
    e.i32(val.value)?;
    Ok(())
}

// Decode IncrementRequest from cbor input stream
#[doc(hidden)]
pub fn decode_increment_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<IncrementRequest, RpcError> {
    let __result = {
        let mut key: Option<String> = None;
        let mut value: Option<i32> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct IncrementRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct IncrementRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => key = Some(d.str()?.to_string()),
                    1 => value = Some(d.i32()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct IncrementRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "key" => key = Some(d.str()?.to_string()),
                    "value" => value = Some(d.i32()?),
                    _ => d.skip()?,
                }
            }
        }
        IncrementRequest {
            key: if let Some(__x) = key {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field IncrementRequest.key (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field IncrementRequest.value (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Parameter to ListAdd operation
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ListAddRequest {
    /// name of the list to modify
    #[serde(rename = "listName")]
    #[serde(default)]
    pub list_name: String,
    /// value to append to the list
    #[serde(default)]
    pub value: String,
}

// Encode ListAddRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_list_add_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ListAddRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.list_name)?;
    e.str(&val.value)?;
    Ok(())
}

// Decode ListAddRequest from cbor input stream
#[doc(hidden)]
pub fn decode_list_add_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ListAddRequest, RpcError> {
    let __result = {
        let mut list_name: Option<String> = None;
        let mut value: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ListAddRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListAddRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => list_name = Some(d.str()?.to_string()),
                    1 => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListAddRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "listName" => list_name = Some(d.str()?.to_string()),
                    "value" => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        ListAddRequest {
            list_name: if let Some(__x) = list_name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListAddRequest.list_name (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListAddRequest.value (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Removes an item from the list. If the item occurred more than once,
/// removes only the first item.
/// Returns true if the item was found.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ListDelRequest {
    /// name of list to modify
    #[serde(rename = "listName")]
    #[serde(default)]
    pub list_name: String,
    #[serde(default)]
    pub value: String,
}

// Encode ListDelRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_list_del_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ListDelRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.list_name)?;
    e.str(&val.value)?;
    Ok(())
}

// Decode ListDelRequest from cbor input stream
#[doc(hidden)]
pub fn decode_list_del_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ListDelRequest, RpcError> {
    let __result = {
        let mut list_name: Option<String> = None;
        let mut value: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ListDelRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListDelRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => list_name = Some(d.str()?.to_string()),
                    1 => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListDelRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "listName" => list_name = Some(d.str()?.to_string()),
                    "value" => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        ListDelRequest {
            list_name: if let Some(__x) = list_name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListDelRequest.list_name (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListDelRequest.value (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ListRangeRequest {
    /// name of list
    #[serde(rename = "listName")]
    #[serde(default)]
    pub list_name: String,
    /// start index of the range, 0-based, inclusive.
    #[serde(default)]
    pub start: i32,
    /// end index of the range, 0-based, inclusive.
    #[serde(default)]
    pub stop: i32,
}

// Encode ListRangeRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_list_range_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ListRangeRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(3)?;
    e.str(&val.list_name)?;
    e.i32(val.start)?;
    e.i32(val.stop)?;
    Ok(())
}

// Decode ListRangeRequest from cbor input stream
#[doc(hidden)]
pub fn decode_list_range_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ListRangeRequest, RpcError> {
    let __result = {
        let mut list_name: Option<String> = None;
        let mut start: Option<i32> = None;
        let mut stop: Option<i32> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ListRangeRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListRangeRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => list_name = Some(d.str()?.to_string()),
                    1 => start = Some(d.i32()?),
                    2 => stop = Some(d.i32()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListRangeRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "listName" => list_name = Some(d.str()?.to_string()),
                    "start" => start = Some(d.i32()?),
                    "stop" => stop = Some(d.i32()?),
                    _ => d.skip()?,
                }
            }
        }
        ListRangeRequest {
            list_name: if let Some(__x) = list_name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListRangeRequest.list_name (#0)".to_string(),
                ));
            },

            start: if let Some(__x) = start {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListRangeRequest.start (#1)".to_string(),
                ));
            },

            stop: if let Some(__x) = stop {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListRangeRequest.stop (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SetAddRequest {
    /// name of the set
    #[serde(rename = "setName")]
    #[serde(default)]
    pub set_name: String,
    /// value to add to the set
    #[serde(default)]
    pub value: String,
}

// Encode SetAddRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_set_add_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &SetAddRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.set_name)?;
    e.str(&val.value)?;
    Ok(())
}

// Decode SetAddRequest from cbor input stream
#[doc(hidden)]
pub fn decode_set_add_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<SetAddRequest, RpcError> {
    let __result = {
        let mut set_name: Option<String> = None;
        let mut value: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct SetAddRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetAddRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => set_name = Some(d.str()?.to_string()),
                    1 => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetAddRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "setName" => set_name = Some(d.str()?.to_string()),
                    "value" => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        SetAddRequest {
            set_name: if let Some(__x) = set_name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetAddRequest.set_name (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetAddRequest.value (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SetDelRequest {
    #[serde(rename = "setName")]
    #[serde(default)]
    pub set_name: String,
    #[serde(default)]
    pub value: String,
}

// Encode SetDelRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_set_del_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &SetDelRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(2)?;
    e.str(&val.set_name)?;
    e.str(&val.value)?;
    Ok(())
}

// Decode SetDelRequest from cbor input stream
#[doc(hidden)]
pub fn decode_set_del_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<SetDelRequest, RpcError> {
    let __result = {
        let mut set_name: Option<String> = None;
        let mut value: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct SetDelRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetDelRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => set_name = Some(d.str()?.to_string()),
                    1 => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetDelRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "setName" => set_name = Some(d.str()?.to_string()),
                    "value" => value = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        SetDelRequest {
            set_name: if let Some(__x) = set_name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetDelRequest.set_name (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetDelRequest.value (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SetRequest {
    /// the key name to change (or create)
    #[serde(default)]
    pub key: String,
    /// the new value
    #[serde(default)]
    pub value: String,
    /// expiration time in seconds 0 for no expiration
    #[serde(default)]
    pub expires: u32,
}

// Encode SetRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_set_request<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &SetRequest,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(3)?;
    e.str(&val.key)?;
    e.str(&val.value)?;
    e.u32(val.expires)?;
    Ok(())
}

// Decode SetRequest from cbor input stream
#[doc(hidden)]
pub fn decode_set_request(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<SetRequest, RpcError> {
    let __result = {
        let mut key: Option<String> = None;
        let mut value: Option<String> = None;
        let mut expires: Option<u32> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct SetRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => key = Some(d.str()?.to_string()),
                    1 => value = Some(d.str()?.to_string()),
                    2 => expires = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SetRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "key" => key = Some(d.str()?.to_string()),
                    "value" => value = Some(d.str()?.to_string()),
                    "expires" => expires = Some(d.u32()?),
                    _ => d.skip()?,
                }
            }
        }
        SetRequest {
            key: if let Some(__x) = key {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetRequest.key (#0)".to_string(),
                ));
            },

            value: if let Some(__x) = value {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetRequest.value (#1)".to_string(),
                ));
            },

            expires: if let Some(__x) = expires {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SetRequest.expires (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// list of strings
pub type StringList = Vec<String>;

// Encode StringList as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_string_list<W>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &StringList,
) -> RpcResult<()>
where
    W: wasmbus_rpc::cbor::Write + 'static,
{
    e.array(val.len() as u64)?;
    for item in val.iter() {
        e.str(item)?;
    }
    Ok(())
}

// Decode StringList from cbor input stream
#[doc(hidden)]
pub fn decode_string_list(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<StringList, RpcError> {
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
/// wasmbus.contractId: wasmcloud:keyvalue
/// wasmbus.providerReceive
#[async_trait]
pub trait KeyValue {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:keyvalue"
    }
    /// Increments a numeric value, returning the new value
    async fn increment(&self, ctx: &Context, arg: &IncrementRequest) -> RpcResult<i32>;
    /// returns whether the store contains the key
    async fn contains<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool>;
    /// Deletes a key, returning true if the key was deleted
    async fn del<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool>;
    /// Gets a value for a specified key. If the key exists,
    /// the return structure contains exists: true and the value,
    /// otherwise the return structure contains exists == false.
    async fn get<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<GetResponse>;
    /// Append a value onto the end of a list. Returns the new list size
    async fn list_add(&self, ctx: &Context, arg: &ListAddRequest) -> RpcResult<u32>;
    /// Deletes a list and its contents
    /// input: list name
    /// output: true if the list existed and was deleted
    async fn list_clear<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool>;
    /// Deletes a value from a list. Returns true if the item was removed.
    async fn list_del(&self, ctx: &Context, arg: &ListDelRequest) -> RpcResult<bool>;
    /// Retrieves a range of values from a list using 0-based indices.
    /// Start and end values are inclusive, for example, (0,10) returns
    /// 11 items if the list contains at least 11 items. If the stop value
    /// is beyond the end of the list, it is treated as the end of the list.
    async fn list_range(&self, ctx: &Context, arg: &ListRangeRequest) -> RpcResult<StringList>;
    /// Sets the value of a key.
    /// expires is an optional number of seconds before the value should be automatically deleted,
    /// or 0 for no expiration.
    async fn set(&self, ctx: &Context, arg: &SetRequest) -> RpcResult<()>;
    /// Add an item into a set. Returns number of items added (1 or 0)
    async fn set_add(&self, ctx: &Context, arg: &SetAddRequest) -> RpcResult<u32>;
    /// Deletes an item from the set. Returns number of items removed from the set (1 or 0)
    async fn set_del(&self, ctx: &Context, arg: &SetDelRequest) -> RpcResult<u32>;
    /// perform intersection of sets and returns values from the intersection.
    /// input: list of sets for performing intersection (at least two)
    /// output: values
    async fn set_intersection(&self, ctx: &Context, arg: &StringList) -> RpcResult<StringList>;
    /// Retrieves all items from a set
    /// input: String
    /// output: set members
    async fn set_query<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<StringList>;
    /// perform union of sets and returns values from the union
    /// input: list of sets for performing union (at least two)
    /// output: union of values
    async fn set_union(&self, ctx: &Context, arg: &StringList) -> RpcResult<StringList>;
    /// clears all values from the set and removes it
    /// input: set name
    /// output: true if the set existed and was deleted
    async fn set_clear<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool>;
}

/// KeyValueReceiver receives messages defined in the KeyValue service trait
#[doc(hidden)]
#[async_trait]
pub trait KeyValueReceiver: MessageDispatch + KeyValue {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Increment" => {
                let value: IncrementRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'IncrementRequest': {}", e)))?;
                let resp = KeyValue::increment(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.Increment",
                    arg: Cow::Owned(buf),
                })
            }
            "Contains" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::contains(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.Contains",
                    arg: Cow::Owned(buf),
                })
            }
            "Del" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::del(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.Del",
                    arg: Cow::Owned(buf),
                })
            }
            "Get" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::get(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.Get",
                    arg: Cow::Owned(buf),
                })
            }
            "ListAdd" => {
                let value: ListAddRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'ListAddRequest': {}", e)))?;
                let resp = KeyValue::list_add(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.ListAdd",
                    arg: Cow::Owned(buf),
                })
            }
            "ListClear" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::list_clear(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.ListClear",
                    arg: Cow::Owned(buf),
                })
            }
            "ListDel" => {
                let value: ListDelRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'ListDelRequest': {}", e)))?;
                let resp = KeyValue::list_del(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.ListDel",
                    arg: Cow::Owned(buf),
                })
            }
            "ListRange" => {
                let value: ListRangeRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'ListRangeRequest': {}", e)))?;
                let resp = KeyValue::list_range(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.ListRange",
                    arg: Cow::Owned(buf),
                })
            }
            "Set" => {
                let value: SetRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'SetRequest': {}", e)))?;
                let _resp = KeyValue::set(self, ctx, &value).await?;
                let buf = Vec::new();
                Ok(Message {
                    method: "KeyValue.Set",
                    arg: Cow::Owned(buf),
                })
            }
            "SetAdd" => {
                let value: SetAddRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'SetAddRequest': {}", e)))?;
                let resp = KeyValue::set_add(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetAdd",
                    arg: Cow::Owned(buf),
                })
            }
            "SetDel" => {
                let value: SetDelRequest = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'SetDelRequest': {}", e)))?;
                let resp = KeyValue::set_del(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetDel",
                    arg: Cow::Owned(buf),
                })
            }
            "SetIntersection" => {
                let value: StringList = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'StringList': {}", e)))?;
                let resp = KeyValue::set_intersection(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetIntersection",
                    arg: Cow::Owned(buf),
                })
            }
            "SetQuery" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::set_query(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetQuery",
                    arg: Cow::Owned(buf),
                })
            }
            "SetUnion" => {
                let value: StringList = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'StringList': {}", e)))?;
                let resp = KeyValue::set_union(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetUnion",
                    arg: Cow::Owned(buf),
                })
            }
            "SetClear" => {
                let value: String = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'String': {}", e)))?;
                let resp = KeyValue::set_clear(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "KeyValue.SetClear",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "KeyValue::{}",
                message.method
            ))),
        }
    }
}

/// KeyValueSender sends messages to a KeyValue service
/// client for sending KeyValue messages
#[derive(Debug)]
pub struct KeyValueSender<T: Transport> {
    transport: T,
}

impl<T: Transport> KeyValueSender<T> {
    /// Constructs a KeyValueSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(target_arch = "wasm32")]
impl KeyValueSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a KeyValue provider
    /// implementing the 'wasmcloud:keyvalue' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:keyvalue", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a KeyValue provider
    /// implementing the 'wasmcloud:keyvalue' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:keyvalue", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> KeyValue for KeyValueSender<T> {
    #[allow(unused)]
    /// Increments a numeric value, returning the new value
    async fn increment(&self, ctx: &Context, arg: &IncrementRequest) -> RpcResult<i32> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.Increment",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: i32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': I32", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// returns whether the store contains the key
    async fn contains<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.Contains",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: bool = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': Boolean", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Deletes a key, returning true if the key was deleted
    async fn del<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.Del",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: bool = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': Boolean", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Gets a value for a specified key. If the key exists,
    /// the return structure contains exists: true and the value,
    /// otherwise the return structure contains exists == false.
    async fn get<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<GetResponse> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.Get",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: GetResponse = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': GetResponse", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Append a value onto the end of a list. Returns the new list size
    async fn list_add(&self, ctx: &Context, arg: &ListAddRequest) -> RpcResult<u32> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.ListAdd",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: u32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': U32", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Deletes a list and its contents
    /// input: list name
    /// output: true if the list existed and was deleted
    async fn list_clear<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.ListClear",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: bool = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': Boolean", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Deletes a value from a list. Returns true if the item was removed.
    async fn list_del(&self, ctx: &Context, arg: &ListDelRequest) -> RpcResult<bool> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.ListDel",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: bool = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': Boolean", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Retrieves a range of values from a list using 0-based indices.
    /// Start and end values are inclusive, for example, (0,10) returns
    /// 11 items if the list contains at least 11 items. If the stop value
    /// is beyond the end of the list, it is treated as the end of the list.
    async fn list_range(&self, ctx: &Context, arg: &ListRangeRequest) -> RpcResult<StringList> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.ListRange",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: StringList = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': StringList", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Sets the value of a key.
    /// expires is an optional number of seconds before the value should be automatically deleted,
    /// or 0 for no expiration.
    async fn set(&self, ctx: &Context, arg: &SetRequest) -> RpcResult<()> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.Set",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        Ok(())
    }

    #[allow(unused)]
    /// Add an item into a set. Returns number of items added (1 or 0)
    async fn set_add(&self, ctx: &Context, arg: &SetAddRequest) -> RpcResult<u32> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetAdd",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: u32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': U32", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Deletes an item from the set. Returns number of items removed from the set (1 or 0)
    async fn set_del(&self, ctx: &Context, arg: &SetDelRequest) -> RpcResult<u32> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetDel",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: u32 = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': U32", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// perform intersection of sets and returns values from the intersection.
    /// input: list of sets for performing intersection (at least two)
    /// output: values
    async fn set_intersection(&self, ctx: &Context, arg: &StringList) -> RpcResult<StringList> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetIntersection",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: StringList = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': StringList", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Retrieves all items from a set
    /// input: String
    /// output: set members
    async fn set_query<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<StringList> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetQuery",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: StringList = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': StringList", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// perform union of sets and returns values from the union
    /// input: list of sets for performing union (at least two)
    /// output: union of values
    async fn set_union(&self, ctx: &Context, arg: &StringList) -> RpcResult<StringList> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetUnion",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: StringList = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': StringList", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// clears all values from the set and removes it
    /// input: set name
    /// output: true if the set existed and was deleted
    async fn set_clear<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<bool> {
        let buf = wasmbus_rpc::common::serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "KeyValue.SetClear",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: bool = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': Boolean", e)))?;
        Ok(value)
    }
}
