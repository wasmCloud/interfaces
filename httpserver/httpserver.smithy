// httpserver.smithy
// Definition of the wasmcloud:httpserver capability contract
//

metadata package = [{
    namespace: "org.wasmcloud.interface.httpserver",
    crate: "wasmcloud_interface_httpserver"
    py_module: "wasmcloud_interface_httpserver",
    doc: "Httpserver: wasmcloud capability contract for http server",
}]

namespace org.wasmcloud.interface.httpserver

use org.wasmcloud.model#codegenRust
use org.wasmcloud.model#n
use org.wasmcloud.model#U16
use org.wasmcloud.model#wasmbus

/// HttpServer is the contract to be implemented by actor
@wasmbus(
    contractId: "wasmcloud:httpserver",
    actorReceive: true )
service HttpServer {
  version: "0.1",
  operations: [ HandleRequest ]
}

operation HandleRequest {
  input: HttpRequest,
  output: HttpResponse,
}

/// HttpRequest contains data sent to actor about the http request
structure HttpRequest {

  /// HTTP method. One of: GET,POST,PUT,DELETE,HEAD,OPTIONS,CONNECT,PATCH,TRACE
  @required
  @n(0)
  method: String,

  /// full request path
  @required
  @n(1)
  path: String,

  /// query string. May be an empty string if there were no query parameters.
  @required
  @n(2)
  queryString: String,

  /// map of request headers (string key, string value)
  @required
  @n(3)
  header: HeaderMap,

  /// Request body as a byte array. May be empty.
  @required
  @n(4)
  body: Blob,
}

/// HttpResponse contains the actor's response to return to the http client
// don't generate Default so we can customize it
@codegenRust( noDeriveDefault: true )
structure HttpResponse {
  /// statusCode is a three-digit number, usually in the range 100-599,
  /// A value of 200 indicates success.
  @required
  @n(0)
  statusCode: U16,

  /// Map of headers (string keys, list of values)
  @required
  @n(1)
  header: HeaderMap,

  /// Body of response as a byte array. May be an empty array.
  @required
  @n(2)
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
