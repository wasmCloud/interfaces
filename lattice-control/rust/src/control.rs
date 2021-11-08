// This file is generated automatically using wasmcloud/weld-codegen and smithy model definitions
//

#![allow(unused_imports, clippy::ptr_arg, clippy::needless_lifetimes)]
use async_trait::async_trait;
use serde::{Deserialize, Serialize};
use std::{borrow::Cow, io::Write, string::ToString};
use wasmbus_rpc::{
    deserialize, serialize, Context, Message, MessageDispatch, RpcError, RpcResult, SendOpts,
    Timestamp, Transport,
};

pub const SMITHY_VERSION: &str = "1.0";

/// One of a potential list of responses to an actor auction
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ActorAuctionAck {
    /// The original actor reference used for the auction
    #[serde(default)]
    pub actor_ref: String,
    /// The host ID of the "bidder" for this auction.
    #[serde(default)]
    pub host_id: String,
}

pub type ActorAuctionAcks = Vec<ActorAuctionAck>;

/// A request to locate suitable hosts for a given actor
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ActorAuctionRequest {
    /// The reference for this actor. Can be any one of the acceptable forms
    /// of uniquely identifying an actor.
    #[serde(default)]
    pub actor_ref: String,
    /// The set of constraints to which any candidate host must conform
    pub constraints: ConstraintMap,
}

/// A summary description of an actor within a host inventory
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ActorDescription {
    /// Actor's 56-character unique ID
    #[serde(default)]
    pub id: String,
    /// Image reference for this actor, if applicable
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub image_ref: Option<String>,
    /// The individual instances of this actor that are running
    pub instances: ActorInstances,
    /// Name of this actor, if one exists
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}

pub type ActorDescriptions = Vec<ActorDescription>;

#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ActorInstance {
    /// The annotations that were used in the start request that produced
    /// this actor instance
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// This instance's unique ID (guid)
    #[serde(default)]
    pub instance_id: String,
    /// The revision number for this actor instance
    pub revision: i32,
}

pub type ActorInstances = Vec<ActorInstance>;

pub type AnnotationMap = std::collections::HashMap<String, String>;

pub type ClaimsList = Vec<ClaimsMap>;

pub type ClaimsMap = std::collections::HashMap<String, String>;

pub type ConfigurationString = String;

pub type ConstraintMap = std::collections::HashMap<String, String>;

/// Standard response for control interface operations
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct CtlOperationAck {
    #[serde(default)]
    pub accepted: bool,
    #[serde(default)]
    pub error: String,
}

/// A response containing the full list of known claims within the lattice
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct GetClaimsResponse {
    pub claims: ClaimsList,
}

/// A summary representation of a host
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct Host {
    #[serde(default)]
    pub id: String,
    /// uptime in seconds
    pub uptime_seconds: u64,
}

/// Describes the known contents of a given host at the time of
/// a query
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct HostInventory {
    /// Actors running on this host.
    pub actors: ActorDescriptions,
    /// The host's unique ID
    #[serde(default)]
    pub host_id: String,
    /// The host's labels
    pub labels: LabelsMap,
    /// Providers running on this host
    pub providers: ProviderDescriptions,
}

pub type Hosts = Vec<Host>;

pub type LabelsMap = std::collections::HashMap<String, String>;

/// A list of link definitions
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct LinkDefinitionList {
    pub links: wasmbus_rpc::core::ActorLinks,
}

/// One of a potential list of responses to a provider auction
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ProviderAuctionAck {
    /// The host ID of the "bidder" for this auction
    #[serde(default)]
    pub host_id: String,
    /// The link name provided for the auction
    #[serde(default)]
    pub link_name: String,
    /// The original provider ref provided for the auction
    #[serde(default)]
    pub provider_ref: String,
}

pub type ProviderAuctionAcks = Vec<ProviderAuctionAck>;

/// A request to locate a suitable host for a capability provider. The
/// provider's unique identity (reference + link name) is used to rule
/// out sites on which the provider is already running.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ProviderAuctionRequest {
    /// The set of constraints to which a suitable target host must conform
    pub constraints: ConstraintMap,
    /// The link name of the provider
    #[serde(default)]
    pub link_name: String,
    /// The reference for the provider. Can be any one of the accepted
    /// forms of uniquely identifying a provider
    #[serde(default)]
    pub provider_ref: String,
}

/// A summary description of a capability provider within a host inventory
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ProviderDescription {
    /// Provider's unique 56-character ID
    #[serde(default)]
    pub id: String,
    /// Image reference for this provider, if applicable
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub image_ref: Option<String>,
    /// Provider's link name
    #[serde(default)]
    pub link_name: String,
    /// Name of the provider, if one exists
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    /// The revision of the provider
    pub revision: i32,
}

