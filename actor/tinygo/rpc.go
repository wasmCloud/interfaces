package actor

// ToProvider constructs a Transport for actor-to-provider calls
func ToProvider(contractId string, linkName string) Transport {
	return Transport{
		Binding:   linkName,
		Namespace: contractId,
	}
}

// ToActor constructs a Transport for actor-to-actor calls
func ToActor(actor_id string) Transport {
	return Transport{
		Binding:   "",
		Namespace: actor_id,
	}
}

// Send sends the rpc Message using a Transport
func (t *Transport) Send(ctx *Context, msg Message) ([]byte, error) {
	r, ok := HostCall(t.Binding, t.Namespace, msg.Method, msg.Arg)
	return r, ok
}

// RpcError is an error type emitted by the rpc infrastructure
type RpcError struct {
	kind    string
	message string
}

// NewRpcError constructs an RpcError
func NewRpcError(kind string, message string) *RpcError {
	return &RpcError{kind: kind, message: message}
}

// Error is RpcError's implementation of the error interface
func (e *RpcError) Error() string {
	return e.message
}
