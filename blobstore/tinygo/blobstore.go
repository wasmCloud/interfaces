// Blobstore: wasmcloud capability contract for storing objects (blobs) in named containers
package blobstore

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// A portion of a file. The `isLast` field indicates whether this chunk
// is the last in a stream. The `offset` field indicates the 0-based offset
// from the start of the file for this chunk.
type Chunk struct {
	ObjectId    ObjectId
	ContainerId ContainerId
	// bytes in this chunk
	Bytes []byte
	// The byte offset within the object for this chunk
	Offset uint64
	// true if this is the last chunk
	IsLast bool
}

// MEncode serializes a Chunk using msgpack
func (o *Chunk) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("objectId")
	o.ObjectId.MEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("bytes")
	encoder.WriteByteArray(o.Bytes)
	encoder.WriteString("offset")
	encoder.WriteUint64(o.Offset)
	encoder.WriteString("isLast")
	encoder.WriteBool(o.IsLast)

	return encoder.CheckError()
}

// MDecodeChunk deserializes a Chunk using msgpack
func MDecodeChunk(d *msgpack.Decoder) (Chunk, error) {
	var val Chunk
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = MDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "bytes":
			val.Bytes, err = d.ReadByteArray()
		case "offset":
			val.Offset, err = d.ReadUint64()
		case "isLast":
			val.IsLast, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Chunk using cbor
func (o *Chunk) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("objectId")
	o.ObjectId.CEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("bytes")
	encoder.WriteByteArray(o.Bytes)
	encoder.WriteString("offset")
	encoder.WriteUint64(o.Offset)
	encoder.WriteString("isLast")
	encoder.WriteBool(o.IsLast)

	return encoder.CheckError()
}