pub type ProviderDescriptions = Vec<ProviderDescription>;

/// A request to remove a link definition and detach the relevant actor
/// from the given provider
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct RemoveLinkDefinitionRequest {
    /// The actor's public key. This cannot be an image reference
    #[serde(default)]
    pub actor_id: String,
    /// The provider's link name
    #[serde(default)]
    pub link_name: String,
    /// The provider's public key. This cannot be an image reference
    #[serde(default)]
    pub provider_id: String,
}

/// A command sent to a specific host instructing it to start the actor
/// indicated by the reference.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct StartActorCommand {
    /// Reference for the actor. Can be any of the acceptable forms of unique identification
    #[serde(default)]
    pub actor_ref: String,
    /// Optional set of annotations used to describe the nature of this actor start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// Host ID on which this actor should start
    #[serde(default)]
    pub host_id: String,
}

/// A command sent to a host requesting a capability provider be started with the
/// given link name and optional configuration.
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct StartProviderCommand {
    /// Optional set of annotations used to describe the nature of this provider start command. For
    /// example, autonomous agents may wish to "tag" start requests as part of a given deployment
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// Optional provider configuration in the form of an opaque string. Many
    /// providers prefer base64-encoded JSON here, though that data should never
    /// exceed 500KB
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub configuration: Option<ConfigurationString>,
    /// The host ID on which to start the provider
    #[serde(default)]
    pub host_id: String,
    /// The link name of the provider to be started
    #[serde(default)]
    pub link_name: String,
    /// The image reference of the provider to be started
    #[serde(default)]
    pub provider_ref: String,
}

/// A command sent to a host to request that instances of a given actor
/// be terminated on that host
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct StopActorCommand {
    /// Reference for this actor. Can be any of the means of uniquely identifying
    /// an actor
    #[serde(default)]
    pub actor_ref: String,
    /// Optional set of annotations used to describe the nature of this
    /// stop request. If supplied, the only instances of this actor with these
    /// annotations will be stopped
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// Optional count. If 0, all instances of this actor will be terminated
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub count: Option<u16>,
    /// The ID of the target host
    #[serde(default)]
    pub host_id: String,
}

/// A request to stop the given provider on the indicated host
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct StopProviderCommand {
    /// Optional set of annotations used to describe the nature of this
    /// stop request
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// Contract ID of the capability provider
    #[serde(default)]
    pub contract_id: String,
    /// Host ID on which to stop the provider
    #[serde(default)]
    pub host_id: String,
    /// Link name for this provider
    #[serde(default)]
    pub link_name: String,
    /// Reference for the capability provider. Can be any of the forms of
    /// uniquely identifying a provider
    #[serde(default)]
    pub provider_ref: String,
}

/// A command instructing a specific host to perform a live update
/// on the indicated actor by supplying a new image reference. Note that
/// live updates are only possible through image references
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct UpdateActorCommand {
    /// The actor's 56-character unique ID
    #[serde(default)]
    pub actor_id: String,
    /// Optional set of annotations used to describe the nature of this
    /// update request. Only actor instances that have matching annotations
    /// will be upgraded, allowing for instance isolation by
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub annotations: Option<AnnotationMap>,
    /// The host ID of the host to perform the live update
    #[serde(default)]
    pub host_id: String,
    /// The new image reference of the upgraded version of this actor
    #[serde(default)]
    pub new_actor_ref: String,
}

