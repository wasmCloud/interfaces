// Messaging: wasmcloud messaging capability provider: publish, request-reply, and subscriptions
package messaging

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// A message to be published
type PubMessage struct {
	// The subject, or topic, of the message
	Subject string `json:"subject"`
	// An optional topic on which the reply should be sent.
	ReplyTo string `json:"replyTo"`
	// The message payload
	Body []byte `json:"body"`
}

// MEncode serializes a PubMessage using msgpack
func (o *PubMessage) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// MDecodePubMessage deserializes a PubMessage using msgpack
func MDecodePubMessage(d *msgpack.Decoder) (PubMessage, error) {
	var val PubMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a PubMessage using cbor
func (o *PubMessage) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// CDecodePubMessage deserializes a PubMessage using cbor
func CDecodePubMessage(d *cbor.Decoder) (PubMessage, error) {
	var val PubMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Reply received from a Request operation
type ReplyMessage struct {
	// The subject, or topic, of the message
	Subject string `json:"subject"`
	// An optional topic on which the reply should be sent.
	ReplyTo string `json:"replyTo"`
	// The message payload
	Body []byte `json:"body"`
}

// MEncode serializes a ReplyMessage using msgpack
func (o *ReplyMessage) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// MDecodeReplyMessage deserializes a ReplyMessage using msgpack
func MDecodeReplyMessage(d *msgpack.Decoder) (ReplyMessage, error) {
	var val ReplyMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ReplyMessage using cbor
func (o *ReplyMessage) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// CDecodeReplyMessage deserializes a ReplyMessage using cbor
func CDecodeReplyMessage(d *cbor.Decoder) (ReplyMessage, error) {
	var val ReplyMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Message sent as part of a request, with timeout
type RequestMessage struct {
	// The subject, or topic, of the message
	Subject string `json:"subject"`
	// The message payload
	Body []byte `json:"body"`
	// A timeout, in milliseconds
	TimeoutMs uint32 `json:"timeoutMs"`
}

// MEncode serializes a RequestMessage using msgpack
func (o *RequestMessage) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)
	encoder.WriteString("timeoutMs")
	encoder.WriteUint32(o.TimeoutMs)

	return encoder.CheckError()
}

// MDecodeRequestMessage deserializes a RequestMessage using msgpack
func MDecodeRequestMessage(d *msgpack.Decoder) (RequestMessage, error) {
	var val RequestMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		case "timeoutMs":
			val.TimeoutMs, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a RequestMessage using cbor
func (o *RequestMessage) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)
	encoder.WriteString("timeoutMs")
	encoder.WriteUint32(o.TimeoutMs)

	return encoder.CheckError()
}

// CDecodeRequestMessage deserializes a RequestMessage using cbor
func CDecodeRequestMessage(d *cbor.Decoder) (RequestMessage, error) {
	var val RequestMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		case "timeoutMs":
			val.TimeoutMs, err = d.ReadUint32()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Message received as part of a subscription
type SubMessage struct {
	// The subject, or topic, of the message
	Subject string `json:"subject"`
	// An optional topic on which the reply should be sent.
	ReplyTo string `json:"replyTo"`
	// The message payload
	Body []byte `json:"body"`
}

// MEncode serializes a SubMessage using msgpack
func (o *SubMessage) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// MDecodeSubMessage deserializes a SubMessage using msgpack
func MDecodeSubMessage(d *msgpack.Decoder) (SubMessage, error) {
	var val SubMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SubMessage using cbor
func (o *SubMessage) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("replyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return encoder.CheckError()
}

// CDecodeSubMessage deserializes a SubMessage using cbor
func CDecodeSubMessage(d *cbor.Decoder) (SubMessage, error) {
	var val SubMessage
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
		case "subject":
			val.Subject, err = d.ReadString()
		case "replyTo":
			val.ReplyTo, err = d.ReadString()
		case "body":
			val.Body, err = d.ReadByteArray()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// The MessageSubscriber interface describes
// an actor interface that receives messages
// sent by the Messaging provider
type MessageSubscriber interface {
	// subscription handler
	HandleMessage(ctx *actor.Context, arg SubMessage) error
}

// MessageSubscriberHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func MessageSubscriberHandler(actor_ MessageSubscriber) actor.Handler {
	return actor.NewHandler("MessageSubscriber", &MessageSubscriberReceiver{}, actor_)
}

// MessageSubscriberContractId returns the capability contract id for this interface
func MessageSubscriberContractId() string { return "wasmcloud:messaging" }

// MessageSubscriberReceiver receives messages defined in the MessageSubscriber service interface
// The MessageSubscriber interface describes
// an actor interface that receives messages
// sent by the Messaging provider
type MessageSubscriberReceiver struct{}

func (r *MessageSubscriberReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(MessageSubscriber)
	switch message.Method {

	case "HandleMessage":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeSubMessage(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.HandleMessage(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "MessageSubscriber.HandleMessage", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "MessageSubscriber."+message.Method)
	}
}

// MessageSubscriberSender sends messages to a MessageSubscriber service
// The MessageSubscriber interface describes
// an actor interface that receives messages
// sent by the Messaging provider
type MessageSubscriberSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorMessageSubscriberSender(actor_id string) *MessageSubscriberSender {
	transport := actor.ToActor(actor_id)
	return &MessageSubscriberSender{transport: transport}
}

// subscription handler
func (s *MessageSubscriberSender) HandleMessage(ctx *actor.Context, arg SubMessage) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "MessageSubscriber.HandleMessage", Arg: buf})
	return nil
}

// The Messaging interface describes a service
// that can deliver messages
type Messaging interface {
	// Publish - send a message
	// The function returns after the message has been sent.
	// If the sender expects to receive an asynchronous reply,
	// the replyTo field should be filled with the
	// subject for the response.
	Publish(ctx *actor.Context, arg PubMessage) error
	// Request - send a message in a request/reply pattern,
	// waiting for a response.
	Request(ctx *actor.Context, arg RequestMessage) (*ReplyMessage, error)
}

// MessagingHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func MessagingHandler(actor_ Messaging) actor.Handler {
	return actor.NewHandler("Messaging", &MessagingReceiver{}, actor_)
}

