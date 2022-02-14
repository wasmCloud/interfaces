# wasmCloud Configuration Service Interface
The wasmCloud Configuration Service is an abstract contract between wasmCloud hosts and a _configuration provider_. This isn't a traditional contract where one side is satisfied by a capability provider. In this case, anything can act as a configuration service so long as it meets the following requirements:
* Listens on a NATS topic with a known prefix and the suffixes `.req` and `.push`
* Replies to requests by comparing the labels and other constraints in the request to stored configuration

This interface isn't used to generate a reusable crate, but rather just provides smithy-based types that allow for localized code generation for use in actors (which can be where the configuration provider is implemented).

It is assumed that this configuration will be transmitted in JSON format