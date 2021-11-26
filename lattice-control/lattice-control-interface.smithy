// control-interface.smithy
//
// Lattice control interface
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [
    {
        namespace: "org.wasmcloud.lattice.control",
        crate: "lattice-control-interface"
     }
]

namespace org.wasmcloud.lattice.control

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#serialization
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
    version: "0.1"
    operations: [AuctionProvider, AuctionActor, GetHosts, 
                 GetHostInventory, GetClaims, StartActor,
                 AdvertiseLink, RemoveLink, GetLinks,
                 UpdateActor, StartProvider, StopProvider,
                 StopActor, StopHost]
}

/// Seek out a list of suitable hosts for a capability provider given
/// a set of host label constraints. Hosts on which this provider is already
/// running will not be among the successful "bidders" in this auction.
operation AuctionProvider {
    input: ProviderAuctionRequest
    output: ProviderAuctionAcks
}

/// Seek out a list of suitable hosts for an actor given a set of host
/// label constraints.
operation AuctionActor {
    input: ActorAuctionRequest
    output: ActorAuctionAcks
}

/// Queries the list of hosts currently visible to the lattice. This is
/// a "gather" operation and so can be influenced by short timeouts,
/// network partition events, etc.
operation GetHosts {
    output: Hosts
}

/// Queries for the contents of a host given the supplied 56-character unique ID
operation GetHostInventory {
    input: String
    output: HostInventory
}

/// Queries the lattice for the list of known/cached claims by taking the response
/// from the first host that answers the query.
operation GetClaims {    
    output: GetClaimsResponse
}

/// Publish a link definition into the lattice, allowing it to be cached and
/// delivered to the appropriate capability provider instances
operation AdvertiseLink {
    input: LinkDefinition
    output: CtlOperationAck
}

/// Requests the removal of a link definition. The definition will be removed
/// from the cache and the relevant capability providers will be given a chance
/// to de-provision any used resources
operation RemoveLink {
    input: RemoveLinkDefinitionRequest
    output: CtlOperationAck
}


/// Instructs a given host to start the indicated actor
operation StartActor {
    input: StartActorCommand
    output: CtlOperationAck
}

/// Requests that a specific host perform a live update on the indicated
/// actor
operation UpdateActor {
    input: UpdateActorCommand
    output: CtlOperationAck
}

/// Queries all current link definitions in the lattice. The first host
/// that receives this response will reply with the contents of the distributed
/// cache
operation GetLinks {
    output: LinkDefinitionList
}

/// Requests that the given host start the indicated capability provider
operation StartProvider {
    input: StartProviderCommand
    output: CtlOperationAck
}

/// Requests that the given capability provider be stopped on the indicated host
operation StopProvider {
    input: StopProviderCommand
    output: CtlOperationAck
}

/// Requests that an actor be stopped on the given host
operation StopActor {
    input: StopActorCommand
    output: CtlOperationAck    
}

/// Requests that the given host be stopped
operation StopHost {
    input: StopHostCommand
    output: CtlOperationAck
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

    /// The reference for the provider. Can be any one of the accepted 
    /// forms of uniquely identifying a provider
    @required   
    @serialization(name: "provider_ref") 
    providerRef: String,

    /// The link name of the provider
    @required  
    @serialization(name: "link_name")  
    linkName: String,

    /// The set of constraints to which a suitable target host must conform
    @required
    @serialization(name:"constraints")
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
    @serialization(name: "provider_ref")
    providerRef: String,

    /// The link name provided for the auction
    @required    
    @serialization(name: "link_name")
    linkName: String,

    /// The host ID of the "bidder" for this auction
    @required  
    @serialization(name: "host_id")  
    hostId: String,
}

/// A request to locate suitable hosts for a given actor
structure ActorAuctionRequest {
    /// The reference for this actor. Can be any one of the acceptable forms
    /// of uniquely identifying an actor.
    @required    
    @serialization(name: "actor_ref")
    actorRef: String,

    /// The set of constraints to which any candidate host must conform
    @required
    @serialization(name: "constraints")
    constraints: ConstraintMap,
}

/// One of a potential list of responses to an actor auction
structure ActorAuctionAck {
    /// The original actor reference used for the auction
    @required    
    @serialization(name: "actor_ref")
    actorRef: String,    

    /// The host ID of the "bidder" for this auction.
    @required    
    @serialization(name: "host_id")
    hostId: String,
}

/// Describes the known contents of a given host at the time of
/// a query
structure HostInventory {

    /// The host's unique ID
    @required    
    @serialization(name: "host_id")
    hostId: String,

    /// The host's labels
    @required
    @serialization(name: "labels")
    labels: LabelsMap,

    /// Actors running on this host.
    @required
    @serialization(name: "actors")
    actors: ActorDescriptions,

    /// Providers running on this host
    @required
    @serialization(name: "providers")
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
    @serialization(name: "id")
    id: String,    

    /// Image reference for this actor, if applicable
    @serialization(name: "image_ref")
    imageRef: String,

    /// Name of this actor, if one exists
    @serialization(name: "name")
    name: String,

    /// The individual instances of this actor that are running
    @required
    @serialization(name: "instances")
    instances: ActorInstances
}

