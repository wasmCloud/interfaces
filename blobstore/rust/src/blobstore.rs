// This file is generated automatically using wasmcloud/weld-codegen 0.3.2

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
pub struct BlobstoreResult {
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
    #[serde(default)]
    pub success: bool,
}

// Encode BlobstoreResult as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_blobstore_result<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &BlobstoreResult,
) -> RpcResult<()> {
    e.map(2)?;
    if let Some(val) = val.error.as_ref() {
        e.str("error")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.str("success")?;
    e.bool(val.success)?;
    Ok(())
}

// Decode BlobstoreResult from cbor input stream
#[doc(hidden)]
pub fn decode_blobstore_result(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<BlobstoreResult, RpcError> {
    let __result = {
        let mut error: Option<Option<String>> = Some(None);
        let mut success: Option<bool> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct BlobstoreResult, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct BlobstoreResult: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    1 => success = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct BlobstoreResult: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "error" => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "success" => success = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        }
        BlobstoreResult {
            error: error.unwrap(),

            success: if let Some(__x) = success {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field BlobstoreResult.success (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct Chunk {
    /// bytes in this chunk
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub bytes: Vec<u8>,
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    /// true if this is the last chunk
    #[serde(rename = "isLast")]
    #[serde(default)]
    pub is_last: bool,
    #[serde(rename = "objectId")]
    pub object_id: ObjectId,
    /// The byte offset within the object for this chunk
    #[serde(default)]
    pub offset: u64,
}

// Encode Chunk as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_chunk<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Chunk,
) -> RpcResult<()> {
    e.map(5)?;
    e.str("bytes")?;
    e.bytes(&val.bytes)?;
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    e.str("isLast")?;
    e.bool(val.is_last)?;
    e.str("objectId")?;
    encode_object_id(e, &val.object_id)?;
    e.str("offset")?;
    e.u64(val.offset)?;
    Ok(())
}

// Decode Chunk from cbor input stream
#[doc(hidden)]
pub fn decode_chunk(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Chunk, RpcError> {
    let __result = {
        let mut bytes: Option<Vec<u8>> = None;
        let mut container_id: Option<ContainerId> = None;
        let mut is_last: Option<bool> = None;
        let mut object_id: Option<ObjectId> = None;
        let mut offset: Option<u64> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct Chunk, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser("decoding struct Chunk: indefinite array not supported".to_string())
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => bytes = Some(d.bytes()?.to_vec()),
                    1 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    2 => is_last = Some(d.bool()?),
                    3 => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    4 => offset = Some(d.u64()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser("decoding struct Chunk: indefinite map not supported".to_string())
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "bytes" => bytes = Some(d.bytes()?.to_vec()),
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "isLast" => is_last = Some(d.bool()?),
                    "objectId" => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    "offset" => offset = Some(d.u64()?),
                    _ => d.skip()?,
                }
            }
        }
        Chunk {
            bytes: if let Some(__x) = bytes {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Chunk.bytes (#0)".to_string(),
                ));
            },

            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Chunk.container_id (#1)".to_string(),
                ));
            },

            is_last: if let Some(__x) = is_last {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Chunk.is_last (#2)".to_string(),
                ));
            },

            object_id: if let Some(__x) = object_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Chunk.object_id (#3)".to_string(),
                ));
            },

            offset: if let Some(__x) = offset {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Chunk.offset (#4)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Response from actor after receiving a download chunk.
/// If cancelDownload is true, the sender will stop sending chunks
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ChunkResponse {
    #[serde(rename = "cancelDownload")]
    #[serde(default)]
    pub cancel_download: bool,
}

// Encode ChunkResponse as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_chunk_response<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ChunkResponse,
) -> RpcResult<()> {
    e.map(1)?;
    e.str("cancelDownload")?;
    e.bool(val.cancel_download)?;
    Ok(())
}

// Decode ChunkResponse from cbor input stream
#[doc(hidden)]
pub fn decode_chunk_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ChunkResponse, RpcError> {
    let __result = {
        let mut cancel_download: Option<bool> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ChunkResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ChunkResponse: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => cancel_download = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ChunkResponse: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "cancelDownload" => cancel_download = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        }
        ChunkResponse {
            cancel_download: if let Some(__x) = cancel_download {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ChunkResponse.cancel_download (#0)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Unique id of a container
pub type ContainerId = String;

// Encode ContainerId as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_container_id<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ContainerId,
) -> RpcResult<()> {
    e.str(val)?;
    Ok(())
}

// Decode ContainerId from cbor input stream
#[doc(hidden)]
pub fn decode_container_id(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ContainerId, RpcError> {
    let __result = { d.str()?.to_string() };
    Ok(__result)
}
pub type ContainerIds = Vec<ContainerId>;

// Encode ContainerIds as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_container_ids<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ContainerIds,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_container_id(e, item)?;
    }
    Ok(())
}

// Decode ContainerIds from cbor input stream
#[doc(hidden)]
pub fn decode_container_ids(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ContainerIds, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<ContainerId> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(
                    decode_container_id(d).map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                )
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<ContainerId> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(
                        decode_container_id(d)
                            .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                    ),
                }
            }
            arr
        }
    };
    Ok(__result)
}
/// A container is a logical grouping of blobs, similar to a directory
/// in a file system. The containerId is its name.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ContainerMetadata {
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    /// creation date (rfc3339)
    #[serde(rename = "createdAt")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
}

