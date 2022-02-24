// wasmcloud-core.smithy
// Core definitions for wasmcloud platform
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [ {
    namespace: "org.wasmcloud.core",
    crate: "wasmbus_rpc::core",
    py_module: "wasmbus_rpc.core",
} ]

namespace org.wasmcloud.core

use org.wasmcloud.model#nonEmptyString
use org.wasmcloud.model#codegenRust
use org.wasmcloud.model#serialization
use org.wasmcloud.model#CapabilityContractId
use org.wasmcloud.model#wasmbusData
use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#n
use org.wasmcloud.model#U64

/// Actor service
@wasmbus(
    actorReceive: true,
)
service Actor {
  version: "0.1",
  operations: [ HealthRequest ]
}

/// Link definition for binding actor to provider
@wasmbusData
structure LinkDefinition {
    /// actor public key
    @required
    @serialization(name:"actor_id")
    @n(0)
    actorId: String,

    /// provider public key
    @required
    @serialization(name:"provider_id")
    @n(1)
    providerId: String,

    /// link name
    @required
    @serialization(name:"link_name")
    @n(2)
    linkName: String,

    /// contract id
    @required
    @serialization(name:"contract_id")
    @n(3)
    contractId: String,

    @required
    @n(4)
    values: LinkSettings,
}


/// Return value from actors and providers for health check status
@wasmbusData
structure HealthCheckResponse {

  /// A flag that indicates the the actor is healthy
  @n(0)
  healthy: Boolean

  /// A message containing additional information about the actors health
  @n(1)
  message: String
}

/// health check request parameter
@wasmbusData
structure HealthCheckRequest { }

/// Perform health check. Called at regular intervals by host
operation HealthRequest {
    input: HealthCheckRequest
    output: HealthCheckResponse
}

/// Settings associated with an actor-provider link
map LinkSettings {
    key: String,
    value: String,
}

/// List of linked actors for a provider
list ActorLinks {
    member: LinkDefinition
}

/// initialization data for a capability provider
@wasmbusData
structure HostData {
    @required
    @serialization(name: "host_id")
    @n(0)
    hostId: String,

    @required
    @serialization(name: "lattice_rpc_prefix")
    @n(1)
    latticeRpcPrefix: String,

    @required
    @serialization(name: "link_name")
    @n(2)
    linkName: String,

    @required
    @serialization(name: "lattice_rpc_user_jwt")
    @n(3)
    latticeRpcUserJwt: String,

    @required
    @serialization(name: "lattice_rpc_user_seed")
    @n(4)
    latticeRpcUserSeed: String,

    @required
    @serialization(name: "lattice_rpc_url")
    @n(5)
    latticeRpcUrl: String,

    @required
    @serialization(name: "provider_key")
    @n(6)
    providerKey: String,

    @required
    @serialization(name: "invocation_seed")
    @n(7)
    invocationSeed: String,

    @required
    @serialization(name: "env_values")
    @n(8)
    envValues: HostEnvValues,

    @required
    @serialization(name: "instance_id")
    @n(9)
    instanceId: String,

    /// initial list of links for provider
    @required
    @serialization(name: "link_definitions")
    @n(10)
    linkDefinitions: ActorLinks,

    /// list of cluster issuers
    @required
    @serialization(name: "cluster_issuers")
    @n(11)
    clusterIssuers: ClusterIssuers,

    /// Optional configuration JSON sent to a given link name of a provider
    /// without an actor context
    @serialization(name:"config_json")
    @n(12)
    configJson: String
}

list ClusterIssuers {
    member: ClusterIssuerKey,
}

@nonEmptyString
string ClusterIssuerKey

/// Environment settings for initializing a capability provider
map HostEnvValues {
    key: String,
    value: String,
}

/// RPC message to capability provider
@wasmbusData
@codegenRust( nonExhaustive: true )
structure Invocation {
    @required
    @n(0)
    origin: WasmCloudEntity,

    @required
    @n(1)
    target: WasmCloudEntity,

    @required
    @n(2)
    operation: String,

    @required
    @n(3)
    msg: Blob,

    @required
    @n(4)
    id: String,

    @required
    @serialization(name: "encoded_claims")
    @n(5)
    encodedClaims: String,

    @required
    @serialization(name: "host_id")
    @n(6)
    hostId: String,

    /// total message size (optional)
    @n(7)
    @serialization(name: "content_length")
    contentLength: U64,
}

@wasmbusData
structure WasmCloudEntity {

    @required
    @serialization(name: "public_key")
    @n(0)
    publicKey: String,

    @required
    @serialization(name: "link_name")
    @n(1)
    linkName: String,

    @required
    @serialization(name: "contract_id")
    @n(2)
    contractId: CapabilityContractId,
}

/// Response to an invocation
@wasmbusData
structure InvocationResponse {

    /// serialize response message
    @required
    @n(0)
    msg: Blob,

    /// id connecting this response to the invocation
    @required
    @serialization(name: "invocation_id")
    @n(1)
    invocationId: String,

    /// optional error message
    @n(2)
    error: String,
}

