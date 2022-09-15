package actor

//go:wasm-module wasmbus
//go:export __guest_request
func guestRequest(operationPtr uintptr, payloadPtr uintptr)

//go:wasm-module wasmbus
//go:export __guest_response
func guestResponse(ptr uintptr, len uint32) //nolint

//go:wasm-module wasmbus
//go:export __guest_error
func guestError(ptr uintptr, len uint32)

//go:wasm-module wasmbus
//go:export __host_call
func hostCall(
	bindingPtr uintptr, bindingLen uint32,
	namespacePtr uintptr, namespaceLen uint32,
	operationPtr uintptr, operationLen uint32,
	payloadPtr uintptr, payloadLen uint32) bool

//go:wasm-module wasmbus
//go:export __host_response_len
func hostResponseLen() uint32

//go:wasm-module wasmbus
//go:export __host_response
func hostResponse(ptr uintptr)

//go:wasm-module wasmbus
//go:export __host_error_len
func hostErrorLen() uint32

//go:wasm-module wasmbus
//go:export __host_error
func hostError(ptr uintptr)

//go:wasm-module wasmbus
//go:export __console_log
func consoleLog(str uintptr, strLen uint32)
