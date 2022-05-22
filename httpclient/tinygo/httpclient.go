// Httpclient: wasmcloud capability contract for http client
package httpclient

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

// http request to be sent through the provider
type HttpRequest struct {
	// http method, defaults to "GET"
	Method string
	Url    string
	// optional headers. defaults to empty
	Headers HeaderMap
	// request body, defaults to empty
	Body []byte
}

// Encode serializes a HttpRequest using msgpack
func (o *HttpRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("method")
	encoder.WriteString(o.Method)
	encoder.WriteString("url")
	encoder.WriteString(o.Url)
	encoder.WriteString("headers")
	o.Headers.Encode(encoder)
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
		case "url":
			val.Url, err = d.ReadString()
		case "headers":
			val.Headers, err = DecodeHeaderMap(d)
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

// response from the http request
type HttpResponse struct {
	// response status code
	StatusCode uint16
	// Case is not guaranteed to be normalized, so
	// actors checking response headers need to do their own
	// case conversion.
	// Example (rust):
	// // check for 'Content-Type' header
	// let content_type:Option<&Vec<String>> = header.iter()
	// .map(|(k,_)| k.to_ascii_lowercase())
	// .find(|(k,_)| k == "content-type")
	// .map(|(_,v)| v);
	Header HeaderMap
	// response body
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

// HttpClient - issue outgoing http requests via an external provider
// To use this capability, the actor must be linked
// with "wasmcloud:httpclient"
type HttpClient interface {
	// Issue outgoing http request
	Request(ctx *actor.Context, arg HttpRequest) (*HttpResponse, error)
}

// HttpClientHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func HttpClientHandler(actor_ HttpClient) actor.Handler {
	return actor.NewHandler("HttpClient", &HttpClientReceiver{}, actor_)
}

// HttpClientContractId returns the capability contract id for this interface
func HttpClientContractId() string { return "wasmcloud:httpclient" }

// HttpClientReceiver receives messages defined in the HttpClient service interface
// HttpClient - issue outgoing http requests via an external provider
// To use this capability, the actor must be linked
// with "wasmcloud:httpclient"
type HttpClientReceiver struct{}

func (r *HttpClientReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(HttpClient)
	switch message.Method {

	case "Request":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeHttpRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Request(ctx, value)
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
			return &actor.Message{Method: "HttpClient.Request", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "HttpClient."+message.Method)
	}
}

// HttpClientSender sends messages to a HttpClient service
// HttpClient - issue outgoing http requests via an external provider
// To use this capability, the actor must be linked
// with "wasmcloud:httpclient"
type HttpClientSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a HttpClient provider
// implementing the 'wasmcloud:httpclient' capability contract, with the "default" link
func NewProviderHttpClient() *HttpClientSender {
	transport := actor.ToProvider("wasmcloud:httpclient", "default")
	return &HttpClientSender{transport: transport}
}

// NewProviderHttpClientLink constructs a client for sending to a HttpClient provider
// implementing the 'wasmcloud:httpclient' capability contract, with the specified link name
func NewProviderHttpClientLink(linkName string) *HttpClientSender {
	transport := actor.ToProvider("wasmcloud:httpclient", linkName)
	return &HttpClientSender{transport: transport}
}

// Issue outgoing http request
func (s *HttpClientSender) Request(ctx *actor.Context, arg HttpRequest) (*HttpResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "HttpClient.Request", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeHttpResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
