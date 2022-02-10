// blobstore.smithy
// A service that stores objects (blobs) in named containers

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [{ 
    namespace: "org.wasmcloud.interface.blobstore", 
    crate: "wasmcloud-interface-blobstore" 
}]

namespace org.wasmcloud.interface.blobstore

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64

/// The BlobStore service, provider side
@wasmbus(
    contractId: "wasmcloud:blobstore",
    providerReceive: true )
service Blobstore {
    version: "0.1",
    operations: [ 
        CreateContainer,
        FindContainer,
        GetContainerInfo,
        ListContainers,
        RemoveContainers,
        ListObjects,
        RemoveObjects,
        StartUpload,
        UploadChunk,
        StartDownload,
        ReceiveChunk,
    ]
}

/// Creates a container by name, returning success if it worked
/// Note that names are not globally unique - just unique within the 
/// "namespace" of the connecting actor and linkdef
operation CreateContainer {
  input: ContainerId,
  output: BlobstoreResult
}

/// Searches for a container and returns success if it exists
@idempotent
operation FindContainer {
  input: ContainerId,
  output: BlobstoreResult
}

/// Retrieves information about the container,
/// Returns no value if the container id is invalid
operation GetContainerInfo {
    input: ContainerId,
    output: ContainerMetadata,
}

/// Returns list of all containers
operation ListContainers {
    output: Containers
}

/// Empty and remove the container(s)
operation RemoveContainers {
    input: ContainerIds
    output: BlobstoreResult
}

/// List the objects in the container
operation ListObjects {
    input: ContainerId,
    output: ObjectList,
}

/// Remove the objects.
/// The objects do not need to be in the same container
operation RemoveObjects {
    input: ObjectList,
    output: BlobstoreResult,
}

/// Requests to start upload of a file/blob to the Blobstore
/// It is recommended to keep chunks under 1MB to not exceed wasm memory allocation
operation StartUpload{
    input: UploadChunkArgs,
    output: BlobstoreResult,
}

/// Uploads a file chunk to a blobstore. This must be called AFTER
/// the StartUpload operation. Is is recommended to keep chunks
/// under 1MB to not exceed wasm memory allocation
operation UploadChunk {
    input: Chunk
    output: BlobstoreResult
}

/// Requests to start a download of a file/blob from the Blobstore
/// It is recommended to keep chunks under 1MB to not exceed wasm memory allocation
operation StartDownload{
    input: DownloadChunkArgs,
    output: DownloadResult,
}

/// Receives a file chunk from a blobstore. This must be called AFTER
/// the StartDownload operation. 
/// It is recommended to keep chunks under 1MB to not exceed wasm memory allocation
operation ReceiveChunk {
    input: DownloadChunkArgs,
    output: DownloadResult,
}


list ContainerIds {
    member: ContainerId
}

list Containers {
    member: ContainerMetadata
}

list ObjectList {
    member: ObjectMetadata
}

/// Unique id of a container
string ContainerId

/// Id of an object, unique to a container
string ObjectId

/// A container is a logical grouping of blobs, similar to a directory
/// in a file system.
structure ContainerMetadata {
    @required
    id: ContainerId,
    @required
    name: String,
}

structure ObjectMetadata {
    /// Object unique id
    @required
    id: ObjectId,

    /// container of the object
    @required
    containerId: ContainerId,

    /// size of the object in bytes
    @required
    size: U64,
}

structure DownloadChunkArgs {
    @required
    objectMetadata: ObjectMetadata,

    @required
    chunkSize: U64,

    @required
    sequenceNumber: U64,
}

structure UploadChunkArgs {
    @required
    objectMetadata: ObjectMetadata,

    @required
    chunkSize: U64,

    @required
    chunk: Chunk,
}

structure DownloadResult {
    @required
    success: Boolean,

    error: String,

    chunk: Chunk,
}

structure Chunk {
    @required
    objectData: ObjectMetadata,

    @required
    sequenceNo: U64,

    @required
    chunkSize: U64,

    @required
    bytes: Blob,

}

structure BlobstoreResult {
    @required
    success: Boolean,

    error: String,
  
}

