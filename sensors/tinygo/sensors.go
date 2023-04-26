// Sensor Measurement Interface
package sensors

import (
	actor "github.com/wasmcloud/actor-tinygo"     //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// List of floating point values
type FloatList []float32

// MEncode serializes a FloatList using msgpack
func (o *FloatList) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteFloat32(item_o)
	}

	return encoder.CheckError()
}

// MDecodeFloatList deserializes a FloatList using msgpack
func MDecodeFloatList(d *msgpack.Decoder) (FloatList, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]float32, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]float32, 0), err
	}
	val := make([]float32, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadFloat32()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a FloatList using cbor
func (o *FloatList) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteFloat32(item_o)
	}

	return encoder.CheckError()
}

// CDecodeFloatList deserializes a FloatList using cbor
func CDecodeFloatList(d *cbor.Decoder) (FloatList, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]float32, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]float32, 0), err
	}
	val := make([]float32, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadFloat32()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Indicates a failure to obtain or produce a measurement
type MeasurementError struct {
	ErrorMessage string `json:"errorMessage"`
	// Network ID (optional)
	NetworkId string `json:"networkId"`
	// Sensor ID associated with the error. Optional
	SensorId string `json:"sensorId"`
	// Type of the sensor on which the error occurred. Optional
	SensorType string `json:"sensorType"`
	// Timestamp when the error occurred
	Timestamp uint64 `json:"timestamp"`
}

// MEncode serializes a MeasurementError using msgpack
func (o *MeasurementError) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("errorMessage")
	encoder.WriteString(o.ErrorMessage)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)
	encoder.WriteString("sensorType")
	encoder.WriteString(o.SensorType)
	encoder.WriteString("timestamp")
	encoder.WriteUint64(o.Timestamp)

	return encoder.CheckError()
}

// MDecodeMeasurementError deserializes a MeasurementError using msgpack
func MDecodeMeasurementError(d *msgpack.Decoder) (MeasurementError, error) {
	var val MeasurementError
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "errorMessage":
			val.ErrorMessage, err = d.ReadString()
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		case "sensorType":
			val.SensorType, err = d.ReadString()
		case "timestamp":
			val.Timestamp, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a MeasurementError using cbor
func (o *MeasurementError) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("errorMessage")
	encoder.WriteString(o.ErrorMessage)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)
	encoder.WriteString("sensorType")
	encoder.WriteString(o.SensorType)
	encoder.WriteString("timestamp")
	encoder.WriteUint64(o.Timestamp)

	return encoder.CheckError()
}

// CDecodeMeasurementError deserializes a MeasurementError using cbor
func CDecodeMeasurementError(d *cbor.Decoder) (MeasurementError, error) {
	var val MeasurementError
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "errorMessage":
			val.ErrorMessage, err = d.ReadString()
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		case "sensorType":
			val.SensorType, err = d.ReadString()
		case "timestamp":
			val.Timestamp, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Represents a single measurement that originates from a sensor
type MeasurementRecord struct {
	// The data contained within the measurement. Must contain at least one value. Measurements
	// should never be delivered with an empty payload
	Data FloatList `json:"data"`
	// The network ID from which the measurement originated. This value is optional
	NetworkId string `json:"networkId"`
	// The ID of the sensor from which the measurement originated.
	SensorId string `json:"sensorId"`
	// The type of the sensor. The value of the sensor type is understood by consumer
	// and producer and is not dictated by the interface or contained in an interface
	// enumeration/union.
	SensorType string `json:"sensorType"`
	// UNIX timestamp indicating when the measurement occurred. Timezone is opaque though recommended
	// that it should be in UTC
	Timestamp uint64 `json:"timestamp"`
}

// MEncode serializes a MeasurementRecord using msgpack
func (o *MeasurementRecord) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("data")
	o.Data.MEncode(encoder)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)
	encoder.WriteString("sensorType")
	encoder.WriteString(o.SensorType)
	encoder.WriteString("timestamp")
	encoder.WriteUint64(o.Timestamp)

	return encoder.CheckError()
}

