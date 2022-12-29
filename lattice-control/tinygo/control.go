package latticecontrol
            import (
                "github.com/wasmcloud/actor-tinygo" //nolint
                msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
                cbor "github.com/wasmcloud/tinygo-cbor" //nolint
            )
// One of a potential list of responses to an actor auction
type ActorAuctionAck struct {
// The original actor reference used for the auction
  ActorRef string 
// The host ID of the "bidder" for this auction.
  HostId string 
}

// MEncode serializes a ActorAuctionAck using msgpack
            func (o *ActorAuctionAck) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)

                return encoder.CheckError()
            }
            
            // MDecodeActorAuctionAck deserializes a ActorAuctionAck using msgpack
            func MDecodeActorAuctionAck(d *msgpack.Decoder) (ActorAuctionAck,error) {
                var val ActorAuctionAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "hostId":
val.HostId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ActorAuctionAck using cbor
            func (o *ActorAuctionAck) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)

                return encoder.CheckError()
            }
            
            // CDecodeActorAuctionAck deserializes a ActorAuctionAck using cbor
            func CDecodeActorAuctionAck(d *cbor.Decoder) (ActorAuctionAck,error) {
                var val ActorAuctionAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "hostId":
val.HostId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ActorAuctionAcks []ActorAuctionAck
// MEncode serializes a ActorAuctionAcks using msgpack
            func (o *ActorAuctionAcks) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeActorAuctionAcks deserializes a ActorAuctionAcks using msgpack
            func MDecodeActorAuctionAcks(d *msgpack.Decoder) (ActorAuctionAcks,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorAuctionAck, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]ActorAuctionAck, 0 ), err }
                        val := make([]ActorAuctionAck, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeActorAuctionAck(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a ActorAuctionAcks using cbor
            func (o *ActorAuctionAcks) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeActorAuctionAcks deserializes a ActorAuctionAcks using cbor
            func CDecodeActorAuctionAcks(d *cbor.Decoder) (ActorAuctionAcks,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorAuctionAck, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]ActorAuctionAck, 0 ), err }
                        val := make([]ActorAuctionAck, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeActorAuctionAck(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// A request to locate suitable hosts for a given actor
type ActorAuctionRequest struct {
// The reference for this actor. Can be any one of the acceptable forms
// of uniquely identifying an actor.
  ActorRef string 
// The set of constraints to which any candidate host must conform
  Constraints ConstraintMap 
// The ID of the lattice on which this request will be performed
  LatticeId string 
}

// MEncode serializes a ActorAuctionRequest using msgpack
            func (o *ActorAuctionRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("constraints")
o.Constraints.MEncode(encoder)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeActorAuctionRequest deserializes a ActorAuctionRequest using msgpack
            func MDecodeActorAuctionRequest(d *msgpack.Decoder) (ActorAuctionRequest,error) {
                var val ActorAuctionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "constraints":
val.Constraints,err = MDecodeConstraintMap(d)
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ActorAuctionRequest using cbor
            func (o *ActorAuctionRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("constraints")
o.Constraints.CEncode(encoder)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeActorAuctionRequest deserializes a ActorAuctionRequest using cbor
            func CDecodeActorAuctionRequest(d *cbor.Decoder) (ActorAuctionRequest,error) {
                var val ActorAuctionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "constraints":
val.Constraints,err = CDecodeConstraintMap(d)
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A summary description of an actor within a host inventory
type ActorDescription struct {
// Actor's 56-character unique ID
  Id string 
// Image reference for this actor, if applicable
  ImageRef string 
// The individual instances of this actor that are running
  Instances ActorInstances 
// Name of this actor, if one exists
  Name string 
}

// MEncode serializes a ActorDescription using msgpack
            func (o *ActorDescription) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("imageRef")
encoder.WriteString(o.ImageRef)
encoder.WriteString("instances")
o.Instances.MEncode(encoder)
encoder.WriteString("name")
encoder.WriteString(o.Name)

                return encoder.CheckError()
            }
            
            // MDecodeActorDescription deserializes a ActorDescription using msgpack
            func MDecodeActorDescription(d *msgpack.Decoder) (ActorDescription,error) {
                var val ActorDescription
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "id":
val.Id,err = d.ReadString()
case "imageRef":
val.ImageRef,err = d.ReadString()
case "instances":
val.Instances,err = MDecodeActorInstances(d)
case "name":
val.Name,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ActorDescription using cbor
            func (o *ActorDescription) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("imageRef")
encoder.WriteString(o.ImageRef)
encoder.WriteString("instances")
o.Instances.CEncode(encoder)
encoder.WriteString("name")
encoder.WriteString(o.Name)

                return encoder.CheckError()
            }
            
            // CDecodeActorDescription deserializes a ActorDescription using cbor
            func CDecodeActorDescription(d *cbor.Decoder) (ActorDescription,error) {
                var val ActorDescription
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "id":
val.Id,err = d.ReadString()
case "imageRef":
val.ImageRef,err = d.ReadString()
case "instances":
val.Instances,err = CDecodeActorInstances(d)
case "name":
val.Name,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ActorDescriptions []ActorDescription
// MEncode serializes a ActorDescriptions using msgpack
            func (o *ActorDescriptions) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeActorDescriptions deserializes a ActorDescriptions using msgpack
            func MDecodeActorDescriptions(d *msgpack.Decoder) (ActorDescriptions,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorDescription, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]ActorDescription, 0 ), err }
                        val := make([]ActorDescription, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeActorDescription(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a ActorDescriptions using cbor
            func (o *ActorDescriptions) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeActorDescriptions deserializes a ActorDescriptions using cbor
            func CDecodeActorDescriptions(d *cbor.Decoder) (ActorDescriptions,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorDescription, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]ActorDescription, 0 ), err }
                        val := make([]ActorDescription, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeActorDescription(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
type ActorInstance struct {
// The annotations that were used in the start request that produced
// this actor instance
  Annotations *AnnotationMap 
// This instance's unique ID (guid)
  InstanceId string 
// The revision number for this actor instance
  Revision int32 
}

// MEncode serializes a ActorInstance using msgpack
            func (o *ActorInstance) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("instanceId")
encoder.WriteString(o.InstanceId)
encoder.WriteString("revision")
encoder.WriteInt32(o.Revision)

                return encoder.CheckError()
            }
            
            // MDecodeActorInstance deserializes a ActorInstance using msgpack
            func MDecodeActorInstance(d *msgpack.Decoder) (ActorInstance,error) {
                var val ActorInstance
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "instanceId":
val.InstanceId,err = d.ReadString()
case "revision":
val.Revision,err = d.ReadInt32()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ActorInstance using cbor
            func (o *ActorInstance) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("instanceId")
encoder.WriteString(o.InstanceId)
encoder.WriteString("revision")
encoder.WriteInt32(o.Revision)

                return encoder.CheckError()
            }
            
            // CDecodeActorInstance deserializes a ActorInstance using cbor
            func CDecodeActorInstance(d *cbor.Decoder) (ActorInstance,error) {
                var val ActorInstance
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "instanceId":
val.InstanceId,err = d.ReadString()
case "revision":
val.Revision,err = d.ReadInt32()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ActorInstances []ActorInstance
// MEncode serializes a ActorInstances using msgpack
            func (o *ActorInstances) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeActorInstances deserializes a ActorInstances using msgpack
            func MDecodeActorInstances(d *msgpack.Decoder) (ActorInstances,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorInstance, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]ActorInstance, 0 ), err }
                        val := make([]ActorInstance, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeActorInstance(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a ActorInstances using cbor
            func (o *ActorInstances) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeActorInstances deserializes a ActorInstances using cbor
            func CDecodeActorInstances(d *cbor.Decoder) (ActorInstances,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ActorInstance, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]ActorInstance, 0 ), err }
                        val := make([]ActorInstance, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeActorInstance(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// A request to advertise/publish a link definition on a given lattice.
type AdvertiseLinkRequest struct {
// The ID of the lattice for this request
  LatticeId string 
  Link actor.LinkDefinition 
}

// MEncode serializes a AdvertiseLinkRequest using msgpack
            func (o *AdvertiseLinkRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("link")
o.Link.MEncode(encoder)

                return encoder.CheckError()
            }
            
            // MDecodeAdvertiseLinkRequest deserializes a AdvertiseLinkRequest using msgpack
            func MDecodeAdvertiseLinkRequest(d *msgpack.Decoder) (AdvertiseLinkRequest,error) {
                var val AdvertiseLinkRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
case "link":
val.Link,err = actor.MDecodeLinkDefinition(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a AdvertiseLinkRequest using cbor
            func (o *AdvertiseLinkRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("link")
o.Link.CEncode(encoder)

                return encoder.CheckError()
            }
            
            // CDecodeAdvertiseLinkRequest deserializes a AdvertiseLinkRequest using cbor
            func CDecodeAdvertiseLinkRequest(d *cbor.Decoder) (AdvertiseLinkRequest,error) {
                var val AdvertiseLinkRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
case "link":
val.Link,err = actor.CDecodeLinkDefinition(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type AnnotationMap map[string]string
// MEncode serializes a AnnotationMap using msgpack
            func (o *AnnotationMap) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // MDecodeAnnotationMap deserializes a AnnotationMap using msgpack
            func MDecodeAnnotationMap(d *msgpack.Decoder) (AnnotationMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,err := d.ReadMapSize()
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// CEncode serializes a AnnotationMap using cbor
            func (o *AnnotationMap) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // CDecodeAnnotationMap deserializes a AnnotationMap using cbor
            func CDecodeAnnotationMap(d *cbor.Decoder) (AnnotationMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported") }
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
type ConfigurationString string
// MEncode serializes a ConfigurationString using msgpack
            func (o *ConfigurationString) MEncode(encoder msgpack.Writer) error {
                encoder.WriteString(string(*o))
                return encoder.CheckError()
            }
            
            // MDecodeConfigurationString deserializes a ConfigurationString using msgpack
            func MDecodeConfigurationString(d *msgpack.Decoder) (ConfigurationString,error) {
                val,err := d.ReadString()
                  if err != nil {
                    return "",err
                  }
                  return ConfigurationString(val),nil
            }
// CEncode serializes a ConfigurationString using cbor
            func (o *ConfigurationString) CEncode(encoder cbor.Writer) error {
                encoder.WriteString(string(*o))
                return encoder.CheckError()
            }
            
            // CDecodeConfigurationString deserializes a ConfigurationString using cbor
            func CDecodeConfigurationString(d *cbor.Decoder) (ConfigurationString,error) {
                val,err := d.ReadString()
                  if err != nil {
                    return "",err
                  }
                  return ConfigurationString(val),nil
            }
type ConstraintMap map[string]string
// MEncode serializes a ConstraintMap using msgpack
            func (o *ConstraintMap) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // MDecodeConstraintMap deserializes a ConstraintMap using msgpack
            func MDecodeConstraintMap(d *msgpack.Decoder) (ConstraintMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,err := d.ReadMapSize()
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// CEncode serializes a ConstraintMap using cbor
            func (o *ConstraintMap) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // CDecodeConstraintMap deserializes a ConstraintMap using cbor
            func CDecodeConstraintMap(d *cbor.Decoder) (ConstraintMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported") }
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
type CtlKVList []KeyValueMap
// MEncode serializes a CtlKVList using msgpack
            func (o *CtlKVList) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeCtlKVList deserializes a CtlKVList using msgpack
            func MDecodeCtlKVList(d *msgpack.Decoder) (CtlKVList,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]KeyValueMap, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]KeyValueMap, 0 ), err }
                        val := make([]KeyValueMap, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeKeyValueMap(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a CtlKVList using cbor
            func (o *CtlKVList) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeCtlKVList deserializes a CtlKVList using cbor
            func CDecodeCtlKVList(d *cbor.Decoder) (CtlKVList,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]KeyValueMap, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]KeyValueMap, 0 ), err }
                        val := make([]KeyValueMap, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeKeyValueMap(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// Standard response for control interface operations
type CtlOperationAck struct {
  Accepted bool 
  Error string 
}

// MEncode serializes a CtlOperationAck using msgpack
            func (o *CtlOperationAck) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("accepted")
encoder.WriteBool(o.Accepted)
encoder.WriteString("error")
encoder.WriteString(o.Error)

                return encoder.CheckError()
            }
            
            // MDecodeCtlOperationAck deserializes a CtlOperationAck using msgpack
            func MDecodeCtlOperationAck(d *msgpack.Decoder) (CtlOperationAck,error) {
                var val CtlOperationAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "accepted":
val.Accepted,err = d.ReadBool()
case "error":
val.Error,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a CtlOperationAck using cbor
            func (o *CtlOperationAck) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("accepted")
encoder.WriteBool(o.Accepted)
encoder.WriteString("error")
encoder.WriteString(o.Error)

                return encoder.CheckError()
            }
            
            // CDecodeCtlOperationAck deserializes a CtlOperationAck using cbor
            func CDecodeCtlOperationAck(d *cbor.Decoder) (CtlOperationAck,error) {
                var val CtlOperationAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "accepted":
val.Accepted,err = d.ReadBool()
case "error":
val.Error,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A request to obtain claims from a given lattice
type GetClaimsRequest struct {
// The ID of the lattice for this request
  LatticeId string 
}

// MEncode serializes a GetClaimsRequest using msgpack
            func (o *GetClaimsRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeGetClaimsRequest deserializes a GetClaimsRequest using msgpack
            func MDecodeGetClaimsRequest(d *msgpack.Decoder) (GetClaimsRequest,error) {
                var val GetClaimsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a GetClaimsRequest using cbor
            func (o *GetClaimsRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeGetClaimsRequest deserializes a GetClaimsRequest using cbor
            func CDecodeGetClaimsRequest(d *cbor.Decoder) (GetClaimsRequest,error) {
                var val GetClaimsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A response containing the full list of known claims within the lattice
type GetClaimsResponse struct {
  Claims CtlKVList 
}

// MEncode serializes a GetClaimsResponse using msgpack
            func (o *GetClaimsResponse) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("claims")
o.Claims.MEncode(encoder)

                return encoder.CheckError()
            }
            
            // MDecodeGetClaimsResponse deserializes a GetClaimsResponse using msgpack
            func MDecodeGetClaimsResponse(d *msgpack.Decoder) (GetClaimsResponse,error) {
                var val GetClaimsResponse
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "claims":
val.Claims,err = MDecodeCtlKVList(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a GetClaimsResponse using cbor
            func (o *GetClaimsResponse) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("claims")
o.Claims.CEncode(encoder)

                return encoder.CheckError()
            }
            
            // CDecodeGetClaimsResponse deserializes a GetClaimsResponse using cbor
            func CDecodeGetClaimsResponse(d *cbor.Decoder) (GetClaimsResponse,error) {
                var val GetClaimsResponse
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "claims":
val.Claims,err = CDecodeCtlKVList(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A request to query the inventory of a given host within a given lattice
type GetHostInventoryRequest struct {
// The public key of the host being targeted for this request
  HostId string 
// The ID of the lattice for this request
  LatticeId string 
}

// MEncode serializes a GetHostInventoryRequest using msgpack
            func (o *GetHostInventoryRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeGetHostInventoryRequest deserializes a GetHostInventoryRequest using msgpack
            func MDecodeGetHostInventoryRequest(d *msgpack.Decoder) (GetHostInventoryRequest,error) {
                var val GetHostInventoryRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a GetHostInventoryRequest using cbor
            func (o *GetHostInventoryRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeGetHostInventoryRequest deserializes a GetHostInventoryRequest using cbor
            func CDecodeGetHostInventoryRequest(d *cbor.Decoder) (GetHostInventoryRequest,error) {
                var val GetHostInventoryRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A request to obtain the list of hosts responding within a given lattice
type GetHostsRequest struct {
// The ID of the lattice for which these credentials will be used
  LatticeId string 
}

// MEncode serializes a GetHostsRequest using msgpack
            func (o *GetHostsRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeGetHostsRequest deserializes a GetHostsRequest using msgpack
            func MDecodeGetHostsRequest(d *msgpack.Decoder) (GetHostsRequest,error) {
                var val GetHostsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a GetHostsRequest using cbor
            func (o *GetHostsRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeGetHostsRequest deserializes a GetHostsRequest using cbor
            func CDecodeGetHostsRequest(d *cbor.Decoder) (GetHostsRequest,error) {
                var val GetHostsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A summary representation of a host
type Host struct {
// Comma-delimited list of valid cluster issuer public keys as known
// to this host
  ClusterIssuers string 
// NATS server host used for the control interface
  CtlHost string 
  Id string 
// JetStream domain (if applicable) in use by this host
  JsDomain string 
// Hash map of label-value pairs for this host
  Labels *KeyValueMap 
// Lattice prefix/ID used by the host
  LatticePrefix string 
// NATS server host used for provider RPC
  ProvRpcHost string 
// NATS server host used for regular RPC
  RpcHost string 
// Human-friendly uptime description
  UptimeHuman string 
// uptime in seconds
  UptimeSeconds uint64 
// Current wasmCloud Host software version
  Version string 
}

// MEncode serializes a Host using msgpack
            func (o *Host) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(11)
encoder.WriteString("clusterIssuers")
encoder.WriteString(o.ClusterIssuers)
encoder.WriteString("ctlHost")
encoder.WriteString(o.CtlHost)
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("jsDomain")
encoder.WriteString(o.JsDomain)
encoder.WriteString("labels")
if o.Labels == nil {
                        encoder.WriteNil()
                    } else {
                        o.Labels.MEncode(encoder)
                    }
encoder.WriteString("latticePrefix")
encoder.WriteString(o.LatticePrefix)
encoder.WriteString("provRpcHost")
encoder.WriteString(o.ProvRpcHost)
encoder.WriteString("rpcHost")
encoder.WriteString(o.RpcHost)
encoder.WriteString("uptimeHuman")
encoder.WriteString(o.UptimeHuman)
encoder.WriteString("uptimeSeconds")
encoder.WriteUint64(o.UptimeSeconds)
encoder.WriteString("version")
encoder.WriteString(o.Version)

                return encoder.CheckError()
            }
            
            // MDecodeHost deserializes a Host using msgpack
            func MDecodeHost(d *msgpack.Decoder) (Host,error) {
                var val Host
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "clusterIssuers":
val.ClusterIssuers,err = d.ReadString()
case "ctlHost":
val.CtlHost,err = d.ReadString()
case "id":
val.Id,err = d.ReadString()
case "jsDomain":
val.JsDomain,err = d.ReadString()
case "labels":
fval,err := MDecodeKeyValueMap(d)
                  if err != nil { return val, err }
                  val.Labels = &fval
case "latticePrefix":
val.LatticePrefix,err = d.ReadString()
case "provRpcHost":
val.ProvRpcHost,err = d.ReadString()
case "rpcHost":
val.RpcHost,err = d.ReadString()
case "uptimeHuman":
val.UptimeHuman,err = d.ReadString()
case "uptimeSeconds":
val.UptimeSeconds,err = d.ReadUint64()
case "version":
val.Version,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a Host using cbor
            func (o *Host) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(11)
encoder.WriteString("clusterIssuers")
encoder.WriteString(o.ClusterIssuers)
encoder.WriteString("ctlHost")
encoder.WriteString(o.CtlHost)
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("jsDomain")
encoder.WriteString(o.JsDomain)
encoder.WriteString("labels")
if o.Labels == nil {
                        encoder.WriteNil()
                    } else {
                        o.Labels.CEncode(encoder)
                    }
encoder.WriteString("latticePrefix")
encoder.WriteString(o.LatticePrefix)
encoder.WriteString("provRpcHost")
encoder.WriteString(o.ProvRpcHost)
encoder.WriteString("rpcHost")
encoder.WriteString(o.RpcHost)
encoder.WriteString("uptimeHuman")
encoder.WriteString(o.UptimeHuman)
encoder.WriteString("uptimeSeconds")
encoder.WriteUint64(o.UptimeSeconds)
encoder.WriteString("version")
encoder.WriteString(o.Version)

                return encoder.CheckError()
            }
            
            // CDecodeHost deserializes a Host using cbor
            func CDecodeHost(d *cbor.Decoder) (Host,error) {
                var val Host
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "clusterIssuers":
val.ClusterIssuers,err = d.ReadString()
case "ctlHost":
val.CtlHost,err = d.ReadString()
case "id":
val.Id,err = d.ReadString()
case "jsDomain":
val.JsDomain,err = d.ReadString()
case "labels":
fval,err := CDecodeKeyValueMap(d)
                  if err != nil { return val, err }
                  val.Labels = &fval
case "latticePrefix":
val.LatticePrefix,err = d.ReadString()
case "provRpcHost":
val.ProvRpcHost,err = d.ReadString()
case "rpcHost":
val.RpcHost,err = d.ReadString()
case "uptimeHuman":
val.UptimeHuman,err = d.ReadString()
case "uptimeSeconds":
val.UptimeSeconds,err = d.ReadUint64()
case "version":
val.Version,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// Describes the known contents of a given host at the time of
// a query
type HostInventory struct {
// Actors running on this host.
  Actors ActorDescriptions 
// The host's unique ID
  HostId string 
// The host's labels
  Labels LabelsMap 
// Providers running on this host
  Providers ProviderDescriptions 
}

// MEncode serializes a HostInventory using msgpack
            func (o *HostInventory) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("actors")
o.Actors.MEncode(encoder)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("labels")
o.Labels.MEncode(encoder)
encoder.WriteString("providers")
o.Providers.MEncode(encoder)

                return encoder.CheckError()
            }
            
            // MDecodeHostInventory deserializes a HostInventory using msgpack
            func MDecodeHostInventory(d *msgpack.Decoder) (HostInventory,error) {
                var val HostInventory
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actors":
val.Actors,err = MDecodeActorDescriptions(d)
case "hostId":
val.HostId,err = d.ReadString()
case "labels":
val.Labels,err = MDecodeLabelsMap(d)
case "providers":
val.Providers,err = MDecodeProviderDescriptions(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a HostInventory using cbor
            func (o *HostInventory) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("actors")
o.Actors.CEncode(encoder)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("labels")
o.Labels.CEncode(encoder)
encoder.WriteString("providers")
o.Providers.CEncode(encoder)

                return encoder.CheckError()
            }
            
            // CDecodeHostInventory deserializes a HostInventory using cbor
            func CDecodeHostInventory(d *cbor.Decoder) (HostInventory,error) {
                var val HostInventory
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actors":
val.Actors,err = CDecodeActorDescriptions(d)
case "hostId":
val.HostId,err = d.ReadString()
case "labels":
val.Labels,err = CDecodeLabelsMap(d)
case "providers":
val.Providers,err = CDecodeProviderDescriptions(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type Hosts []Host
// MEncode serializes a Hosts using msgpack
            func (o *Hosts) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeHosts deserializes a Hosts using msgpack
            func MDecodeHosts(d *msgpack.Decoder) (Hosts,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]Host, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]Host, 0 ), err }
                        val := make([]Host, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeHost(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a Hosts using cbor
            func (o *Hosts) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeHosts deserializes a Hosts using cbor
            func CDecodeHosts(d *cbor.Decoder) (Hosts,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]Host, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]Host, 0 ), err }
                        val := make([]Host, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeHost(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
type KeyValueMap map[string]string
// MEncode serializes a KeyValueMap using msgpack
            func (o *KeyValueMap) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // MDecodeKeyValueMap deserializes a KeyValueMap using msgpack
            func MDecodeKeyValueMap(d *msgpack.Decoder) (KeyValueMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,err := d.ReadMapSize()
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// CEncode serializes a KeyValueMap using cbor
            func (o *KeyValueMap) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // CDecodeKeyValueMap deserializes a KeyValueMap using cbor
            func CDecodeKeyValueMap(d *cbor.Decoder) (KeyValueMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported") }
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
type LabelsMap map[string]string
// MEncode serializes a LabelsMap using msgpack
            func (o *LabelsMap) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // MDecodeLabelsMap deserializes a LabelsMap using msgpack
            func MDecodeLabelsMap(d *msgpack.Decoder) (LabelsMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,err := d.ReadMapSize()
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// CEncode serializes a LabelsMap using cbor
            func (o *LabelsMap) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        encoder.WriteString(val_o)
                    }        
                    
                return encoder.CheckError()
            }
            
            // CDecodeLabelsMap deserializes a LabelsMap using cbor
            func CDecodeLabelsMap(d *cbor.Decoder) (LabelsMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]string, 0), err
                        }
                       	size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported") }
                        if err != nil { return make(map[string]string, 0),err }
                        val := make(map[string]string, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := d.ReadString()
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// A list of link definitions
type LinkDefinitionList struct {
  Links actor.ActorLinks 
}

// MEncode serializes a LinkDefinitionList using msgpack
            func (o *LinkDefinitionList) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("links")
o.Links.MEncode(encoder)

                return encoder.CheckError()
            }
            
            // MDecodeLinkDefinitionList deserializes a LinkDefinitionList using msgpack
            func MDecodeLinkDefinitionList(d *msgpack.Decoder) (LinkDefinitionList,error) {
                var val LinkDefinitionList
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "links":
val.Links,err = actor.MDecodeActorLinks(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a LinkDefinitionList using cbor
            func (o *LinkDefinitionList) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(1)
encoder.WriteString("links")
o.Links.CEncode(encoder)

                return encoder.CheckError()
            }
            
            // CDecodeLinkDefinitionList deserializes a LinkDefinitionList using cbor
            func CDecodeLinkDefinitionList(d *cbor.Decoder) (LinkDefinitionList,error) {
                var val LinkDefinitionList
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "links":
val.Links,err = actor.CDecodeActorLinks(d)
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// One of a potential list of responses to a provider auction
type ProviderAuctionAck struct {
// The host ID of the "bidder" for this auction
  HostId string 
// The link name provided for the auction
  LinkName string 
// The original provider ref provided for the auction
  ProviderRef string 
}

// MEncode serializes a ProviderAuctionAck using msgpack
            func (o *ProviderAuctionAck) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // MDecodeProviderAuctionAck deserializes a ProviderAuctionAck using msgpack
            func MDecodeProviderAuctionAck(d *msgpack.Decoder) (ProviderAuctionAck,error) {
                var val ProviderAuctionAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ProviderAuctionAck using cbor
            func (o *ProviderAuctionAck) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // CDecodeProviderAuctionAck deserializes a ProviderAuctionAck using cbor
            func CDecodeProviderAuctionAck(d *cbor.Decoder) (ProviderAuctionAck,error) {
                var val ProviderAuctionAck
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ProviderAuctionAcks []ProviderAuctionAck
// MEncode serializes a ProviderAuctionAcks using msgpack
            func (o *ProviderAuctionAcks) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeProviderAuctionAcks deserializes a ProviderAuctionAcks using msgpack
            func MDecodeProviderAuctionAcks(d *msgpack.Decoder) (ProviderAuctionAcks,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ProviderAuctionAck, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]ProviderAuctionAck, 0 ), err }
                        val := make([]ProviderAuctionAck, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeProviderAuctionAck(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a ProviderAuctionAcks using cbor
            func (o *ProviderAuctionAcks) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeProviderAuctionAcks deserializes a ProviderAuctionAcks using cbor
            func CDecodeProviderAuctionAcks(d *cbor.Decoder) (ProviderAuctionAcks,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ProviderAuctionAck, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]ProviderAuctionAck, 0 ), err }
                        val := make([]ProviderAuctionAck, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeProviderAuctionAck(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// A request to locate a suitable host for a capability provider. The
// provider's unique identity (reference + link name) is used to rule
// out sites on which the provider is already running.
type ProviderAuctionRequest struct {
// The set of constraints to which a suitable target host must conform
  Constraints ConstraintMap 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// The link name of the provider
  LinkName string 
// The reference for the provider. Can be any one of the accepted
// forms of uniquely identifying a provider
  ProviderRef string 
}

// MEncode serializes a ProviderAuctionRequest using msgpack
            func (o *ProviderAuctionRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("constraints")
o.Constraints.MEncode(encoder)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // MDecodeProviderAuctionRequest deserializes a ProviderAuctionRequest using msgpack
            func MDecodeProviderAuctionRequest(d *msgpack.Decoder) (ProviderAuctionRequest,error) {
                var val ProviderAuctionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "constraints":
val.Constraints,err = MDecodeConstraintMap(d)
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ProviderAuctionRequest using cbor
            func (o *ProviderAuctionRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("constraints")
o.Constraints.CEncode(encoder)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // CDecodeProviderAuctionRequest deserializes a ProviderAuctionRequest using cbor
            func CDecodeProviderAuctionRequest(d *cbor.Decoder) (ProviderAuctionRequest,error) {
                var val ProviderAuctionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "constraints":
val.Constraints,err = CDecodeConstraintMap(d)
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A summary description of a capability provider within a host inventory
type ProviderDescription struct {
// The annotations that were used in the start request that produced
// this provider instance
  Annotations *AnnotationMap 
// Provider's unique 56-character ID
  Id string 
// Image reference for this provider, if applicable
  ImageRef string 
// Provider's link name
  LinkName string 
// Name of the provider, if one exists
  Name string 
// The revision of the provider
  Revision int32 
}

// MEncode serializes a ProviderDescription using msgpack
            func (o *ProviderDescription) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("imageRef")
encoder.WriteString(o.ImageRef)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("name")
encoder.WriteString(o.Name)
encoder.WriteString("revision")
encoder.WriteInt32(o.Revision)

                return encoder.CheckError()
            }
            
            // MDecodeProviderDescription deserializes a ProviderDescription using msgpack
            func MDecodeProviderDescription(d *msgpack.Decoder) (ProviderDescription,error) {
                var val ProviderDescription
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "id":
val.Id,err = d.ReadString()
case "imageRef":
val.ImageRef,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "name":
val.Name,err = d.ReadString()
case "revision":
val.Revision,err = d.ReadInt32()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ProviderDescription using cbor
            func (o *ProviderDescription) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("id")
encoder.WriteString(o.Id)
encoder.WriteString("imageRef")
encoder.WriteString(o.ImageRef)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("name")
encoder.WriteString(o.Name)
encoder.WriteString("revision")
encoder.WriteInt32(o.Revision)

                return encoder.CheckError()
            }
            
            // CDecodeProviderDescription deserializes a ProviderDescription using cbor
            func CDecodeProviderDescription(d *cbor.Decoder) (ProviderDescription,error) {
                var val ProviderDescription
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "id":
val.Id,err = d.ReadString()
case "imageRef":
val.ImageRef,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "name":
val.Name,err = d.ReadString()
case "revision":
val.Revision,err = d.ReadInt32()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ProviderDescriptions []ProviderDescription
// MEncode serializes a ProviderDescriptions using msgpack
            func (o *ProviderDescriptions) MEncode(encoder msgpack.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.MEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // MDecodeProviderDescriptions deserializes a ProviderDescriptions using msgpack
            func MDecodeProviderDescriptions(d *msgpack.Decoder) (ProviderDescriptions,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ProviderDescription, 0), err
                        }
                       	size,err := d.ReadArraySize()
                        if err != nil { return make([]ProviderDescription, 0 ), err }
                        val := make([]ProviderDescription, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := MDecodeProviderDescription(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
// CEncode serializes a ProviderDescriptions using cbor
            func (o *ProviderDescriptions) CEncode(encoder cbor.Writer) error {
                
                    encoder.WriteArraySize(uint32(len(*o)))
                    for _,item_o := range *o {
                        item_o.CEncode(encoder)
                    }
                    
                return encoder.CheckError()
            }
            
            // CDecodeProviderDescriptions deserializes a ProviderDescriptions using cbor
            func CDecodeProviderDescriptions(d *cbor.Decoder) (ProviderDescriptions,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make([]ProviderDescription, 0), err
                        }
                       	size,indef,err := d.ReadArraySize()
                if err != nil && indef { err = cbor.NewReadError("indefinite arrays not supported") }
                        if err != nil { return make([]ProviderDescription, 0 ), err }
                        val := make([]ProviderDescription, size)
                        for i := uint32(0); i < size; i++ {
                           item,err := CDecodeProviderDescription(d)
                           if err != nil { return val, err }
                           val = append(val,item)
                        }
                        return val,nil
            }
type RegistryCredential struct {
  Password string 
// The type of the registry (either "oci" or "bindle")
  RegistryType string 
// If supplied, token authentication will be used for the registry
  Token string 
// If supplied, username and password will be used for HTTP Basic authentication
  Username string 
}

// MEncode serializes a RegistryCredential using msgpack
            func (o *RegistryCredential) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("password")
encoder.WriteString(o.Password)
encoder.WriteString("registryType")
encoder.WriteString(o.RegistryType)
encoder.WriteString("token")
encoder.WriteString(o.Token)
encoder.WriteString("username")
encoder.WriteString(o.Username)

                return encoder.CheckError()
            }
            
            // MDecodeRegistryCredential deserializes a RegistryCredential using msgpack
            func MDecodeRegistryCredential(d *msgpack.Decoder) (RegistryCredential,error) {
                var val RegistryCredential
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "password":
val.Password,err = d.ReadString()
case "registryType":
val.RegistryType,err = d.ReadString()
case "token":
val.Token,err = d.ReadString()
case "username":
val.Username,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a RegistryCredential using cbor
            func (o *RegistryCredential) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("password")
encoder.WriteString(o.Password)
encoder.WriteString("registryType")
encoder.WriteString(o.RegistryType)
encoder.WriteString("token")
encoder.WriteString(o.Token)
encoder.WriteString("username")
encoder.WriteString(o.Username)

                return encoder.CheckError()
            }
            
            // CDecodeRegistryCredential deserializes a RegistryCredential using cbor
            func CDecodeRegistryCredential(d *cbor.Decoder) (RegistryCredential,error) {
                var val RegistryCredential
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "password":
val.Password,err = d.ReadString()
case "registryType":
val.RegistryType,err = d.ReadString()
case "token":
val.Token,err = d.ReadString()
case "username":
val.Username,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A set of credentials to be used for fetching from specific registries
type RegistryCredentialMap map[string]RegistryCredential
// MEncode serializes a RegistryCredentialMap using msgpack
            func (o *RegistryCredentialMap) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        val_o.MEncode(encoder)
                    }        
                    
                return encoder.CheckError()
            }
            
            // MDecodeRegistryCredentialMap deserializes a RegistryCredentialMap using msgpack
            func MDecodeRegistryCredentialMap(d *msgpack.Decoder) (RegistryCredentialMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]RegistryCredential, 0), err
                        }
                       	size,err := d.ReadMapSize()
                        if err != nil { return make(map[string]RegistryCredential, 0),err }
                        val := make(map[string]RegistryCredential, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := MDecodeRegistryCredential(d)
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// CEncode serializes a RegistryCredentialMap using cbor
            func (o *RegistryCredentialMap) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(uint32(len(*o)))
                    for key_o,val_o := range *o {
                        encoder.WriteString(key_o)
                        val_o.CEncode(encoder)
                    }        
                    
                return encoder.CheckError()
            }
            
            // CDecodeRegistryCredentialMap deserializes a RegistryCredentialMap using cbor
            func CDecodeRegistryCredentialMap(d *cbor.Decoder) (RegistryCredentialMap,error) {
                isNil,err := d.IsNextNil()
                        if err != nil || isNil {
                       		return make(map[string]RegistryCredential, 0), err
                        }
                       	size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported") }
                        if err != nil { return make(map[string]RegistryCredential, 0),err }
                        val := make(map[string]RegistryCredential, size)
                        for i := uint32(0); i < size; i++ {
                           k,_ := d.ReadString()
                           v,err := CDecodeRegistryCredential(d)
                           if err != nil { return val, err }
                           val[k] = v
                        }
                        return val,nil
            }
// A request to remove a link definition and detach the relevant actor
// from the given provider
type RemoveLinkDefinitionRequest struct {
// The actor's public key. This cannot be an image reference
  ActorId string 
// The provider contract
  ContractId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// The provider's link name
  LinkName string 
}

// MEncode serializes a RemoveLinkDefinitionRequest using msgpack
            func (o *RemoveLinkDefinitionRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("contractId")
encoder.WriteString(o.ContractId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)

                return encoder.CheckError()
            }
            
            // MDecodeRemoveLinkDefinitionRequest deserializes a RemoveLinkDefinitionRequest using msgpack
            func MDecodeRemoveLinkDefinitionRequest(d *msgpack.Decoder) (RemoveLinkDefinitionRequest,error) {
                var val RemoveLinkDefinitionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "contractId":
val.ContractId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a RemoveLinkDefinitionRequest using cbor
            func (o *RemoveLinkDefinitionRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(4)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("contractId")
encoder.WriteString(o.ContractId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)

                return encoder.CheckError()
            }
            
            // CDecodeRemoveLinkDefinitionRequest deserializes a RemoveLinkDefinitionRequest using cbor
            func CDecodeRemoveLinkDefinitionRequest(d *cbor.Decoder) (RemoveLinkDefinitionRequest,error) {
                var val RemoveLinkDefinitionRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "contractId":
val.ContractId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type ScaleActorCommand struct {
// Public Key ID of the actor to scale
  ActorId string 
// Reference for the actor. Can be any of the acceptable forms of unique identification
  ActorRef string 
// Optional set of annotations used to describe the nature of this actor scale command. For
// example, autonomous agents may wish to "tag" scale requests as part of a given deployment
  Annotations *AnnotationMap 
// The target number of actors
  Count uint16 
// Host ID on which to scale this actor
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
}

// MEncode serializes a ScaleActorCommand using msgpack
            func (o *ScaleActorCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeScaleActorCommand deserializes a ScaleActorCommand using msgpack
            func MDecodeScaleActorCommand(d *msgpack.Decoder) (ScaleActorCommand,error) {
                var val ScaleActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "actorRef":
val.ActorRef,err = d.ReadString()
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a ScaleActorCommand using cbor
            func (o *ScaleActorCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeScaleActorCommand deserializes a ScaleActorCommand using cbor
            func CDecodeScaleActorCommand(d *cbor.Decoder) (ScaleActorCommand,error) {
                var val ScaleActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "actorRef":
val.ActorRef,err = d.ReadString()
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// Represents a request to set/store the credentials that correspond to a given lattice ID.
type SetLatticeCredentialsRequest struct {
// If there is a JS domain required for communicating with the underlying KV metadata
// bucket for this lattice, then that should be supplied in this parameter. Otherwise,
// leave it blank
  JsDomain string 
// The ID of the lattice for which these credentials will be used
  LatticeId string 
// If natsUrl is supplied, then the capability provider will use this URL (and port) for
// establishing a connection for the given lattice.
  NatsUrl string 
// If supplied, contains the user JWT to be used for authenticating against NATS to allow
// access to the indicated lattice. If not supplied, the capability provider will assume/set
// anonymous access for this lattice.
  UserJwt string 
// If userJwt is supplied, user seed must also be supplied and is the seed key used for user
// authentication against NATS for this lattice.
  UserSeed string 
}

// MEncode serializes a SetLatticeCredentialsRequest using msgpack
            func (o *SetLatticeCredentialsRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("jsDomain")
encoder.WriteString(o.JsDomain)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("natsUrl")
encoder.WriteString(o.NatsUrl)
encoder.WriteString("userJwt")
encoder.WriteString(o.UserJwt)
encoder.WriteString("userSeed")
encoder.WriteString(o.UserSeed)

                return encoder.CheckError()
            }
            
            // MDecodeSetLatticeCredentialsRequest deserializes a SetLatticeCredentialsRequest using msgpack
            func MDecodeSetLatticeCredentialsRequest(d *msgpack.Decoder) (SetLatticeCredentialsRequest,error) {
                var val SetLatticeCredentialsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "jsDomain":
val.JsDomain,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "natsUrl":
val.NatsUrl,err = d.ReadString()
case "userJwt":
val.UserJwt,err = d.ReadString()
case "userSeed":
val.UserSeed,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a SetLatticeCredentialsRequest using cbor
            func (o *SetLatticeCredentialsRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("jsDomain")
encoder.WriteString(o.JsDomain)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("natsUrl")
encoder.WriteString(o.NatsUrl)
encoder.WriteString("userJwt")
encoder.WriteString(o.UserJwt)
encoder.WriteString("userSeed")
encoder.WriteString(o.UserSeed)

                return encoder.CheckError()
            }
            
            // CDecodeSetLatticeCredentialsRequest deserializes a SetLatticeCredentialsRequest using cbor
            func CDecodeSetLatticeCredentialsRequest(d *cbor.Decoder) (SetLatticeCredentialsRequest,error) {
                var val SetLatticeCredentialsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "jsDomain":
val.JsDomain,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "natsUrl":
val.NatsUrl,err = d.ReadString()
case "userJwt":
val.UserJwt,err = d.ReadString()
case "userSeed":
val.UserSeed,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
type SetRegistryCredentialsRequest struct {
  Credentials *RegistryCredentialMap 
// The ID of the lattice on which this request will be performed
  LatticeId string 
}

// MEncode serializes a SetRegistryCredentialsRequest using msgpack
            func (o *SetRegistryCredentialsRequest) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("credentials")
if o.Credentials == nil {
                        encoder.WriteNil()
                    } else {
                        o.Credentials.MEncode(encoder)
                    }
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeSetRegistryCredentialsRequest deserializes a SetRegistryCredentialsRequest using msgpack
            func MDecodeSetRegistryCredentialsRequest(d *msgpack.Decoder) (SetRegistryCredentialsRequest,error) {
                var val SetRegistryCredentialsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "credentials":
fval,err := MDecodeRegistryCredentialMap(d)
                  if err != nil { return val, err }
                  val.Credentials = &fval
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a SetRegistryCredentialsRequest using cbor
            func (o *SetRegistryCredentialsRequest) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(2)
encoder.WriteString("credentials")
if o.Credentials == nil {
                        encoder.WriteNil()
                    } else {
                        o.Credentials.CEncode(encoder)
                    }
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeSetRegistryCredentialsRequest deserializes a SetRegistryCredentialsRequest using cbor
            func CDecodeSetRegistryCredentialsRequest(d *cbor.Decoder) (SetRegistryCredentialsRequest,error) {
                var val SetRegistryCredentialsRequest
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "credentials":
fval,err := CDecodeRegistryCredentialMap(d)
                  if err != nil { return val, err }
                  val.Credentials = &fval
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A command sent to a specific host instructing it to start the actor
// indicated by the reference.
type StartActorCommand struct {
// Reference for the actor. This can be either a bindle or OCI reference
  ActorRef string 
// Optional set of annotations used to describe the nature of this actor start command. For
// example, autonomous agents may wish to "tag" start requests as part of a given deployment
  Annotations *AnnotationMap 
// The number of actors to start
// A zero value will be interpreted as 1.
  Count uint16 
// Host ID on which this actor should start
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
}

// MEncode serializes a StartActorCommand using msgpack
            func (o *StartActorCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeStartActorCommand deserializes a StartActorCommand using msgpack
            func MDecodeStartActorCommand(d *msgpack.Decoder) (StartActorCommand,error) {
                var val StartActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a StartActorCommand using cbor
            func (o *StartActorCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorRef")
encoder.WriteString(o.ActorRef)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeStartActorCommand deserializes a StartActorCommand using cbor
            func CDecodeStartActorCommand(d *cbor.Decoder) (StartActorCommand,error) {
                var val StartActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorRef":
val.ActorRef,err = d.ReadString()
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A command sent to a host requesting a capability provider be started with the
// given link name and optional configuration.
type StartProviderCommand struct {
// Optional set of annotations used to describe the nature of this provider start command. For
// example, autonomous agents may wish to "tag" start requests as part of a given deployment
  Annotations *AnnotationMap 
// Optional provider configuration in the form of an opaque string. Many
// providers prefer base64-encoded JSON here, though that data should never
// exceed 500KB
  Configuration ConfigurationString 
// The host ID on which to start the provider
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// The link name of the provider to be started
  LinkName string 
// The image reference of the provider to be started
  ProviderRef string 
}

// MEncode serializes a StartProviderCommand using msgpack
            func (o *StartProviderCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("configuration")
o.Configuration.MEncode(encoder)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // MDecodeStartProviderCommand deserializes a StartProviderCommand using msgpack
            func MDecodeStartProviderCommand(d *msgpack.Decoder) (StartProviderCommand,error) {
                var val StartProviderCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "configuration":
val.Configuration,err = MDecodeConfigurationString(d)
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a StartProviderCommand using cbor
            func (o *StartProviderCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("configuration")
o.Configuration.CEncode(encoder)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerRef")
encoder.WriteString(o.ProviderRef)

                return encoder.CheckError()
            }
            
            // CDecodeStartProviderCommand deserializes a StartProviderCommand using cbor
            func CDecodeStartProviderCommand(d *cbor.Decoder) (StartProviderCommand,error) {
                var val StartProviderCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "configuration":
val.Configuration,err = CDecodeConfigurationString(d)
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerRef":
val.ProviderRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A command sent to a host to request that instances of a given actor
// be terminated on that host
type StopActorCommand struct {
// The public key of the actor to stop
  ActorId string 
// Optional set of annotations used to describe the nature of this
// stop request. If supplied, the only instances of this actor with these
// annotations will be stopped
  Annotations *AnnotationMap 
// The number of actors to stop
// A zero value means stop all actors
  Count uint16 
// The ID of the target host
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
}

// MEncode serializes a StopActorCommand using msgpack
            func (o *StopActorCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // MDecodeStopActorCommand deserializes a StopActorCommand using msgpack
            func MDecodeStopActorCommand(d *msgpack.Decoder) (StopActorCommand,error) {
                var val StopActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a StopActorCommand using cbor
            func (o *StopActorCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("count")
encoder.WriteUint16(o.Count)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)

                return encoder.CheckError()
            }
            
            // CDecodeStopActorCommand deserializes a StopActorCommand using cbor
            func CDecodeStopActorCommand(d *cbor.Decoder) (StopActorCommand,error) {
                var val StopActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "count":
val.Count,err = d.ReadUint16()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A command sent to request that the given host purge and stop
type StopHostCommand struct {
// The ID of the target host
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// An optional timeout, in seconds
  Timeout uint64 
}

// MEncode serializes a StopHostCommand using msgpack
            func (o *StopHostCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("timeout")
encoder.WriteUint64(o.Timeout)

                return encoder.CheckError()
            }
            
            // MDecodeStopHostCommand deserializes a StopHostCommand using msgpack
            func MDecodeStopHostCommand(d *msgpack.Decoder) (StopHostCommand,error) {
                var val StopHostCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "timeout":
val.Timeout,err = d.ReadUint64()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a StopHostCommand using cbor
            func (o *StopHostCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(3)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("timeout")
encoder.WriteUint64(o.Timeout)

                return encoder.CheckError()
            }
            
            // CDecodeStopHostCommand deserializes a StopHostCommand using cbor
            func CDecodeStopHostCommand(d *cbor.Decoder) (StopHostCommand,error) {
                var val StopHostCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "timeout":
val.Timeout,err = d.ReadUint64()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A request to stop the given provider on the indicated host
type StopProviderCommand struct {
// Optional set of annotations used to describe the nature of this
// stop request
  Annotations *AnnotationMap 
// Contract ID of the capability provider
  ContractId string 
// Host ID on which to stop the provider
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// Link name for this provider
  LinkName string 
// The public key of the capability provider to stop
  ProviderId string 
}

// MEncode serializes a StopProviderCommand using msgpack
            func (o *StopProviderCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("contractId")
encoder.WriteString(o.ContractId)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerId")
encoder.WriteString(o.ProviderId)

                return encoder.CheckError()
            }
            
            // MDecodeStopProviderCommand deserializes a StopProviderCommand using msgpack
            func MDecodeStopProviderCommand(d *msgpack.Decoder) (StopProviderCommand,error) {
                var val StopProviderCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "contractId":
val.ContractId,err = d.ReadString()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerId":
val.ProviderId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a StopProviderCommand using cbor
            func (o *StopProviderCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(6)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("contractId")
encoder.WriteString(o.ContractId)
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("linkName")
encoder.WriteString(o.LinkName)
encoder.WriteString("providerId")
encoder.WriteString(o.ProviderId)

                return encoder.CheckError()
            }
            
            // CDecodeStopProviderCommand deserializes a StopProviderCommand using cbor
            func CDecodeStopProviderCommand(d *cbor.Decoder) (StopProviderCommand,error) {
                var val StopProviderCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "contractId":
val.ContractId,err = d.ReadString()
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "linkName":
val.LinkName,err = d.ReadString()
case "providerId":
val.ProviderId,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// A command instructing a specific host to perform a live update
// on the indicated actor by supplying a new image reference. Note that
// live updates are only possible through image references
type UpdateActorCommand struct {
// The actor's 56-character unique ID
  ActorId string 
// Optional set of annotations used to describe the nature of this
// update request. Only actor instances that have matching annotations
// will be upgraded, allowing for instance isolation by
  Annotations *AnnotationMap 
// The host ID of the host to perform the live update
  HostId string 
// The ID of the lattice on which this request will be performed
  LatticeId string 
// The new image reference of the upgraded version of this actor
  NewActorRef string 
}

// MEncode serializes a UpdateActorCommand using msgpack
            func (o *UpdateActorCommand) MEncode(encoder msgpack.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.MEncode(encoder)
                    }
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("newActorRef")
encoder.WriteString(o.NewActorRef)

                return encoder.CheckError()
            }
            
            // MDecodeUpdateActorCommand deserializes a UpdateActorCommand using msgpack
            func MDecodeUpdateActorCommand(d *msgpack.Decoder) (UpdateActorCommand,error) {
                var val UpdateActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,err := d.ReadMapSize()
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "annotations":
fval,err := MDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "newActorRef":
val.NewActorRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// CEncode serializes a UpdateActorCommand using cbor
            func (o *UpdateActorCommand) CEncode(encoder cbor.Writer) error {
                encoder.WriteMapSize(5)
encoder.WriteString("actorId")
encoder.WriteString(o.ActorId)
encoder.WriteString("annotations")
if o.Annotations == nil {
                        encoder.WriteNil()
                    } else {
                        o.Annotations.CEncode(encoder)
                    }
encoder.WriteString("hostId")
encoder.WriteString(o.HostId)
encoder.WriteString("latticeId")
encoder.WriteString(o.LatticeId)
encoder.WriteString("newActorRef")
encoder.WriteString(o.NewActorRef)

                return encoder.CheckError()
            }
            
            // CDecodeUpdateActorCommand deserializes a UpdateActorCommand using cbor
            func CDecodeUpdateActorCommand(d *cbor.Decoder) (UpdateActorCommand,error) {
                var val UpdateActorCommand
            isNil,err := d.IsNextNil()
            if err != nil || isNil { 
                return val,err 
            }
            size,indef,err := d.ReadMapSize()
                if err != nil && indef { err = cbor.NewReadError("indefinite maps not supported")}
            if err != nil { return val,err }
            for i := uint32(0); i < size; i++ {
                field,err := d.ReadString()
                if err != nil { return val,err }
                switch field {
case "actorId":
val.ActorId,err = d.ReadString()
case "annotations":
fval,err := CDecodeAnnotationMap(d)
                  if err != nil { return val, err }
                  val.Annotations = &fval
case "hostId":
val.HostId,err = d.ReadString()
case "latticeId":
val.LatticeId,err = d.ReadString()
case "newActorRef":
val.NewActorRef,err = d.ReadString()
 default: 
                err = d.Skip()
            }
            if err != nil {
                return val, err
            }
            }
            return val,nil
            }
// Lattice Controller - Describes the interface used for actors
// to communicate with a lattice controller, enabling developers
// to deploy actors that can manipulate the lattice in which they're
// running.
type LatticeController interface {
// Seek out a list of suitable hosts for a capability provider given
// a set of host label constraints. Hosts on which this provider is already
// running will not be among the successful "bidders" in this auction.
AuctionProvider(ctx *actor.Context, arg ProviderAuctionRequest) (*ProviderAuctionAcks, error)
// Seek out a list of suitable hosts for an actor given a set of host
// label constraints.
AuctionActor(ctx *actor.Context, arg ActorAuctionRequest) (*ActorAuctionAcks, error)
// Queries the list of hosts currently visible to the lattice. This is
// a "gather" operation and so can be influenced by short timeouts,
// network partition events, etc. The sole input to this query is the
// lattice ID on which the request takes place.
GetHosts(ctx *actor.Context, arg string) (*Hosts, error)
// Queries for the contents of a host given the supplied 56-character unique ID
GetHostInventory(ctx *actor.Context, arg GetHostInventoryRequest) (*HostInventory, error)
// Queries the lattice for the list of known/cached claims by taking the response
// from the first host that answers the query. The sole input to this request is
// the lattice ID on which the request takes place.
GetClaims(ctx *actor.Context, arg string) (*GetClaimsResponse, error)
// Instructs a given host to scale the indicated actor
ScaleActor(ctx *actor.Context, arg ScaleActorCommand) (*CtlOperationAck, error)
// Instructs a given host to start the indicated actor
StartActor(ctx *actor.Context, arg StartActorCommand) (*CtlOperationAck, error)
// Publish a link definition into the lattice, allowing it to be cached and
// delivered to the appropriate capability provider instances
AdvertiseLink(ctx *actor.Context, arg AdvertiseLinkRequest) (*CtlOperationAck, error)
// Requests the removal of a link definition. The definition will be removed
// from the cache and the relevant capability providers will be given a chance
// to de-provision any used resources
RemoveLink(ctx *actor.Context, arg RemoveLinkDefinitionRequest) (*CtlOperationAck, error)
// Queries all current link definitions in the specified lattice. The first host
// that receives this response will reply with the contents of the distributed
// cache
GetLinks(ctx *actor.Context, arg string) (*LinkDefinitionList, error)
// Requests that a specific host perform a live update on the indicated
// actor
UpdateActor(ctx *actor.Context, arg UpdateActorCommand) (*CtlOperationAck, error)
// Requests that the given host start the indicated capability provider
StartProvider(ctx *actor.Context, arg StartProviderCommand) (*CtlOperationAck, error)
// Requests that the given capability provider be stopped on the indicated host
StopProvider(ctx *actor.Context, arg StopProviderCommand) (*CtlOperationAck, error)
// Requests that an actor be stopped on the given host
StopActor(ctx *actor.Context, arg StopActorCommand) (*CtlOperationAck, error)
StopHost(ctx *actor.Context, arg StopHostCommand) (*CtlOperationAck, error)
// Instructs the provider to store the NATS credentials/URL for a given lattice. This is
// designed to allow a single capability provider (or multiple instances of the same) to manage
// multiple lattices, reducing overhead and making it easier to support secure multi-tenancy of
// lattices.
SetLatticeCredentials(ctx *actor.Context, arg SetLatticeCredentialsRequest) (*CtlOperationAck, error)
// Instructs all listening hosts to use the enclosed credential map for
// authentication to secure artifact (OCI/bindle) registries. Any host that
// receives this message will _delete_ its previous credential map and replace
// it with the enclosed. The credential map for a lattice can be purged by sending
// this message with an empty map
SetRegistryCredentials(ctx *actor.Context, arg SetRegistryCredentialsRequest) error
}


        // LatticeControllerHandler is called by an actor during `main` to generate a dispatch handler
        // The output of this call should be passed into `actor.RegisterHandlers`
        func LatticeControllerHandler(actor_ LatticeController) actor.Handler {
            return actor.NewHandler("LatticeController", &LatticeControllerReceiver{}, actor_)
        }
// LatticeControllerContractId returns the capability contract id for this interface
                func LatticeControllerContractId() string { return "wasmcloud:latticecontrol" } 
                
// LatticeControllerReceiver receives messages defined in the LatticeController service interface
// Lattice Controller - Describes the interface used for actors
// to communicate with a lattice controller, enabling developers
// to deploy actors that can manipulate the lattice in which they're
// running.
type LatticeControllerReceiver struct {}
func (r* LatticeControllerReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
                svc_,_ := svc.(LatticeController)
                switch message.Method {
                 
case "AuctionProvider" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeProviderAuctionRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.AuctionProvider (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.AuctionProvider", Arg: buf }, nil
                    }
case "AuctionActor" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeActorAuctionRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.AuctionActor (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.AuctionActor", Arg: buf }, nil
                    }
case "GetHosts" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := d.ReadString()
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.GetHosts (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.GetHosts", Arg: buf }, nil
                    }
case "GetHostInventory" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeGetHostInventoryRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.GetHostInventory (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.GetHostInventory", Arg: buf }, nil
                    }
case "GetClaims" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := d.ReadString()
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.GetClaims (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.GetClaims", Arg: buf }, nil
                    }
case "ScaleActor" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeScaleActorCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.ScaleActor (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.ScaleActor", Arg: buf }, nil
                    }
case "StartActor" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeStartActorCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.StartActor (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.StartActor", Arg: buf }, nil
                    }
case "AdvertiseLink" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeAdvertiseLinkRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.AdvertiseLink (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.AdvertiseLink", Arg: buf }, nil
                    }
case "RemoveLink" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeRemoveLinkDefinitionRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.RemoveLink (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.RemoveLink", Arg: buf }, nil
                    }
case "GetLinks" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := d.ReadString()
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.GetLinks (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.GetLinks", Arg: buf }, nil
                    }
case "UpdateActor" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeUpdateActorCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.UpdateActor (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.UpdateActor", Arg: buf }, nil
                    }
case "StartProvider" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeStartProviderCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.StartProvider (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.StartProvider", Arg: buf }, nil
                    }
case "StopProvider" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeStopProviderCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.StopProvider (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.StopProvider", Arg: buf }, nil
                    }
case "StopActor" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeStopActorCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.StopActor (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.StopActor", Arg: buf }, nil
                    }
case "StopHost" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeStopHostCommand(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.StopHost (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.StopHost", Arg: buf }, nil
                    }
case "SetLatticeCredentials" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeSetLatticeCredentialsRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
resp, err := svc_.SetLatticeCredentials (ctx, value)
                if err != nil { 
                    return nil,err
                }

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    resp.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    encoder := msgpack.NewEncoder(buf)
            	    enc := &encoder
                    resp.MEncode(enc)
 return &actor.Message { Method: "LatticeController.SetLatticeCredentials", Arg: buf }, nil
                    }
case "SetRegistryCredentials" : {

                        d := msgpack.NewDecoder(message.Arg)
                        value,err_ := MDecodeSetRegistryCredentialsRequest(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        
err := svc_.SetRegistryCredentials (ctx, value)
                if err != nil { 
                    return nil,err
                }
buf := make([]byte, 0)
 return &actor.Message { Method: "LatticeController.SetRegistryCredentials", Arg: buf }, nil
                    }
default: 
                   return nil, actor.NewRpcError("MethodNotHandled", "LatticeController." + message.Method)
               }
            }
            
// LatticeControllerSender sends messages to a LatticeController service
// Lattice Controller - Describes the interface used for actors
// to communicate with a lattice controller, enabling developers
// to deploy actors that can manipulate the lattice in which they're
// running.
type LatticeControllerSender struct { transport actor.Transport }
            
            
                // NewProvider constructs a client for sending to a LatticeController provider
                // implementing the 'wasmcloud:latticecontrol' capability contract, with the "default" link
                func NewProviderLatticeController() *LatticeControllerSender {
                    transport := actor.ToProvider("wasmcloud:latticecontrol", "default")
                    return &LatticeControllerSender { transport: transport }
                }

                // NewProviderLatticeControllerLink constructs a client for sending to a LatticeController provider
                // implementing the 'wasmcloud:latticecontrol' capability contract, with the specified link name
                func NewProviderLatticeControllerLink(linkName string) *LatticeControllerSender {
                    transport :=  actor.ToProvider("wasmcloud:latticecontrol", linkName)
                    return &LatticeControllerSender { transport: transport }
                }
                
// Seek out a list of suitable hosts for a capability provider given
// a set of host label constraints. Hosts on which this provider is already
// running will not be among the successful "bidders" in this auction.
func (s *LatticeControllerSender) AuctionProvider(ctx *actor.Context, arg ProviderAuctionRequest) (*ProviderAuctionAcks, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.AuctionProvider", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeProviderAuctionAcks(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Seek out a list of suitable hosts for an actor given a set of host
// label constraints.
func (s *LatticeControllerSender) AuctionActor(ctx *actor.Context, arg ActorAuctionRequest) (*ActorAuctionAcks, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.AuctionActor", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeActorAuctionAcks(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Queries the list of hosts currently visible to the lattice. This is
// a "gather" operation and so can be influenced by short timeouts,
// network partition events, etc. The sole input to this query is the
// lattice ID on which the request takes place.
func (s *LatticeControllerSender) GetHosts(ctx *actor.Context, arg string) (*Hosts, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    size_enc.WriteString(arg) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    enc.WriteString(arg)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.GetHosts", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeHosts(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Queries for the contents of a host given the supplied 56-character unique ID
func (s *LatticeControllerSender) GetHostInventory(ctx *actor.Context, arg GetHostInventoryRequest) (*HostInventory, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.GetHostInventory", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeHostInventory(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Queries the lattice for the list of known/cached claims by taking the response
// from the first host that answers the query. The sole input to this request is
// the lattice ID on which the request takes place.
func (s *LatticeControllerSender) GetClaims(ctx *actor.Context, arg string) (*GetClaimsResponse, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    size_enc.WriteString(arg) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    enc.WriteString(arg)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.GetClaims", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeGetClaimsResponse(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Instructs a given host to scale the indicated actor
func (s *LatticeControllerSender) ScaleActor(ctx *actor.Context, arg ScaleActorCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.ScaleActor", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Instructs a given host to start the indicated actor
func (s *LatticeControllerSender) StartActor(ctx *actor.Context, arg StartActorCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.StartActor", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Publish a link definition into the lattice, allowing it to be cached and
// delivered to the appropriate capability provider instances
func (s *LatticeControllerSender) AdvertiseLink(ctx *actor.Context, arg AdvertiseLinkRequest) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.AdvertiseLink", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Requests the removal of a link definition. The definition will be removed
// from the cache and the relevant capability providers will be given a chance
// to de-provision any used resources
func (s *LatticeControllerSender) RemoveLink(ctx *actor.Context, arg RemoveLinkDefinitionRequest) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.RemoveLink", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Queries all current link definitions in the specified lattice. The first host
// that receives this response will reply with the contents of the distributed
// cache
func (s *LatticeControllerSender) GetLinks(ctx *actor.Context, arg string) (*LinkDefinitionList, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    size_enc.WriteString(arg) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    enc.WriteString(arg)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.GetLinks", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeLinkDefinitionList(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Requests that a specific host perform a live update on the indicated
// actor
func (s *LatticeControllerSender) UpdateActor(ctx *actor.Context, arg UpdateActorCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.UpdateActor", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Requests that the given host start the indicated capability provider
func (s *LatticeControllerSender) StartProvider(ctx *actor.Context, arg StartProviderCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.StartProvider", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Requests that the given capability provider be stopped on the indicated host
func (s *LatticeControllerSender) StopProvider(ctx *actor.Context, arg StopProviderCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.StopProvider", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Requests that an actor be stopped on the given host
func (s *LatticeControllerSender) StopActor(ctx *actor.Context, arg StopActorCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.StopActor", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
func (s *LatticeControllerSender) StopHost(ctx *actor.Context, arg StopHostCommand) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.StopHost", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Instructs the provider to store the NATS credentials/URL for a given lattice. This is
// designed to allow a single capability provider (or multiple instances of the same) to manage
// multiple lattices, reducing overhead and making it easier to support secure multi-tenancy of
// lattices.
func (s *LatticeControllerSender) SetLatticeCredentials(ctx *actor.Context, arg SetLatticeCredentialsRequest) (*CtlOperationAck, error) {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
out_buf,_ := s.transport.Send(ctx, actor.Message{ Method: "LatticeController.SetLatticeCredentials", Arg:buf })
d := msgpack.NewDecoder(out_buf)
                        resp,err_ := MDecodeCtlOperationAck(&d)
                        if err_ != nil { 
                            return nil,err_
                        }
                        return &resp,nil
                     }
// Instructs all listening hosts to use the enclosed credential map for
// authentication to secure artifact (OCI/bindle) registries. Any host that
// receives this message will _delete_ its previous credential map and replace
// it with the enclosed. The credential map for a lattice can be purged by sending
// this message with an empty map
func (s *LatticeControllerSender) SetRegistryCredentials(ctx *actor.Context, arg SetRegistryCredentialsRequest) error {

            	    var sizer msgpack.Sizer
            	    size_enc := &sizer
            	    arg.MEncode(size_enc) 
            	    buf := make([]byte, sizer.Len())
            	    
            	    var encoder = msgpack.NewEncoder(buf)
            	    enc := &encoder
                    arg.MEncode(enc)
            	
s.transport.Send(ctx, actor.Message{ Method: "LatticeController.SetRegistryCredentials", Arg:buf })
return nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.5.0
