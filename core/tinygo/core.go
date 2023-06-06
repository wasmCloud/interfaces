// wasmcloud platform core data structures
package actor

import (
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// List of linked actors for a provider
type ActorLinks []LinkDefinition

// MEncode serializes a ActorLinks using msgpack
func (o *ActorLinks) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeActorLinks deserializes a ActorLinks using msgpack
func MDecodeActorLinks(d *msgpack.Decoder) (ActorLinks, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]LinkDefinition, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]LinkDefinition, 0), err
	}
	val := make([]LinkDefinition, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeLinkDefinition(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ActorLinks using cbor
func (o *ActorLinks) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeActorLinks deserializes a ActorLinks using cbor
func CDecodeActorLinks(d *cbor.Decoder) (ActorLinks, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]LinkDefinition, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]LinkDefinition, 0), err
	}
	val := make([]LinkDefinition, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeLinkDefinition(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type ClusterIssuerKey string

// MEncode serializes a ClusterIssuerKey using msgpack
func (o *ClusterIssuerKey) MEncode(encoder msgpack.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// MDecodeClusterIssuerKey deserializes a ClusterIssuerKey using msgpack
func MDecodeClusterIssuerKey(d *msgpack.Decoder) (ClusterIssuerKey, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ClusterIssuerKey(val), nil
}

// CEncode serializes a ClusterIssuerKey using cbor
func (o *ClusterIssuerKey) CEncode(encoder cbor.Writer) error {
	encoder.WriteString(string(*o))
	return encoder.CheckError()
}

// CDecodeClusterIssuerKey deserializes a ClusterIssuerKey using cbor
func CDecodeClusterIssuerKey(d *cbor.Decoder) (ClusterIssuerKey, error) {
	val, err := d.ReadString()
	if err != nil {
		return "", err
	}
	return ClusterIssuerKey(val), nil
}

type ClusterIssuers []ClusterIssuerKey

// MEncode serializes a ClusterIssuers using msgpack
func (o *ClusterIssuers) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeClusterIssuers deserializes a ClusterIssuers using msgpack
func MDecodeClusterIssuers(d *msgpack.Decoder) (ClusterIssuers, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ClusterIssuerKey, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]ClusterIssuerKey, 0), err
	}
	val := make([]ClusterIssuerKey, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeClusterIssuerKey(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a ClusterIssuers using cbor
func (o *ClusterIssuers) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeClusterIssuers deserializes a ClusterIssuers using cbor
func CDecodeClusterIssuers(d *cbor.Decoder) (ClusterIssuers, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]ClusterIssuerKey, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]ClusterIssuerKey, 0), err
	}
	val := make([]ClusterIssuerKey, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeClusterIssuerKey(d)
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

// MEncode serializes a HealthCheckRequest using msgpack
func (o *HealthCheckRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// MDecodeHealthCheckRequest deserializes a HealthCheckRequest using msgpack
func MDecodeHealthCheckRequest(d *msgpack.Decoder) (HealthCheckRequest, error) {
	var val HealthCheckRequest
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

// CEncode serializes a HealthCheckRequest using cbor
func (o *HealthCheckRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(0)

	return encoder.CheckError()
}

// CDecodeHealthCheckRequest deserializes a HealthCheckRequest using cbor
func CDecodeHealthCheckRequest(d *cbor.Decoder) (HealthCheckRequest, error) {
	var val HealthCheckRequest
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

// Return value from actors and providers for health check status
type HealthCheckResponse struct {
	// A flag that indicates the the actor is healthy
	Healthy bool
	// A message containing additional information about the actors health
	Message string
}

// MEncode serializes a HealthCheckResponse using msgpack
func (o *HealthCheckResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("healthy")
	encoder.WriteBool(o.Healthy)
	encoder.WriteString("message")
	encoder.WriteString(o.Message)

	return encoder.CheckError()
}

// MDecodeHealthCheckResponse deserializes a HealthCheckResponse using msgpack
func MDecodeHealthCheckResponse(d *msgpack.Decoder) (HealthCheckResponse, error) {
	var val HealthCheckResponse
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
		case "healthy":
			val.Healthy, err = d.ReadBool()
		case "message":
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

// CEncode serializes a HealthCheckResponse using cbor
func (o *HealthCheckResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("healthy")
	encoder.WriteBool(o.Healthy)
	encoder.WriteString("message")
	encoder.WriteString(o.Message)

	return encoder.CheckError()
}

// CDecodeHealthCheckResponse deserializes a HealthCheckResponse using cbor
func CDecodeHealthCheckResponse(d *cbor.Decoder) (HealthCheckResponse, error) {
	var val HealthCheckResponse
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
		case "healthy":
			val.Healthy, err = d.ReadBool()
		case "message":
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
	// True if structured logging is enabled for the host. Providers should use the same setting as the host.
	StructuredLogging bool
}

// MEncode serializes a HostData using msgpack
func (o *HostData) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(15)
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
	o.EnvValues.MEncode(encoder)
	encoder.WriteString("instance_id")
	encoder.WriteString(o.InstanceId)
	encoder.WriteString("link_definitions")
	o.LinkDefinitions.MEncode(encoder)
	encoder.WriteString("cluster_issuers")
	o.ClusterIssuers.MEncode(encoder)
	encoder.WriteString("config_json")
	encoder.WriteString(o.ConfigJson)
	encoder.WriteString("default_rpc_timeout_ms")
	encoder.WriteUint64(o.DefaultRpcTimeoutMs)
	encoder.WriteString("structured_logging")
	encoder.WriteBool(o.StructuredLogging)

	return encoder.CheckError()
}

// MDecodeHostData deserializes a HostData using msgpack
func MDecodeHostData(d *msgpack.Decoder) (HostData, error) {
	var val HostData
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
			val.EnvValues, err = MDecodeHostEnvValues(d)
		case "instance_id":
			val.InstanceId, err = d.ReadString()
		case "link_definitions":
			val.LinkDefinitions, err = MDecodeActorLinks(d)
		case "cluster_issuers":
			val.ClusterIssuers, err = MDecodeClusterIssuers(d)
		case "config_json":
			val.ConfigJson, err = d.ReadString()
		case "default_rpc_timeout_ms":
			val.DefaultRpcTimeoutMs, err = d.ReadUint64()
		case "structured_logging":
			val.StructuredLogging, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a HostData using cbor
func (o *HostData) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(15)
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
	o.EnvValues.CEncode(encoder)
	encoder.WriteString("instance_id")
	encoder.WriteString(o.InstanceId)
	encoder.WriteString("link_definitions")
	o.LinkDefinitions.CEncode(encoder)
	encoder.WriteString("cluster_issuers")
	o.ClusterIssuers.CEncode(encoder)
	encoder.WriteString("config_json")
	encoder.WriteString(o.ConfigJson)
	encoder.WriteString("default_rpc_timeout_ms")
	encoder.WriteUint64(o.DefaultRpcTimeoutMs)
	encoder.WriteString("structured_logging")
	encoder.WriteBool(o.StructuredLogging)

	return encoder.CheckError()
}

// CDecodeHostData deserializes a HostData using cbor
func CDecodeHostData(d *cbor.Decoder) (HostData, error) {
	var val HostData
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
			val.EnvValues, err = CDecodeHostEnvValues(d)
		case "instance_id":
			val.InstanceId, err = d.ReadString()
		case "link_definitions":
			val.LinkDefinitions, err = CDecodeActorLinks(d)
		case "cluster_issuers":
			val.ClusterIssuers, err = CDecodeClusterIssuers(d)
		case "config_json":
			val.ConfigJson, err = d.ReadString()
		case "default_rpc_timeout_ms":
			val.DefaultRpcTimeoutMs, err = d.ReadUint64()
		case "structured_logging":
			val.StructuredLogging, err = d.ReadBool()
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

// MEncode serializes a HostEnvValues using msgpack
func (o *HostEnvValues) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// MDecodeHostEnvValues deserializes a HostEnvValues using msgpack
func MDecodeHostEnvValues(d *msgpack.Decoder) (HostEnvValues, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

// CEncode serializes a HostEnvValues using cbor
func (o *HostEnvValues) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// CDecodeHostEnvValues deserializes a HostEnvValues using cbor
func CDecodeHostEnvValues(d *cbor.Decoder) (HostEnvValues, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
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
	TraceContext *TraceContext
}

// MEncode serializes a Invocation using msgpack
func (o *Invocation) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(9)
	encoder.WriteString("origin")
	o.Origin.MEncode(encoder)
	encoder.WriteString("target")
	o.Target.MEncode(encoder)
	encoder.WriteString("operation")
	encoder.WriteString(o.Operation)
	encoder.WriteString("msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("id")
	encoder.WriteString(o.Id)
	encoder.WriteString("encoded_claims")
	encoder.WriteString(o.EncodedClaims)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("traceContext")
	if o.TraceContext == nil {
		encoder.WriteNil()
	} else {
		o.TraceContext.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeInvocation deserializes a Invocation using msgpack
func MDecodeInvocation(d *msgpack.Decoder) (Invocation, error) {
	var val Invocation
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
		case "origin":
			val.Origin, err = MDecodeWasmCloudEntity(d)
		case "target":
			val.Target, err = MDecodeWasmCloudEntity(d)
		case "operation":
			val.Operation, err = d.ReadString()
		case "msg":
			val.Msg, err = d.ReadByteArray()
		case "id":
			val.Id, err = d.ReadString()
		case "encoded_claims":
			val.EncodedClaims, err = d.ReadString()
		case "host_id":
			val.HostId, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		case "traceContext":
			fval, err := MDecodeTraceContext(d)
			if err != nil {
				return val, err
			}
			val.TraceContext = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Invocation using cbor
func (o *Invocation) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(9)
	encoder.WriteString("origin")
	o.Origin.CEncode(encoder)
	encoder.WriteString("target")
	o.Target.CEncode(encoder)
	encoder.WriteString("operation")
	encoder.WriteString(o.Operation)
	encoder.WriteString("msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("id")
	encoder.WriteString(o.Id)
	encoder.WriteString("encoded_claims")
	encoder.WriteString(o.EncodedClaims)
	encoder.WriteString("host_id")
	encoder.WriteString(o.HostId)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)
	encoder.WriteString("traceContext")
	if o.TraceContext == nil {
		encoder.WriteNil()
	} else {
		o.TraceContext.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeInvocation deserializes a Invocation using cbor
func CDecodeInvocation(d *cbor.Decoder) (Invocation, error) {
	var val Invocation
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
		case "origin":
			val.Origin, err = CDecodeWasmCloudEntity(d)
		case "target":
			val.Target, err = CDecodeWasmCloudEntity(d)
		case "operation":
			val.Operation, err = d.ReadString()
		case "msg":
			val.Msg, err = d.ReadByteArray()
		case "id":
			val.Id, err = d.ReadString()
		case "encoded_claims":
			val.EncodedClaims, err = d.ReadString()
		case "host_id":
			val.HostId, err = d.ReadString()
		case "content_length":
			val.ContentLength, err = d.ReadUint64()
		case "traceContext":
			fval, err := CDecodeTraceContext(d)
			if err != nil {
				return val, err
			}
			val.TraceContext = &fval
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

// MEncode serializes a InvocationResponse using msgpack
func (o *InvocationResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("invocation_id")
	encoder.WriteString(o.InvocationId)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)

	return encoder.CheckError()
}

// MDecodeInvocationResponse deserializes a InvocationResponse using msgpack
func MDecodeInvocationResponse(d *msgpack.Decoder) (InvocationResponse, error) {
	var val InvocationResponse
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
		case "msg":
			val.Msg, err = d.ReadByteArray()
		case "invocation_id":
			val.InvocationId, err = d.ReadString()
		case "error":
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

// CEncode serializes a InvocationResponse using cbor
func (o *InvocationResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("msg")
	encoder.WriteByteArray(o.Msg)
	encoder.WriteString("invocation_id")
	encoder.WriteString(o.InvocationId)
	encoder.WriteString("error")
	encoder.WriteString(o.Error)
	encoder.WriteString("content_length")
	encoder.WriteUint64(o.ContentLength)

	return encoder.CheckError()
}

// CDecodeInvocationResponse deserializes a InvocationResponse using cbor
func CDecodeInvocationResponse(d *cbor.Decoder) (InvocationResponse, error) {
	var val InvocationResponse
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
		case "msg":
			val.Msg, err = d.ReadByteArray()
		case "invocation_id":
			val.InvocationId, err = d.ReadString()
		case "error":
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

// MEncode serializes a LinkDefinition using msgpack
func (o *LinkDefinition) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("actor_id")
	encoder.WriteString(o.ActorId)
	encoder.WriteString("provider_id")
	encoder.WriteString(o.ProviderId)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	encoder.WriteString(o.ContractId)
	encoder.WriteString("values")
	o.Values.MEncode(encoder)

	return encoder.CheckError()
}

// MDecodeLinkDefinition deserializes a LinkDefinition using msgpack
func MDecodeLinkDefinition(d *msgpack.Decoder) (LinkDefinition, error) {
	var val LinkDefinition
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
		case "actor_id":
			val.ActorId, err = d.ReadString()
		case "provider_id":
			val.ProviderId, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = d.ReadString()
		case "values":
			val.Values, err = MDecodeLinkSettings(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a LinkDefinition using cbor
func (o *LinkDefinition) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("actor_id")
	encoder.WriteString(o.ActorId)
	encoder.WriteString("provider_id")
	encoder.WriteString(o.ProviderId)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	encoder.WriteString(o.ContractId)
	encoder.WriteString("values")
	o.Values.CEncode(encoder)

	return encoder.CheckError()
}

// CDecodeLinkDefinition deserializes a LinkDefinition using cbor
func CDecodeLinkDefinition(d *cbor.Decoder) (LinkDefinition, error) {
	var val LinkDefinition
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
		case "actor_id":
			val.ActorId, err = d.ReadString()
		case "provider_id":
			val.ProviderId, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = d.ReadString()
		case "values":
			val.Values, err = CDecodeLinkSettings(d)
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

// MEncode serializes a LinkSettings using msgpack
func (o *LinkSettings) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// MDecodeLinkSettings deserializes a LinkSettings using msgpack
func MDecodeLinkSettings(d *msgpack.Decoder) (LinkSettings, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

// CEncode serializes a LinkSettings using cbor
func (o *LinkSettings) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// CDecodeLinkSettings deserializes a LinkSettings using cbor
func CDecodeLinkSettings(d *cbor.Decoder) (LinkSettings, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
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

// MEncode serializes a TraceContext using msgpack
func (o *TraceContext) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// MDecodeTraceContext deserializes a TraceContext using msgpack
func MDecodeTraceContext(d *msgpack.Decoder) (TraceContext, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
		v, err := d.ReadString()
		if err != nil {
			return val, err
		}
		val[k] = v
	}
	return val, nil
}

// CEncode serializes a TraceContext using cbor
func (o *TraceContext) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(uint32(len(*o)))
	for key_o, val_o := range *o {
		encoder.WriteString(key_o)
		encoder.WriteString(val_o)
	}

	return encoder.CheckError()
}

// CDecodeTraceContext deserializes a TraceContext using cbor
func CDecodeTraceContext(d *cbor.Decoder) (TraceContext, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make(map[string]string, 0), err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return make(map[string]string, 0), err
	}
	val := make(map[string]string, size)
	for i := uint32(0); i < size; i++ {
		k, _ := d.ReadString()
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

// MEncode serializes a WasmCloudEntity using msgpack
func (o *WasmCloudEntity) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("public_key")
	encoder.WriteString(o.PublicKey)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	o.ContractId.MEncode(encoder)

	return encoder.CheckError()
}

// MDecodeWasmCloudEntity deserializes a WasmCloudEntity using msgpack
func MDecodeWasmCloudEntity(d *msgpack.Decoder) (WasmCloudEntity, error) {
	var val WasmCloudEntity
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
		case "public_key":
			val.PublicKey, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = MDecodeCapabilityContractId(d)
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a WasmCloudEntity using cbor
func (o *WasmCloudEntity) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("public_key")
	encoder.WriteString(o.PublicKey)
	encoder.WriteString("link_name")
	encoder.WriteString(o.LinkName)
	encoder.WriteString("contract_id")
	o.ContractId.CEncode(encoder)

	return encoder.CheckError()
}

// CDecodeWasmCloudEntity deserializes a WasmCloudEntity using cbor
func CDecodeWasmCloudEntity(d *cbor.Decoder) (WasmCloudEntity, error) {
	var val WasmCloudEntity
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
		case "public_key":
			val.PublicKey, err = d.ReadString()
		case "link_name":
			val.LinkName, err = d.ReadString()
		case "contract_id":
			val.ContractId, err = CDecodeCapabilityContractId(d)
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
			value, err_ := MDecodeHealthCheckRequest(&d)
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

// Perform health check. Called at regular intervals by host
func (s *ActorSender) HealthRequest(ctx *Context, arg HealthCheckRequest) (*HealthCheckResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, Message{Method: "Actor.HealthRequest", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeHealthCheckResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.5.0
