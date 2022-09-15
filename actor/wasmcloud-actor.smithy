metadata package = [ {
    namespace: "org.wasmcloud.actor",
    crate: "wasmbus_rpc::actor",
    py_module: "wasmbus_rpc.actor",
    doc: "wasmcloud platform core actor",
} ]

namespace org.wasmcloud.actor

use org.wasmcloud.model#wasmbus
use org.wasmcloud.core#HealthCheckRequest
use org.wasmcloud.core#HealthCheckResponse
use org.wasmcloud.model#I64
use org.wasmcloud.model#U32

/// Actor service
@wasmbus(
    actorReceive: true,
)
service Actor {
  version: "0.1",
  operations: [ HealthRequest ]
}

operation HealthRequest {
    input: HealthCheckRequest
    output: HealthCheckResponse
}

structure Context {}

structure Document {}

structure Timestamp {
    Sec: I64,
    Nsec: U32,
}

structure Transport {
    binding: String,
    namespace: String,
}

structure Message {
    Method: String,
    Arg: Blob,
}