/// Lattice Controller - Describes the interface used for actors
/// to communicate with a lattice controller, enabling developers
/// to deploy actors that can manipulate the lattice in which they're
/// running.
#[async_trait]
pub trait LatticeController {
    /// Seek out a list of suitable hosts for a capability provider given
    /// a set of host label constraints. Hosts on which this provider is already
    /// running will not be among the successful "bidders" in this auction.
    async fn auction_provider(
        &self,
        ctx: &Context,
        arg: &ProviderAuctionRequest,
    ) -> RpcResult<ProviderAuctionAcks>;
    /// Seek out a list of suitable hosts for an actor given a set of host
    /// label constraints.
    async fn auction_actor(
        &self,
        ctx: &Context,
        arg: &ActorAuctionRequest,
    ) -> RpcResult<ActorAuctionAcks>;
    /// Queries the list of hosts currently visible to the lattice. This is
    /// a "gather" operation and so can be influenced by short timeouts,
    /// network partition events, etc.
    async fn get_hosts(&self, ctx: &Context) -> RpcResult<Hosts>;
    /// Queries for the contents of a host given the supplied 56-character unique ID
    async fn get_host_inventory<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<HostInventory>;
    /// Queries the lattice for the list of known/cached claims by taking the response
    /// from the first host that answers the query.
    async fn get_claims(&self, ctx: &Context) -> RpcResult<GetClaimsResponse>;
    /// Instructs a given host to start the indicated actor
    async fn start_actor(
        &self,
        ctx: &Context,
        arg: &StartActorCommand,
    ) -> RpcResult<CtlOperationAck>;
    /// Publish a link definition into the lattice, allowing it to be cached and
    /// delivered to the appropriate capability provider instances
    async fn advertise_link(
        &self,
        ctx: &Context,
        arg: &wasmbus_rpc::core::LinkDefinition,
    ) -> RpcResult<CtlOperationAck>;
    /// Requests the removal of a link definition. The definition will be removed
    /// from the cache and the relevant capability providers will be given a chance
    /// to de-provision any used resources
    async fn remove_link(
        &self,
        ctx: &Context,
        arg: &RemoveLinkDefinitionRequest,
    ) -> RpcResult<CtlOperationAck>;
    /// Queries all current link definitions in the lattice. The first host
    /// that receives this response will reply with the contents of the distributed
    /// cache
    async fn get_links(&self, ctx: &Context) -> RpcResult<LinkDefinitionList>;
    /// Requests that a specific host perform a live update on the indicated
    /// actor
    async fn update_actor(
        &self,
        ctx: &Context,
        arg: &UpdateActorCommand,
    ) -> RpcResult<CtlOperationAck>;
    /// Requests that the given host start the indicated capability provider
    async fn start_provider(
        &self,
        ctx: &Context,
        arg: &StartProviderCommand,
    ) -> RpcResult<CtlOperationAck>;
    /// Requests that the given capability provider be stopped on the indicated host
    async fn stop_provider(
        &self,
        ctx: &Context,
        arg: &StopProviderCommand,
    ) -> RpcResult<CtlOperationAck>;
    /// Requests that an actor be stopped on the given host
    async fn stop_actor(&self, ctx: &Context, arg: &StopActorCommand)
        -> RpcResult<CtlOperationAck>;
}

/// LatticeControllerReceiver receives messages defined in the LatticeController service trait
/// Lattice Controller - Describes the interface used for actors
/// to communicate with a lattice controller, enabling developers
/// to deploy actors that can manipulate the lattice in which they're
/// running.
#[doc(hidden)]
#[async_trait]
pub trait LatticeControllerReceiver: MessageDispatch + LatticeController {
    async fn dispatch(&self, ctx: &Context, message: &Message<'_>) -> RpcResult<Message<'_>> {
        match message.method {
            "AuctionProvider" => {
                let value: ProviderAuctionRequest = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::auction_provider(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.AuctionProvider",
                    arg: Cow::Owned(buf),
                })
            }
            "AuctionActor" => {
                let value: ActorAuctionRequest = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::auction_actor(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.AuctionActor",
                    arg: Cow::Owned(buf),
                })
            }
            "GetHosts" => {
                let resp = LatticeController::get_hosts(self, ctx).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.GetHosts",
                    arg: Cow::Owned(buf),
                })
            }
            "GetHostInventory" => {
                let value: String = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::get_host_inventory(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.GetHostInventory",
                    arg: Cow::Owned(buf),
                })
            }
            "GetClaims" => {
                let resp = LatticeController::get_claims(self, ctx).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.GetClaims",
                    arg: Cow::Owned(buf),
                })
            }
            "StartActor" => {
                let value: StartActorCommand = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::start_actor(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.StartActor",
                    arg: Cow::Owned(buf),
                })
            }
            "AdvertiseLink" => {
                let value: wasmbus_rpc::core::LinkDefinition = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::advertise_link(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.AdvertiseLink",
                    arg: Cow::Owned(buf),
                })
            }
            "RemoveLink" => {
                let value: RemoveLinkDefinitionRequest = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::remove_link(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.RemoveLink",
                    arg: Cow::Owned(buf),
                })
            }
            "GetLinks" => {
                let resp = LatticeController::get_links(self, ctx).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.GetLinks",
                    arg: Cow::Owned(buf),
                })
            }
            "UpdateActor" => {
                let value: UpdateActorCommand = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::update_actor(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.UpdateActor",
                    arg: Cow::Owned(buf),
                })
            }
            "StartProvider" => {
                let value: StartProviderCommand = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::start_provider(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.StartProvider",
                    arg: Cow::Owned(buf),
                })
            }
            "StopProvider" => {
                let value: StopProviderCommand = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::stop_provider(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.StopProvider",
                    arg: Cow::Owned(buf),
                })
            }
            "StopActor" => {
                let value: StopActorCommand = deserialize(message.arg.as_ref())
                    .map_err(|e| RpcError::Deser(format!("message '{}': {}", message.method, e)))?;
                let resp = LatticeController::stop_actor(self, ctx, &value).await?;
                let buf = serialize(&resp)?;
                Ok(Message {
                    method: "LatticeController.StopActor",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "LatticeController::{}",
                message.method
            ))),
        }
    }
}

