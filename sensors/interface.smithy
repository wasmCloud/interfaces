// Definition of the sensor measurement interface `wasmcloud:sensors`

// Provide helpful metadata used by code generators
metadata package = [{
    namespace: "org.wasmcloud.interface.sensors",
    crate: "wasmcloud_interface_sensors",    
    doc: "Sensor Measurement Interface",
}]

namespace org.wasmcloud.interface.sensors

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64

@wasmbus(
    contractId: "wasmcloud:sensors",
    actorReceive: true )
service Sensors {
  version: "0.1.0",
  operations: [
    HandleMeasurement
  ]
}

/// Handles a measurement delivered by some provider of this service. Measurements are delivered
/// to a component without need for a reply
operation HandleMeasurement {
    input: MeasurementRecord
}

/// Represents a single measurement that originates from a sensor
structure MeasurementRecord {
    /// The network ID from which the measurement originated. This value is optional, if supplied
    /// then a 0 should be considered a meaningful network ID and not the absence of a network
    networkId: U32,

    /// The ID of the sensor from which the measurement originated. 
    @required
    sensorId: U32,

    /// The type of the sensor. The value of the sensor type is understood by consumer
    /// and producer and is not dictated by the interface or contained in an interface
    /// enumeration/union.
    @required
    sensorType: U32,

    /// The data contained within the measurement. Must contain at least one value. Measurements
    /// should never be delivered with an empty payload
    @required
    data: FloatList,

    /// UNIX timestamp indicating when the measurement occurred. Timezone is opaque though recommended
    /// that it should be in UTC
    @required
    timestamp: U64
}

/// List of floating point values
list FloatList {
  member: Float
}