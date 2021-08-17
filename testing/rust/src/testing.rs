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

/// A map of test options.
/// Keys may be test case names, or other keys meaningful for the test.
/// Values are utf8 strings containing serialized json, with contents specific to the test
pub type OptMap = std::collections::HashMap<String, String>;

/// list of regex patterns
pub type PatternList = Vec<String>;

/// Options passed to all test cases
#[derive(Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
pub struct TestOptions {
    /// additional test configuration, optional
    /// Keys may be test case names, or other keys meaningful for the test.
    /// Values are serialized json, with contents specific to the test
    pub options: OptMap,
    /// List of regex patterns for test names to run
    /// Default is ".*", to run all tests.
    pub patterns: PatternList,
}

#[derive(Default, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
pub struct TestResult {
    /// test case name
    #[serde(default)]
    pub name: String,
    /// true if the test case passed
    #[serde(default)]
    pub pass: bool,
    /// (optional) more detailed results, if available.
    /// data is snap-compressed json
    /// failed tests should have a firsts-level key called "error".
    #[serde(rename = "snapData")]
    #[serde(with = "serde_bytes")]
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub snap_data: Option<Vec<u8>>,
}

pub type TestResults = Vec<TestResult>;

/// Test api for testable actors and providers
/// wasmbus.contractId: wasmcloud:testing
/// wasmbus.providerReceive
/// wasmbus.actorReceive
#[async_trait]
pub trait Testing {
    /// Begin tests
    async fn start(&self, ctx: &Context, arg: &TestOptions) -> RpcResult<TestResults>;
}

/// TestingReceiver receives messages defined in the Testing service trait
/// Test api for testable actors and providers
#[async_trait]
pub trait TestingReceiver: MessageDispatch + Testing {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "Start" => {
                let value: TestOptions = deserialize(message.arg.as_ref()).map_err(|e| {
                    RpcError::Deser(format!(
                        "deserialization for message '{}': {}",
                        message.method, e
                    ))
                })?;
                let resp = Testing::start(self, ctx, &value).await?;
                let buf = Cow::Owned(serialize(&resp)?);
                Ok(Message {
                    method: "Testing.Start",
                    arg: buf,
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
#[derive(Debug)]
pub struct TestingSender<'send, T> {
    transport: &'send T,
}

impl<'send, T: Transport> TestingSender<'send, T> {
    pub fn new(transport: &'send T) -> Self {
        TestingSender { transport }
    }
}

#[async_trait]
impl<'send, T: Transport + std::marker::Sync + std::marker::Send> Testing
    for TestingSender<'send, T>
{
    #[allow(unused)]
    /// Begin tests
    async fn start(&self, ctx: &Context, arg: &TestOptions) -> RpcResult<TestResults> {
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Start",
                    arg: Cow::Borrowed(&arg),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "Start", e)))?;
        Ok(value)
    }
}
