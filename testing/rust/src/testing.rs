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
#[derive(Clone, Debug, Deserialize, Eq, PartialEq, Serialize)]
pub struct TestOptions {
    /// additional test configuration, optional
    /// Keys may be test case names, or other keys meaningful for the test.
    /// Values are serialized json, with contents specific to the test
    pub options: OptMap,
    /// List of regex patterns for test names to run
    /// Default is ".*", to run all tests.
    pub patterns: PatternList,
}

#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
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
                let value: TestOptions = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
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
        let arg = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "Testing.Start",
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
