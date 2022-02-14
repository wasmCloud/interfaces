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
use org.wasmcloud.model#n
use org.wasmcloud.model#codegenRust

/// The BlobStore service, provider side
@wasmbus(
    contractId: "wasmcloud:blobstore",
    providerReceive: true,
    protocol: "2" )
service Blobstore {
    version: "0.1",
    operations: [
        // container operations
        ContainerExists,
        CreateContainer,
        GetContainerInfo,
        ListContainers,
        RemoveContainers,
        
        // object operations
        ObjectExists,
        ListObjects,
        RemoveObjects,
        PutObject,
        GetObject,
        PutChunk,
    ]
}

/// The BlobStore service, actor side
@wasmbus(
    contractId: "wasmcloud:blobstore",
    actorReceive: true )
service ChunkReceiver {
    version: "0.1",
    operations: [
        ReceiveChunk,
    ]
}

/// Creates a container by name, returning success if it worked
/// Note that container names may not be globally unique - just unique within the
/// "namespace" of the connecting actor and linkdef
@idempotent
operation CreateContainer {
  input: ContainerId,
}

/// Returns whether the container exists
@readonly
operation ContainerExists {
  input: ContainerId,
  output: Boolean,
}

/// Returns whether the object exists
@readonly
operation ObjectExists {
  input: ContainerObject,
  output: Boolean,
}

/// Returns list of container ids
@readonly
operation ListContainers {
    output: ContainersInfo,
}

/// Empty and remove the container(s)
/// The MultiResult list contains one entry for each container
/// that was not successfully removed, with the 'key' value representing the container name.
/// If the MultiResult list is empty, all container removals succeeded.
operation RemoveContainers {
    input: ContainerIds,
    output: MultiResult,
}

/// Retrieves information about the container.
/// Returns error if the container id is invalid or not found.
@readonly
operation GetContainerInfo {
    input: ContainerId,
    output: ContainerMetadata,
}

/// Lists the objects in the container.
/// If the container exists and is empty, the returned `objects` list is empty.
/// Parameters of the request may be used to limit the object names returned
/// with an optional start value, end value, and maximum number of items.
/// The provider may limit the number of items returned. If the list is truncated, 
/// the response contains a `continuation` token that may be submitted in
/// a subsequent ListObjects request.
@readonly
operation ListObjects {
    input: ListObjectsRequest,
    output: ListObjectsResponse,
}

/// Removes the objects. In the event any of the objects cannot be removed,
/// the operation continues until all requested deletions have been attempted.
/// The MultiRequest includes a list of errors, one for each deletion request
/// that did not succeed. If the list is empty, all removals succeeded.
operation RemoveObjects {
    input: RemoveObjectsRequest,
    output: MultiResult,
}

/// Requests to start upload of a file/blob to the Blobstore.
/// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
operation PutObject {
    input: PutObjectRequest,
    output: PutObjectResponse,
}

/// Uploads a file chunk to a blobstore. This must be called AFTER PutObject
/// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
operation PutChunk {
    input: PutChunkRequest,
}

/// Requests to retrieve an object. If the object is large, the provider
/// may split the response into multiple parts
/// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
@readonly
operation GetObject {
    input: GetObjectRequest,
    output: GetObjectResponse,
}

/// Receives a file chunk from a blobstore.
/// A blobstore provider invokes this operation on actors in response to the GetObject request.
/// If the response sets cancelDownload, the provider will stop downloading chunks
operation ReceiveChunk {
    input: Chunk,
    output: ChunkResponse,
}

/// Response from actor after receiving a download chunk.
structure ChunkResponse {

    /// If set and `true`, the sender will stop sending chunks, 
    cancelDownload: Boolean
}

/// Parameter to list_objects. 
structure ListObjectsRequest {

    /// Name of the container to search
    @required
    containerId: String,

    /// Request object names starting with this value. (Optional)
    startWith: String,

    /// Continuation token passed in ListObjectsResponse.
    /// If set, `startWith` is ignored. (Optional)
    continuation: String,

    /// Last item to return (inclusive terminator) (Optional)
    endWith: String,

    /// Optionally, stop returning items before returning this value.
    /// (exclusive terminator)
    /// If startFrom is "a" and endBefore is "b", and items are ordered
    /// alphabetically, then only items beginning with "a" would be returned.
    /// (Optional)
    endBefore: String,

    /// maximum number of items to return. If not specified, provider 
    /// will return an initial set of up to 1000 items. if maxItems > 1000,
    /// the provider implementation may return fewer items than requested.
    /// (Optional)
    maxItems: U32,
}

