// Messaging: wasmcloud messaging capability provider: publish, request-reply, and subscriptions
package messaging

import (
	"github.com/wasmcloud/actor-tinygo"   //nolint
	"github.com/wasmcloud/tinygo-msgpack" //nolint
)

// A message to be published
type PubMessage struct {
	// The subject, or topic, of the message
	Subject string
	// An optional topic on which the reply should be sent.
	ReplyTo string
	// The message payload
	Body []byte
}

// Encode serializes a PubMessage using msgpack
func (o *PubMessage) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("Subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("ReplyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("Body")
	encoder.WriteByteArray(o.Body)

	return nil
}

// Decode deserializes a PubMessage using msgpack
func DecodePubMessage(d msgpack.Decoder) (PubMessage, error) {
	var val PubMessage
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
		case "Subject":
			val.Subject, err = d.ReadString()
		case "ReplyTo":
			val.ReplyTo, err = d.ReadString()
		case "Body":
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
	Subject string
	// An optional topic on which the reply should be sent.
	ReplyTo string
	// The message payload
	Body []byte
}

// Encode serializes a ReplyMessage using msgpack
func (o *ReplyMessage) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("Subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("ReplyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("Body")
	encoder.WriteByteArray(o.Body)

	return nil
}

// Decode deserializes a ReplyMessage using msgpack
func DecodeReplyMessage(d msgpack.Decoder) (ReplyMessage, error) {
	var val ReplyMessage
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
		case "Subject":
			val.Subject, err = d.ReadString()
		case "ReplyTo":
			val.ReplyTo, err = d.ReadString()
		case "Body":
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
	Subject string
	// The message payload
	Body []byte
	// A timeout, in milliseconds
	TimeoutMs uint32
}

// Encode serializes a RequestMessage using msgpack
func (o *RequestMessage) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("Subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("Body")
	encoder.WriteByteArray(o.Body)
	encoder.WriteString("TimeoutMs")
	encoder.WriteUint32(o.TimeoutMs)

	return nil
}

// Decode deserializes a RequestMessage using msgpack
func DecodeRequestMessage(d msgpack.Decoder) (RequestMessage, error) {
	var val RequestMessage
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
		case "Subject":
			val.Subject, err = d.ReadString()
		case "Body":
			val.Body, err = d.ReadByteArray()
		case "TimeoutMs":
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
	Subject string
	// An optional topic on which the reply should be sent.
	ReplyTo string
	// The message payload
	Body []byte
}

// Encode serializes a SubMessage using msgpack
func (o *SubMessage) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("Subject")
	encoder.WriteString(o.Subject)
	encoder.WriteString("ReplyTo")
	encoder.WriteString(o.ReplyTo)
	encoder.WriteString("Body")
	encoder.WriteByteArray(o.Body)

	return nil
}

// Decode deserializes a SubMessage using msgpack
func DecodeSubMessage(d msgpack.Decoder) (SubMessage, error) {
	var val SubMessage
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
		case "Subject":
			val.Subject, err = d.ReadString()
		case "ReplyTo":
			val.ReplyTo, err = d.ReadString()
		case "Body":
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
func MessageSubscriberHandler() actor.Handler {
	return actor.NewHandler("MessageSubscriber", MessageSubscriberReceiver{})
}

// MessageSubscriberReceiver receives messages defined in the MessageSubscriber service interface
// The MessageSubscriber interface describes
// an actor interface that receives messages
// sent by the Messaging provider
type MessageSubscriberReceiver struct{}

func (r *MessageSubscriberReceiver) dispatch(ctx *actor.Context, svc MessageSubscriber, message *actor.Message) (*actor.Message, error) {
	switch message.Method {
	case "HandleMessage":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeSubMessage(d)
			if err_ != nil {
				return nil, err_
			}

			err := svc.HandleMessage(ctx, value)
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

// subscription handler
func (s *MessageSubscriberSender) HandleMessage(ctx *actor.Context, arg SubMessage) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

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
func MessagingHandler() actor.Handler {
	return actor.NewHandler("Messaging", MessagingReceiver{})
}

// MessagingReceiver receives messages defined in the Messaging service interface
// The Messaging interface describes a service
// that can deliver messages
type MessagingReceiver struct{}

func (r *MessagingReceiver) dispatch(ctx *actor.Context, svc Messaging, message *actor.Message) (*actor.Message, error) {
	switch message.Method {
	case "Publish":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodePubMessage(d)
			if err_ != nil {
				return nil, err_
			}

			err := svc.Publish(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Messaging.Publish", Arg: buf}, nil
		}
	case "Request":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeRequestMessage(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.Request(ctx, value)
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

// Publish - send a message
// The function returns after the message has been sent.
// If the sender expects to receive an asynchronous reply,
// the replyTo field should be filled with the
// subject for the response.
func (s *MessagingSender) Publish(ctx *actor.Context, arg PubMessage) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Messaging.Publish", Arg: buf})
	return nil
}

// Request - send a message in a request/reply pattern,
// waiting for a response.
func (s *MessagingSender) Request(ctx *actor.Context, arg RequestMessage) (*ReplyMessage, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Messaging.Request", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeReplyMessage(d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
