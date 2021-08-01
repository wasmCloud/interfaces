//! org.wasmcloud.interace.httpserver

mod httpserver;
pub use httpserver::*;

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