structure ActorInstance {
    /// This instance's unique ID (guid)
    @required
    @serialization(name: "instance_id")
    instanceId: String

    /// The revision number for this actor instance
    @required
    @serialization(name: "revision")
    revision: I32,

    /// The annotations that were used in the start request that produced
    /// this actor instance
    @serialization(name: "annotations")
    annotations: AnnotationMap
}

list ActorInstances {
    member: ActorInstance
}

/// A summary description of a capability provider within a host inventory
structure ProviderDescription {

    /// Provider's unique 56-character ID
    @required
    @serialization(name: "id")
    id: String,

    /// Provider's link name
    @required    
    @serialization(name: "link_name")
    linkName: String,

    /// Image reference for this provider, if applicable
    @serialization(name: "image_ref")
    imageRef: String,

    /// Name of the provider, if one exists
    name: String,

    /// The revision of the provider
    @required
    revision: I32,
}


/// A command sent to a specific host instructing it to start the actor
/// indicated by the reference.
structure StartActorCommand {
    /// Reference for the actor. Can be any of the acceptable forms of unique identification
    @required    
    @serialization(name: "actor_ref")
    actorRef: String,

    /// Host ID on which this actor should start
    @required
    @serialization(name: "host_id")
    hostId: String,

    /// Optional set of annotations used to describe the nature of this actor start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment    
    annotations: AnnotationMap
}

/// A command sent to a host requesting a capability provider be started with the 
/// given link name and optional configuration.
structure StartProviderCommand {
    /// The host ID on which to start the provider
    @required
    @serialization(name: "host_id")
    hostId: String,

    /// The image reference of the provider to be started
    @required    
    @serialization(name: "provider_ref")
    providerRef: String,

    /// The link name of the provider to be started
    @required    
    @serialization(name: "link_name")
    linkName: String,

    /// Optional set of annotations used to describe the nature of this provider start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment    
    annotations: AnnotationMap


    /// Optional provider configuration in the form of an opaque string. Many
    /// providers prefer base64-encoded JSON here, though that data should never
    /// exceed 500KB
    configuration: ConfigurationString
}

/// A command sent to a host to request that instances of a given actor
/// be terminated on that host
structure StopActorCommand {
    /// The ID of the target host
    @required  
    @serialization(name: "host_id")  
    hostId: String,

    /// Reference for this actor. Can be any of the means of uniquely identifying
    /// an actor
    @required    
    @serialization(name:"actor_ref")
    actorRef: String,

    /// Optional count. If 0, all instances of this actor will be terminated
    count: U16,

    /// Optional set of annotations used to describe the nature of this
    /// stop request. If supplied, the only instances of this actor with these
    /// annotations will be stopped
    annotations: AnnotationMap

}

/// A request to stop the given provider on the indicated host
structure StopProviderCommand {
    /// Host ID on which to stop the provider
    @required    
    @serialization(name: "host_id")
    hostId: String,

    /// Reference for the capability provider. Can be any of the forms of 
    /// uniquely identifying a provider
    @required    
    @serialization(name: "provider_ref")
    providerRef: String,

    /// Link name for this provider
    @required    
    @serialization(name: "link_name")
    linkName: String,

    /// Contract ID of the capability provider
    @required    
    @serialization(name: "contract_id")
    contractId: String,

    /// Optional set of annotations used to describe the nature of this
    /// stop request
    annotations: AnnotationMap
}

/// A command sent to request that the given host purge and stop
structure StopHostCommand {
    /// The ID of the target host
    @required  
    @serialization(name: "host_id")  
    hostId: String,

    /// An optional timeout, in seconds
    timeout: U64
}

/// A command instructing a specific host to perform a live update
/// on the indicated actor by supplying a new image reference. Note that
/// live updates are only possible through image references
structure UpdateActorCommand {
    /// The host ID of the host to perform the live update
    @required   
    @serialization(name: "host_id") 
    hostId: String,

    /// The actor's 56-character unique ID
    @required    
    @serialization(name: "actor_id")
    actorId: String,

    /// The new image reference of the upgraded version of this actor
    @required    
    @serialization(name: "new_actor_ref")
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
    @serialization(name: "uptime_seconds")
    uptimeSeconds: U64
}

/// A response containing the full list of known claims within the lattice
structure GetClaimsResponse {
    @required
    claims: ClaimsList
}

list ClaimsList {
    member: ClaimsMap,
}

map ClaimsMap {
    key: String,
    value: String,
}

/// A request to remove a link definition and detach the relevant actor
/// from the given provider
structure RemoveLinkDefinitionRequest {
     /// The actor's public key. This cannot be an image reference
    @required    
    @serialization(name: "actor_id")
    actorId: String,

    /// The provider contract
    @required    
    @serialization(name: "contract_id")
    contractId: String,

    /// The provider's link name
    @required    
    @serialization(name: "link_name")
    linkName: String,
}