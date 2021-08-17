//! org.wasmcloud.interface.testing

mod testing;
use serde::Serialize;
use serde_json::json;
pub use testing::*;
use wasmbus_rpc::RpcResult;

impl Default for TestOptions {
    fn default() -> TestOptions {
        TestOptions {
            patterns: vec![".*".to_string()],
            options: std::collections::HashMap::default(),
        }
    }
}

pub type NamedResult<'nam, T> = (&'nam str, RpcResult<T>);

// convert empty RpcResult into a testResult
impl<'nam, T: Serialize> From<NamedResult<'nam, T>> for TestResult {
    fn from(name_res: NamedResult<'nam, T>) -> TestResult {
        match name_res.1 {
            Ok(res) => {
                // TODO: if serialization of data fails, it doesn't change
                // the test result, but serialization errors should be logged.
                // Logging requires us to have logging set up
                // (we might be running in an actor)
                let data = match serde_json::to_vec(&res) {
                    Ok(v) => serde_json::to_vec(&json!({ "data": v })).unwrap_or_default(),
                    Err(_) => b"".to_vec(),
                };
                TestResult {
                    name: name_res.0.to_string(),
                    pass: true,
                    snap_data: Some(data),
                }
            }
            Err(e) => {
                let data = serde_json::to_vec(&json!(
                    {
                       "error": e.to_string(),
                    }
                ))
                .ok();
                TestResult {
                    name: name_res.0.to_string(),
                    pass: false,
                    snap_data: data,
                }
            }
        }
    }
}