// MDecodeMeasurementRecord deserializes a MeasurementRecord using msgpack
func MDecodeMeasurementRecord(d *msgpack.Decoder) (MeasurementRecord, error) {
	var val MeasurementRecord
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = MDecodeFloatList(d)
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		case "sensorType":
			val.SensorType, err = d.ReadString()
		case "timestamp":
			val.Timestamp, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a MeasurementRecord using cbor
func (o *MeasurementRecord) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(5)
	encoder.WriteString("data")
	o.Data.CEncode(encoder)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)
	encoder.WriteString("sensorType")
	encoder.WriteString(o.SensorType)
	encoder.WriteString("timestamp")
	encoder.WriteUint64(o.Timestamp)

	return encoder.CheckError()
}

// CDecodeMeasurementRecord deserializes a MeasurementRecord using cbor
func CDecodeMeasurementRecord(d *cbor.Decoder) (MeasurementRecord, error) {
	var val MeasurementRecord
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = CDecodeFloatList(d)
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		case "sensorType":
			val.SensorType, err = d.ReadString()
		case "timestamp":
			val.Timestamp, err = d.ReadUint64()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// A query for the latest data value from the sensor. Note that sensor type is not supplied
// in order to query the latest value
type SensorQuery struct {
	// If supplied, indicates the network ID for the sensor being queried
	NetworkId string `json:"networkId"`
	// Indicates the ID of the sensor being queried
	SensorId string `json:"sensorId"`
}

// MEncode serializes a SensorQuery using msgpack
func (o *SensorQuery) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)

	return encoder.CheckError()
}

// MDecodeSensorQuery deserializes a SensorQuery using msgpack
func MDecodeSensorQuery(d *msgpack.Decoder) (SensorQuery, error) {
	var val SensorQuery
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SensorQuery using cbor
func (o *SensorQuery) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("networkId")
	encoder.WriteString(o.NetworkId)
	encoder.WriteString("sensorId")
	encoder.WriteString(o.SensorId)

	return encoder.CheckError()
}