/// Respose to list_objects.
/// If `isLast` is false, the list was truncated by the provider,
/// and the remainder of the objects can be requested with another
/// request using the `continuation` token.
structure ListObjectsResponse {

    /// set of objects returned
    @required
    objects: ObjectsInfo,

    /// Indicates if the item list is complete, or the last item
    /// in a multi-part response.
    @required
    isLast: Boolean,

    /// If `isLast` is false, this value can be used in the `continuation` field
    /// of a `ListObjectsRequest`.
    /// Clients should not attempt to interpret this field: it may or may not
    /// be a real key or object name, and may be obfuscated by the provider.
    continuation: String,
}


/// parameter to removeObjects
structure RemoveObjectsRequest {

    /// name of container
    @required
    containerId: ContainerId,

    /// list of object names to be removed
    @required
    objects: ObjectIds,
}

/// Name of a container
string ContainerId

/// Name of an object within a container
string ObjectId

/// Metadata for a container.
structure ContainerMetadata {
    /// Container name
    @required
    containerId: ContainerId,

    /// Creation date, if available
    createdAt: Timestamp,
}

/// Parameter for PutObject operation
structure PutObjectRequest {

    /// File path and initial data
    @required
    chunk: Chunk,

    /// A MIME type of the object
    /// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17
    contentType: String,

    /// Specifies what content encodings have been applied to the object 
    /// and thus what decoding mechanisms must be applied to obtain the media-type 
    /// referenced by the contentType field. For more information, 
    /// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11.
    contentEncoding: String,
}

/// Response to PutObject operation
structure PutObjectResponse {

    /// If this is a multipart upload, `streamId` must be returned
    /// with subsequent PutChunk requests
    streamId: String,
}

/// Parameter to PutChunk operation
structure PutChunkRequest {

    /// upload chunk from the file.
    /// if chunk.isLast is set, this will be the last chunk uploaded
    @required
    chunk: Chunk,

    /// This value should be set to the `streamId` returned from the initial PutObject.
    streamId: String,

    /// If set, the receiving provider should cancel the upload process 
    /// and remove the file.
    cancelAndRemove: Boolean,
}

structure ObjectMetadata {
    /// Object identifier that is unique within its container.
    /// Naming of objects is determined by the capability provider.
    /// An object id could be a path, hash of object contents, or some other unique identifier.
    @required
    objectId: ObjectId,

    /// container of the object
    @required
    containerId: ContainerId,

    /// size of the object in bytes
    @required
    size: U64,

    /// date object was last modified
    lastModified: Timestamp,
}

/// Parameter to GetObject
structure GetObjectRequest {

    /// object to download
    @required
    objectId: ObjectId,

    /// object's container
    @required
    containerId: ContainerId,

    /// optional size requested
    /// The provider will not return a chunk larger than this amount,
    /// but may return a smaller chunk.
    chunkSize: U64,

    /// Requested start of object to retrieve.
    /// The first byte is at offset 0. Range values are inclusive.
    /// If rangeStart is beyond the end of the file,
    /// an empty chunk will be returned with isLast == true
    rangeStart: U64,

    /// Requested end of object to retrieve. Defaults to the object's size.
    /// It is not an error for rangeEnd to be greater than the object size.
    /// Range values are inclusive.
    rangeEnd: U64,
}

/// Response to GetObject
structure GetObjectResponse {

    /// indication whether the request was successful
    @required
    success: Boolean

    /// If success is false, this may contain an error
    error: String

    /// The provider may begin the download by returning a first chunk
    initialChunk: Chunk

    /// Length of the content. (for multi-part downloads, this may not
    /// be the same as the length of the initial chunk)
    @required
    contentLength: u64

    /// A standard MIME type describing the format of the object data.
    contentType: String

    /// Specifies what content encodings have been applied to the object 
    /// and thus what decoding mechanisms must be applied to obtain the media-type 
    // referenced by the contenType field.
    contentEncoding: String
}

/// Combination of container id and object id
structure ContainerObject {
    @required
    containerId: ContainerId,

    @required
    objectId: ObjectId,
}

/// A portion of a file. The `isLast` field indicates whether this chunk
/// is the last in a stream. The `offset` field indicates the 0-based offset
/// from the start of the file for this chunk.
structure Chunk {
    @required
    objectId: ObjectId,

    @required
    containerId: ContainerId,

    /// bytes in this chunk
    @required
    bytes: Blob,

    /// The byte offset within the object for this chunk
    @required
    offset: U64,

    /// true if this is the last chunk
    @required
    isLast: Boolean,
}

/// Result of input item
structure ItemResult {
    @required
    key: String,

    /// whether the item succeeded or failed
    @required
    success: Boolean

    /// optional error message for failures
    error: String
}

/// result for an operation on a list of inputs
list MultiResult {
    member : ItemResult
}

/// list of object names
list ObjectIds {
    member: ObjectId
}

/// list of container names
list ContainerIds {
    member: ContainerId
}

/// list of container metadata objects
list ContainersInfo {
    member: ContainerMetadata
}

/// list of object metadata objects
list ObjectsInfo {
    member: ObjectMetadata
}
