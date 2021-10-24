// messaging.smithy
//
// Interface for messaging provider
// supports publish, request-reply, and subscriptions
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [{
    namespace: "org.wasmcloud.interface.messaging",
    crate: "wasmcloud_interface_messaging",
    py_module: "wasmcloud_interface_messaging",
}]

namespace org.wasmcloud.interface.messaging

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#n
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64

/// The Messaging interface describes a service
/// that can deliver messages
@wasmbus(
    contractId: "wasmcloud:messaging",
    providerReceive: true )
service Messaging {
  version: "0.1",
  operations: [ Publish, Request ]
}

/// The MessageSubscriber interface describes
/// an actor interface that receives messages
/// sent by the Messaging provider
@wasmbus(
    contractId: "wasmcloud:messaging",
    actorReceive: true )
service MessageSubscriber {
  version: "0.1",
  operations: [ HandleMessage ]
}

/// Publish - send a message
/// The function returns after the message has been sent.
/// If the sender expects to receive an asynchronous reply,
/// the replyTo field should be filled with the
/// subject for the response.
operation Publish {
    input: PubMessage
}

/// A message to be published
structure PubMessage {
    /// The subject, or topic, of the message
    @required
    @n(0)
    subject: String,

    /// An optional topic on which the reply should be sent.
    @n(1)
    replyTo: String,

    /// The message payload
    @required
    @n(2)
    body: Blob,
}

/// Reply received from a Request operation
structure ReplyMessage {
    /// The subject, or topic, of the message
    @required
    @n(0)
    subject: String,

    /// An optional topic on which the reply should be sent.
    @n(1)
    replyTo: String,

    /// The message payload
    @required
    @n(2)
    body: Blob,
}

/// Message received as part of a subscription
structure SubMessage {
    /// The subject, or topic, of the message
    @required
    @n(0)
    subject: String,

    /// An optional topic on which the reply should be sent.
    @n(1)
    replyTo: String,

    /// The message payload
    @required
    @n(2)
    body: Blob,
}

/// Request - send a message in a request/reply pattern,
/// waiting for a response.
operation Request {
    input: RequestMessage
    output: ReplyMessage
}

/// Message sent as part of a request, with timeout
structure RequestMessage {

    /// The subject, or topic, of the message
    @required
    @n(0)
    subject: String,

    /// The message payload
    @required
    @n(1)
    body: Blob,

    /// A timeout, in milliseconds
    @required
    @n(2)
    timeoutMs: u32,
}

/// subscription handler
operation HandleMessage {
    input: SubMessage
}