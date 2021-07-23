// numbergen.smithy
//
// built-in capability provider for number generation
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [
    {
        namespace: "org.wasmcloud.interface.numbergen",
        crate: "wasmcloud_interface_numbergen"
     }
]

namespace org.wasmcloud.interface.numbergen

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64

@wasmbus(
    contractId: "wasmcloud:builtin:numbergen",
    providerReceive: true )
service NumberGen {
  version: "0.1",
  operations: [ GenerateGuid, RandomInRange, Random32 ]
}

///
/// GenerateGuid - return a 128-bit guid in the form 123e4567-e89b-12d3-a456-426655440000
/// These guids are known as "version 4", meaning all bits are random or pseudo-random.
///
operation GenerateGuid {
    output: String
}


/// Request a random integer within a range
/// The result will will be in the range [min,max), i.e., >= min and < max.
operation RandomInRange {
    input: RangeLimit
    output: U32
}

/// Request a 32-bit random number
operation Random32 {
    output: U32
}

/// Input range for RandomInRange. Result will be >= min and < max
/// Example:
///    random_in_range(RangeLimit{0,5}) returns one the values, 0, 1, 2, 3, or 4.
structure RangeLimit {
    @required
    min: U32,
    @required
    max: U32,
}

