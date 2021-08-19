//! actor interface for 'wasmcloud:httpclient' capability
//!

mod httpclient;
pub use httpclient::*;

impl Default for HttpRequest {
    /// constructs a default request with GET method
    fn default() -> HttpRequest {
        HttpRequest {
            method: "GET".to_string(),
            url: String::default(),
            headers: HeaderMap::default(),
            body: Vec::default(),
        }
    }
}

impl HttpRequest {
    /// convenience method to create HttpRequest with GET method and url
    pub fn get(url: &str) -> HttpRequest {
        HttpRequest {
            url: url.to_string(),
            ..Default::default()
        }
    }

    /// convenience method to create HttpRequest with POST method, url and body
    pub fn post(url: &str, body: Vec<u8>) -> HttpRequest {
        HttpRequest {
            method: "POST".to_string(),
            url: url.to_string(),
            body,
            ..Default::default()
        }
    }

    /// convenience method to create HttpRequest with POST method, url and body
    pub fn put(url: &str, body: Vec<u8>) -> HttpRequest {
        HttpRequest {
            method: "PUT".to_string(),
            url: url.to_string(),
            body,
            ..Default::default()
        }
    }
}

impl Default for HttpResponse {
    /// constructs a default response with status 200, empty body, and no headers
    fn default() -> HttpResponse {
        HttpResponse {
            status_code: 200,
            header: HeaderMap::default(),
            body: Vec::default(),
        }
    }
}
