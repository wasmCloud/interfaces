// Keyvalue: wasmcloud capability contract for key-value store
package keyvalue

import (
	"github.com/wasmcloud/actor-tinygo"   //nolint
	"github.com/wasmcloud/tinygo-msgpack" //nolint
)

// Response to get request
type GetResponse struct {
	// the value, if it existed
	Value string
	// whether or not the value existed
	Exists bool
}

// Encode serializes a GetResponse using msgpack
func (o *GetResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)
	encoder.WriteString("Exists")
	encoder.WriteBool(o.Exists)

	return nil
}

// Decode deserializes a GetResponse using msgpack
func DecodeGetResponse(d msgpack.Decoder) (GetResponse, error) {
	var val GetResponse
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "Value":
			val.Value, err = d.ReadString()
		case "Exists":
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

// Encode serializes a IncrementRequest using msgpack
func (o *IncrementRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Key")
	encoder.WriteString(o.Key)
	encoder.WriteString("Value")
	encoder.WriteInt32(o.Value)

	return nil
}

// Decode deserializes a IncrementRequest using msgpack
func DecodeIncrementRequest(d msgpack.Decoder) (IncrementRequest, error) {
	var val IncrementRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "Key":
			val.Key, err = d.ReadString()
		case "Value":
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

// Encode serializes a ListAddRequest using msgpack
func (o *ListAddRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("ListName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)

	return nil
}

// Decode deserializes a ListAddRequest using msgpack
func DecodeListAddRequest(d msgpack.Decoder) (ListAddRequest, error) {
	var val ListAddRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "ListName":
			val.ListName, err = d.ReadString()
		case "Value":
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

// Encode serializes a ListDelRequest using msgpack
func (o *ListDelRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("ListName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)

	return nil
}

// Decode deserializes a ListDelRequest using msgpack
func DecodeListDelRequest(d msgpack.Decoder) (ListDelRequest, error) {
	var val ListDelRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "ListName":
			val.ListName, err = d.ReadString()
		case "Value":
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

// Encode serializes a ListRangeRequest using msgpack
func (o *ListRangeRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ListName")
	encoder.WriteString(o.ListName)
	encoder.WriteString("Start")
	encoder.WriteInt32(o.Start)
	encoder.WriteString("Stop")
	encoder.WriteInt32(o.Stop)

	return nil
}

// Decode deserializes a ListRangeRequest using msgpack
func DecodeListRangeRequest(d msgpack.Decoder) (ListRangeRequest, error) {
	var val ListRangeRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "ListName":
			val.ListName, err = d.ReadString()
		case "Start":
			val.Start, err = d.ReadInt32()
		case "Stop":
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

// Encode serializes a SetAddRequest using msgpack
func (o *SetAddRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("SetName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)

	return nil
}

// Decode deserializes a SetAddRequest using msgpack
func DecodeSetAddRequest(d msgpack.Decoder) (SetAddRequest, error) {
	var val SetAddRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "SetName":
			val.SetName, err = d.ReadString()
		case "Value":
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

// Encode serializes a SetDelRequest using msgpack
func (o *SetDelRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("SetName")
	encoder.WriteString(o.SetName)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)

	return nil
}

// Decode deserializes a SetDelRequest using msgpack
func DecodeSetDelRequest(d msgpack.Decoder) (SetDelRequest, error) {
	var val SetDelRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "SetName":
			val.SetName, err = d.ReadString()
		case "Value":
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

// Encode serializes a SetRequest using msgpack
func (o *SetRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("Key")
	encoder.WriteString(o.Key)
	encoder.WriteString("Value")
	encoder.WriteString(o.Value)
	encoder.WriteString("Expires")
	encoder.WriteUint32(o.Expires)

	return nil
}

// Decode deserializes a SetRequest using msgpack
func DecodeSetRequest(d msgpack.Decoder) (SetRequest, error) {
	var val SetRequest
	isNil, err := d.IsNextNil()
	if err != nil {
		return val, err
	}
	if isNil {
		return val, nil
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
		case "Key":
			val.Key, err = d.ReadString()
		case "Value":
			val.Value, err = d.ReadString()
		case "Expires":
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

// Encode serializes a StringList using msgpack
func (o *StringList) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteString(item_o)
	}

	return nil
}

// Decode deserializes a StringList using msgpack
func DecodeStringList(d msgpack.Decoder) (StringList, error) {
	isNil, err := d.IsNextNil()
	if err == nil && isNil {
		d.Skip()
		return make([]string, 0), nil
	}
	size, err := d.ReadArraySize()
	if err != nil {
		size = 0
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
func KeyValueHandler() actor.Handler {
	return actor.NewHandler("KeyValue", KeyValueReceiver{})
}

// KeyValueReceiver receives messages defined in the KeyValue service interface
type KeyValueReceiver struct{}

func (r *KeyValueReceiver) dispatch(ctx *actor.Context, svc KeyValue, message *actor.Message) (*actor.Message, error) {
	switch message.Method {
	case "Increment":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeIncrementRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.Increment(ctx, value)
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

			resp, err := svc.Contains(ctx, value)
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

			resp, err := svc.Del(ctx, value)
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

			resp, err := svc.Get(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)
			return &actor.Message{Method: "KeyValue.Get", Arg: buf}, nil
		}
	case "ListAdd":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeListAddRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.ListAdd(ctx, value)
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

			resp, err := svc.ListClear(ctx, value)
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
			value, err_ := DecodeListDelRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.ListDel(ctx, value)
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
			value, err_ := DecodeListRangeRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.ListRange(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)
			return &actor.Message{Method: "KeyValue.ListRange", Arg: buf}, nil
		}
	case "Set":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeSetRequest(d)
			if err_ != nil {
				return nil, err_
			}

			err := svc.Set(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "KeyValue.Set", Arg: buf}, nil
		}
	case "SetAdd":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeSetAddRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetAdd(ctx, value)
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
			value, err_ := DecodeSetDelRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetDel(ctx, value)
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
			value, err_ := DecodeStringList(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetIntersection(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)
			return &actor.Message{Method: "KeyValue.SetIntersection", Arg: buf}, nil
		}
	case "SetQuery":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetQuery(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)
			return &actor.Message{Method: "KeyValue.SetQuery", Arg: buf}, nil
		}
	case "SetUnion":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeStringList(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetUnion(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.Encode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.Encode(enc)
			return &actor.Message{Method: "KeyValue.SetUnion", Arg: buf}, nil
		}
	case "SetClear":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadString()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.SetClear(ctx, value)
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

// Increments a numeric value, returning the new value
func (s *KeyValueSender) Increment(ctx *actor.Context, arg IncrementRequest) (int32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
	resp, err_ := DecodeGetResponse(d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Append a value onto the end of a list. Returns the new list size
func (s *KeyValueSender) ListAdd(ctx *actor.Context, arg ListAddRequest) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.ListRange", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeStringList(d)
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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	s.transport.Send(ctx, actor.Message{Method: "KeyValue.Set", Arg: buf})
	return nil
}

// Add an item into a set. Returns number of items added (1 or 0)
func (s *KeyValueSender) SetAdd(ctx *actor.Context, arg SetAddRequest) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetIntersection", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeStringList(d)
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
	resp, err_ := DecodeStringList(d)
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
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "KeyValue.SetUnion", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeStringList(d)
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

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
