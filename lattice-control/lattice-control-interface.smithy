// control-interface.smithy
//
// Lattice control interface
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [
    {
        namespace: "org.wasmcloud.lattice.control",
        crate: "wasmcloud_interface_lattice_control"
     }
]

namespace org.wasmcloud.lattice.control

use org.wasmcloud.model#wasmbus
use org.wasmcloud.core#LinkDefinition
use org.wasmcloud.core#ActorLinks
use org.wasmcloud.model#I32
use org.wasmcloud.model#U16
use org.wasmcloud.model#U64

@length(min:1, max:500000)
string ConfigurationString

/// Lattice Controller - Describes the interface used for actors
/// to communicate with a lattice controller, enabling developers
/// to deploy actors that can manipulate the lattice in which they're
/// running.
@wasmbus(
    contractId: "wasmcloud:latticecontrol",
    providerReceive: true )
service LatticeController {
    version: "0.1",
    operations: [AuctionProvider, AuctionActor, GetHosts, 
                 GetHostInventory, GetClaims, ScaleActor,
                 StartActor, AdvertiseLink, RemoveLink,
                 GetLinks, UpdateActor, StartProvider,
                 StopProvider, StopActor, StopHost, 
                 SetLatticeCredentials,
                 SetRegistryCredentials]
}

/// Seek out a list of suitable hosts for a capability provider given
/// a set of host label constraints. Hosts on which this provider is already
/// running will not be among the successful "bidders" in this auction.
operation AuctionProvider {
    input: ProviderAuctionRequest,
    output: ProviderAuctionAcks
}

/// Seek out a list of suitable hosts for an actor given a set of host
/// label constraints.
operation AuctionActor {
    input: ActorAuctionRequest,
    output: ActorAuctionAcks
}

/// Queries the list of hosts currently visible to the lattice. This is
/// a "gather" operation and so can be influenced by short timeouts,
/// network partition events, etc. The sole input to this query is the 
/// lattice ID on which the request takes place.
operation GetHosts {
    input: String,
    output: Hosts
}

/// Queries for the contents of a host given the supplied 56-character unique ID
operation GetHostInventory {
    input: GetHostInventoryRequest,
    output: HostInventory
}

/// Queries the lattice for the list of known/cached claims by taking the response
/// from the first host that answers the query. The sole input to this request is
/// the lattice ID on which the request takes place.
operation GetClaims {    
    input: String,
    output: GetClaimsResponse
}

/// Publish a link definition into the lattice, allowing it to be cached and
/// delivered to the appropriate capability provider instances
operation AdvertiseLink {
    input: AdvertiseLinkRequest,
    output: CtlOperationAck
}

/// Requests the removal of a link definition. The definition will be removed
/// from the cache and the relevant capability providers will be given a chance
/// to de-provision any used resources
operation RemoveLink {
    input: RemoveLinkDefinitionRequest,
    output: CtlOperationAck
}


/// Instructs a given host to start the indicated actor
operation StartActor {
    input: StartActorCommand,
    output: CtlOperationAck
}

/// Instructs a given host to scale the indicated actor
operation ScaleActor {
    input: ScaleActorCommand,
    output: CtlOperationAck
}

/// Requests that a specific host perform a live update on the indicated
/// actor
operation UpdateActor {
    input: UpdateActorCommand,
    output: CtlOperationAck
}

/// Queries all current link definitions in the specified lattice. The first host
/// that receives this response will reply with the contents of the distributed
/// cache
operation GetLinks {
    input: String,
    output: LinkDefinitionList
}

/// Requests that the given host start the indicated capability provider
operation StartProvider {
    input: StartProviderCommand,
    output: CtlOperationAck
}

/// Requests that the given capability provider be stopped on the indicated host
operation StopProvider {
    input: StopProviderCommand,
    output: CtlOperationAck
}

/// Requests that an actor be stopped on the given host
operation StopActor {
    input: StopActorCommand,
    output: CtlOperationAck
}

