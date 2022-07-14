[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-blobstore.svg)](https://crates.io/crates/wasmcloud-interface-blobstore)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=blobstore%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/blobstore/tinygo)
# wasmCloud Blobstore Interface

The blobstore interface abstracts a service (capability provider) that can manage containers and objects. Actors that use this interface must have the capability contract `wasmcloud:blobstore` in their claims list (`wash claims sign --blob_store`).

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:blobstore` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [blobstore-s3](https://github.com/wasmCloud/capability-providers/tree/main/blobstore-s3) | wasmCloud | An AWS S3 implementation of a blobstore that manages S3 buckets and objects
| ⚠️ WIP: [blobstore-fs](https://github.com/wasmCloud/capability-providers/pull/154) | wasmCloud | An implementation that manages folders and files on a filesystem

## Example Usage (🦀 Rust)
Create a container in a blobstore:
```rust
use std::result::Result;
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_blobstore::{Blobstore, BlobstoreSender};
async fn create_container(ctx: &Context, container_name: &str) -> Result<(), RpcError> {
    let blobstore = BlobstoreSender::new();
    blobstore
        .create_container(ctx, &container_name.to_string())
        .await
}

```
Uploading an object (image bytes) to a blobstore:
```rust
use std::result::Result;
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_blobstore::{
    Blobstore, BlobstoreSender, Chunk, PutObjectRequest, PutObjectResponse,
};
async fn upload_bytes(ctx: &Context, image_bytes: &[u8]) -> Result<PutObjectResponse, RpcError> {
    BlobstoreSender::new()
        .put_object(
            ctx,
            &PutObjectRequest {
                chunk: Chunk {
                    container_id: "myfolder".to_string(),
                    object_id: "myobjectname".to_string(),
                    bytes: image_bytes.to_vec(),
                    offset: 0,
                    is_last: true,
                },
                content_type: Some("image/png".to_string()),
                ..Default::default()
            },
        )
        .await
}
```

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)
Coming Soon