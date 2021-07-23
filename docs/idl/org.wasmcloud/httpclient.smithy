// httpclient.smithy
// definition of http client capability contract

metadata package = [
    {
        namespace: "org.wasmcloud.interface.httpclient",
        crate: "wasmcloud_interface_httpclient"
     }
]

namespace org.wasmcloud.interface.httpclient

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U16

/// HttpClient - issue outgoing http requests via an external provider
/// To use this capability, the actor must be linked
/// with "wasmcloud:httpclient"
@wasmbus(
    contractId: "wasmcloud:httpclient",
    providerReceive: true )
service HttpClient {
  version: "0.1",
  operations: [ Request ]
}

/// Issue outgoing http request
operation Request {
    input: HttpRequest
    output: HttpResponse
}

/// http request to be sent through the provider
structure HttpRequest {

    @required
    method: String,

    @required
    url: String,

    @required
    headers: HeaderMap,

    @required
    body: Blob,
}

/// response from the http request
structure HttpResponse {
    /// response status code
    @required
    statusCode: U16,

    /// Case is not guaranteed to be normalized, so
    /// actors checking response headers need to do their own
    /// case conversion.
    /// Example (rust):
    ///   // check for 'Content-Type' header
    ///   let content_type:Option<&Vec<String>> = header.iter()
    ///          .map(|(k,_)| k.to_ascii_lowercase())
    ///          .find(|(k,_)| k == "content-type")
    ///          .map(|(_,v)| v);
    @required
    header: HeaderMap,

    /// response body
    @required
    body: Blob,
}

/// map data structure for holding http headers
///
map HeaderMap {
    key: String,
    value: HeaderValues,
}

list HeaderValues {
    member: String
}
