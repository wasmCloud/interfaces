//! org.wasmcloud.interace.httpserver

mod httpserver;
pub use httpserver::*;

impl Default for HttpResponse {
    fn default() -> HttpResponse {
        HttpResponse {
            status_code: 200,
            body: Vec::default(),
            header: HeaderMap::default(),
        }
    }
}