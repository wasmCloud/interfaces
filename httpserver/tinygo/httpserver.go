// Httpserver: wasmcloud capability contract for http server
package httpserver

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// map data structure for holding http headers
//
type HeaderMap map[string]HeaderValues

// Encode serializes a HeaderMap using msgpack
func (o *HeaderMap) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		val_o.Encode(encoder)
	}

	return nil
}

// Decode deserializes a HeaderMap using msgpack
func DecodeHeaderMap(d *msgpack.Decoder) (HeaderMap, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]HeaderValues, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]HeaderValues, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := DecodeHeaderValues(d)
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

type HeaderValues []string

// Encode serializes a HeaderValues using msgpack
func (o *HeaderValues) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteString(item_o)
	}

	return nil
}

// Decode deserializes a HeaderValues using msgpack
func DecodeHeaderValues(d *msgpack.Decoder) (HeaderValues, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]string, 0), err
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

// HttpRequest contains data sent to actor about the http request
type HttpRequest struct {
	// HTTP method. One of: GET,POST,PUT,DELETE,HEAD,OPTIONS,CONNECT,PATCH,TRACE
	Method string
	// full request path
	Path string
	// query string. May be an empty string if there were no query parameters.
	QueryString string
	// map of request headers (string key, string value)
	Header HeaderMap
	// Request body as a byte array. May be empty.
	Body []byte
}

// Encode serializes a HttpRequest using msgpack
func (o *HttpRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("method")
	encoder.WriteString(o.Method)
	encoder.WriteString("path")
	encoder.WriteString(o.Path)
	encoder.WriteString("queryString")
	encoder.WriteString(o.QueryString)
	encoder.WriteString("header")
	o.Header.Encode(encoder)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return nil
}

// Decode deserializes a HttpRequest using msgpack
func DecodeHttpRequest(d *msgpack.Decoder) (HttpRequest, error) {
	var val HttpRequest
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
		case "method":
			val.Method, err = d.ReadString()
		case "path":
			val.Path, err = d.ReadString()
		case "queryString":
			val.QueryString, err = d.ReadString()
		case "header":
			val.Header, err = DecodeHeaderMap(d)
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

// HttpResponse contains the actor's response to return to the http client
type HttpResponse struct {
	// statusCode is a three-digit number, usually in the range 100-599,
	// A value of 200 indicates success.
	StatusCode uint16
	// Map of headers (string keys, list of values)
	Header HeaderMap
	// Body of response as a byte array. May be an empty array.
	Body []byte
}

// Encode serializes a HttpResponse using msgpack
func (o *HttpResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("statusCode")
	encoder.WriteUint16(o.StatusCode)
	encoder.WriteString("header")
	o.Header.Encode(encoder)
	encoder.WriteString("body")
	encoder.WriteByteArray(o.Body)

	return nil
}

// Decode deserializes a HttpResponse using msgpack
func DecodeHttpResponse(d *msgpack.Decoder) (HttpResponse, error) {
	var val HttpResponse
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
		case "statusCode":
			val.StatusCode, err = d.ReadUint16()
		case "header":
			val.Header, err = DecodeHeaderMap(d)
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

// HttpServer is the contract to be implemented by actor
type HttpServer interface {
	HandleRequest(ctx *actor.Context, arg HttpRequest) (*HttpResponse, error)
}

// HttpServerHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func HttpServerHandler(actor_ HttpServer) actor.Handler {
	return actor.NewHandler("HttpServer", &HttpServerReceiver{}, actor_)
}

// HttpServerContractId returns the capability contract id for this interface
func HttpServerContractId() string { return "wasmcloud:httpserver" }

// HttpServerReceiver receives messages defined in the HttpServer service interface
// HttpServer is the contract to be implemented by actor
type HttpServerReceiver struct{}

func (r *HttpServerReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(HttpServer)
	switch message.Method {

	case "HandleRequest":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeHttpRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.HandleRequest(ctx, value)
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
			return &actor.Message{Method: "HttpServer.HandleRequest", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "HttpServer."+message.Method)
	}
}

// HttpServerSender sends messages to a HttpServer service
// HttpServer is the contract to be implemented by actor
type HttpServerSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorHttpServerSender(actor_id string) *HttpServerSender {
	transport := actor.ToActor(actor_id)
	return &HttpServerSender{transport: transport}
}

func (s *HttpServerSender) HandleRequest(ctx *actor.Context, arg HttpRequest) (*HttpResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "HttpServer.HandleRequest", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeHttpResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
