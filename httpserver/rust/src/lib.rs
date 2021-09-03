//! org.wasmcloud.interface.httpserver

mod httpserver;
pub use httpserver::*;
use serde::Serialize;
use std::string::ToString;
use wasmbus_rpc::RpcError;

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
    pub fn json<T>(payload: T, status_code: u16) -> Result<HttpResponse, RpcError>
    where
        T: Serialize,
    {
        let body = serde_json::to_string(&payload)
            .map_err(|e| RpcError::Ser(e.to_string()))?
            .into_bytes();
        Ok(HttpResponse {
            body,
            status_code,
            ..Default::default()
        })
    }

    /// Handy shortcut for creating a 404/Not Found response
    pub fn not_found() -> HttpResponse {
        HttpResponse {
            status_code: 404,
            ..Default::default()
        }
    }

    /// Useful shortcut for creating a 500/Internal Server Error response
    pub fn internal_server_error<T: ToString>(msg: T) -> HttpResponse {
        HttpResponse {
            status_code: 500,
            body: msg.to_string().as_bytes().into(),
            ..Default::default()
        }
    }

    /// Shortcut for creating a 400/Bad Request response
    pub fn bad_request<T: ToString>(msg: T) -> HttpResponse {
        HttpResponse {
            status_code: 400,
            body: msg.to_string().as_bytes().into(),
            ..Default::default()
        }
    }
}

#[cfg(test)]
mod test {

    use super::HttpResponse;
    use serde::{Serialize, Serializer};

    #[test]
    fn http_response_constructors() {
        let r = HttpResponse::default();
        assert_eq!(r.status_code, 200);
        assert_eq!(r.body.len(), 0usize);

        let r = HttpResponse::bad_request("Bad");
        assert_eq!(r.status_code, 400);
        assert_eq!(&r.body, b"Bad");

        let r = HttpResponse::internal_server_error("Oops!");
        assert_eq!(r.status_code, 500);
        assert_eq!(&r.body, b"Oops!");

        let r = HttpResponse::not_found();
        assert_eq!(r.status_code, 404);
        assert_eq!(r.body.len(), 0usize);

        let mut obj = std::collections::HashMap::new();
        obj.insert("list", vec![1, 2, 3]);
        let r = HttpResponse::json(&obj, 201);
        assert!(r.is_ok());
        let r = r.unwrap();
        assert_eq!(r.status_code, 201);
        assert_eq!(&r.body, br#"{"list":[1,2,3]}"#);
    }

    struct Thing {
        value: i32,
    }
    impl Serialize for Thing {
        fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
        where
            S: Serializer,
        {
            match self.value {
                13 => Err(serde::ser::Error::custom("you choose poorly")),
                _ => serializer.serialize_i32(self.value),
            }
        }
    }

    #[test]
    fn json_err() {
        let r = HttpResponse::json(&{ Thing { value: 0 } }, 200);
        assert!(r.is_ok());

        let r = HttpResponse::json(&{ Thing { value: 13 } }, 200);
        assert!(r.is_err());
    }
}
