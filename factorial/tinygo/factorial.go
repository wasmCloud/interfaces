// Factorial: A simple service that calculates the factorial of a whole number
package factorial

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// The Factorial service has a single method, calculate, which
// calculates the factorial of its whole number parameter.
type Factorial interface {
	// Calculates the factorial (n!) of the input parameter
	Calculate(ctx *actor.Context, arg uint32) (uint64, error)
}

// FactorialHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func FactorialHandler(actor_ Factorial) actor.Handler {
	return actor.NewHandler("Factorial", &FactorialReceiver{}, actor_)
}

// FactorialReceiver receives messages defined in the Factorial service interface
// The Factorial service has a single method, calculate, which
// calculates the factorial of its whole number parameter.
type FactorialReceiver struct{}

func (r *FactorialReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Factorial)
	switch message.Method {

	case "Calculate":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := d.ReadUint32()
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Calculate(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteUint64(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteUint64(resp)
			return &actor.Message{Method: "Factorial.Calculate", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Factorial."+message.Method)
	}
}

// FactorialSender sends messages to a Factorial service
// The Factorial service has a single method, calculate, which
// calculates the factorial of its whole number parameter.
type FactorialSender struct{ transport actor.Transport }

// Calculates the factorial (n!) of the input parameter
func (s *FactorialSender) Calculate(ctx *actor.Context, arg uint32) (uint64, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	size_enc.WriteUint32(arg)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	enc.WriteUint32(arg)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Factorial.Calculate", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint64()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