operation StopHost {
    input: StopHostCommand,
    output: CtlOperationAck
}

/// Instructs all listening hosts to use the enclosed credential map for
/// authentication to secure artifact (OCI/bindle) registries. Any host that
/// receives this message will _delete_ its previous credential map and replace
/// it with the enclosed. The credential map for a lattice can be purged by sending
/// this message with an empty map
operation SetRegistryCredentials {
    input: SetRegistryCredentialsRequest
}

/// Instructs the provider to store the NATS credentials/URL for a given lattice. This is
/// designed to allow a single capability provider (or multiple instances of the same) to manage
/// multiple lattices, reducing overhead and making it easier to support secure multi-tenancy of
/// lattices.
operation SetLatticeCredentials {
    input: SetLatticeCredentialsRequest,
    output: CtlOperationAck
}

/// A request to advertise/publish a link definition on a given lattice.
structure AdvertiseLinkRequest {
    /// The ID of the lattice for this request
    @required    
    latticeId: String,

    @required
    link: LinkDefinition
}

/// A request to obtain claims from a given lattice
structure GetClaimsRequest {
    /// The ID of the lattice for this request
    @required    
    latticeId: String,
}

/// A request to query the inventory of a given host within a given lattice
structure GetHostInventoryRequest {
    /// The ID of the lattice for this request
    @required    
    latticeId: String,

    /// The public key of the host being targeted for this request
    @required    
    hostId: String
}

/// A request to obtain the list of hosts responding within a given lattice
structure GetHostsRequest {
    /// The ID of the lattice for which these credentials will be used
    @required    
    latticeId: String,
}

/// Represents a request to set/store the credentials that correspond to a given lattice ID. 
structure SetLatticeCredentialsRequest {
    /// The ID of the lattice for which these credentials will be used
    @required    
    latticeId: String,

    /// If supplied, contains the user JWT to be used for authenticating against NATS to allow
    /// access to the indicated lattice. If not supplied, the capability provider will assume/set
    /// anonymous access for this lattice.    
    userJwt: String,

    /// If userJwt is supplied, user seed must also be supplied and is the seed key used for user
    /// authentication against NATS for this lattice.    
    userSeed: String,

    /// If natsUrl is supplied, then the capability provider will use this URL (and port) for 
    /// establishing a connection for the given lattice.    
    natsUrl: String,

    /// If there is a JS domain required for communicating with the underlying KV metadata
    /// bucket for this lattice, then that should be supplied in this parameter. Otherwise,
    /// leave it blank
    jsDomain: String
}

list ProviderAuctionAcks {
    member: ProviderAuctionAck
}

list ActorAuctionAcks {
    member: ActorAuctionAck
}

list Hosts {
    member: Host
}

/// A request to locate a suitable host for a capability provider. The
/// provider's unique identity (reference + link name) is used to rule
/// out sites on which the provider is already running.
structure ProviderAuctionRequest {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The reference for the provider. Can be any one of the accepted 
    /// forms of uniquely identifying a provider
    @required       
    providerRef: String,

    /// The link name of the provider
    @required      
    linkName: String,

    /// The set of constraints to which a suitable target host must conform
    @required    
    constraints: ConstraintMap,
}

map ConstraintMap {
    key: String,
    value: String,
}

map AnnotationMap {
    key: String,
    value: String
}

/// One of a potential list of responses to a provider auction
structure ProviderAuctionAck {
    /// The original provider ref provided for the auction
    @required        
    providerRef: String,

    /// The link name provided for the auction
    @required        
    linkName: String,

    /// The host ID of the "bidder" for this auction
    @required      
    hostId: String,
}

/// A request to locate suitable hosts for a given actor
structure ActorAuctionRequest {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The reference for this actor. Can be any one of the acceptable forms
    /// of uniquely identifying an actor.
    @required        
    actorRef: String,

    /// The set of constraints to which any candidate host must conform
    @required    
    constraints: ConstraintMap,
}

