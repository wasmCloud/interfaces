// Keyvalue: wasmcloud capability contract for key-value store
package keyvalue

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// Response to get request
type GetResponse struct {
	// the value, if it existed
	Value string
	// whether or not the value existed
	Exists bool
}

// MEncode serializes a GetResponse using msgpack
func (o *GetResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)
	encoder.WriteString("exists")
	encoder.WriteBool(o.Exists)

	return encoder.CheckError()
}

// MDecodeGetResponse deserializes a GetResponse using msgpack
func MDecodeGetResponse(d *msgpack.Decoder) (GetResponse, error) {
	var val GetResponse
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
		case "value":
			val.Value, err = d.ReadString()
		case "exists":
			val.Exists, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a GetResponse using cbor
func (o *GetResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)
	encoder.WriteString("exists")
	encoder.WriteBool(o.Exists)

	return encoder.CheckError()
}

// CDecodeGetResponse deserializes a GetResponse using cbor
func CDecodeGetResponse(d *cbor.Decoder) (GetResponse, error) {
	var val GetResponse
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
		case "value":
			val.Value, err = d.ReadString()
		case "exists":
			val.Exists, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type IncrementRequest struct {
	// name of value to increment
	Key string
	// amount to add to value
	Value int32
}

// MEncode serializes a IncrementRequest using msgpack
func (o *IncrementRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("value")
	encoder.WriteInt32(o.Value)

	return encoder.CheckError()
}

// MDecodeIncrementRequest deserializes a IncrementRequest using msgpack
func MDecodeIncrementRequest(d *msgpack.Decoder) (IncrementRequest, error) {
	var val IncrementRequest
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
		case "value":
			val.Value, err = d.ReadInt32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a IncrementRequest using cbor
func (o *IncrementRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("value")
	encoder.WriteInt32(o.Value)

	return encoder.CheckError()
}

// CDecodeIncrementRequest deserializes a IncrementRequest using cbor
func CDecodeIncrementRequest(d *cbor.Decoder) (IncrementRequest, error) {
	var val IncrementRequest
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
		case "value":
			val.Value, err = d.ReadInt32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Parameter to ListAdd operation
type ListAddRequest struct {
	// name of the list to modify
	ListName string
	// value to append to the list
	Value string
}

// MEncode serializes a ListAddRequest using msgpack
func (o *ListAddRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// MDecodeListAddRequest deserializes a ListAddRequest using msgpack
func MDecodeListAddRequest(d *msgpack.Decoder) (ListAddRequest, error) {
	var val ListAddRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ListAddRequest using cbor
func (o *ListAddRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// CDecodeListAddRequest deserializes a ListAddRequest using cbor
func CDecodeListAddRequest(d *cbor.Decoder) (ListAddRequest, error) {
	var val ListAddRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Removes an item from the list. If the item occurred more than once,
// removes only the first item.
// Returns true if the item was found.
type ListDelRequest struct {
	// name of list to modify
	ListName string
	Value    string
}

// MEncode serializes a ListDelRequest using msgpack
func (o *ListDelRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// MDecodeListDelRequest deserializes a ListDelRequest using msgpack
func MDecodeListDelRequest(d *msgpack.Decoder) (ListDelRequest, error) {
	var val ListDelRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ListDelRequest using cbor
func (o *ListDelRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// CDecodeListDelRequest deserializes a ListDelRequest using cbor
func CDecodeListDelRequest(d *cbor.Decoder) (ListDelRequest, error) {
	var val ListDelRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type ListRangeRequest struct {
	// name of list
	ListName string
	// start index of the range, 0-based, inclusive.
	Start int32
	// end index of the range, 0-based, inclusive.
	Stop int32
}

// MEncode serializes a ListRangeRequest using msgpack
func (o *ListRangeRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("start")
	encoder.WriteInt32(o.Start)
	encoder.WriteString("stop")
	encoder.WriteInt32(o.Stop)

	return encoder.CheckError()
}

// MDecodeListRangeRequest deserializes a ListRangeRequest using msgpack
func MDecodeListRangeRequest(d *msgpack.Decoder) (ListRangeRequest, error) {
	var val ListRangeRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "start":
			val.Start, err = d.ReadInt32()
		case "stop":
			val.Stop, err = d.ReadInt32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ListRangeRequest using cbor
func (o *ListRangeRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("listName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("start")
	encoder.WriteInt32(o.Start)
	encoder.WriteString("stop")
	encoder.WriteInt32(o.Stop)

	return encoder.CheckError()
}

// CDecodeListRangeRequest deserializes a ListRangeRequest using cbor
func CDecodeListRangeRequest(d *cbor.Decoder) (ListRangeRequest, error) {
	var val ListRangeRequest
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
		case "listName":
			val.ListName, err = d.ReadString()
		case "start":
			val.Start, err = d.ReadInt32()
		case "stop":
			val.Stop, err = d.ReadInt32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type SetAddRequest struct {
	// name of the set
	SetName string
	// value to add to the set
	Value string
}

// MEncode serializes a SetAddRequest using msgpack
func (o *SetAddRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("setName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// MDecodeSetAddRequest deserializes a SetAddRequest using msgpack
func MDecodeSetAddRequest(d *msgpack.Decoder) (SetAddRequest, error) {
	var val SetAddRequest
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
		case "setName":
			val.SetName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SetAddRequest using cbor
func (o *SetAddRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("setName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// CDecodeSetAddRequest deserializes a SetAddRequest using cbor
func CDecodeSetAddRequest(d *cbor.Decoder) (SetAddRequest, error) {
	var val SetAddRequest
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
		case "setName":
			val.SetName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type SetDelRequest struct {
	SetName string
	Value   string
}

// MEncode serializes a SetDelRequest using msgpack
func (o *SetDelRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("setName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// MDecodeSetDelRequest deserializes a SetDelRequest using msgpack
func MDecodeSetDelRequest(d *msgpack.Decoder) (SetDelRequest, error) {
	var val SetDelRequest
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
		case "setName":
			val.SetName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SetDelRequest using cbor
func (o *SetDelRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("setName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)

	return encoder.CheckError()
}

// CDecodeSetDelRequest deserializes a SetDelRequest using cbor
func CDecodeSetDelRequest(d *cbor.Decoder) (SetDelRequest, error) {
	var val SetDelRequest
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
		case "setName":
			val.SetName, err = d.ReadString()
		case "value":
			val.Value, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type SetRequest struct {
	// the key name to change (or create)
	Key string
	// the new value
	Value string
	// expiration time in seconds 0 for no expiration
	Expires uint32
}

// MEncode serializes a SetRequest using msgpack
func (o *SetRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)
	encoder.WriteString("expires")
	encoder.WriteUint32(o.Expires)

	return encoder.CheckError()
}

// MDecodeSetRequest deserializes a SetRequest using msgpack
func MDecodeSetRequest(d *msgpack.Decoder) (SetRequest, error) {
	var val SetRequest
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
		case "value":
			val.Value, err = d.ReadString()
		case "expires":
			val.Expires, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SetRequest using cbor
func (o *SetRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("key")
	encoder.WriteString(o.Key)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)
	encoder.WriteString("expires")
	encoder.WriteUint32(o.Expires)

	return encoder.CheckError()
}

// CDecodeSetRequest deserializes a SetRequest using cbor
func CDecodeSetRequest(d *cbor.Decoder) (SetRequest, error) {
	var val SetRequest
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
		case "value":
			val.Value, err = d.ReadString()
		case "expires":
			val.Expires, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// list of strings
type StringList []string

// MEncode serializes a StringList using msgpack
func (o *StringList) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteString(item_o)
	}

	return encoder.CheckError()
}

// MDecodeStringList deserializes a StringList using msgpack
func MDecodeStringList(d *msgpack.Decoder) (StringList, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]string, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]string, 0), err
	}
	val := make([]string, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a StringList using cbor
func (o *StringList) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteString(item_o)
	}

	return encoder.CheckError()
}

// CDecodeStringList deserializes a StringList using cbor
func CDecodeStringList(d *cbor.Decoder) (StringList, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]string, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]string, 0), err
	}
	val := make([]string, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type KeyValue interface {
	// Increments a numeric value, returning the new value
	Increment(ctx *actor.Context, arg IncrementRequest) (int32, error)
	// returns whether the store contains the key
	Contains(ctx *actor.Context, arg string) (bool, error)
	// Deletes a key, returning true if the key was deleted
	Del(ctx *actor.Context, arg string) (bool, error)
	// Gets a value for a specified key. If the key exists,
	// the return structure contains exists: true and the value,
	// otherwise the return structure contains exists == false.
	Get(ctx *actor.Context, arg string) (*GetResponse, error)
	// Append a value onto the end of a list. Returns the new list size
	ListAdd(ctx *actor.Context, arg ListAddRequest) (uint32, error)
	// Deletes a list and its contents
	// input: list name
	// output: true if the list existed and was deleted
	ListClear(ctx *actor.Context, arg string) (bool, error)
	// Deletes a value from a list. Returns true if the item was removed.
	ListDel(ctx *actor.Context, arg ListDelRequest) (bool, error)
	// Retrieves a range of values from a list using 0-based indices.
	// Start and end values are inclusive, for example, (0,10) returns
	// 11 items if the list contains at least 11 items. If the stop value
	// is beyond the end of the list, it is treated as the end of the list.
	ListRange(ctx *actor.Context, arg ListRangeRequest) (*StringList, error)
	// Sets the value of a key.
	// expires is an optional number of seconds before the value should be automatically deleted,
	// or 0 for no expiration.
	Set(ctx *actor.Context, arg SetRequest) error
	// Add an item into a set. Returns number of items added (1 or 0)
	SetAdd(ctx *actor.Context, arg SetAddRequest) (uint32, error)
	// Deletes an item from the set. Returns number of items removed from the set (1 or 0)
	SetDel(ctx *actor.Context, arg SetDelRequest) (uint32, error)
	// perform intersection of sets and returns values from the intersection.
	// input: list of sets for performing intersection (at least two)
	// output: values
	SetIntersection(ctx *actor.Context, arg StringList) (*StringList, error)
	// Retrieves all items from a set
	// input: String
	// output: set members
	SetQuery(ctx *actor.Context, arg string) (*StringList, error)
	// perform union of sets and returns values from the union
	// input: list of sets for performing union (at least two)
	// output: union of values
	SetUnion(ctx *actor.Context, arg StringList) (*StringList, error)
	// clears all values from the set and removes it
	// input: set name
	// output: true if the set existed and was deleted
	SetClear(ctx *actor.Context, arg string) (bool, error)
}

// KeyValueHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func KeyValueHandler(actor_ KeyValue) actor.Handler {
	return actor.NewHandler("KeyValue", &KeyValueReceiver{}, actor_)
}

// KeyValueContractId returns the capability contract id for this interface
func KeyValueContractId() string { return "wasmcloud:keyvalue" }

// KeyValueReceiver receives messages defined in the KeyValue service interface
type KeyValueReceiver struct{}

func (r *KeyValueReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(KeyValue)
	switch message.Method {

	case "Increment":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeIncrementRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Increment(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteInt32(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteInt32(resp)
			return &actor.Message{Method: "KeyValue.Increment", Arg: buf}, nil
		}
	case "Contains":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Contains(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "KeyValue.Contains", Arg: buf}, nil
		}
	case "Del":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Del(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "KeyValue.Del", Arg: buf}, nil
		}
	case "Get":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Get(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "KeyValue.Get", Arg: buf}, nil
		}
	case "ListAdd":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeListAddRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ListAdd(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteUint32(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteUint32(resp)
			return &actor.Message{Method: "KeyValue.ListAdd", Arg: buf}, nil
		}
	case "ListClear":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ListClear(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "KeyValue.ListClear", Arg: buf}, nil
		}
	case "ListDel":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeListDelRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ListDel(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "KeyValue.ListDel", Arg: buf}, nil
		}
	case "ListRange":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeListRangeRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.ListRange(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "KeyValue.ListRange", Arg: buf}, nil
		}
	case "Set":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeSetRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.Set(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "KeyValue.Set", Arg: buf}, nil
		}
	case "SetAdd":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeSetAddRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetAdd(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteUint32(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteUint32(resp)
			return &actor.Message{Method: "KeyValue.SetAdd", Arg: buf}, nil
		}
	case "SetDel":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeSetDelRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetDel(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteUint32(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteUint32(resp)
			return &actor.Message{Method: "KeyValue.SetDel", Arg: buf}, nil
		}
	case "SetIntersection":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeStringList(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetIntersection(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "KeyValue.SetIntersection", Arg: buf}, nil
		}
	case "SetQuery":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetQuery(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "KeyValue.SetQuery", Arg: buf}, nil
		}
	case "SetUnion":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeStringList(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetUnion(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "KeyValue.SetUnion", Arg: buf}, nil
		}
	case "SetClear":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.SetClear(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteBool(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteBool(resp)
			return &actor.Message{Method: "KeyValue.SetClear", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "KeyValue."+message.Method)
	}
}

// KeyValueSender sends messages to a KeyValue service
type KeyValueSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a KeyValue provider
// implementing the 'wasmcloud:keyvalue' capability contract, with the "default" link
func NewProviderKeyValue() *KeyValueSender {
	transport := actor.ToProvider("wasmcloud:keyvalue", "default")
	return &KeyValueSender{transport: transport}
}

// NewProviderKeyValueLink constructs a client for sending to a KeyValue provider
// implementing the 'wasmcloud:keyvalue' capability contract, with the specified link name
func NewProviderKeyValueLink(linkName string) *KeyValueSender {
	transport := actor.ToProvider("wasmcloud:keyvalue", linkName)
	return &KeyValueSender{transport: transport}
}

// Increments a numeric value, returning the new value
func (s *KeyValueSender) Increment(ctx *actor.Context, arg IncrementRequest) (int32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.Increment", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadInt32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// returns whether the store contains the key
func (s *KeyValueSender) Contains(ctx *actor.Context, arg string) (bool, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.Contains", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Deletes a key, returning true if the key was deleted
func (s *KeyValueSender) Del(ctx *actor.Context, arg string) (bool, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.Del", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Gets a value for a specified key. If the key exists,
// the return structure contains exists: true and the value,
// otherwise the return structure contains exists == false.
func (s *KeyValueSender) Get(ctx *actor.Context, arg string) (*GetResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.Get", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeGetResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Append a value onto the end of a list. Returns the new list size
func (s *KeyValueSender) ListAdd(ctx *actor.Context, arg ListAddRequest) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.ListAdd", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// Deletes a list and its contents
// input: list name
// output: true if the list existed and was deleted
func (s *KeyValueSender) ListClear(ctx *actor.Context, arg string) (bool, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.ListClear", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Deletes a value from a list. Returns true if the item was removed.
func (s *KeyValueSender) ListDel(ctx *actor.Context, arg ListDelRequest) (bool, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.ListDel", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// Retrieves a range of values from a list using 0-based indices.
// Start and end values are inclusive, for example, (0,10) returns
// 11 items if the list contains at least 11 items. If the stop value
// is beyond the end of the list, it is treated as the end of the list.
func (s *KeyValueSender) ListRange(ctx *actor.Context, arg ListRangeRequest) (*StringList, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.ListRange", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeStringList(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Sets the value of a key.
// expires is an optional number of seconds before the value should be automatically deleted,
// or 0 for no expiration.
func (s *KeyValueSender) Set(ctx *actor.Context, arg SetRequest) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "KeyValue.Set", Arg: buf})
	return nil
}

// Add an item into a set. Returns number of items added (1 or 0)
func (s *KeyValueSender) SetAdd(ctx *actor.Context, arg SetAddRequest) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetAdd", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// Deletes an item from the set. Returns number of items removed from the set (1 or 0)
func (s *KeyValueSender) SetDel(ctx *actor.Context, arg SetDelRequest) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetDel", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// perform intersection of sets and returns values from the intersection.
// input: list of sets for performing intersection (at least two)
// output: values
func (s *KeyValueSender) SetIntersection(ctx *actor.Context, arg StringList) (*StringList, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetIntersection", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeStringList(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Retrieves all items from a set
// input: String
// output: set members
func (s *KeyValueSender) SetQuery(ctx *actor.Context, arg string) (*StringList, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetQuery", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeStringList(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// perform union of sets and returns values from the union
// input: list of sets for performing union (at least two)
// output: union of values
func (s *KeyValueSender) SetUnion(ctx *actor.Context, arg StringList) (*StringList, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetUnion", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeStringList(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// clears all values from the set and removes it
// input: set name
// output: true if the set existed and was deleted
func (s *KeyValueSender) SetClear(ctx *actor.Context, arg string) (bool, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteString(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteString(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetClear", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadBool()
	if err_ != nil {
		return false, err_
	}
	return resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.5