// MessagingContractId returns the capability contract id for this interface
func MessagingContractId() string { return "wasmcloud:messaging" }

// MessagingReceiver receives messages defined in the Messaging service interface
// The Messaging interface describes a service
// that can deliver messages
type MessagingReceiver struct{}

func (r *MessagingReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Messaging)
	switch message.Method {

	case "Publish":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodePubMessage(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.Publish(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Messaging.Publish", Arg: buf}, nil
		}
	case "Request":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeRequestMessage(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Request(ctx, value)
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
			return &actor.Message{Method: "Messaging.Request", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Messaging."+message.Method)
	}
}

// MessagingSender sends messages to a Messaging service
// The Messaging interface describes a service
// that can deliver messages
type MessagingSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a Messaging provider
// implementing the 'wasmcloud:messaging' capability contract, with the "default" link
func NewProviderMessaging() *MessagingSender {
	transport := actor.ToProvider("wasmcloud:messaging", "default")
	return &MessagingSender{transport: transport}
}

// NewProviderMessagingLink constructs a client for sending to a Messaging provider
// implementing the 'wasmcloud:messaging' capability contract, with the specified link name
func NewProviderMessagingLink(linkName string) *MessagingSender {
	transport := actor.ToProvider("wasmcloud:messaging", linkName)
	return &MessagingSender{transport: transport}
}

// Publish - send a message
// The function returns after the message has been sent.
// If the sender expects to receive an asynchronous reply,
// the replyTo field should be filled with the
// subject for the response.
func (s *MessagingSender) Publish(ctx *actor.Context, arg PubMessage) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Messaging.Publish", Arg: buf})
	return nil
}

// Request - send a message in a request/reply pattern,
// waiting for a response.
func (s *MessagingSender) Request(ctx *actor.Context, arg RequestMessage) (*ReplyMessage, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Messaging.Request", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeReplyMessage(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.5.1
