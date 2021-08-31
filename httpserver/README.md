# HTTP Server Interface
This is the interface for an HTTP Server capability with the contract ID `wasmcloud:httpserver`

This folder contains 
- Model definition for `wasmcloud:httpserver`
- Generated documentation (in HTML)
- Generated Rust library (in Rust)

Any Rust actor or capability provider using `wasmcloud:httpserver` should rely upon this library. A capability provider implements the trait `HttpServerReceiver`.

## Implementations
The following is a list of known implementations of the HTTP server interface. Feel free to submit a PR if you know of others.

| Name | Vendor | Description |
| :---: | :---: | :--- |
| [Default Server](https://github.com/wasmCloud/capability-providers/tree/main/httpserver-rs) | wasmCloud | wasmCloud Default HTTP Server Provider