// Encode ContainerMetadata as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_container_metadata<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ContainerMetadata,
) -> RpcResult<()> {
    e.map(2)?;
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    if let Some(val) = val.created_at.as_ref() {
        e.str("createdAt")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode ContainerMetadata from cbor input stream
#[doc(hidden)]
pub fn decode_container_metadata(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ContainerMetadata, RpcError> {
    let __result = {
        let mut container_id: Option<ContainerId> = None;
        let mut created_at: Option<Option<String>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ContainerMetadata, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ContainerMetadata: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    1 => {
                        created_at = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ContainerMetadata: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "createdAt" => {
                        created_at = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        ContainerMetadata {
            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ContainerMetadata.container_id (#0)".to_string(),
                ));
            },
            created_at: created_at.unwrap(),
        }
    };
    Ok(__result)
}
/// Combination of container id and object id
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ContainerObject {
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    #[serde(rename = "objectId")]
    pub object_id: ObjectId,
}

// Encode ContainerObject as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_container_object<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ContainerObject,
) -> RpcResult<()> {
    e.map(2)?;
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    e.str("objectId")?;
    encode_object_id(e, &val.object_id)?;
    Ok(())
}

// Decode ContainerObject from cbor input stream
#[doc(hidden)]
pub fn decode_container_object(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ContainerObject, RpcError> {
    let __result = {
        let mut container_id: Option<ContainerId> = None;
        let mut object_id: Option<ObjectId> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ContainerObject, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ContainerObject: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    1 => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ContainerObject: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "objectId" => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        }
        ContainerObject {
            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ContainerObject.container_id (#0)".to_string(),
                ));
            },

            object_id: if let Some(__x) = object_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ContainerObject.object_id (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
pub type ContainersInfo = Vec<ContainerMetadata>;

// Encode ContainersInfo as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_containers_info<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ContainersInfo,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_container_metadata(e, item)?;
    }
    Ok(())
}

// Decode ContainersInfo from cbor input stream
#[doc(hidden)]
pub fn decode_containers_info(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ContainersInfo, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<ContainerMetadata> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(
                    decode_container_metadata(d)
                        .map_err(|e| format!("decoding 'ContainerMetadata': {}", e))?,
                )
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<ContainerMetadata> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(
                        decode_container_metadata(d)
                            .map_err(|e| format!("decoding 'ContainerMetadata': {}", e))?,
                    ),
                }
            }
            arr
        }
    };
    Ok(__result)
}
/// Parameter to GetObject
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct GetObjectRequest {
    /// optional size requested
    /// The provider will not return a chunk larger than this amount,
    /// but may return a smaller chunk.
    #[serde(rename = "chunkSize")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub chunk_size: Option<u64>,
    /// object's container
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    /// object to download
    #[serde(rename = "objectId")]
    pub object_id: ObjectId,
    /// Requested end of object to retrieve. Defaults to the object's size.
    /// It is not an error for regionEnd to be greater than the object size.
    #[serde(rename = "regionEnd")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub region_end: Option<u64>,
    /// Requested start of object to retrieve. Defaults to 0
    /// If regionStart is beyond the end of the file,
    /// an empty chunk will be returned with isLast == true
    #[serde(rename = "regionStart")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub region_start: Option<u64>,
}