/// One of a potential list of responses to an actor auction
structure ActorAuctionAck {
    /// The original actor reference used for the auction
    @required        
    actorRef: String,    

    /// The host ID of the "bidder" for this auction.
    @required        
    hostId: String,
}

/// Describes the known contents of a given host at the time of
/// a query
structure HostInventory {    

    /// The host's unique ID
    @required        
    hostId: String,

    /// The host's labels
    @required    
    labels: LabelsMap,

    /// Actors running on this host.
    @required    
    actors: ActorDescriptions,

    /// Providers running on this host
    @required    
    providers: ProviderDescriptions,
}

map LabelsMap {
    key: String,
    value: String,
}

list ActorDescriptions {
    member: ActorDescription,
}

list ProviderDescriptions {
    member: ProviderDescription,
}

/// A summary description of an actor within a host inventory
structure ActorDescription {

    /// Actor's 56-character unique ID
    @required    
    id: String,    

    /// Image reference for this actor, if applicable    
    imageRef: String,

    /// Name of this actor, if one exists    
    name: String,

    /// The individual instances of this actor that are running
    @required    
    instances: ActorInstances
}

structure ActorInstance {
    /// This instance's unique ID (guid)
    @required    
    instanceId: String,

    /// The revision number for this actor instance
    @required    
    revision: I32,

    /// The annotations that were used in the start request that produced
    /// this actor instance    
    annotations: AnnotationMap
}

list ActorInstances {
    member: ActorInstance
}

/// A summary description of a capability provider within a host inventory
structure ProviderDescription {

    /// Provider's unique 56-character ID
    @required    
    id: String,

    /// Provider's link name
    @required        
    linkName: String,

    /// Image reference for this provider, if applicable    
    imageRef: String,

    /// Name of the provider, if one exists
    name: String,

    /// The revision of the provider
    @required
    revision: I32,

    /// The annotations that were used in the start request that produced
    /// this provider instance
    annotations: AnnotationMap
}


/// A command sent to a specific host instructing it to start the actor
/// indicated by the reference.
structure StartActorCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,    

    /// Reference for the actor. This can be either a bindle or OCI reference
    @required        
    actorRef: String,

    /// Host ID on which this actor should start
    @required    
    hostId: String,

    /// Optional set of annotations used to describe the nature of this actor start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment    
    annotations: AnnotationMap,

    /// The number of actors to start
    /// A zero value will be interpreted as 1.
    @required
    count: U16,
}

/// A command sent to a host requesting a capability provider be started with the 
/// given link name and optional configuration.
structure StartProviderCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The host ID on which to start the provider
    @required    
    hostId: String,

    /// The image reference of the provider to be started
    @required        
    providerRef: String,

    /// The link name of the provider to be started
    @required        
    linkName: String,

    /// Optional set of annotations used to describe the nature of this provider start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment    
    annotations: AnnotationMap,


    /// Optional provider configuration in the form of an opaque string. Many
    /// providers prefer base64-encoded JSON here, though that data should never
    /// exceed 500KB
    configuration: ConfigurationString
}

structure ScaleActorCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// Reference for the actor. Can be any of the acceptable forms of unique identification
    @required    
    actorRef: String,

    /// Public Key ID of the actor to scale
    @required    
    actorId: String,

    /// Host ID on which to scale this actor
    @required    
    hostId: String,

    /// Optional set of annotations used to describe the nature of this actor scale command. For
    /// example, autonomous agents may wish to "tag" scale requests as part of a given deployment
    annotations: AnnotationMap,

    /// The target number of actors
    @required
    count: U16,
}

/// A command sent to a host to request that instances of a given actor
/// be terminated on that host
structure StopActorCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The ID of the target host
    @required      
    hostId: String,

    /// The public key of the actor to stop
    @required        
    actorId: String,

    /// The number of actors to stop
    /// A zero value means stop all actors
    @required
    count: U16,

    /// Optional set of annotations used to describe the nature of this
    /// stop request. If supplied, the only instances of this actor with these
    /// annotations will be stopped
    annotations: AnnotationMap
}

