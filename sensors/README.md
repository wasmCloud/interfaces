# Sensor Measurement Interface
This is an interface through which WebAssembly components can interact with sensors of the type often used in IoT applications, though _any_ source capable of emitting measurements can use this interface.

## Goals
The following is a list of goals desired by this interface:

* To define an interface by which data emitters can deliver data to consumers within the domain of sensors and measurements
* To define an interface that can be used equally well by constained-device, IoT, and cloud applications

## Non-Goals
The following is a list of non-goals (problems out of scope) of this interface:

* This interface does **not** declare a finite set of sensor types
* This interface does **not** assign _any_ meaning to the data values delivered with any measurement
* This interface does **not** issue any requirement or recommendation with regard to the _medium_ by which measurements are gathered
* This interface does **not** deal with data aggregation, sliding windows, rolling computations, or other features commonly associated with CEP (complex event processing) applications. Such logic is left to the consumer of the interface.
* This interface does **not** deal with the frequency of measurement deliveries (sampling rate). This too is an out-of-band arrangement between measurement provider and consumer.

## Design Discussion
The following sections contain a detailed description of the design of the _sensor interface specification_.

### Glossary of Terms
Before getting into the expected usage scenarios and main design of the interface, it may be helpful to describe some terms that will be used throughout the interface definition because these terms have slightly different or even conflicting meanings in other domains.

* **Sensor Network** - A logical association of a group of sensors, identified by a unique numeric identifier. This is the only optional value in the measurement record.
* **Sensor** - An abstraction of a _source_ of a measurement, identified by a numeric identifier
* **Sensor Type** - A classification of a sensor that can be used to give application code information on how data from that sensor should be processed. There is no universal standard of type IDs. Type identifiers are agreed upon by users of the interface and not defined by the interface itself.
* **Measurement** - A discrete bundle of data emitted by a sensor, within a sensor network, at a given time
* **Timestamp** - A floating point numeric UNIX timestamp value. It is _strongly recommended_ that this value always be in UTC, but an interface implementation can choose to disregard this. This number should be in _milliseconds_ resolution.

### Sensor Networks
A sensor network is a _logical_ grouping of sensors. This logical grouping can be used to group sensors by distance (near, far, etc), by region, or any other arbitrary association. Providing redundancy or "backup" sensor data can also be accomplished by reusing sensor IDs across different sensor networks. A complicated IoT device like a car could consist of dozens or even hundreds of sensor networks, all at the discretion of the manufacturer, while other usages require no sensor network at all.

### Sensor Identification
Identifying sensors is a design and architecture concern. It might be tempting to assign an ID to a sensor for `temperature`, and another for `velocity` and another for something like barometric pressure, and so on. However, this isn't always practical or possible.

For example, a group of sensors may contain multiple sensors of the same type with but with unique, meaningful identities. A quadcopter drone may have a sensor that reports wind speed and direction at each of the four motor locations. That data is collected and used to produce compensatory instructions sent to the motor controller(s). To accommodate this, you would use a different ID for each of the wind sensors, all of which would have the same _sensor type_. The sensor type gives the code a clue as to how to interpret the data in the measurement payload, while the sensor ID itself can either be treated as "just a unique value", or a value that might be something like `ROTOR_1`, `ROTOR_2`, etc.

Another example of using the sensor ID to provide meaningful value might be in a vehicle that has obstacle detection sensors. Let's say that it has an array of 8 of those sensors. The IDs of those sensors could indicate the sensor's position code (`RIGHT_FRONT`, `MIDDLE_FRONT`, etc) and provide the application logic with the metadata necessary to take appropriate action.

Sensor, type, and network IDs are all opaque strings with meaning agreed upon between component and provider.

### Sampling Rates
Sampling rates, both awareness and control thereof, are outside the scope of this interface. It is entirely up to an implementor of this interface to determine the rate at which measurements are delivered.

