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

/// A map of test options.
/// Keys may be test case names, or other keys meaningful for the test.
/// Values are utf8 strings containing serialized json, with contents specific to the test
pub type OptMap = std::collections::HashMap<String, String>;

// Encode OptMap as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_opt_map<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &OptMap,
) -> RpcResult<()> {
    e.map(val.len() as u64)?;
    for (k, v) in val {
        e.str(k)?;
        e.str(v)?;
    }
    Ok(())
}

// Decode OptMap from cbor input stream
#[doc(hidden)]
pub fn decode_opt_map(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<OptMap, RpcError> {
    let __result = {
        {
            let mut m: std::collections::HashMap<String, String> =
                std::collections::HashMap::default();
            if let Some(n) = d.map()? {
                for _ in 0..(n as usize) {
                    let k = d.str()?.to_string();
                    let v = d.str()?.to_string();
                    m.insert(k, v);
                }
            } else {
                return Err(RpcError::Deser("indefinite maps not supported".to_string()));
            }
            m
        }
    };
    Ok(__result)
}
/// list of regex patterns
pub type PatternList = Vec<String>;

// Encode PatternList as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_pattern_list<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PatternList,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        e.str(item)?;
    }
    Ok(())
}

// Decode PatternList from cbor input stream
#[doc(hidden)]
pub fn decode_pattern_list(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<PatternList, RpcError> {
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
/// Options passed to all test cases
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct TestOptions {
    /// List of regex patterns for test names to run
    /// Default is ".*", to run all tests.
    pub patterns: PatternList,
    /// additional test configuration, optional
    /// Keys may be test case names, or other keys meaningful for the test.
    /// Values are serialized json, with contents specific to the test
    pub options: OptMap,
}

// Encode TestOptions as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_test_options<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &TestOptions,
) -> RpcResult<()> {
    e.array(2)?;
    encode_pattern_list(e, &val.patterns)?;
    encode_opt_map(e, &val.options)?;
    Ok(())
}

// Decode TestOptions from cbor input stream
#[doc(hidden)]
pub fn decode_test_options(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<TestOptions, RpcError> {
    let __result = {
        let mut patterns: Option<PatternList> = None;
        let mut options: Option<OptMap> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct TestOptions, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct TestOptions: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        patterns = Some(
                            decode_pattern_list(d)
                                .map_err(|e| format!("decoding 'PatternList': {}", e))?,
                        )
                    }
                    1 => {
                        options = Some(
                            decode_opt_map(d).map_err(|e| format!("decoding 'OptMap': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct TestOptions: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "patterns" => {
                        patterns = Some(
                            decode_pattern_list(d)
                                .map_err(|e| format!("decoding 'PatternList': {}", e))?,
                        )
                    }
                    "options" => {
                        options = Some(
                            decode_opt_map(d).map_err(|e| format!("decoding 'OptMap': {}", e))?,
                        )
                    }
                    _ => d.skip()?,
                }
            }
        }
        TestOptions {
            patterns: if let Some(__x) = patterns {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field TestOptions.patterns (#0)".to_string(),
                ));
            },

            options: if let Some(__x) = options {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field TestOptions.options (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct TestResult {
    /// test case name
    #[serde(default)]
    pub name: String,
    /// true if the test case passed
    #[serde(default)]
    pub passed: bool,
    /// (optional) more detailed results, if available.
    /// data is snap-compressed json
    /// failed tests should have a firsts-level key called "error".
    #[serde(rename = "snapData")]
    #[serde(with = "serde_bytes")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub snap_data: Option<Vec<u8>>,
}

// Encode TestResult as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_test_result<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &TestResult,
) -> RpcResult<()> {
    e.array(3)?;
    e.str(&val.name)?;
    e.bool(val.passed)?;
    if let Some(val) = val.snap_data.as_ref() {
        e.bytes(val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode TestResult from cbor input stream
#[doc(hidden)]
pub fn decode_test_result(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<TestResult, RpcError> {
    let __result = {
        let mut name: Option<String> = None;
        let mut passed: Option<bool> = None;
        let mut snap_data: Option<Option<Vec<u8>>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct TestResult, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct TestResult: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => name = Some(d.str()?.to_string()),
                    1 => passed = Some(d.bool()?),
                    2 => {
                        snap_data = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.bytes()?.to_vec()))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct TestResult: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "name" => name = Some(d.str()?.to_string()),
                    "passed" => passed = Some(d.bool()?),
                    "snapData" => {
                        snap_data = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.bytes()?.to_vec()))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        TestResult {
            name: if let Some(__x) = name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field TestResult.name (#0)".to_string(),
                ));
            },

            passed: if let Some(__x) = passed {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field TestResult.passed (#1)".to_string(),
                ));
            },
            snap_data: snap_data.unwrap(),
        }
    };
    Ok(__result)
}
pub type TestResults = Vec<TestResult>;

// Encode TestResults as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_test_results<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &TestResults,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_test_result(e, item)?;
    }
    Ok(())
}

// Decode TestResults from cbor input stream
#[doc(hidden)]
pub fn decode_test_results(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<TestResults, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<TestResult> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(
                    decode_test_result(d).map_err(|e| format!("decoding 'TestResult': {}", e))?,
                )
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<TestResult> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(
                        decode_test_result(d)
                            .map_err(|e| format!("decoding 'TestResult': {}", e))?,
                    ),
                }
            }
            arr
        }
    };
    Ok(__result)
}
/// Test api for testable actors and providers
/// wasmbus.contractId: wasmcloud:testing
/// wasmbus.providerReceive
/// wasmbus.actorReceive
#[async_trait]
pub trait Testing {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:testing"
    }
    /// Begin tests
    async fn start(&self, ctx: &Context, arg: &TestOptions) -> RpcResult<TestResults>;
}