// CDecodeChunk deserializes a Chunk using cbor
func CDecodeChunk(d *cbor.Decoder) (Chunk, error) {
	var val Chunk
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = CDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "bytes":
			val.Bytes, err = d.ReadByteArray()
		case "offset":
			val.Offset, err = d.ReadUint64()
		case "isLast":
			val.IsLast, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Response from actor after receiving a download chunk.
type ChunkResponse struct {
	// If set and `true`, the sender will stop sending chunks,
	CancelDownload bool
}

// MEncode serializes a ChunkResponse using msgpack
func (o *ChunkResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("cancelDownload")
	encoder.WriteBool(o.CancelDownload)

	return encoder.CheckError()
}

// MDecodeChunkResponse deserializes a ChunkResponse using msgpack
func MDecodeChunkResponse(d *msgpack.Decoder) (ChunkResponse, error) {
	var val ChunkResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "cancelDownload":
			val.CancelDownload, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ChunkResponse using cbor
func (o *ChunkResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("cancelDownload")
	encoder.WriteBool(o.CancelDownload)

	return encoder.CheckError()
}

// CDecodeChunkResponse deserializes a ChunkResponse using cbor
func CDecodeChunkResponse(d *cbor.Decoder) (ChunkResponse, error) {
	var val ChunkResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "cancelDownload":
			val.CancelDownload, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Name of a container
type ContainerId string

// MEncode serializes a ContainerId using msgpack
func (o *ContainerId) MEncode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// MDecodeContainerId deserializes a ContainerId using msgpack
func MDecodeContainerId(d *msgpack.Decoder) (ContainerId, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ContainerId(val), nil
}

// CEncode serializes a ContainerId using cbor
func (o *ContainerId) CEncode(encoder cbor.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// CDecodeContainerId deserializes a ContainerId using cbor
func CDecodeContainerId(d *cbor.Decoder) (ContainerId, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ContainerId(val), nil
}

// list of container names
type ContainerIds []ContainerId

// MEncode serializes a ContainerIds using msgpack
func (o *ContainerIds) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeContainerIds deserializes a ContainerIds using msgpack
func MDecodeContainerIds(d *msgpack.Decoder) (ContainerIds, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ContainerId, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ContainerId, 0), err
	}
	val := make([]ContainerId, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeContainerId(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ContainerIds using cbor
func (o *ContainerIds) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeContainerIds deserializes a ContainerIds using cbor
func CDecodeContainerIds(d *cbor.Decoder) (ContainerIds, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ContainerId, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ContainerId, 0), err
	}
	val := make([]ContainerId, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeContainerId(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Metadata for a container.
type ContainerMetadata struct {
	// Container name
	ContainerId ContainerId
	// Creation date, if available
	CreatedAt *actor.Timestamp
}

// MEncode serializes a ContainerMetadata using msgpack
func (o *ContainerMetadata) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("createdAt")
	if o.CreatedAt == nil {
		encoder.WriteNil()
	} else {
		o.CreatedAt.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeContainerMetadata deserializes a ContainerMetadata using msgpack
func MDecodeContainerMetadata(d *msgpack.Decoder) (ContainerMetadata, error) {
	var val ContainerMetadata
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "createdAt":
			fval, err := actor.MDecodeTimestamp(d)
			if err != nil {
				return val, err
			}
			val.CreatedAt = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ContainerMetadata using cbor
func (o *ContainerMetadata) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("createdAt")
	if o.CreatedAt == nil {
		encoder.WriteNil()
	} else {
		o.CreatedAt.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeContainerMetadata deserializes a ContainerMetadata using cbor
func CDecodeContainerMetadata(d *cbor.Decoder) (ContainerMetadata, error) {
	var val ContainerMetadata
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "createdAt":
			fval, err := actor.CDecodeTimestamp(d)
			if err != nil {
				return val, err
			}
			val.CreatedAt = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Combination of container id and object id
type ContainerObject struct {
	ContainerId ContainerId
	ObjectId    ObjectId
}

// MEncode serializes a ContainerObject using msgpack
func (o *ContainerObject) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("objectId")
	o.ObjectId.MEncode(encoder)

	return encoder.CheckError()
}

// MDecodeContainerObject deserializes a ContainerObject using msgpack
func MDecodeContainerObject(d *msgpack.Decoder) (ContainerObject, error) {
	var val ContainerObject
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "objectId":
			val.ObjectId, err = MDecodeObjectId(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ContainerObject using cbor
func (o *ContainerObject) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("objectId")
	o.ObjectId.CEncode(encoder)

	return encoder.CheckError()
}

// CDecodeContainerObject deserializes a ContainerObject using cbor
func CDecodeContainerObject(d *cbor.Decoder) (ContainerObject, error) {
	var val ContainerObject
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "objectId":
			val.ObjectId, err = CDecodeObjectId(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// list of container metadata objects
type ContainersInfo []ContainerMetadata

// MEncode serializes a ContainersInfo using msgpack
func (o *ContainersInfo) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeContainersInfo deserializes a ContainersInfo using msgpack
func MDecodeContainersInfo(d *msgpack.Decoder) (ContainersInfo, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ContainerMetadata, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ContainerMetadata, 0), err
	}
	val := make([]ContainerMetadata, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeContainerMetadata(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ContainersInfo using cbor
func (o *ContainersInfo) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeContainersInfo deserializes a ContainersInfo using cbor
func CDecodeContainersInfo(d *cbor.Decoder) (ContainersInfo, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ContainerMetadata, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ContainerMetadata, 0), err
	}
	val := make([]ContainerMetadata, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeContainerMetadata(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Parameter to GetObject
type GetObjectRequest struct {
	// object to download
	ObjectId ObjectId
	// object's container
	ContainerId ContainerId
	// Requested start of object to retrieve.
	// The first byte is at offset 0. Range values are inclusive.
	// If rangeStart is beyond the end of the file,
	// an empty chunk will be returned with isLast == true
	RangeStart uint64
	// Requested end of object to retrieve. Defaults to the object's size.
	// It is not an error for rangeEnd to be greater than the object size.
	// Range values are inclusive.
	RangeEnd uint64
}

// MEncode serializes a GetObjectRequest using msgpack
func (o *GetObjectRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("objectId")
	o.ObjectId.MEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("rangeStart")
	encoder.WriteUint64(o.RangeStart)
	encoder.WriteString("rangeEnd")
	encoder.WriteUint64(o.RangeEnd)

	return encoder.CheckError()
}

// MDecodeGetObjectRequest deserializes a GetObjectRequest using msgpack
func MDecodeGetObjectRequest(d *msgpack.Decoder) (GetObjectRequest, error) {
	var val GetObjectRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = MDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "rangeStart":
			val.RangeStart, err = d.ReadUint64()
		case "rangeEnd":
			val.RangeEnd, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a GetObjectRequest using cbor
func (o *GetObjectRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("objectId")
	o.ObjectId.CEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("rangeStart")
	encoder.WriteUint64(o.RangeStart)
	encoder.WriteString("rangeEnd")
	encoder.WriteUint64(o.RangeEnd)

	return encoder.CheckError()
}

// CDecodeGetObjectRequest deserializes a GetObjectRequest using cbor
func CDecodeGetObjectRequest(d *cbor.Decoder) (GetObjectRequest, error) {
	var val GetObjectRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = CDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "rangeStart":
			val.RangeStart, err = d.ReadUint64()
		case "rangeEnd":
			val.RangeEnd, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Response to GetObject
type GetObjectResponse struct {
	// indication whether the request was successful
	Success bool
	// If success is false, this may contain an error
	Error string
	// The provider may begin the download by returning a first chunk
	InitialChunk *Chunk
	// Length of the content. (for multi-part downloads, this may not
	// be the same as the length of the initial chunk)
	ContentLength uint64
	// A standard MIME type describing the format of the object data.
	ContentType string
	// Specifies what content encodings have been applied to the object
	// and thus what decoding mechanisms must be applied to obtain the media-type
	ContentEncoding string
}

// MEncode serializes a GetObjectResponse using msgpack
func (o *GetObjectResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)
	encoder.WriteString("initialChunk")
	if o.InitialChunk == nil {
		encoder.WriteNil()
	} else {
		o.InitialChunk.MEncode(encoder)
	}
	encoder.WriteString("contentLength")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// MDecodeGetObjectResponse deserializes a GetObjectResponse using msgpack
func MDecodeGetObjectResponse(d *msgpack.Decoder) (GetObjectResponse, error) {
	var val GetObjectResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "success":
			val.Success, err = d.ReadBool()
		case "error":
			val.Error, err = d.ReadString()
		case "initialChunk":
			fval, err := MDecodeChunk(d)
			if err != nil {
				return val, err
			}
			val.InitialChunk = &fval
		case "contentLength":
			val.ContentLength, err = d.ReadUint64()
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a GetObjectResponse using cbor
func (o *GetObjectResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)
	encoder.WriteString("initialChunk")
	if o.InitialChunk == nil {
		encoder.WriteNil()
	} else {
		o.InitialChunk.CEncode(encoder)
	}
	encoder.WriteString("contentLength")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// CDecodeGetObjectResponse deserializes a GetObjectResponse using cbor
func CDecodeGetObjectResponse(d *cbor.Decoder) (GetObjectResponse, error) {
	var val GetObjectResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "success":
			val.Success, err = d.ReadBool()
		case "error":
			val.Error, err = d.ReadString()
		case "initialChunk":
			fval, err := CDecodeChunk(d)
			if err != nil {
				return val, err
			}
			val.InitialChunk = &fval
		case "contentLength":
			val.ContentLength, err = d.ReadUint64()
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Result of input item
type ItemResult struct {
	Key string
	// whether the item succeeded or failed
	Success bool
	// optional error message for failures
	Error string
}

// MEncode serializes a ItemResult using msgpack
func (o *ItemResult) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)

	return encoder.CheckError()
}

// MDecodeItemResult deserializes a ItemResult using msgpack
func MDecodeItemResult(d *msgpack.Decoder) (ItemResult, error) {
	var val ItemResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "key":
			val.Key, err = d.ReadString()
		case "success":
			val.Success, err = d.ReadBool()
		case "error":
			val.Error, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ItemResult using cbor
func (o *ItemResult) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)

	return encoder.CheckError()
}

// CDecodeItemResult deserializes a ItemResult using cbor
func CDecodeItemResult(d *cbor.Decoder) (ItemResult, error) {
	var val ItemResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "key":
			val.Key, err = d.ReadString()
		case "success":
			val.Success, err = d.ReadBool()
		case "error":
			val.Error, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Parameter to list_objects.
type ListObjectsRequest struct {
	// Name of the container to search
	ContainerId string
	// Request object names starting with this value. (Optional)
	StartWith string
	// Continuation token passed in ListObjectsResponse.
	// If set, `startWith` is ignored. (Optional)
	Continuation string
	// Last item to return (inclusive terminator) (Optional)
	EndWith string
	// Optionally, stop returning items before returning this value.
	// (exclusive terminator)
	// If startFrom is "a" and endBefore is "b", and items are ordered
	// alphabetically, then only items beginning with "a" would be returned.
	// (Optional)
	EndBefore string
	// maximum number of items to return. If not specified, provider
	// will return an initial set of up to 1000 items. if maxItems > 1000,
	// the provider implementation may return fewer items than requested.
	// (Optional)
	MaxItems uint32
}

// MEncode serializes a ListObjectsRequest using msgpack
func (o *ListObjectsRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("containerId")
	encoder.WriteString(o.ContainerId)
	encoder.WriteString("startWith")
	encoder.WriteString(o.StartWith)
	encoder.WriteString("continuation")
	encoder.WriteString(o.Continuation)
	encoder.WriteString("endWith")
	encoder.WriteString(o.EndWith)
	encoder.WriteString("endBefore")
	encoder.WriteString(o.EndBefore)
	encoder.WriteString("maxItems")
	encoder.WriteUint32(o.MaxItems)

	return encoder.CheckError()
}

// MDecodeListObjectsRequest deserializes a ListObjectsRequest using msgpack
func MDecodeListObjectsRequest(d *msgpack.Decoder) (ListObjectsRequest, error) {
	var val ListObjectsRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = d.ReadString()
		case "startWith":
			val.StartWith, err = d.ReadString()
		case "continuation":
			val.Continuation, err = d.ReadString()
		case "endWith":
			val.EndWith, err = d.ReadString()
		case "endBefore":
			val.EndBefore, err = d.ReadString()
		case "maxItems":
			val.MaxItems, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ListObjectsRequest using cbor
func (o *ListObjectsRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("containerId")
	encoder.WriteString(o.ContainerId)
	encoder.WriteString("startWith")
	encoder.WriteString(o.StartWith)
	encoder.WriteString("continuation")
	encoder.WriteString(o.Continuation)
	encoder.WriteString("endWith")
	encoder.WriteString(o.EndWith)
	encoder.WriteString("endBefore")
	encoder.WriteString(o.EndBefore)
	encoder.WriteString("maxItems")
	encoder.WriteUint32(o.MaxItems)

	return encoder.CheckError()
}

// CDecodeListObjectsRequest deserializes a ListObjectsRequest using cbor
func CDecodeListObjectsRequest(d *cbor.Decoder) (ListObjectsRequest, error) {
	var val ListObjectsRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = d.ReadString()
		case "startWith":
			val.StartWith, err = d.ReadString()
		case "continuation":
			val.Continuation, err = d.ReadString()
		case "endWith":
			val.EndWith, err = d.ReadString()
		case "endBefore":
			val.EndBefore, err = d.ReadString()
		case "maxItems":
			val.MaxItems, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Respose to list_objects.
// If `isLast` is false, the list was truncated by the provider,
// and the remainder of the objects can be requested with another
// request using the `continuation` token.
type ListObjectsResponse struct {
	// set of objects returned
	Objects ObjectsInfo
	// Indicates if the item list is complete, or the last item
	// in a multi-part response.
	IsLast bool
	// If `isLast` is false, this value can be used in the `continuation` field
	// of a `ListObjectsRequest`.
	// Clients should not attempt to interpret this field: it may or may not
	// be a real key or object name, and may be obfuscated by the provider.
	Continuation string
}

// MEncode serializes a ListObjectsResponse using msgpack
func (o *ListObjectsResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("objects")
	o.Objects.MEncode(encoder)
	encoder.WriteString("isLast")
	encoder.WriteBool(o.IsLast)
	encoder.WriteString("continuation")
	encoder.WriteString(o.Continuation)

	return encoder.CheckError()
}

// MDecodeListObjectsResponse deserializes a ListObjectsResponse using msgpack
func MDecodeListObjectsResponse(d *msgpack.Decoder) (ListObjectsResponse, error) {
	var val ListObjectsResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objects":
			val.Objects, err = MDecodeObjectsInfo(d)
		case "isLast":
			val.IsLast, err = d.ReadBool()
		case "continuation":
			val.Continuation, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ListObjectsResponse using cbor
func (o *ListObjectsResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("objects")
	o.Objects.CEncode(encoder)
	encoder.WriteString("isLast")
	encoder.WriteBool(o.IsLast)
	encoder.WriteString("continuation")
	encoder.WriteString(o.Continuation)

	return encoder.CheckError()
}

// CDecodeListObjectsResponse deserializes a ListObjectsResponse using cbor
func CDecodeListObjectsResponse(d *cbor.Decoder) (ListObjectsResponse, error) {
	var val ListObjectsResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objects":
			val.Objects, err = CDecodeObjectsInfo(d)
		case "isLast":
			val.IsLast, err = d.ReadBool()
		case "continuation":
			val.Continuation, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// result for an operation on a list of inputs
type MultiResult []ItemResult

// MEncode serializes a MultiResult using msgpack
func (o *MultiResult) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeMultiResult deserializes a MultiResult using msgpack
func MDecodeMultiResult(d *msgpack.Decoder) (MultiResult, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ItemResult, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ItemResult, 0), err
	}
	val := make([]ItemResult, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeItemResult(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a MultiResult using cbor
func (o *MultiResult) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeMultiResult deserializes a MultiResult using cbor
func CDecodeMultiResult(d *cbor.Decoder) (MultiResult, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ItemResult, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ItemResult, 0), err
	}
	val := make([]ItemResult, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeItemResult(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Name of an object within a container
type ObjectId string

// MEncode serializes a ObjectId using msgpack
func (o *ObjectId) MEncode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// MDecodeObjectId deserializes a ObjectId using msgpack
func MDecodeObjectId(d *msgpack.Decoder) (ObjectId, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ObjectId(val), nil
}

// CEncode serializes a ObjectId using cbor
func (o *ObjectId) CEncode(encoder cbor.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// CDecodeObjectId deserializes a ObjectId using cbor
func CDecodeObjectId(d *cbor.Decoder) (ObjectId, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ObjectId(val), nil
}

// list of object names
type ObjectIds []ObjectId

// MEncode serializes a ObjectIds using msgpack
func (o *ObjectIds) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeObjectIds deserializes a ObjectIds using msgpack
func MDecodeObjectIds(d *msgpack.Decoder) (ObjectIds, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ObjectId, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ObjectId, 0), err
	}
	val := make([]ObjectId, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeObjectId(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ObjectIds using cbor
func (o *ObjectIds) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeObjectIds deserializes a ObjectIds using cbor
func CDecodeObjectIds(d *cbor.Decoder) (ObjectIds, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ObjectId, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ObjectId, 0), err
	}
	val := make([]ObjectId, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeObjectId(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type ObjectMetadata struct {
	// Object identifier that is unique within its container.
	// Naming of objects is determined by the capability provider.
	// An object id could be a path, hash of object contents, or some other unique identifier.
	ObjectId ObjectId
	// container of the object
	ContainerId ContainerId
	// size of the object in bytes
	ContentLength uint64
	// date object was last modified
	LastModified *actor.Timestamp
	// A MIME type of the object
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17
	// Provider implementations _may_ return None for this field for metadata
	// returned from ListObjects
	ContentType string
	// Specifies what content encodings have been applied to the object
	// and thus what decoding mechanisms must be applied to obtain the media-type
	// referenced by the contentType field. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11.
	// Provider implementations _may_ return None for this field for metadata
	// returned from ListObjects
	ContentEncoding string
}

// MEncode serializes a ObjectMetadata using msgpack
func (o *ObjectMetadata) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("objectId")
	o.ObjectId.MEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("contentLength")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("lastModified")
	if o.LastModified == nil {
		encoder.WriteNil()
	} else {
		o.LastModified.MEncode(encoder)
	}
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// MDecodeObjectMetadata deserializes a ObjectMetadata using msgpack
func MDecodeObjectMetadata(d *msgpack.Decoder) (ObjectMetadata, error) {
	var val ObjectMetadata
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = MDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "contentLength":
			val.ContentLength, err = d.ReadUint64()
		case "lastModified":
			fval, err := actor.MDecodeTimestamp(d)
			if err != nil {
				return val, err
			}
			val.LastModified = &fval
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ObjectMetadata using cbor
func (o *ObjectMetadata) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(6)
	encoder.WriteString("objectId")
	o.ObjectId.CEncode(encoder)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("contentLength")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("lastModified")
	if o.LastModified == nil {
		encoder.WriteNil()
	} else {
		o.LastModified.CEncode(encoder)
	}
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// CDecodeObjectMetadata deserializes a ObjectMetadata using cbor
func CDecodeObjectMetadata(d *cbor.Decoder) (ObjectMetadata, error) {
	var val ObjectMetadata
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "objectId":
			val.ObjectId, err = CDecodeObjectId(d)
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "contentLength":
			val.ContentLength, err = d.ReadUint64()
		case "lastModified":
			fval, err := actor.CDecodeTimestamp(d)
			if err != nil {
				return val, err
			}
			val.LastModified = &fval
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// list of object metadata objects
type ObjectsInfo []ObjectMetadata

// MEncode serializes a ObjectsInfo using msgpack
func (o *ObjectsInfo) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeObjectsInfo deserializes a ObjectsInfo using msgpack
func MDecodeObjectsInfo(d *msgpack.Decoder) (ObjectsInfo, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ObjectMetadata, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ObjectMetadata, 0), err
	}
	val := make([]ObjectMetadata, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeObjectMetadata(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ObjectsInfo using cbor
func (o *ObjectsInfo) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeObjectsInfo deserializes a ObjectsInfo using cbor
func CDecodeObjectsInfo(d *cbor.Decoder) (ObjectsInfo, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ObjectMetadata, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ObjectMetadata, 0), err
	}
	val := make([]ObjectMetadata, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeObjectMetadata(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Parameter to PutChunk operation
type PutChunkRequest struct {
	// upload chunk from the file.
	// if chunk.isLast is set, this will be the last chunk uploaded
	Chunk Chunk
	// This value should be set to the `streamId` returned from the initial PutObject.
	StreamId string
	// If set, the receiving provider should cancel the upload process
	// and remove the file.
	CancelAndRemove bool
}

// MEncode serializes a PutChunkRequest using msgpack
func (o *PutChunkRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("chunk")
	o.Chunk.MEncode(encoder)
	encoder.WriteString("streamId")
	encoder.WriteString(o.StreamId)
	encoder.WriteString("cancelAndRemove")
	encoder.WriteBool(o.CancelAndRemove)

	return encoder.CheckError()
}

// MDecodePutChunkRequest deserializes a PutChunkRequest using msgpack
func MDecodePutChunkRequest(d *msgpack.Decoder) (PutChunkRequest, error) {
	var val PutChunkRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "chunk":
			val.Chunk, err = MDecodeChunk(d)
		case "streamId":
			val.StreamId, err = d.ReadString()
		case "cancelAndRemove":
			val.CancelAndRemove, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a PutChunkRequest using cbor
func (o *PutChunkRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("chunk")
	o.Chunk.CEncode(encoder)
	encoder.WriteString("streamId")
	encoder.WriteString(o.StreamId)
	encoder.WriteString("cancelAndRemove")
	encoder.WriteBool(o.CancelAndRemove)

	return encoder.CheckError()
}

// CDecodePutChunkRequest deserializes a PutChunkRequest using cbor
func CDecodePutChunkRequest(d *cbor.Decoder) (PutChunkRequest, error) {
	var val PutChunkRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "chunk":
			val.Chunk, err = CDecodeChunk(d)
		case "streamId":
			val.StreamId, err = d.ReadString()
		case "cancelAndRemove":
			val.CancelAndRemove, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Parameter for PutObject operation
type PutObjectRequest struct {
	// File path and initial data
	Chunk Chunk
	// A MIME type of the object
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17
	ContentType string
	// Specifies what content encodings have been applied to the object
	// and thus what decoding mechanisms must be applied to obtain the media-type
	// referenced by the contentType field. For more information,
	// see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11.
	ContentEncoding string
}

// MEncode serializes a PutObjectRequest using msgpack
func (o *PutObjectRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("chunk")
	o.Chunk.MEncode(encoder)
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// MDecodePutObjectRequest deserializes a PutObjectRequest using msgpack
func MDecodePutObjectRequest(d *msgpack.Decoder) (PutObjectRequest, error) {
	var val PutObjectRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "chunk":
			val.Chunk, err = MDecodeChunk(d)
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a PutObjectRequest using cbor
func (o *PutObjectRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("chunk")
	o.Chunk.CEncode(encoder)
	encoder.WriteString("contentType")
	encoder.WriteString(o.ContentType)
	encoder.WriteString("contentEncoding")
	encoder.WriteString(o.ContentEncoding)

	return encoder.CheckError()
}

// CDecodePutObjectRequest deserializes a PutObjectRequest using cbor
func CDecodePutObjectRequest(d *cbor.Decoder) (PutObjectRequest, error) {
	var val PutObjectRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "chunk":
			val.Chunk, err = CDecodeChunk(d)
		case "contentType":
			val.ContentType, err = d.ReadString()
		case "contentEncoding":
			val.ContentEncoding, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Response to PutObject operation
type PutObjectResponse struct {
	// If this is a multipart upload, `streamId` must be returned
	// with subsequent PutChunk requests
	StreamId string
}

// MEncode serializes a PutObjectResponse using msgpack
func (o *PutObjectResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("streamId")
	encoder.WriteString(o.StreamId)

	return encoder.CheckError()
}

// MDecodePutObjectResponse deserializes a PutObjectResponse using msgpack
func MDecodePutObjectResponse(d *msgpack.Decoder) (PutObjectResponse, error) {
	var val PutObjectResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "streamId":
			val.StreamId, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a PutObjectResponse using cbor
func (o *PutObjectResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("streamId")
	encoder.WriteString(o.StreamId)

	return encoder.CheckError()
}

// CDecodePutObjectResponse deserializes a PutObjectResponse using cbor
func CDecodePutObjectResponse(d *cbor.Decoder) (PutObjectResponse, error) {
	var val PutObjectResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "streamId":
			val.StreamId, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// parameter to removeObjects
type RemoveObjectsRequest struct {
	// name of container
	ContainerId ContainerId
	// list of object names to be removed
	Objects ObjectIds
}

// MEncode serializes a RemoveObjectsRequest using msgpack
func (o *RemoveObjectsRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.MEncode(encoder)
	encoder.WriteString("objects")
	o.Objects.MEncode(encoder)

	return encoder.CheckError()
}

// MDecodeRemoveObjectsRequest deserializes a RemoveObjectsRequest using msgpack
func MDecodeRemoveObjectsRequest(d *msgpack.Decoder) (RemoveObjectsRequest, error) {
	var val RemoveObjectsRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = MDecodeContainerId(d)
		case "objects":
			val.Objects, err = MDecodeObjectIds(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a RemoveObjectsRequest using cbor
func (o *RemoveObjectsRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("containerId")
	o.ContainerId.CEncode(encoder)
	encoder.WriteString("objects")
	o.Objects.CEncode(encoder)

	return encoder.CheckError()
}

// CDecodeRemoveObjectsRequest deserializes a RemoveObjectsRequest using cbor
func CDecodeRemoveObjectsRequest(d *cbor.Decoder) (RemoveObjectsRequest, error) {
	var val RemoveObjectsRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "containerId":
			val.ContainerId, err = CDecodeContainerId(d)
		case "objects":
			val.Objects, err = CDecodeObjectIds(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// The BlobStore service, provider side
type Blobstore interface {
	// Returns whether the container exists
	ContainerExists(ctx *actor.Context, arg ContainerId) (bool, error)
	// Creates a container by name, returning success if it worked
	// Note that container names may not be globally unique - just unique within the
	// "namespace" of the connecting actor and linkdef
	CreateContainer(ctx *actor.Context, arg ContainerId) error
	// Retrieves information about the container.
	// Returns error if the container id is invalid or not found.
	GetContainerInfo(ctx *actor.Context, arg ContainerId) (*ContainerMetadata, error)
	// Returns list of container ids
	ListContainers(ctx *actor.Context) (*ContainersInfo, error)
	// Empty and remove the container(s)
	// The MultiResult list contains one entry for each container
	// that was not successfully removed, with the 'key' value representing the container name.
	// If the MultiResult list is empty, all container removals succeeded.
	RemoveContainers(ctx *actor.Context, arg ContainerIds) (*MultiResult, error)
	// Returns whether the object exists
	ObjectExists(ctx *actor.Context, arg ContainerObject) (bool, error)
	// Retrieves information about the object.
	// Returns error if the object id is invalid or not found.
	GetObjectInfo(ctx *actor.Context, arg ContainerObject) (*ObjectMetadata, error)
	// Lists the objects in the container.
	// If the container exists and is empty, the returned `objects` list is empty.
	// Parameters of the request may be used to limit the object names returned
	// with an optional start value, end value, and maximum number of items.
	// The provider may limit the number of items returned. If the list is truncated,
	// the response contains a `continuation` token that may be submitted in
	// a subsequent ListObjects request.
	//
	// Optional object metadata fields (i.e., `contentType` and `contentEncoding`) may not be
	// filled in for ListObjects response. To get complete object metadata, use GetObjectInfo.
	ListObjects(ctx *actor.Context, arg ListObjectsRequest) (*ListObjectsResponse, error)
	// Removes the objects. In the event any of the objects cannot be removed,
	// the operation continues until all requested deletions have been attempted.
	// The MultiRequest includes a list of errors, one for each deletion request
	// that did not succeed. If the list is empty, all removals succeeded.
	RemoveObjects(ctx *actor.Context, arg RemoveObjectsRequest) (*MultiResult, error)
	// Requests to start upload of a file/blob to the Blobstore.
	// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
	PutObject(ctx *actor.Context, arg PutObjectRequest) (*PutObjectResponse, error)
	// Requests to retrieve an object. If the object is large, the provider
	// may split the response into multiple parts
	// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
	GetObject(ctx *actor.Context, arg GetObjectRequest) (*GetObjectResponse, error)
	// Uploads a file chunk to a blobstore. This must be called AFTER PutObject
	// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
	PutChunk(ctx *actor.Context, arg PutChunkRequest) error
}

// BlobstoreHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func BlobstoreHandler(actor_ Blobstore) actor.Handler {
	return actor.NewHandler("Blobstore", &BlobstoreReceiver{}, actor_)
}

// BlobstoreContractId returns the capability contract id for this interface
func BlobstoreContractId() string { return "wasmcloud:blobstore" }

// BlobstoreReceiver receives messages defined in the Blobstore service interface
// The BlobStore service, provider side
type BlobstoreReceiver struct{}

func (r *BlobstoreReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Blobstore)
	switch message.Method {

	case "ContainerExists":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerId(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ContainerExists(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "Blobstore.ContainerExists", Arg: buf}, nil
		}
	case "CreateContainer":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerId(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.CreateContainer(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Blobstore.CreateContainer", Arg: buf}, nil
		}
	case "GetContainerInfo":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerId(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.GetContainerInfo(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.GetContainerInfo", Arg: buf}, nil
		}
	case "ListContainers":
		{
			resp, err := svc_.ListContainers(ctx)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.ListContainers", Arg: buf}, nil
		}
	case "RemoveContainers":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerIds(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.RemoveContainers(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.RemoveContainers", Arg: buf}, nil
		}
	case "ObjectExists":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerObject(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ObjectExists(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "Blobstore.ObjectExists", Arg: buf}, nil
		}
	case "GetObjectInfo":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeContainerObject(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.GetObjectInfo(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.GetObjectInfo", Arg: buf}, nil
		}
	case "ListObjects":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeListObjectsRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ListObjects(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.ListObjects", Arg: buf}, nil
		}
	case "RemoveObjects":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeRemoveObjectsRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.RemoveObjects(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.RemoveObjects", Arg: buf}, nil
		}
	case "PutObject":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodePutObjectRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.PutObject(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.PutObject", Arg: buf}, nil
		}
	case "GetObject":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeGetObjectRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.GetObject(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "Blobstore.GetObject", Arg: buf}, nil
		}
	case "PutChunk":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodePutChunkRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.PutChunk(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Blobstore.PutChunk", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Blobstore."+message.Method)
	}
}

// BlobstoreSender sends messages to a Blobstore service
// The BlobStore service, provider side
type BlobstoreSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a Blobstore provider
// implementing the 'wasmcloud:blobstore' capability contract, with the "default" link
func NewProviderBlobstore() *BlobstoreSender {
	transport := actor.ToProvider("wasmcloud:blobstore", "default")
	return &BlobstoreSender{transport: transport}
}

// NewProviderBlobstoreLink constructs a client for sending to a Blobstore provider
// implementing the 'wasmcloud:blobstore' capability contract, with the specified link name
func NewProviderBlobstoreLink(linkName string) *BlobstoreSender {
	transport := actor.ToProvider("wasmcloud:blobstore", linkName)
	return &BlobstoreSender{transport: transport}
}

// Returns whether the container exists
func (s *BlobstoreSender) ContainerExists(ctx *actor.Context, arg ContainerId) (bool, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.ContainerExists", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Creates a container by name, returning success if it worked
// Note that container names may not be globally unique - just unique within the
// "namespace" of the connecting actor and linkdef
func (s *BlobstoreSender) CreateContainer(ctx *actor.Context, arg ContainerId) error {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Blobstore.CreateContainer", Arg: buf})
	return nil
}

// Retrieves information about the container.
// Returns error if the container id is invalid or not found.
func (s *BlobstoreSender) GetContainerInfo(ctx *actor.Context, arg ContainerId) (*ContainerMetadata, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.GetContainerInfo", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeContainerMetadata(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Returns list of container ids
func (s *BlobstoreSender) ListContainers(ctx *actor.Context) (*ContainersInfo, error) {
	buf := make([]byte, 0)
	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.ListContainers", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeContainersInfo(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Empty and remove the container(s)
// The MultiResult list contains one entry for each container
// that was not successfully removed, with the 'key' value representing the container name.
// If the MultiResult list is empty, all container removals succeeded.
func (s *BlobstoreSender) RemoveContainers(ctx *actor.Context, arg ContainerIds) (*MultiResult, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.RemoveContainers", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeMultiResult(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Returns whether the object exists
func (s *BlobstoreSender) ObjectExists(ctx *actor.Context, arg ContainerObject) (bool, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.ObjectExists", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Retrieves information about the object.
// Returns error if the object id is invalid or not found.
func (s *BlobstoreSender) GetObjectInfo(ctx *actor.Context, arg ContainerObject) (*ObjectMetadata, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.GetObjectInfo", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeObjectMetadata(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Lists the objects in the container.
// If the container exists and is empty, the returned `objects` list is empty.
// Parameters of the request may be used to limit the object names returned
// with an optional start value, end value, and maximum number of items.
// The provider may limit the number of items returned. If the list is truncated,
// the response contains a `continuation` token that may be submitted in
// a subsequent ListObjects request.
//
// Optional object metadata fields (i.e., `contentType` and `contentEncoding`) may not be
// filled in for ListObjects response. To get complete object metadata, use GetObjectInfo.
func (s *BlobstoreSender) ListObjects(ctx *actor.Context, arg ListObjectsRequest) (*ListObjectsResponse, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.ListObjects", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeListObjectsResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Removes the objects. In the event any of the objects cannot be removed,
// the operation continues until all requested deletions have been attempted.
// The MultiRequest includes a list of errors, one for each deletion request
// that did not succeed. If the list is empty, all removals succeeded.
func (s *BlobstoreSender) RemoveObjects(ctx *actor.Context, arg RemoveObjectsRequest) (*MultiResult, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.RemoveObjects", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeMultiResult(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Requests to start upload of a file/blob to the Blobstore.
// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
func (s *BlobstoreSender) PutObject(ctx *actor.Context, arg PutObjectRequest) (*PutObjectResponse, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.PutObject", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodePutObjectResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Requests to retrieve an object. If the object is large, the provider
// may split the response into multiple parts
// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
func (s *BlobstoreSender) GetObject(ctx *actor.Context, arg GetObjectRequest) (*GetObjectResponse, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Blobstore.GetObject", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeGetObjectResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Uploads a file chunk to a blobstore. This must be called AFTER PutObject
// It is recommended to keep chunks under 1MB to avoid exceeding nats default message size
func (s *BlobstoreSender) PutChunk(ctx *actor.Context, arg PutChunkRequest) error {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Blobstore.PutChunk", Arg: buf})
	return nil
}

// The BlobStore service, actor side
type ChunkReceiver interface {
	// Receives a file chunk from a blobstore.
	// A blobstore provider invokes this operation on actors in response to the GetObject request.
	// If the response sets cancelDownload, the provider will stop downloading chunks
	ReceiveChunk(ctx *actor.Context, arg Chunk) (*ChunkResponse, error)
}

// ChunkReceiverHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func ChunkReceiverHandler(actor_ ChunkReceiver) actor.Handler {
	return actor.NewHandler("ChunkReceiver", &ChunkReceiverReceiver{}, actor_)
}

// ChunkReceiverContractId returns the capability contract id for this interface
func ChunkReceiverContractId() string { return "wasmcloud:blobstore" }

// ChunkReceiverReceiver receives messages defined in the ChunkReceiver service interface
// The BlobStore service, actor side
type ChunkReceiverReceiver struct{}

func (r *ChunkReceiverReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(ChunkReceiver)
	switch message.Method {

	case "ReceiveChunk":
		{

			d := cbor.NewDecoder(message.Arg)
			value, err_ := CDecodeChunk(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ReceiveChunk(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer cbor.Sizer
			size_enc := &sizer
			resp.CEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := cbor.NewEncoder(buf)
			enc := &encoder
			resp.CEncode(enc)
			return &actor.Message{Method: "ChunkReceiver.ReceiveChunk", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "ChunkReceiver."+message.Method)
	}
}

// ChunkReceiverSender sends messages to a ChunkReceiver service
// The BlobStore service, actor side
type ChunkReceiverSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorChunkReceiverSender(actor_id string) *ChunkReceiverSender {
	transport := actor.ToActor(actor_id)
	return &ChunkReceiverSender{transport: transport}
}

// Receives a file chunk from a blobstore.
// A blobstore provider invokes this operation on actors in response to the GetObject request.
// If the response sets cancelDownload, the provider will stop downloading chunks
func (s *ChunkReceiverSender) ReceiveChunk(ctx *actor.Context, arg Chunk) (*ChunkResponse, error) {

	var sizer cbor.Sizer
	size_enc := &sizer
	arg.CEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = cbor.NewEncoder(buf)
	enc := &encoder
	arg.CEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "ChunkReceiver.ReceiveChunk", Arg: buf})
	d := cbor.NewDecoder(out_buf)
	resp, err_ := CDecodeChunkResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.5
