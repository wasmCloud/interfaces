// config-service.smithy
// Declaration of data types used for information exchange between a wasmCloud host
// and a wasmCloud configuration service

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [{
    namespace: "org.wasmcloud.interface.configservice",
    crate: "wasmcloud_interface_configservice"    
}]

namespace org.wasmcloud.interface.configservice

/// A request made by a wasmCloud host at startup that contains the host
/// profile, like its labels. These labels are used by the configuration service
/// to decide which configuration profile to include in the reply
structure ConfigurationRequest {
    @required
    labels: LabelMap
}

map LabelMap {
    key: String,
    value: String,
}

/// Contains the configuration profile as deemed appropriate by the configuration
/// service. A host should react accordingly to this profile upon startup and during
/// real-time configuration change notifications on the "push" topic. It's noteworthy
/// that this configuration profile does NOT contain any connection information for NATS,
/// as allowing that to change at runtime would be a security and reliability risk
structure HostConfigurationProfile {
    /// The optional list of capability providers a host should automatically start
    autoStartProviders: ProviderReferenceList

    /// The optional list of actors that a host should automatically start
    autoStartActors: ActorReferenceList

    /// A set of credentials the host can use for fetching artifacts
    registryCredentials: RegistryCredentialMap
}

list ActorReferenceList {
    member: String
}

list ProviderReferenceList {
    member: ProviderReference,    
}

structure ProviderReference {
    imageReference: String,
    linkName: String
}

/// A set of credentials to be used for fetching from specific registries
map RegistryCredentialMap {
    /// The key of this map is the OCI/BINDLE URL without the artifact reference. Credentials
    /// are matched via substring comparison on the URL of an artifact.
    key: String
    value: RegistryCredential
}

structure RegistryCredential {
    /// If supplied, token authentication will be used for the registry
    token: String,
    /// If supplied, username and password will be used for HTTP Basic authentication
    username: String,
    password: String,
    /// The type of the registry (either "oci" or "bindle")
    @required
    registryType: String
}