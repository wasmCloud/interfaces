// wasmcloud platform core data structures
package actor

import (
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// List of linked actors for a provider
type ActorLinks []LinkDefinition

// Encode serializes a ActorLinks using msgpack
func (o *ActorLinks) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.Encode(encoder)
	}

	return nil
}

// Decode deserializes a ActorLinks using msgpack
func DecodeActorLinks(d msgpack.Decoder) (ActorLinks, error) {
	isNil, err := d.IsNextNil()
	if isNil {
		if err != nil {
			err = d.Skip()
		}
		return make([]LinkDefinition, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		size = 0
	}
	val := make([]LinkDefinition, size)
	for i := uint32(0); i < size; i++ {
		item, err := DecodeLinkDefinition(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type ClusterIssuerKey string

// Encode serializes a ClusterIssuerKey using msgpack
func (o *ClusterIssuerKey) Encode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return nil
}

// Decode deserializes a ClusterIssuerKey using msgpack
func DecodeClusterIssuerKey(d msgpack.Decoder) (ClusterIssuerKey, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ClusterIssuerKey(val), nil
}

type ClusterIssuers []ClusterIssuerKey

// Encode serializes a ClusterIssuers using msgpack
func (o *ClusterIssuers) Encode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.Encode(encoder)
	}

	return nil
}

// Decode deserializes a ClusterIssuers using msgpack
func DecodeClusterIssuers(d msgpack.Decoder) (ClusterIssuers, error) {
	isNil, err := d.IsNextNil()
	if isNil {
		if err != nil {
			err = d.Skip()
		}
		return make([]ClusterIssuerKey, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		size = 0
	}
	val := make([]ClusterIssuerKey, size)
	for i := uint32(0); i < size; i++ {
		item, err := DecodeClusterIssuerKey(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// health check request parameter
type HealthCheckRequest struct {
}

// Encode serializes a HealthCheckRequest using msgpack
func (o *HealthCheckRequest) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(0)

	return nil
}

// Decode deserializes a HealthCheckRequest using msgpack
func DecodeHealthCheckRequest(d msgpack.Decoder) (HealthCheckRequest, error) {
	var val HealthCheckRequest
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
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Return value from actors and providers for health check status
type HealthCheckResponse struct {
	// A flag that indicates the the actor is healthy
	Healthy bool
	// A message containing additional information about the actors health
	Message string
}

// Encode serializes a HealthCheckResponse using msgpack
func (o *HealthCheckResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("Healthy")
	encoder.WriteBool(o.Healthy)
	encoder.WriteString("Message")
	encoder.WriteString(o.Message)

	return nil
}

// Decode deserializes a HealthCheckResponse using msgpack
func DecodeHealthCheckResponse(d msgpack.Decoder) (HealthCheckResponse, error) {
	var val HealthCheckResponse
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
		case "Healthy":
			val.Healthy, err = d.ReadBool()
		case "Message":
			val.Message, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// initialization data for a capability provider
type HostData struct {
	HostId             string
	LatticeRpcPrefix   string
	LinkName           string
	LatticeRpcUserJwt  string
	LatticeRpcUserSeed string
	LatticeRpcUrl      string
	ProviderKey        string
	InvocationSeed     string
	EnvValues          HostEnvValues
	InstanceId         string
	// initial list of links for provider
	LinkDefinitions ActorLinks
	// list of cluster issuers
	ClusterIssuers ClusterIssuers
	// Optional configuration JSON sent to a given link name of a provider
	// without an actor context
	ConfigJson string
	// Host-wide default RPC timeout for rpc messages, in milliseconds.  Defaults to 2000.
	DefaultRpcTimeoutMs uint64
}

// Encode serializes a HostData using msgpack
func (o *HostData) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(14)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("lattice_rpc_prefix")
	encoder.WriteString(o.LatticeRpcPrefix)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("lattice_rpc_user_jwt")
	encoder.WriteString(o.LatticeRpcUserJwt)
	encoder.WriteString("lattice_rpc_user_seed")
	encoder.WriteString(o.LatticeRpcUserSeed)
	encoder.WriteString("lattice_rpc_url")
	encoder.WriteString(o.LatticeRpcUrl)
	encoder.WriteString("provider_key")
	encoder.WriteString(o.ProviderKey)
	encoder.WriteString("invocation_seed")
	encoder.WriteString(o.InvocationSeed)
	encoder.WriteString("env_values")
	o.EnvValues.Encode(encoder)
	encoder.WriteString("instance_id")
	encoder.WriteString(o.InstanceId)
	encoder.WriteString("link_definitions")
	o.LinkDefinitions.Encode(encoder)
	encoder.WriteString("cluster_issuers")
	o.ClusterIssuers.Encode(encoder)
	encoder.WriteString("config_json")
	encoder.WriteString(o.ConfigJson)
	encoder.WriteString("default_rpc_timeout_ms")
	encoder.WriteUint64(o.DefaultRpcTimeoutMs)

	return nil
}

// Decode deserializes a HostData using msgpack
func DecodeHostData(d msgpack.Decoder) (HostData, error) {
	var val HostData
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
		case "host_id":
			val.HostId, err = d.ReadString()
		case "lattice_rpc_prefix":
			val.LatticeRpcPrefix, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "lattice_rpc_user_jwt":
			val.LatticeRpcUserJwt, err = d.ReadString()
		case "lattice_rpc_user_seed":
			val.LatticeRpcUserSeed, err = d.ReadString()
		case "lattice_rpc_url":
			val.LatticeRpcUrl, err = d.ReadString()
		case "provider_key":
			val.ProviderKey, err = d.ReadString()
		case "invocation_seed":
			val.InvocationSeed, err = d.ReadString()
		case "env_values":
			val.EnvValues, err = DecodeHostEnvValues(d)
		case "instance_id":
			val.InstanceId, err = d.ReadString()
		case "link_definitions":
			val.LinkDefinitions, err = DecodeActorLinks(d)
		case "cluster_issuers":
			val.ClusterIssuers, err = DecodeClusterIssuers(d)
		case "config_json":
			val.ConfigJson, err = d.ReadString()
		case "default_rpc_timeout_ms":
			val.DefaultRpcTimeoutMs, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Environment settings for initializing a capability provider
type HostEnvValues map[string]string

// Encode serializes a HostEnvValues using msgpack
func (o *HostEnvValues) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}

// Decode deserializes a HostEnvValues using msgpack
func DecodeHostEnvValues(d msgpack.Decoder) (HostEnvValues, error) {
	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		err = d.Skip()
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

// RPC message to capability provider
type Invocation struct {
	Origin        WasmCloudEntity
	Target        WasmCloudEntity
	Operation     string
	Msg           []byte
	Id            string
	EncodedClaims string
	HostId        string
	// total message size (optional)
	ContentLength uint64
	// Open Telemetry tracing support
	TraceContext TraceContext
}

// Encode serializes a Invocation using msgpack
func (o *Invocation) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(9)
	encoder.WriteString("Origin")
	o.Origin.Encode(encoder)
	encoder.WriteString("Target")
	o.Target.Encode(encoder)
	encoder.WriteString("Operation")
	encoder.WriteString(o.Operation)
	encoder.WriteString("Msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("Id")
	encoder.WriteString(o.Id)
	encoder.WriteString("encoded_claims")
	encoder.WriteString(o.EncodedClaims)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("TraceContext")
	if o.TraceContext == nil {
		encoder.WriteNil()
	} else {
		o.TraceContext.Encode(encoder)
	}

	return nil
}

// Decode deserializes a Invocation using msgpack
func DecodeInvocation(d msgpack.Decoder) (Invocation, error) {
	var val Invocation
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
		case "Origin":
			val.Origin, err = DecodeWasmCloudEntity(d)
		case "Target":
			val.Target, err = DecodeWasmCloudEntity(d)
		case "Operation":
			val.Operation, err = d.ReadString()
		case "Msg":
			val.Msg, err = d.ReadByteArray()
		case "Id":
			val.Id, err = d.ReadString()
		case "encoded_claims":
			val.EncodedClaims, err = d.ReadString()
		case "host_id":
			val.HostId, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		case "TraceContext":
			val.TraceContext, err = DecodeTraceContext(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Response to an invocation
type InvocationResponse struct {
	// serialize response message
	Msg []byte
	// id connecting this response to the invocation
	InvocationId string
	// optional error message
	Error string
	// total message size (optional)
	ContentLength uint64
}

// Encode serializes a InvocationResponse using msgpack
func (o *InvocationResponse) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("Msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("invocation_id")
	encoder.WriteString(o.InvocationId)
	encoder.WriteString("Error")
	encoder.WriteString(o.Error)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)

	return nil
}

// Decode deserializes a InvocationResponse using msgpack
func DecodeInvocationResponse(d msgpack.Decoder) (InvocationResponse, error) {
	var val InvocationResponse
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
		case "Msg":
			val.Msg, err = d.ReadByteArray()
		case "invocation_id":
			val.InvocationId, err = d.ReadString()
		case "Error":
			val.Error, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Link definition for binding actor to provider
type LinkDefinition struct {
	// actor public key
	ActorId string
	// provider public key
	ProviderId string
	// link name
	LinkName string
	// contract id
	ContractId string
	Values     LinkSettings
}

// Encode serializes a LinkDefinition using msgpack
func (o *LinkDefinition) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("actor_id")
	encoder.WriteString(o.ActorId)
	encoder.WriteString("provider_id")
	encoder.WriteString(o.ProviderId)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	encoder.WriteString(o.ContractId)
	encoder.WriteString("Values")
	o.Values.Encode(encoder)

	return nil
}

// Decode deserializes a LinkDefinition using msgpack
func DecodeLinkDefinition(d msgpack.Decoder) (LinkDefinition, error) {
	var val LinkDefinition
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
		case "actor_id":
			val.ActorId, err = d.ReadString()
		case "provider_id":
			val.ProviderId, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = d.ReadString()
		case "Values":
			val.Values, err = DecodeLinkSettings(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil

}

// Settings associated with an actor-provider link
type LinkSettings map[string]string

// Encode serializes a LinkSettings using msgpack
func (o *LinkSettings) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}

// Decode deserializes a LinkSettings using msgpack
func DecodeLinkSettings(d msgpack.Decoder) (LinkSettings, error) {
	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		err = d.Skip()
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

// Environment settings for initializing a capability provider
type TraceContext map[string]string

// Encode serializes a TraceContext using msgpack
func (o *TraceContext) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return nil
}

// Decode deserializes a TraceContext using msgpack
func DecodeTraceContext(d msgpack.Decoder) (TraceContext, error) {
	isNil, err := d.IsNextNil()
	if err != nil && isNil {
		err = d.Skip()
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		size = 0
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, err := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

type WasmCloudEntity struct {
	PublicKey  string
	LinkName   string
	ContractId CapabilityContractId
}

// Encode serializes a WasmCloudEntity using msgpack
func (o *WasmCloudEntity) Encode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("public_key")
	encoder.WriteString(o.PublicKey)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	o.ContractId.Encode(encoder)

	return nil
}

// Decode deserializes a WasmCloudEntity using msgpack
func DecodeWasmCloudEntity(d msgpack.Decoder) (WasmCloudEntity, error) {
	var val WasmCloudEntity
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
		case "public_key":
			val.PublicKey, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = DecodeCapabilityContractId(d)
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
	// Perform health check. Called at regular intervals by host
	HealthRequest(ctx *Context, arg HealthCheckRequest) (*HealthCheckResponse, error)
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
			value, err_ := DecodeHealthCheckRequest(d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.HealthRequest(ctx, value)
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
			return &Message{Method: "Actor.HealthRequest", Arg: buf}, nil
		}
	default:
		return nil, NewRpcError("MethodNotHandled", "Actor."+message.Method)
	}
}

// ActorSender sends messages to a Actor service
// Actor service
type ActorSender struct{ transport Transport }

// Perform health check. Called at regular intervals by host
func (s *ActorSender) HealthRequest(ctx *Context, arg HealthCheckRequest) (*HealthCheckResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.Encode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.Encode(enc)

	out_buf, _ := s.transport.Send(ctx, Message{Method: "Actor.HealthRequest", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := DecodeHealthCheckResponse(d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.4