/// LatticeControllerSender sends messages to a LatticeController service
/// Lattice Controller - Describes the interface used for actors
/// to communicate with a lattice controller, enabling developers
/// to deploy actors that can manipulate the lattice in which they're
/// running.
/// client for sending LatticeController messages
#[derive(Debug)]
pub struct LatticeControllerSender<T: Transport> {
    transport: T,
}

impl<T: Transport> LatticeControllerSender<T> {
    /// Constructs a LatticeControllerSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> LatticeController
    for LatticeControllerSender<T>
{
    #[allow(unused)]
    /// Seek out a list of suitable hosts for a capability provider given
    /// a set of host label constraints. Hosts on which this provider is already
    /// running will not be among the successful "bidders" in this auction.
    async fn auction_provider(
        &self,
        ctx: &Context,
        arg: &ProviderAuctionRequest,
    ) -> RpcResult<ProviderAuctionAcks> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.AuctionProvider",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "AuctionProvider", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Seek out a list of suitable hosts for an actor given a set of host
    /// label constraints.
    async fn auction_actor(
        &self,
        ctx: &Context,
        arg: &ActorAuctionRequest,
    ) -> RpcResult<ActorAuctionAcks> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.AuctionActor",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "AuctionActor", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Queries the list of hosts currently visible to the lattice. This is
    /// a "gather" operation and so can be influenced by short timeouts,
    /// network partition events, etc.
    async fn get_hosts(&self, ctx: &Context) -> RpcResult<Hosts> {
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.GetHosts",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "GetHosts", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Queries for the contents of a host given the supplied 56-character unique ID
    async fn get_host_inventory<TS: ToString + ?Sized + std::marker::Sync>(
        &self,
        ctx: &Context,
        arg: &TS,
    ) -> RpcResult<HostInventory> {
        let buf = serialize(&arg.to_string())?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.GetHostInventory",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "GetHostInventory", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Queries the lattice for the list of known/cached claims by taking the response
    /// from the first host that answers the query.
    async fn get_claims(&self, ctx: &Context) -> RpcResult<GetClaimsResponse> {
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.GetClaims",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "GetClaims", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Instructs a given host to start the indicated actor
    async fn start_actor(
        &self,
        ctx: &Context,
        arg: &StartActorCommand,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.StartActor",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "StartActor", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Publish a link definition into the lattice, allowing it to be cached and
    /// delivered to the appropriate capability provider instances
    async fn advertise_link(
        &self,
        ctx: &Context,
        arg: &wasmbus_rpc::core::LinkDefinition,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.AdvertiseLink",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "AdvertiseLink", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Requests the removal of a link definition. The definition will be removed
    /// from the cache and the relevant capability providers will be given a chance
    /// to de-provision any used resources
    async fn remove_link(
        &self,
        ctx: &Context,
        arg: &RemoveLinkDefinitionRequest,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.RemoveLink",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "RemoveLink", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Queries all current link definitions in the lattice. The first host
    /// that receives this response will reply with the contents of the distributed
    /// cache
    async fn get_links(&self, ctx: &Context) -> RpcResult<LinkDefinitionList> {
        let buf = *b"";
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.GetLinks",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "GetLinks", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Requests that a specific host perform a live update on the indicated
    /// actor
    async fn update_actor(
        &self,
        ctx: &Context,
        arg: &UpdateActorCommand,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.UpdateActor",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "UpdateActor", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Requests that the given host start the indicated capability provider
    async fn start_provider(
        &self,
        ctx: &Context,
        arg: &StartProviderCommand,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.StartProvider",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "StartProvider", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Requests that the given capability provider be stopped on the indicated host
    async fn stop_provider(
        &self,
        ctx: &Context,
        arg: &StopProviderCommand,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.StopProvider",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "StopProvider", e)))?;
        Ok(value)
    }
    #[allow(unused)]
    /// Requests that an actor be stopped on the given host
    async fn stop_actor(
        &self,
        ctx: &Context,
        arg: &StopActorCommand,
    ) -> RpcResult<CtlOperationAck> {
        let buf = serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "LatticeController.StopActor",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;
        let value = deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("response to {}: {}", "StopActor", e)))?;
        Ok(value)
    }
}
