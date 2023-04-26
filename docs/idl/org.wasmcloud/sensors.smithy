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
service SensorListener {
  version: "0.1.0",
  operations: [
    HandleMeasurement,
    HandleError
  ]
}

@wasmbus(
    contractId: "wasmcloud:sensors",
    providerReceive: true )
service Sensors {
  version: "0.1.0",
  operations: [
    QuerySensor
  ]
}

/// Query supplied to a provider of this interface to obtain the most recent (if any) value of a sensor
operation QuerySensor {
  input: SensorQuery,
  output: SensorQueryResult
}

/// Handles a measurement delivered by some provider of this service. Measurements are delivered
/// to a component without need for a reply
operation HandleMeasurement {
  input: MeasurementRecord
}

/// Handles an error in the sensor. This is not to be confused with a value delivery with a 
/// value within error limits. This indicates a failure to obtain or produce a measurement.
operation HandleError {
  input: MeasurementError
}

/// A query for the latest data value from the sensor. Note that sensor type is not supplied
/// in order to query the latest value
structure SensorQuery {
  /// If supplied, indicates the network ID for the sensor being queried
  networkId: String,

  /// Indicates the ID of the sensor being queried
  @required
  sensorId: String
}

/// Results of a sensor query
structure SensorQueryResult {
  /// An optional result containing the most recent measurement record for the given
  /// sensor
  result: MeasurementRecord
}

/// Indicates a failure to obtain or produce a measurement
structure MeasurementError {
  /// Network ID (optional)
  networkId: String,

  /// Timestamp when the error occurred
  @required
  timestamp: U64,

  /// Sensor ID associated with the error. Optional  
  sensorId: String,

  /// Type of the sensor on which the error occurred. Optional
  sensorType: String,

  @required
  errorMessage: String
}

/// Represents a single measurement that originates from a sensor
structure MeasurementRecord {
    /// The network ID from which the measurement originated. This value is optional
    networkId: String,

    /// The ID of the sensor from which the measurement originated. 
    @required
    sensorId: String,

    /// The type of the sensor. The value of the sensor type is understood by consumer
    /// and producer and is not dictated by the interface or contained in an interface
    /// enumeration/union.
    @required
    sensorType: String,

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