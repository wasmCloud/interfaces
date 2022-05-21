// logging.smithy
//
// built-in capability provider for logging
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [{
    namespace: "org.wasmcloud.interface.logging",
    crate: "wasmcloud_interface_logging",
    py_module: "wasmcloud_interface_logging",
    doc: "Logging: wasmcloud built-in logging capability provider",
}]

namespace org.wasmcloud.interface.logging

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#n
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64

@wasmbus(
    contractId: "wasmcloud:builtin:logging",
    providerReceive: true )
service Logging {
  version: "0.1",
  operations: [ WriteLog ]
}

///
/// WriteLog - log a text message
///
operation WriteLog {
    input: LogEntry
}

structure LogEntry {
    /// severity level: debug,info,warn,error
    @required
    @n(0)
    level: String,

    /// message to log
    @required
    @n(1)
    text: String,
}
