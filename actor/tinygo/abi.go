package actor

// HostCall performs a "host call" - a request from the actor to either invoke another actor
// or invoke an operation on a linked capability provider.
//
// binding - this is either the link name (e.g. "default") or, in the case of an actor-to-actor call, the public key or call alias of the target
//
// namespace - this is the contract ID when invoking a provider, empty string for actor-to-actor calls
//
// operation - the name of the operation to invoke
func HostCall(binding, namespace, operation string, payload []byte) ([]byte, error) {
	result := hostCall(
		stringToPointer(binding), uint32(len(binding)),
		stringToPointer(namespace), uint32(len(namespace)),
		stringToPointer(operation), uint32(len(operation)),
		bytesToPointer(payload), uint32(len(payload)),
	)
	if !result {
		errorLen := hostErrorLen()
		message := make([]byte, errorLen)
		hostError(bytesToPointer(message))

		return nil, &HostError{message: string(message)}
	}

	responseLen := hostResponseLen()
	response := make([]byte, responseLen)
	hostResponse(bytesToPointer(response))

	return response, nil
}

// HostError is returned by the host from HostCall
type HostError struct {
	message string
}

// Error implements the error interface for HostError
func (e *HostError) Error() string {
	return "Host error: " + e.message
}
