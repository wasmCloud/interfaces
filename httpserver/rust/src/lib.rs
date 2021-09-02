//! org.wasmcloud.interface.httpserver

mod httpserver;
pub use httpserver::*;
use serde::Serialize;
use std::collections::HashMap;

impl Default for HttpResponse {
    /// create default HttpResponse with status 200, empty body and empty header
    fn default() -> HttpResponse {
        HttpResponse {
            status_code: 200,
            body: Vec::default(),
            header: HeaderMap::default(),
        }
    }
}

impl HttpResponse {
    /// Creates a response with a given status code and serializes the given payload as JSON
    pub fn json<T>(payload: T, status_code: u16) -> HttpResponse
    where
        T: Serialize,
    {
        HttpResponse {
            body: serde_json::to_string(&payload).unwrap().into_bytes(),
            header: HashMap::new(),
            status_code,
        }
    }

    /// Handy shortcut for creating a 404/Not Found response
    pub fn not_found() -> HttpResponse {
        HttpResponse {
            status_code: 404,
            ..Default::default()
        }
    }

    /// Useful shortcut for creating a 200/OK response
    pub fn ok() -> HttpResponse {
        HttpResponse {
            status_code: 200,
            ..Default::default()
        }
    }

    /// Useful shortcut for creating a 500/Internal Server Error response
    pub fn internal_server_error(msg: &str) -> HttpResponse {
        HttpResponse {
            status_code: 500,
            body: msg.to_string().as_bytes().into(),
            ..Default::default()
        }
    }

    /// Shortcut for creating a 400/Bad Request response
    pub fn bad_request(msg: &str) -> HttpResponse {
        HttpResponse {
            status_code: 400,
            body: msg.to_string().as_bytes().into(),
            ..Default::default()
        }
    }
}
