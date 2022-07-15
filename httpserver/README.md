[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-httpserver.svg)](https://crates.io/crates/wasmcloud-interface-httpserver)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=httpserver%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/httpserver/tinygo)
# wasmCloud HTTP Server Interface
This is the interface for an HTTP Server capability with the contract ID `wasmcloud:httpserver`

This folder contains 
- Model definition for `wasmcloud:httpserver`
- Generated documentation (in HTML)
- Generated Rust library (in Rust)

Any Rust actor or capability provider using `wasmcloud:httpserver` should rely upon this library. A capability provider implements the trait `HttpServerReceiver`.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:httpserver` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [HTTPServer](https://github.com/wasmCloud/capability-providers/tree/main/httpserver-rs) | wasmCloud | wasmCloud HTTP Server implementation using the highly scalable [warp](https://docs.rs/warp/latest/warp/) web server.


## Example Usage (🦀 Rust)
Implementing the `HttpServer.HandleRequest` operation
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_httpserver::{HttpRequest, HttpResponse, HttpServer, HttpServerReceiver};

#[derive(Debug, Default, Actor, HealthResponder)]
#[services(Actor, HttpServer)]
struct HelloActor {}

#[async_trait]
impl HttpServer for HelloActor {
    async fn handle_request(&self, _ctx: &Context, _req: &HttpRequest) -> RpcResult<HttpResponse> {
        Ok(HttpResponse {
            body: "Hello World".as_bytes().to_owned(),
            ..Default::default()
        })
    }
}
```
