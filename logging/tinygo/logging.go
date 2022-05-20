// Logging: wasmcloud built-in logging capability provider
package logging

import (
	"github.com/wasmcloud/tinygo-msgpack"    //nolint
	"github.com/wasmcloud/actor-tinygo" //nolint
)

type LogEntry struct {
	// severity level: debug,info,warn,error
	Level string
	// message to log
	Text string
}

func (o *LogEntry) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Level")
	encoder.WriteString(o.Level)
	encoder.WriteString("Text")
	encoder.WriteString(o.Text)

	return nil
}
func DecodeLogEntry(d msgpack.Decoder) (LogEntry, error) {

	var val LogEntry
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

		case "Level":
			val.Level, err = d.ReadString()
		case "Text":
			val.Text, err = d.ReadString()
		default:
			err = d.Skip()
			if err != nil {
				return val, err
			}
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

type Logging interface {
	//
	// WriteLog - log a text message
	//
	WriteLog(ctx *actor.Context, arg LogEntry) error
}

// LoggingContractId returns the capability contract id for this interface
func LoggingContractId() string { return "wasmcloud:builtin:logging" }

// LoggingReceiver receives messages defined in the Logging service interface
type LoggingReceiver struct{}

func (r *LoggingReceiver) dispatch(ctx *actor.Context, svc Logging, message *actor.Message) (*actor.Message, error) {
	switch message.Method {
	case "WriteLog":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeLogEntry(d)
			if err_ != nil {
				return nil, err_
			}

			err := svc.WriteLog(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "Logging.WriteLog", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Logging."+message.Method)
	}
}

// LoggingSender sends messages to a Logging service
type LoggingSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a Logging provider
// implementing the 'wasmcloud:builtin:logging' capability contract, with the "default" link
func NewProviderLogging() *LoggingSender {
	transport := actor.ToProvider("wasmcloud:builtin:logging", "default")
	return &LoggingSender{transport: transport}
}

// NewProviderLoggingLink constructs a client for sending to a Logging provider
// implementing the 'wasmcloud:builtin:logging' capability contract, with the specified link name
func NewProviderLoggingLink(linkName string) *LoggingSender {
	transport := actor.ToProvider("wasmcloud:builtin:logging", linkName)
	return &LoggingSender{transport: transport}
}

//
// WriteLog - log a text message
//
func (s *LoggingSender) WriteLog(ctx *actor.Context, arg LogEntry) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	s.transport.Send(ctx, actor.Message{Method: "Logging.WriteLog", Arg: buf})
	return nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
