// wasmcloud platform core actor
package actor

import (
	core "github.com/wasmcloud/interfaces/core/tinygo" //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"            //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack"      //nolint
)

type Context struct {
}

// MEncode serializes a Context using msgpack
func (o *Context) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// MDecodeContext deserializes a Context using msgpack
func MDecodeContext(d *msgpack.Decoder) (Context, error) {
	var val Context
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
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Context using cbor
func (o *Context) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// CDecodeContext deserializes a Context using cbor
func CDecodeContext(d *cbor.Decoder) (Context, error) {
	var val Context
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
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Document struct {
}

// MEncode serializes a Document using msgpack
func (o *Document) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// MDecodeDocument deserializes a Document using msgpack
func MDecodeDocument(d *msgpack.Decoder) (Document, error) {
	var val Document
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
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Document using cbor
func (o *Document) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// CDecodeDocument deserializes a Document using cbor
func CDecodeDocument(d *cbor.Decoder) (Document, error) {
	var val Document
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
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Message struct {
	Arg    []byte `json:"Arg"`
	Method string `json:"Method"`
}

// MEncode serializes a Message using msgpack
func (o *Message) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Arg")
	encoder.WriteByteArray(o.Arg)
	encoder.WriteString("Method")
	encoder.WriteString(o.Method)

	return encoder.CheckError()
}

// MDecodeMessage deserializes a Message using msgpack
func MDecodeMessage(d *msgpack.Decoder) (Message, error) {
	var val Message
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
		case "Arg":
			val.Arg, err = d.ReadByteArray()
		case "Method":
			val.Method, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Message using cbor
func (o *Message) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Arg")
	encoder.WriteByteArray(o.Arg)
	encoder.WriteString("Method")
	encoder.WriteString(o.Method)

	return encoder.CheckError()
}

// CDecodeMessage deserializes a Message using cbor
func CDecodeMessage(d *cbor.Decoder) (Message, error) {
	var val Message
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
		case "Arg":
			val.Arg, err = d.ReadByteArray()
		case "Method":
			val.Method, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Timestamp struct {
	Nsec uint32 `json:"Nsec"`
	Sec  int64  `json:"Sec"`
}

// MEncode serializes a Timestamp using msgpack
func (o *Timestamp) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Nsec")
	encoder.WriteUint32(o.Nsec)
	encoder.WriteString("Sec")
	encoder.WriteInt64(o.Sec)

	return encoder.CheckError()
}

// MDecodeTimestamp deserializes a Timestamp using msgpack
func MDecodeTimestamp(d *msgpack.Decoder) (Timestamp, error) {
	var val Timestamp
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
		case "Nsec":
			val.Nsec, err = d.ReadUint32()
		case "Sec":
			val.Sec, err = d.ReadInt64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Timestamp using cbor
func (o *Timestamp) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Nsec")
	encoder.WriteUint32(o.Nsec)
	encoder.WriteString("Sec")
	encoder.WriteInt64(o.Sec)

	return encoder.CheckError()
}

// CDecodeTimestamp deserializes a Timestamp using cbor
func CDecodeTimestamp(d *cbor.Decoder) (Timestamp, error) {
	var val Timestamp
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
		case "Nsec":
			val.Nsec, err = d.ReadUint32()
		case "Sec":
			val.Sec, err = d.ReadInt64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Transport struct {
	Binding   string `json:"binding"`
	Namespace string `json:"namespace"`
}

// MEncode serializes a Transport using msgpack
func (o *Transport) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("binding")
	encoder.WriteString(o.Binding)
	encoder.WriteString("namespace")
	encoder.WriteString(o.Namespace)

	return encoder.CheckError()
}

// MDecodeTransport deserializes a Transport using msgpack
func MDecodeTransport(d *msgpack.Decoder) (Transport, error) {
	var val Transport
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
		case "binding":
			val.Binding, err = d.ReadString()
		case "namespace":
			val.Namespace, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Transport using cbor
func (o *Transport) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("binding")
	encoder.WriteString(o.Binding)
	encoder.WriteString("namespace")
	encoder.WriteString(o.Namespace)

	return encoder.CheckError()
}

// CDecodeTransport deserializes a Transport using cbor
func CDecodeTransport(d *cbor.Decoder) (Transport, error) {
	var val Transport
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
		case "binding":
			val.Binding, err = d.ReadString()
		case "namespace":
			val.Namespace, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Actor service
type Actor interface {
	HealthRequest(ctx *Context, arg core.HealthCheckRequest) (*core.HealthCheckResponse, error)
}

// ActorHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func ActorHandler(actor_ Actor) Handler {
	return NewHandler("Actor", &ActorReceiver{}, actor_)
}

// ActorReceiver receives messages defined in the Actor service interface
// Actor service
type ActorReceiver struct{}

func (r *ActorReceiver) Dispatch(ctx *Context, svc interface{}, message *Message) (*Message, error) {
	svc_, _ := svc.(Actor)
	switch message.Method {

	case "HealthRequest":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := core.MDecodeHealthCheckRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.HealthRequest(ctx, value)
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
			return &Message{Method: "Actor.HealthRequest", Arg: buf}, nil
		}
	default:
		return nil, NewRpcError("MethodNotHandled", "Actor."+message.Method)
	}
}

// ActorSender sends messages to a Actor service
// Actor service
type ActorSender struct{ transport Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorActorSender(actor_id string) *ActorSender {
	transport := ToActor(actor_id)
	return &ActorSender{transport: transport}
}

func (s *ActorSender) HealthRequest(ctx *Context, arg core.HealthCheckRequest) (*core.HealthCheckResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, Message{Method: "Actor.HealthRequest", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := core.MDecodeHealthCheckResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.5.1
