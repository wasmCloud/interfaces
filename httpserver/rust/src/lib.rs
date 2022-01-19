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
        let mut header = HeaderMap::new();
        header.insert(
            "Content-Type".to_string(),
            vec!["application/json".to_string()],
        );
        Ok(HttpResponse {
            body,
            status_code,
            header,
        })
    }

    /// Creates a response with a given status code, JSON-serialized payload, and headers specified by the header argument. Automatically includes the appropriate Content-Type header
    ///
    /// # Arguments
    ///
    /// * `payload` - Any struct implementing the Serialize trait from serde/serde_json
    /// * `status_code` - A 16-bit unsigned integer representing the outbound HTTP status code, e.g. 200 for successful interactions, 404 for not found, etc..
    /// * `headers` - A std::collections::HashMap from a String to a Vec() of Strings. The Content-Type header is ignored, for Content-Type will always default to application/json
    ///
    pub fn json_with_headers<T>(
        payload: T,
        status_code: u16,
        headers: std::collections::HashMap<String, Vec<String>>,
    ) -> Result<HttpResponse, RpcError>
    where
        T: Serialize,
    {
        let body = serde_json::to_string(&payload)
            .map_err(|e| RpcError::Ser(e.to_string()))?
            .into_bytes();

        let mut fixed_header = headers.clone();
        fixed_header.retain(|k, _| k.to_lowercase() != "content-type");
        fixed_header.insert(
            "Content-Type".to_string(),
            vec!["application/json".to_string()],
        );
        Ok(HttpResponse {
            body,
            status_code,
            header: fixed_header,
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

        let mut obj: std::collections::HashMap<&str, Vec<u16>> = std::collections::HashMap::new();
        obj.insert("list", vec![1, 2, 3]);
        let r = HttpResponse::json(&obj, 201);
        assert!(r.is_ok());
        let r = r.unwrap();
        assert_eq!(r.status_code, 201);
        assert_eq!(&r.body, br#"{"list":[1,2,3]}"#);
        let content_type = &r.header.get("Content-Type").unwrap();
        assert_eq!(**content_type, vec!("application/json".to_string()));

        let mut header: std::collections::HashMap<String, Vec<String>> =
            std::collections::HashMap::new();
        header.insert("X-something-one".to_owned(), vec!["foo".to_owned()]);
        header.insert("X-something-two".to_owned(), vec!["bar".to_owned()]);
        let r = HttpResponse::json_with_headers(&obj, 200, header);
        assert!(r.is_ok());
        let r = r.unwrap();
        assert_eq!(
            *r.header.get("X-something-one").unwrap(),
            vec!("foo".to_owned())
        );
        assert_eq!(
            *r.header.get("X-something-two").unwrap(),
            vec!("bar".to_owned())
        );
        assert_eq!(
            *r.header.get("Content-Type").unwrap(),
            vec!("application/json".to_owned())
        );
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