### What about Blob?
If you are interested in receiving byte arrays/blobs for your application, then you might be better served using the message broker interface. That interface is specifically designed around receiving opaque binary payloads and can be satisfied by an IoT-friendly broker like MQTT or even something lower level.

A measurement is a list of 1 or more floating point values and should not be interpreted as a blob.

## Measurement Examples
The following are a few illustrative examples of how measurements might be structured to provide additional context around the structure of this interface.

### Simple Single Measurement Application
Let's say you're building an application that monitors a single value from a single sensor. This could be something connected to a voltage monitor like GPIO, a thermistor, or countless other devices. A common (and fun) IoT hobby project is to connect a meat thermometer and/or ambient air temperature sensor and create a "smart grill". Such a usage scenario would have no need for the concept of a sensor network, and the sensor ID would always be the same. Each measurement would be a single value, the current temperature as detected by the device.

In the sample data below, the network ID, sensor ID, and sensor type never change while each measurement contains a single value, the temperature (in Fahrenheit, though obviously that is at the discretion of the sensor).

| Timestamp | Network ID | Sensor ID | Sensor Type | Measurement |
| :-------- | :--------: | :-------: | :---------: | :---------- |
| 1672597308100 | - | `t1` | `temp` | [`185.0`] |
| 1672597391100 | - | `t1` | `temp` | [`200.0`] |
| ... | - | - | - | ... |

Another interesting example might be a badged entry system where you want to receive a measurement from the gate sensor each time a badge is swiped. Here, rather than having a single sensor and then differentiating "in" and "out" via the data, you typically see the entrance gate and exit gate having different sensor IDs, and the payload contains the ID of the card swiped. The payload could also include whether the access was granted or denied: 

| Timestamp | Network ID | Sensor ID | Sensor Type | Measurement |
| :-------- | :--------: | :-------: | :---------: | :---------- |
| 1672597308100 | `bldg-2` | `f1-in` | `badge-access` | [`1.0`, `12384231.0`] |
| 1672597391100 | `bldg-2` | `f1-out` | `badge-access`| [`0.0`, `12319312.3`] |
| ... | - | - | - | ... |

In the above sample, we might use network `bldg-2` to identify the building in which the badge swipe sensors reside. The first measurement indicates an access grant (the door unlocked), whereas the second measurement indicates an access denial (0.0 == false).

### Multi-Valued Measurements
Not all sensors deliver a single value with a single measurement payload. In many cases the sensor will deliver multiple values. Sometimes these values are just an array, and other times the position of each data element is meaningful. 

Take, for example, a sensor that delivers _location_ data in measurement payloads. There, the meaasurement data would be an array of 3 elements, containing the `latitude`, `longitude`, and `altitude`, respectively. As with all participating sensors, the meaning of the measurement data array contents is known only to the sensor and the WebAssembly component responding to those measurements. The contract itself remains oblivious to this, only delivering an array of floating point numbers for each measurement.

| Timestamp | Network ID | Sensor ID | Sensor Type | Measurement |
| :-------- | :--------: | :-------: | :---------: | :---------- |
| 1672597308 | - | `s1303` | `position` | [`42.36`, `-71.05`, `175.0`] |


## Interface Specification
The interface, at its core, contains a definition for a measurement reading record, and calls for 3 functions:

* `deliver` - delivers a reading to the consumer
* `error` - indicates a sensor error
* `latest-reading` - allows a WebAssembly component to query the most recent (if any) value on a given sensor


⚠️ **NOTE** It is important to distinguish between error types. A value recorded by a sensor that is outside normal operating range will be _delivered_ with an appropriate value and should be handled by standard logic. A _failure_ to acquire a value, or some other detected failure in the sensor, will be indicated with an `error` call.

The following are some files describing this interface:

* [Smithy](./interface.smithy) - A codegen-ready interface definition
* [WIT](./interface.wit) - A _sample_ illustrating what a WIT file for this interface _might_ look like
