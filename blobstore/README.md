# Interface for the Blobstore, wasmcloud:blobstore

This is an interface for accessing a blobstore similar to AWS S3 and Azure Blob storage. It can also be implemented with a file system provider as backend. 

A blobstore consists of a number of _containers_ which each may contain a number of _objects_. 

The interface consists of the following operations:

* CreateContainer, input argument is an id for a container which needs to be system-wide unique. 
* FindContainer, input argument container id.
* GetContainerInfo, input argument container id. Returns metadata (if any) about a container.
* ListContainers, no input argument. Returns a list of all containers.
* RemoveContainers, removes the containers in the input list.
* ListObjects, returns a list of objects in container specified as input.
* RemoveObjects, removes the objects sent as input. Objects are specified with object/container id. They may be in different containers.
* StartUpload, initiates and upload of an object to the blobstore including the first chunk.
* UploadChunk, if objects is larger than a specified configurable size (capability provider implementation dependent), then this operation sends additional chunks to upload.
* StartDownload, initiates download of an object, including the first chunk.
* ReceiveChunk, operation which will return additional chunks, if any.

If not specified otherwise, each operation returns a status if it was succesful or not.