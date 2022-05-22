// wasmcloud core data models for messaging and code generation
package actor

import (
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// Capability contract id, e.g. 'wasmcloud:httpserver'
// This declaration supports code generations and is not part of an actor or provider sdk
type CapabilityContractId string

// Encode serializes a CapabilityContractId using msgpack
func (o *CapabilityContractId) Encode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return nil
}

// Decode deserializes a CapabilityContractId using msgpack
func DecodeCapabilityContractId(d msgpack.Decoder) (CapabilityContractId, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return CapabilityContractId(val), nil
}

// 32-bit float
type F32 float32

// 64-bit float aka double
type F64 float64

// signed 16-bit int
type I16 int16

// signed 32-bit int
type I32 int32

// signed 64-bit int
type I64 int64

// signed byte
type I8 int8

// list of identifiers
// This declaration supports code generations and is not part of an actor or provider sdk
type IdentifierList []string

// Encode serializes a IdentifierList using msgpack
func (o *IdentifierList) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteString(item_o)
	}

	return nil
}

// Decode deserializes a IdentifierList using msgpack
func DecodeIdentifierList(d msgpack.Decoder) (IdentifierList, error) {
	isNil, err := d.IsNextNil()
	if isNil {
		if err != nil {
			err = d.Skip()
		}
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

// unsigned 16-bit int
type U16 int16

// unsigned 32-bit int
type U32 int32

// unsigned 64-bit int
type U64 int64

// unsigned byte
type U8 int8

// Unit type
type Unit struct {
}

// Encode serializes a Unit using msgpack
func (o *Unit) Encode(encoder msgpack.Writer) error {
	encoder.WriteNil()
	return nil
}

// Decode deserializes a Unit using msgpack
func DecodeUnit(d msgpack.Decoder) (Unit, error) {
	_ = d.Skip()
	return Unit{}, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