// Encode GetObjectRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_get_object_request<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &GetObjectRequest,
) -> RpcResult<()> {
    e.map(5)?;
    if let Some(val) = val.chunk_size.as_ref() {
        e.str("chunkSize")?;
        e.u64(*val)?;
    } else {
        e.null()?;
    }
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    e.str("objectId")?;
    encode_object_id(e, &val.object_id)?;
    if let Some(val) = val.region_end.as_ref() {
        e.str("regionEnd")?;
        e.u64(*val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.region_start.as_ref() {
        e.str("regionStart")?;
        e.u64(*val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode GetObjectRequest from cbor input stream
#[doc(hidden)]
pub fn decode_get_object_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<GetObjectRequest, RpcError> {
    let __result = {
        let mut chunk_size: Option<Option<u64>> = Some(None);
        let mut container_id: Option<ContainerId> = None;
        let mut object_id: Option<ObjectId> = None;
        let mut region_end: Option<Option<u64>> = Some(None);
        let mut region_start: Option<Option<u64>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct GetObjectRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetObjectRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        chunk_size = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    1 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    2 => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    3 => {
                        region_end = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    4 => {
                        region_start = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetObjectRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "chunkSize" => {
                        chunk_size = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "objectId" => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    "regionEnd" => {
                        region_end = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    "regionStart" => {
                        region_start = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        GetObjectRequest {
            chunk_size: chunk_size.unwrap(),

            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field GetObjectRequest.container_id (#1)".to_string(),
                ));
            },

            object_id: if let Some(__x) = object_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field GetObjectRequest.object_id (#2)".to_string(),
                ));
            },
            region_end: region_end.unwrap(),
            region_start: region_start.unwrap(),
        }
    };
    Ok(__result)
}
/// Response to GetObject
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct GetObjectResponse {
    /// If success is false, this may contain an error
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
    /// The provider may begin the download by returning a first chunk
    #[serde(rename = "initialChunk")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub initial_chunk: Option<Chunk>,
    /// indication whether the request was successful
    #[serde(default)]
    pub success: bool,
}

// Encode GetObjectResponse as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_get_object_response<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &GetObjectResponse,
) -> RpcResult<()> {
    e.map(3)?;
    if let Some(val) = val.error.as_ref() {
        e.str("error")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.initial_chunk.as_ref() {
        e.str("initialChunk")?;
        encode_chunk(e, val)?;
    } else {
        e.null()?;
    }
    e.str("success")?;
    e.bool(val.success)?;
    Ok(())
}

// Decode GetObjectResponse from cbor input stream
#[doc(hidden)]
pub fn decode_get_object_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<GetObjectResponse, RpcError> {
    let __result = {
        let mut error: Option<Option<String>> = Some(None);
        let mut initial_chunk: Option<Option<Chunk>> = Some(None);
        let mut success: Option<bool> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct GetObjectResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetObjectResponse: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    1 => {
                        initial_chunk = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?,
                            ))
                        }
                    }
                    2 => success = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct GetObjectResponse: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "error" => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "initialChunk" => {
                        initial_chunk = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?,
                            ))
                        }
                    }
                    "success" => success = Some(d.bool()?),
                    _ => d.skip()?,
                }
            }
        }
        GetObjectResponse {
            error: error.unwrap(),
            initial_chunk: initial_chunk.unwrap(),

            success: if let Some(__x) = success {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field GetObjectResponse.success (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
pub type Items = Vec<String>;

// Encode Items as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_items<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Items,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        e.str(item)?;
    }
    Ok(())
}

// Decode Items from cbor input stream
#[doc(hidden)]
pub fn decode_items(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Items, RpcError> {
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
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ListObjectsRequest {
    /// the container to search
    #[serde(rename = "containerId")]
    #[serde(default)]
    pub container_id: String,
    /// Optional continuation token passed in ListObjectsResponse.
    /// If set, `startWith` is ignored.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub continuation: Option<String>,
    /// Optionally, stop returning items before returning this value.
    /// (exclusive terminator)
    /// If startFrom is "a" and endBefore is "b", and items are ordered
    /// alphabetically, then only items beginning with "a" would be returned.
    #[serde(rename = "endBefore")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub end_before: Option<String>,
    /// Optional last item to return (inclusive terminator)
    #[serde(rename = "endWith")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub end_with: Option<String>,
    /// maximum number of items to return. If not specified, provider
    /// will return an initial set of up to 1000 items. if maxItems > 1000,
    /// the provider implementation may return fewer items than requested.
    #[serde(rename = "maxItems")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub max_items: Option<u64>,
    /// Request object names starting with this value
    #[serde(rename = "startWith")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub start_with: Option<String>,
}

// Encode ListObjectsRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_list_objects_request<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ListObjectsRequest,
) -> RpcResult<()> {
    e.map(6)?;
    e.str("containerId")?;
    e.str(&val.container_id)?;
    if let Some(val) = val.continuation.as_ref() {
        e.str("continuation")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.end_before.as_ref() {
        e.str("endBefore")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.end_with.as_ref() {
        e.str("endWith")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.max_items.as_ref() {
        e.str("maxItems")?;
        e.u64(*val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.start_with.as_ref() {
        e.str("startWith")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode ListObjectsRequest from cbor input stream
#[doc(hidden)]
pub fn decode_list_objects_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ListObjectsRequest, RpcError> {
    let __result = {
        let mut container_id: Option<String> = None;
        let mut continuation: Option<Option<String>> = Some(None);
        let mut end_before: Option<Option<String>> = Some(None);
        let mut end_with: Option<Option<String>> = Some(None);
        let mut max_items: Option<Option<u64>> = Some(None);
        let mut start_with: Option<Option<String>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ListObjectsRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListObjectsRequest: indefinite array not supported"
                        .to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => container_id = Some(d.str()?.to_string()),
                    1 => {
                        continuation = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    2 => {
                        end_before = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    3 => {
                        end_with = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    4 => {
                        max_items = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    5 => {
                        start_with = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListObjectsRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "containerId" => container_id = Some(d.str()?.to_string()),
                    "continuation" => {
                        continuation = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "endBefore" => {
                        end_before = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "endWith" => {
                        end_with = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "maxItems" => {
                        max_items = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.u64()?))
                        }
                    }
                    "startWith" => {
                        start_with = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        ListObjectsRequest {
            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListObjectsRequest.container_id (#0)".to_string(),
                ));
            },
            continuation: continuation.unwrap(),
            end_before: end_before.unwrap(),
            end_with: end_with.unwrap(),
            max_items: max_items.unwrap(),
            start_with: start_with.unwrap(),
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ListObjectsResponse {
    /// If `isLast` is false, this value can be used in the `continuation` field
    /// of a `ListObjectsRequest`.
    /// This field may be obfuscated and may not be a real key or object name.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub continuation: Option<String>,
    /// Indicates if the item list is complete, or the last item
    /// in a multi-part response.
    #[serde(rename = "isLast")]
    #[serde(default)]
    pub is_last: bool,
    /// set of objects returned
    pub objects: ObjectsInfo,
}

// Encode ListObjectsResponse as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_list_objects_response<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ListObjectsResponse,
) -> RpcResult<()> {
    e.map(3)?;
    if let Some(val) = val.continuation.as_ref() {
        e.str("continuation")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.str("isLast")?;
    e.bool(val.is_last)?;
    e.str("objects")?;
    encode_objects_info(e, &val.objects)?;
    Ok(())
}

// Decode ListObjectsResponse from cbor input stream
#[doc(hidden)]
pub fn decode_list_objects_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ListObjectsResponse, RpcError> {
    let __result = {
        let mut continuation: Option<Option<String>> = Some(None);
        let mut is_last: Option<bool> = None;
        let mut objects: Option<ObjectsInfo> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ListObjectsResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListObjectsResponse: indefinite array not supported"
                        .to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        continuation = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    1 => is_last = Some(d.bool()?),
                    2 => {
                        objects = Some(
                            decode_objects_info(d)
                                .map_err(|e| format!("decoding 'ObjectsInfo': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ListObjectsResponse: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "continuation" => {
                        continuation = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "isLast" => is_last = Some(d.bool()?),
                    "objects" => {
                        objects = Some(
                            decode_objects_info(d)
                                .map_err(|e| format!("decoding 'ObjectsInfo': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        }
        ListObjectsResponse {
            continuation: continuation.unwrap(),

            is_last: if let Some(__x) = is_last {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListObjectsResponse.is_last (#1)".to_string(),
                ));
            },

            objects: if let Some(__x) = objects {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ListObjectsResponse.objects (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// Unique id of an object
pub type ObjectId = String;

// Encode ObjectId as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_object_id<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ObjectId,
) -> RpcResult<()> {
    e.str(val)?;
    Ok(())
}

// Decode ObjectId from cbor input stream
#[doc(hidden)]
pub fn decode_object_id(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<ObjectId, RpcError> {
    let __result = { d.str()?.to_string() };
    Ok(__result)
}
pub type ObjectIds = Vec<ObjectId>;

// Encode ObjectIds as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_object_ids<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ObjectIds,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_object_id(e, item)?;
    }
    Ok(())
}

// Decode ObjectIds from cbor input stream
#[doc(hidden)]
pub fn decode_object_ids(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<ObjectIds, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<ObjectId> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(decode_object_id(d).map_err(|e| format!("decoding 'ObjectId': {}", e))?)
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<ObjectId> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(
                        decode_object_id(d).map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                    ),
                }
            }
            arr
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ObjectMetadata {
    /// container of the object
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    /// date object was last modified (rfc3339)
    #[serde(rename = "lastModified")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    /// Object identifier that is unique within its container.
    /// Naming of objects is determined by the capability provider.
    /// An object id could be a path, hash of object contents, or some other unique identifier.
    #[serde(rename = "objectId")]
    pub object_id: ObjectId,
    /// size of the object in bytes
    #[serde(default)]
    pub size: u64,
}

// Encode ObjectMetadata as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_object_metadata<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ObjectMetadata,
) -> RpcResult<()> {
    e.map(4)?;
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    if let Some(val) = val.last_modified.as_ref() {
        e.str("lastModified")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    e.str("objectId")?;
    encode_object_id(e, &val.object_id)?;
    e.str("size")?;
    e.u64(val.size)?;
    Ok(())
}

// Decode ObjectMetadata from cbor input stream
#[doc(hidden)]
pub fn decode_object_metadata(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ObjectMetadata, RpcError> {
    let __result = {
        let mut container_id: Option<ContainerId> = None;
        let mut last_modified: Option<Option<String>> = Some(None);
        let mut object_id: Option<ObjectId> = None;
        let mut size: Option<u64> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ObjectMetadata, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ObjectMetadata: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    1 => {
                        last_modified = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    2 => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    3 => size = Some(d.u64()?),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ObjectMetadata: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "lastModified" => {
                        last_modified = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "objectId" => {
                        object_id = Some(
                            decode_object_id(d)
                                .map_err(|e| format!("decoding 'ObjectId': {}", e))?,
                        )
                    }
                    "size" => size = Some(d.u64()?),
                    _ => d.skip()?,
                }
            }
        }
        ObjectMetadata {
            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ObjectMetadata.container_id (#0)".to_string(),
                ));
            },
            last_modified: last_modified.unwrap(),

            object_id: if let Some(__x) = object_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ObjectMetadata.object_id (#2)".to_string(),
                ));
            },

            size: if let Some(__x) = size {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ObjectMetadata.size (#3)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
pub type ObjectsInfo = Vec<ObjectMetadata>;

// Encode ObjectsInfo as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_objects_info<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ObjectsInfo,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_object_metadata(e, item)?;
    }
    Ok(())
}

// Decode ObjectsInfo from cbor input stream
#[doc(hidden)]
pub fn decode_objects_info(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ObjectsInfo, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<ObjectMetadata> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(
                    decode_object_metadata(d)
                        .map_err(|e| format!("decoding 'ObjectMetadata': {}", e))?,
                )
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<ObjectMetadata> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(
                        decode_object_metadata(d)
                            .map_err(|e| format!("decoding 'ObjectMetadata': {}", e))?,
                    ),
                }
            }
            arr
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PutChunkRequest {
    /// If set, the receiving provider should cancel the upload process
    /// and remove the file.
    #[serde(rename = "cancelAndRemove")]
    #[serde(default)]
    pub cancel_and_remove: bool,
    /// upload chunk from the file.
    /// if chunk.isLast is set, this will be the last chunk uploaded
    pub chunk: Chunk,
    /// streamId returned from PutObject
    #[serde(rename = "streamId")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub stream_id: Option<String>,
}

// Encode PutChunkRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_put_chunk_request<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PutChunkRequest,
) -> RpcResult<()> {
    e.map(3)?;
    e.str("cancelAndRemove")?;
    e.bool(val.cancel_and_remove)?;
    e.str("chunk")?;
    encode_chunk(e, &val.chunk)?;
    if let Some(val) = val.stream_id.as_ref() {
        e.str("streamId")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode PutChunkRequest from cbor input stream
#[doc(hidden)]
pub fn decode_put_chunk_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<PutChunkRequest, RpcError> {
    let __result = {
        let mut cancel_and_remove: Option<bool> = None;
        let mut chunk: Option<Chunk> = None;
        let mut stream_id: Option<Option<String>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct PutChunkRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutChunkRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => cancel_and_remove = Some(d.bool()?),
                    1 => {
                        chunk =
                            Some(decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?)
                    }
                    2 => {
                        stream_id = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutChunkRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "cancelAndRemove" => cancel_and_remove = Some(d.bool()?),
                    "chunk" => {
                        chunk =
                            Some(decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?)
                    }
                    "streamId" => {
                        stream_id = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        PutChunkRequest {
            cancel_and_remove: if let Some(__x) = cancel_and_remove {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field PutChunkRequest.cancel_and_remove (#0)".to_string(),
                ));
            },

            chunk: if let Some(__x) = chunk {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field PutChunkRequest.chunk (#1)".to_string(),
                ));
            },
            stream_id: stream_id.unwrap(),
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PutObjectRequest {
    /// File path and initial data
    pub chunk: Chunk,
}

// Encode PutObjectRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_put_object_request<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PutObjectRequest,
) -> RpcResult<()> {
    e.map(1)?;
    e.str("chunk")?;
    encode_chunk(e, &val.chunk)?;
    Ok(())
}

// Decode PutObjectRequest from cbor input stream
#[doc(hidden)]
pub fn decode_put_object_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<PutObjectRequest, RpcError> {
    let __result = {
        let mut chunk: Option<Chunk> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct PutObjectRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutObjectRequest: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        chunk =
                            Some(decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?)
                    }
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutObjectRequest: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "chunk" => {
                        chunk =
                            Some(decode_chunk(d).map_err(|e| format!("decoding 'Chunk': {}", e))?)
                    }
                    _ => d.skip()?,
                }
            }
        }
        PutObjectRequest {
            chunk: if let Some(__x) = chunk {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field PutObjectRequest.chunk (#0)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PutObjectResponse {
    /// If this is a multipart upload, this streamId value must be returned
    /// with subsequent PutChunk requests
    #[serde(rename = "streamId")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub stream_id: Option<String>,
}

// Encode PutObjectResponse as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_put_object_response<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PutObjectResponse,
) -> RpcResult<()> {
    e.map(1)?;
    if let Some(val) = val.stream_id.as_ref() {
        e.str("streamId")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode PutObjectResponse from cbor input stream
#[doc(hidden)]
pub fn decode_put_object_response(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<PutObjectResponse, RpcError> {
    let __result = {
        let mut stream_id: Option<Option<String>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct PutObjectResponse, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutObjectResponse: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        stream_id = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PutObjectResponse: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "streamId" => {
                        stream_id = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        PutObjectResponse {
            stream_id: stream_id.unwrap(),
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RemoveObjectsRequest {
    #[serde(rename = "containerId")]
    pub container_id: ContainerId,
    pub objects: ObjectIds,
}

// Encode RemoveObjectsRequest as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_remove_objects_request<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &RemoveObjectsRequest,
) -> RpcResult<()> {
    e.map(2)?;
    e.str("containerId")?;
    encode_container_id(e, &val.container_id)?;
    e.str("objects")?;
    encode_object_ids(e, &val.objects)?;
    Ok(())
}

// Decode RemoveObjectsRequest from cbor input stream
#[doc(hidden)]
pub fn decode_remove_objects_request(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<RemoveObjectsRequest, RpcError> {
    let __result = {
        let mut container_id: Option<ContainerId> = None;
        let mut objects: Option<ObjectIds> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct RemoveObjectsRequest, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct RemoveObjectsRequest: indefinite array not supported"
                        .to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    1 => {
                        objects = Some(
                            decode_object_ids(d)
                                .map_err(|e| format!("decoding 'ObjectIds': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct RemoveObjectsRequest: indefinite map not supported"
                        .to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "containerId" => {
                        container_id = Some(
                            decode_container_id(d)
                                .map_err(|e| format!("decoding 'ContainerId': {}", e))?,
                        )
                    }
                    "objects" => {
                        objects = Some(
                            decode_object_ids(d)
                                .map_err(|e| format!("decoding 'ObjectIds': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        }
        RemoveObjectsRequest {
            container_id: if let Some(__x) = container_id {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RemoveObjectsRequest.container_id (#0)".to_string(),
                ));
            },

            objects: if let Some(__x) = objects {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field RemoveObjectsRequest.objects (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// The BlobStore service, provider side
/// wasmbus.contractId: wasmcloud:blobstore
/// wasmbus.providerReceive
#[async_trait]
pub trait Blobstore {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:blobstore"
    }
    /// Returns whether the container exists
    async fn container_exists(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<BlobstoreResult>;
    /// Creates a container by name, returning success if it worked
    /// Note that container names may not be globally unique - just unique within the
    /// "namespace" of the connecting actor and linkdef
    async fn create_container(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<BlobstoreResult>;
    /// Retrieves information about the container,
    /// Returns no value if the container id is invalid
    async fn get_container_info(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<ContainerMetadata>;
    /// Returns list of container ids
    async fn list_containers(&self, ctx: &Context) -> RpcResult<ContainersInfo>;
    /// Empty and remove the container(s)
    async fn remove_containers(
        &self,
        ctx: &Context,
        arg: &ContainerIds,
    ) -> RpcResult<BlobstoreResult>;
    /// Returns whether the object exists
    async fn object_exists(
        &self,
        ctx: &Context,
        arg: &ContainerObject,
    ) -> RpcResult<BlobstoreResult>;
    /// List the objects in the container
    async fn list_objects(
        &self,
        ctx: &Context,
        arg: &ListObjectsRequest,
    ) -> RpcResult<ListObjectsResponse>;
    /// Remove the objects.
    /// The objects do not need to be in the same container
    async fn remove_objects(
        &self,
        ctx: &Context,
        arg: &RemoveObjectsRequest,
    ) -> RpcResult<BlobstoreResult>;
    /// Requests to start upload of a file/blob to the Blobstore
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn put_object(
        &self,
        ctx: &Context,
        arg: &PutObjectRequest,
    ) -> RpcResult<PutObjectResponse>;
    /// Requests to retrieve an object. If the object is large, the provider
    /// may split the response into multiple parts
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn get_object(
        &self,
        ctx: &Context,
        arg: &GetObjectRequest,
    ) -> RpcResult<GetObjectResponse>;
    /// Uploads a file chunk to a blobstore. This must be called AFTER PutObject
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn put_chunk(&self, ctx: &Context, arg: &PutChunkRequest) -> RpcResult<BlobstoreResult>;
}

/// BlobstoreReceiver receives messages defined in the Blobstore service trait
/// The BlobStore service, provider side
#[doc(hidden)]
#[async_trait]
pub trait BlobstoreReceiver: MessageDispatch + Blobstore {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "ContainerExists" => {
                let value: ContainerId =
                    wasmbus_rpc::common::decode(&message.arg, &decode_container_id)
                        .map_err(|e| RpcError::Deser(format!("'ContainerId': {}", e)))?;
                let resp = Blobstore::container_exists(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.ContainerExists",
                    arg: Cow::Owned(buf),
                })
            }
            "CreateContainer" => {
                let value: ContainerId =
                    wasmbus_rpc::common::decode(&message.arg, &decode_container_id)
                        .map_err(|e| RpcError::Deser(format!("'ContainerId': {}", e)))?;
                let resp = Blobstore::create_container(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.CreateContainer",
                    arg: Cow::Owned(buf),
                })
            }
            "GetContainerInfo" => {
                let value: ContainerId =
                    wasmbus_rpc::common::decode(&message.arg, &decode_container_id)
                        .map_err(|e| RpcError::Deser(format!("'ContainerId': {}", e)))?;
                let resp = Blobstore::get_container_info(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_container_metadata(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.GetContainerInfo",
                    arg: Cow::Owned(buf),
                })
            }
            "ListContainers" => {
                let resp = Blobstore::list_containers(self, ctx).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_containers_info(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.ListContainers",
                    arg: Cow::Owned(buf),
                })
            }
            "RemoveContainers" => {
                let value: ContainerIds =
                    wasmbus_rpc::common::decode(&message.arg, &decode_container_ids)
                        .map_err(|e| RpcError::Deser(format!("'ContainerIds': {}", e)))?;
                let resp = Blobstore::remove_containers(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.RemoveContainers",
                    arg: Cow::Owned(buf),
                })
            }
            "ObjectExists" => {
                let value: ContainerObject =
                    wasmbus_rpc::common::decode(&message.arg, &decode_container_object)
                        .map_err(|e| RpcError::Deser(format!("'ContainerObject': {}", e)))?;
                let resp = Blobstore::object_exists(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.ObjectExists",
                    arg: Cow::Owned(buf),
                })
            }
            "ListObjects" => {
                let value: ListObjectsRequest =
                    wasmbus_rpc::common::decode(&message.arg, &decode_list_objects_request)
                        .map_err(|e| RpcError::Deser(format!("'ListObjectsRequest': {}", e)))?;
                let resp = Blobstore::list_objects(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_list_objects_response(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.ListObjects",
                    arg: Cow::Owned(buf),
                })
            }
            "RemoveObjects" => {
                let value: RemoveObjectsRequest =
                    wasmbus_rpc::common::decode(&message.arg, &decode_remove_objects_request)
                        .map_err(|e| RpcError::Deser(format!("'RemoveObjectsRequest': {}", e)))?;
                let resp = Blobstore::remove_objects(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.RemoveObjects",
                    arg: Cow::Owned(buf),
                })
            }
            "PutObject" => {
                let value: PutObjectRequest =
                    wasmbus_rpc::common::decode(&message.arg, &decode_put_object_request)
                        .map_err(|e| RpcError::Deser(format!("'PutObjectRequest': {}", e)))?;
                let resp = Blobstore::put_object(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_put_object_response(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.PutObject",
                    arg: Cow::Owned(buf),
                })
            }
            "GetObject" => {
                let value: GetObjectRequest =
                    wasmbus_rpc::common::decode(&message.arg, &decode_get_object_request)
                        .map_err(|e| RpcError::Deser(format!("'GetObjectRequest': {}", e)))?;
                let resp = Blobstore::get_object(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_get_object_response(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.GetObject",
                    arg: Cow::Owned(buf),
                })
            }
            "PutChunk" => {
                let value: PutChunkRequest =
                    wasmbus_rpc::common::decode(&message.arg, &decode_put_chunk_request)
                        .map_err(|e| RpcError::Deser(format!("'PutChunkRequest': {}", e)))?;
                let resp = Blobstore::put_chunk(self, ctx, &value).await?;
                let mut e = wasmbus_rpc::cbor::vec_encoder();
                encode_blobstore_result(&mut e, &resp)?;
                let buf = e.into_inner();
                Ok(Message {
                    method: "Blobstore.PutChunk",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "Blobstore::{}",
                message.method
            ))),
        }
    }
}

/// BlobstoreSender sends messages to a Blobstore service
/// The BlobStore service, provider side
/// client for sending Blobstore messages
#[derive(Debug)]
pub struct BlobstoreSender<T: Transport> {
    transport: T,
}

impl<T: Transport> BlobstoreSender<T> {
    /// Constructs a BlobstoreSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(target_arch = "wasm32")]
impl BlobstoreSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a Blobstore provider
    /// implementing the 'wasmcloud:blobstore' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:blobstore", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a Blobstore provider
    /// implementing the 'wasmcloud:blobstore' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:blobstore", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> Blobstore for BlobstoreSender<T> {
    #[allow(unused)]
    /// Returns whether the container exists
    async fn container_exists(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_container_id(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.ContainerExists",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Creates a container by name, returning success if it worked
    /// Note that container names may not be globally unique - just unique within the
    /// "namespace" of the connecting actor and linkdef
    async fn create_container(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_container_id(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.CreateContainer",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Retrieves information about the container,
    /// Returns no value if the container id is invalid
    async fn get_container_info(
        &self,
        ctx: &Context,
        arg: &ContainerId,
    ) -> RpcResult<ContainerMetadata> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_container_id(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.GetContainerInfo",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ContainerMetadata =
            wasmbus_rpc::common::decode(&resp, &decode_container_metadata)
                .map_err(|e| RpcError::Deser(format!("'{}': ContainerMetadata", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Returns list of container ids
    async fn list_containers(&self, ctx: &Context) -> RpcResult<ContainersInfo> {
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.ListContainers",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ContainersInfo = wasmbus_rpc::common::decode(&resp, &decode_containers_info)
            .map_err(|e| RpcError::Deser(format!("'{}': ContainersInfo", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Empty and remove the container(s)
    async fn remove_containers(
        &self,
        ctx: &Context,
        arg: &ContainerIds,
    ) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_container_ids(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.RemoveContainers",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Returns whether the object exists
    async fn object_exists(
        &self,
        ctx: &Context,
        arg: &ContainerObject,
    ) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_container_object(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.ObjectExists",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// List the objects in the container
    async fn list_objects(
        &self,
        ctx: &Context,
        arg: &ListObjectsRequest,
    ) -> RpcResult<ListObjectsResponse> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_list_objects_request(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.ListObjects",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ListObjectsResponse =
            wasmbus_rpc::common::decode(&resp, &decode_list_objects_response)
                .map_err(|e| RpcError::Deser(format!("'{}': ListObjectsResponse", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Remove the objects.
    /// The objects do not need to be in the same container
    async fn remove_objects(
        &self,
        ctx: &Context,
        arg: &RemoveObjectsRequest,
    ) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_remove_objects_request(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.RemoveObjects",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Requests to start upload of a file/blob to the Blobstore
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn put_object(
        &self,
        ctx: &Context,
        arg: &PutObjectRequest,
    ) -> RpcResult<PutObjectResponse> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_put_object_request(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.PutObject",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: PutObjectResponse =
            wasmbus_rpc::common::decode(&resp, &decode_put_object_response)
                .map_err(|e| RpcError::Deser(format!("'{}': PutObjectResponse", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Requests to retrieve an object. If the object is large, the provider
    /// may split the response into multiple parts
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn get_object(
        &self,
        ctx: &Context,
        arg: &GetObjectRequest,
    ) -> RpcResult<GetObjectResponse> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_get_object_request(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.GetObject",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: GetObjectResponse =
            wasmbus_rpc::common::decode(&resp, &decode_get_object_response)
                .map_err(|e| RpcError::Deser(format!("'{}': GetObjectResponse", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Uploads a file chunk to a blobstore. This must be called AFTER PutObject
    /// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
    async fn put_chunk(&self, ctx: &Context, arg: &PutChunkRequest) -> RpcResult<BlobstoreResult> {
        let mut e = wasmbus_rpc::cbor::vec_encoder();
        encode_put_chunk_request(&mut e, arg)?;
        let buf = e.into_inner();
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Blobstore.PutChunk",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: BlobstoreResult =
            wasmbus_rpc::common::decode(&resp, &decode_blobstore_result)
                .map_err(|e| RpcError::Deser(format!("'{}': BlobstoreResult", e)))?;
        Ok(value)
    }
}

/// The BlobStore service, actor side
/// wasmbus.contractId: wasmcloud:blobstore
/// wasmbus.actorReceive
#[async_trait]
pub trait ChunkReceiver {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:blobstore"
    }
    /// Receives a file chunk from a blobstore.
    /// A blobstore provider invokes this operation on actors in response to the GetObject request.
    /// If the response sets cancelDownload, the provider will stop downloading chunks
    async fn receive_chunk(&self, ctx: &Context, arg: &Chunk) -> RpcResult<ChunkResponse>;
}

/// ChunkReceiverReceiver receives messages defined in the ChunkReceiver service trait
/// The BlobStore service, actor side
#[doc(hidden)]
#[async_trait]
pub trait ChunkReceiverReceiver: MessageDispatch + ChunkReceiver {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "ReceiveChunk" => {
                let value: Chunk = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'Chunk': {}", e)))?;
                let resp = ChunkReceiver::receive_chunk(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "ChunkReceiver.ReceiveChunk",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "ChunkReceiver::{}",
                message.method
            ))),
        }
    }
}

/// ChunkReceiverSender sends messages to a ChunkReceiver service
/// The BlobStore service, actor side
/// client for sending ChunkReceiver messages
#[derive(Debug)]
pub struct ChunkReceiverSender<T: Transport> {
    transport: T,
}

impl<T: Transport> ChunkReceiverSender<T> {
    /// Constructs a ChunkReceiverSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(not(target_arch = "wasm32"))]
impl<'send> ChunkReceiverSender<wasmbus_rpc::provider::ProviderTransport<'send>> {
    /// Constructs a Sender using an actor's LinkDefinition,
    /// Uses the provider's HostBridge for rpc
    pub fn for_actor(ld: &'send wasmbus_rpc::core::LinkDefinition) -> Self {
        Self {
            transport: wasmbus_rpc::provider::ProviderTransport::new(ld, None),
        }
    }
}
#[cfg(target_arch = "wasm32")]
impl ChunkReceiverSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for actor-to-actor messaging
    /// using the recipient actor's public key
    pub fn to_actor(actor_id: &str) -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_actor(actor_id.to_string()).unwrap();
        Self { transport }
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> ChunkReceiver
    for ChunkReceiverSender<T>
{
    #[allow(unused)]
    /// Receives a file chunk from a blobstore.
    /// A blobstore provider invokes this operation on actors in response to the GetObject request.
    /// If the response sets cancelDownload, the provider will stop downloading chunks
    async fn receive_chunk(&self, ctx: &Context, arg: &Chunk) -> RpcResult<ChunkResponse> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "ChunkReceiver.ReceiveChunk",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ChunkResponse = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': ChunkResponse", e)))?;
        Ok(value)
    }
}