/// TestingReceiver receives messages defined in the Testing service trait
/// Test api for testable actors and providers
#[doc(hidden)]
#[async_trait]
pub trait TestingReceiver: MessageDispatch + Testing {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Start" => {
                let value: TestOptions = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'TestOptions': {}", e)))?;
                let resp = Testing::start(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "Testing.Start",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "Testing::{}",
                message.method
            ))),
        }
    }
}

/// TestingSender sends messages to a Testing service
/// Test api for testable actors and providers
/// client for sending Testing messages
#[derive(Debug)]
pub struct TestingSender<T: Transport> {
    transport: T,
}

impl<T: Transport> TestingSender<T> {
    /// Constructs a TestingSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(not(target_arch = "wasm32"))]
impl<'send> TestingSender<wasmbus_rpc::provider::ProviderTransport<'send>> {
    /// Constructs a Sender using an actor's LinkDefinition,
    /// Uses the provider's HostBridge for rpc
    pub fn for_actor(ld: &'send wasmbus_rpc::core::LinkDefinition) -> Self {
        Self {
            transport: wasmbus_rpc::provider::ProviderTransport::new(ld, None),
        }
    }
}
#[cfg(target_arch = "wasm32")]
impl TestingSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for actor-to-actor messaging
    /// using the recipient actor's public key
    pub fn to_actor(actor_id: &str) -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_actor(actor_id.to_string()).unwrap();
        Self { transport }
    }
}

#[cfg(target_arch = "wasm32")]
impl TestingSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a Testing provider
    /// implementing the 'wasmcloud:testing' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:testing", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a Testing provider
    /// implementing the 'wasmcloud:testing' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:testing", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> Testing for TestingSender<T> {
    #[allow(unused)]
    /// Begin tests
    async fn start(&self, ctx: &Context, arg: &TestOptions) -> RpcResult<TestResults> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Testing.Start",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: TestResults = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': TestResults", e)))?;
        Ok(value)
    }
}