// CDecodeSensorQuery deserializes a SensorQuery using cbor
func CDecodeSensorQuery(d *cbor.Decoder) (SensorQuery, error) {
	var val SensorQuery
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "networkId":
			val.NetworkId, err = d.ReadString()
		case "sensorId":
			val.SensorId, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Results of a sensor query
type SensorQueryResult struct {
	// An optional result containing the most recent measurement record for the given
	// sensor
	Result *MeasurementRecord `json:"result"`
}

// MEncode serializes a SensorQueryResult using msgpack
func (o *SensorQueryResult) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("result")
	if o.Result == nil {
		encoder.WriteNil()
	} else {
		o.Result.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeSensorQueryResult deserializes a SensorQueryResult using msgpack
func MDecodeSensorQueryResult(d *msgpack.Decoder) (SensorQueryResult, error) {
	var val SensorQueryResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "result":
			fval, err := MDecodeMeasurementRecord(d)
			if err != nil {
				return val, err
			}
			val.Result = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SensorQueryResult using cbor
func (o *SensorQueryResult) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("result")
	if o.Result == nil {
		encoder.WriteNil()
	} else {
		o.Result.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeSensorQueryResult deserializes a SensorQueryResult using cbor
func CDecodeSensorQueryResult(d *cbor.Decoder) (SensorQueryResult, error) {
	var val SensorQueryResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "result":
			fval, err := CDecodeMeasurementRecord(d)
			if err != nil {
				return val, err
			}
			val.Result = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type SensorListener interface {
	// Handles a measurement delivered by some provider of this service. Measurements are delivered
	// to a component without need for a reply
	HandleMeasurement(ctx *actor.Context, arg MeasurementRecord) error
	// Handles an error in the sensor. This is not to be confused with a value delivery with a
	// value within error limits. This indicates a failure to obtain or produce a measurement.
	HandleError(ctx *actor.Context, arg MeasurementError) error
}

// SensorListenerHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func SensorListenerHandler(actor_ SensorListener) actor.Handler {
	return actor.NewHandler("SensorListener", &SensorListenerReceiver{}, actor_)
}

// SensorListenerContractId returns the capability contract id for this interface
func SensorListenerContractId() string { return "wasmcloud:sensors" }

// SensorListenerReceiver receives messages defined in the SensorListener service interface
type SensorListenerReceiver struct{}

func (r *SensorListenerReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(SensorListener)
	switch message.Method {

	case "HandleMeasurement":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeMeasurementRecord(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.HandleMeasurement(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "SensorListener.HandleMeasurement", Arg: buf}, nil
		}
	case "HandleError":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeMeasurementError(&d)
			if err_ != nil {
				return nil, err_
			}

			err := svc_.HandleError(ctx, value)
			if err != nil {
				return nil, err
			}
			buf := make([]byte, 0)
			return &actor.Message{Method: "SensorListener.HandleError", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "SensorListener."+message.Method)
	}
}

// SensorListenerSender sends messages to a SensorListener service
type SensorListenerSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorSensorListenerSender(actor_id string) *SensorListenerSender {
	transport := actor.ToActor(actor_id)
	return &SensorListenerSender{transport: transport}
}

// Handles a measurement delivered by some provider of this service. Measurements are delivered
// to a component without need for a reply
func (s *SensorListenerSender) HandleMeasurement(ctx *actor.Context, arg MeasurementRecord) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "SensorListener.HandleMeasurement", Arg: buf})
	return nil
}

// Handles an error in the sensor. This is not to be confused with a value delivery with a
// value within error limits. This indicates a failure to obtain or produce a measurement.
func (s *SensorListenerSender) HandleError(ctx *actor.Context, arg MeasurementError) error {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	s.transport.Send(ctx, actor.Message{Method: "SensorListener.HandleError", Arg: buf})
	return nil
}

type Sensors interface {
	// Query supplied to a provider of this interface to obtain the most recent (if any) value of a sensor
	QuerySensor(ctx *actor.Context, arg SensorQuery) (*SensorQueryResult, error)
}

// SensorsHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func SensorsHandler(actor_ Sensors) actor.Handler {
	return actor.NewHandler("Sensors", &SensorsReceiver{}, actor_)
}

// SensorsContractId returns the capability contract id for this interface
func SensorsContractId() string { return "wasmcloud:sensors" }

// SensorsReceiver receives messages defined in the Sensors service interface
type SensorsReceiver struct{}

func (r *SensorsReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Sensors)
	switch message.Method {

	case "QuerySensor":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeSensorQuery(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.QuerySensor(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "Sensors.QuerySensor", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Sensors."+message.Method)
	}
}

// SensorsSender sends messages to a Sensors service
type SensorsSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a Sensors provider
// implementing the 'wasmcloud:sensors' capability contract, with the "default" link
func NewProviderSensors() *SensorsSender {
	transport := actor.ToProvider("wasmcloud:sensors", "default")
	return &SensorsSender{transport: transport}
}

// NewProviderSensorsLink constructs a client for sending to a Sensors provider
// implementing the 'wasmcloud:sensors' capability contract, with the specified link name
func NewProviderSensorsLink(linkName string) *SensorsSender {
	transport := actor.ToProvider("wasmcloud:sensors", linkName)
	return &SensorsSender{transport: transport}
}

// Query supplied to a provider of this interface to obtain the most recent (if any) value of a sensor
func (s *SensorsSender) QuerySensor(ctx *actor.Context, arg SensorQuery) (*SensorQueryResult, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Sensors.QuerySensor", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeSensorQueryResult(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.6.0
