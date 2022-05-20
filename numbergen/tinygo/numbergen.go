// Numbergen: wasmcloud built-in capability provider for number generation
package numbergen

import (
	"github.com/wasmcloud/tinygo-msgpack"    //nolint
	"github.com/wasmcloud/actor-tinygo" //nolint
)

// Input range for RandomInRange, inclusive. Result will be >= min and <= max
// Example:
// random_in_range(RangeLimit{0,4}) returns one the values, 0, 1, 2, 3, or 4.
type RangeLimit struct {
	Min uint32
	Max uint32
}

func (o *RangeLimit) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Min")
	encoder.WriteUint32(o.Min)
	encoder.WriteString("Max")
	encoder.WriteUint32(o.Max)

	return nil
}
func DecodeRangeLimit(d msgpack.Decoder) (RangeLimit, error) {

	var val RangeLimit
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

		case "Min":
			val.Min, err = d.ReadUint32()
		case "Max":
			val.Max, err = d.ReadUint32()
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

type NumberGen interface {
	//
	// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
	// These guids are known as "version 4", meaning all bits are random or pseudo-random.
	//
	GenerateGuid(ctx *actor.Context) (string, error)
	// Request a random integer within a range
	// The result will will be in the range [min,max), i.e., >= min and < max.
	RandomInRange(ctx *actor.Context, arg RangeLimit) (uint32, error)
	// Request a 32-bit random number
	Random32(ctx *actor.Context) (uint32, error)
}

// NumberGenContractId returns the capability contract id for this interface
func NumberGenContractId() string { return "wasmcloud:builtin:numbergen" }

// NumberGenReceiver receives messages defined in the NumberGen service interface
type NumberGenReceiver struct{}

func (r *NumberGenReceiver) dispatch(ctx *actor.Context, svc NumberGen, message *actor.Message) (*actor.Message, error) {
	switch message.Method {
	case "GenerateGuid":
		{
			resp, err := svc.GenerateGuid(ctx)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			size_enc.WriteString(resp)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			enc.WriteString(resp)

			return &actor.Message{Method: "NumberGen.GenerateGuid", Arg: buf}, nil
		}
	case "RandomInRange":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := DecodeRangeLimit(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc.RandomInRange(ctx, value)
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

			return &actor.Message{Method: "NumberGen.RandomInRange", Arg: buf}, nil
		}
	case "Random32":
		{
			resp, err := svc.Random32(ctx)
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

			return &actor.Message{Method: "NumberGen.Random32", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "NumberGen."+message.Method)
	}
}

// NumberGenSender sends messages to a NumberGen service
type NumberGenSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a NumberGen provider
// implementing the 'wasmcloud:builtin:numbergen' capability contract, with the "default" link
func NewProviderNumberGen() *NumberGenSender {
	transport := actor.ToProvider("wasmcloud:builtin:numbergen", "default")
	return &NumberGenSender{transport: transport}
}

// NewProviderNumberGenLink constructs a client for sending to a NumberGen provider
// implementing the 'wasmcloud:builtin:numbergen' capability contract, with the specified link name
func NewProviderNumberGenLink(linkName string) *NumberGenSender {
	transport := actor.ToProvider("wasmcloud:builtin:numbergen", linkName)
	return &NumberGenSender{transport: transport}
}

//
// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
// These guids are known as "version 4", meaning all bits are random or pseudo-random.
//
func (s *NumberGenSender) GenerateGuid(ctx *actor.Context) (string, error) {
	buf := make([]byte, 0)
	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "NumberGen.GenerateGuid", Arg: buf})

	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadString()
	if err_ != nil {
		return "", err_
	}
	return resp, nil
}

// Request a random integer within a range
// The result will will be in the range [min,max), i.e., >= min and < max.
func (s *NumberGenSender) RandomInRange(ctx *actor.Context, arg RangeLimit) (uint32, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "NumberGen.RandomInRange", Arg: buf})

	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// Request a 32-bit random number
func (s *NumberGenSender) Random32(ctx *actor.Context) (uint32, error) {
	buf := make([]byte, 0)
	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "NumberGen.Random32", Arg: buf})

	d := msgpack.NewDecoder(out_buf)
	resp, err_ := d.ReadUint32()
	if err_ != nil {
		return 0, err_
	}
	return resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