/// A request to stop the given provider on the indicated host
structure StopProviderCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// Host ID on which to stop the provider
    @required        
    hostId: String,

    /// The public key of the capability provider to stop
    @required        
    providerId: String,

    /// Link name for this provider
    @required        
    linkName: String,

    /// Contract ID of the capability provider
    @required        
    contractId: String,

    /// Optional set of annotations used to describe the nature of this
    /// stop request
    annotations: AnnotationMap
}

/// A command sent to request that the given host purge and stop
structure StopHostCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The ID of the target host
    @required      
    hostId: String,

    /// An optional timeout, in seconds
    timeout: U64
}

/// A command instructing a specific host to perform a live update
/// on the indicated actor by supplying a new image reference. Note that
/// live updates are only possible through image references
structure UpdateActorCommand {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

    /// The host ID of the host to perform the live update
    @required       
    hostId: String,

    /// The actor's 56-character unique ID
    @required        
    actorId: String,

    /// The new image reference of the upgraded version of this actor
    @required        
    newActorRef: String,

    /// Optional set of annotations used to describe the nature of this
    /// update request. Only actor instances that have matching annotations 
    /// will be upgraded, allowing for instance isolation by 
    // autonomous agent deployment spec, for example.
    annotations: AnnotationMap
}

/// Standard response for control interface operations
structure CtlOperationAck {
    @required
    accepted: Boolean,
    @required
    error: String
}

/// A list of link definitions
structure LinkDefinitionList {
    @required
    links: ActorLinks
}

/// A summary representation of a host
structure Host {
    @required
    id: String,

    /// uptime in seconds
    @required        
    uptimeSeconds: U64,

    /// Human-friendly uptime description    
    uptimeHuman: String,

    /// Hash map of label-value pairs for this host
    labels: KeyValueMap,

    /// Current wasmCloud Host software version
    version: String,

    /// Comma-delimited list of valid cluster issuer public keys as known
    /// to this host    
    clusterIssuers: String,

    /// JetStream domain (if applicable) in use by this host    
    jsDomain: String,

    /// NATS server host used for the control interface    
    ctlHost: String,

    /// NATS server host used for provider RPC    
    provRpcHost: String,

    /// NATS server host used for regular RPC    
    rpcHost: String,

    /// Lattice prefix/ID used by the host    
    latticePrefix: String
}

/// A response containing the full list of known claims within the lattice
structure GetClaimsResponse {
    @required
    claims: CtlKVList
}

list CtlKVList {
    member: KeyValueMap,
}

map KeyValueMap {
    key: String,
    value: String,
}

/// A request to remove a link definition and detach the relevant actor
/// from the given provider
structure RemoveLinkDefinitionRequest {
    /// The ID of the lattice on which this request will be performed
    @required    
    latticeId: String,

     /// The actor's public key. This cannot be an image reference
    @required        
    actorId: String,

    /// The provider contract
    @required        
    contractId: String,

    /// The provider's link name
    @required        
    linkName: String,
}

structure SetRegistryCredentialsRequest {
    /// The ID of the lattice on which this request will be performed
    @required
    latticeId: String,

    credentials: RegistryCredentialMap
}

/// A set of credentials to be used for fetching from specific registries
map RegistryCredentialMap {
    /// The key of this map is the OCI/BINDLE URL without the artifact reference. Credentials
    /// are matched via substring comparison on the URL of an artifact.
    key: String,
    value: RegistryCredential
}

structure RegistryCredential {
    /// If supplied, token authentication will be used for the registry
    token: String,
    /// If supplied, username and password will be used for HTTP Basic authentication
    username: String,
    @sensitive
    password: String,
    /// The type of the registry (either "oci" or "bindle")
    @required
    registryType: String
}

