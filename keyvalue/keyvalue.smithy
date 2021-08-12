// key-value.smithy
// Definition of a key-value store and the 'wasmcloud:keyvalue' capability contract
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [ { namespace: "org.wasmcloud.interface.keyvalue", crate: "wasmcloud-interface-keyvalue" } ]

namespace org.wasmcloud.interface.keyvalue

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#I32

@wasmbus(
    contractId: "wasmcloud:keyvalue",
    providerReceive: true,
)
service KeyValue {
  version: "0.1.1",
  operations: [
    Increment, Contains, Del, Get,
    ListAdd, ListClear, ListDel, ListRange,
    Set, , SetAdd, SetDel, SetIntersection, SetQuery, SetUnion, SetClear,
  ]
}

/// Gets a value for a specified key. If the key exists,
/// the return structure contains exists: true and the value,
/// otherwise the return structure contains exists == false.
@readonly
operation Get {
  input: String,
  output: GetResponse,
}

/// Response to get request
structure GetResponse {
    /// the value, if it existed
    @required
    value: String,
    /// whether or not the value existed
    @required
    exists: Boolean,
}

/// Sets the value of a key.
/// expires is an optional number of seconds before the value should be automatically deleted,
/// or 0 for no expiration.
operation Set {
  input: SetRequest,
}

structure SetRequest {
    /// the key name to change (or create)
    @required
    key: String,

    /// the new value
    @required
    value: String,

    /// expiration time in seconds 0 for no expiration
    @required
    expires: U32,
}

/// Deletes a key, returning true if the key was deleted
operation Del {
  input: String,
  output: Boolean,
}

/// Increments a numeric value, returning the new value
operation Increment {
  input: IncrementRequest,
  output: I32
}

structure IncrementRequest {
    /// name of value to increment
  @required
  key: String,
  /// amount to add to value
  @required
  value: I32,
}

/// list of strings
list StringList {
  member: String
}

/// Append a value onto the end of a list. Returns the new list size
operation ListAdd {
  input: ListAddRequest,
  output: U32
}

/// Parameter to ListAdd operation
structure ListAddRequest {
    /// name of the list to modify
    @required
    listName: String,

    /// value to append to the list
    @required
    value: String,
}

/// Deletes a value from a list. Returns true if the item was removed.
operation ListDel{
  input: ListDelRequest,
  output: Boolean
}

/// Removes an item from the list. If the item occurred more than once,
/// removes only the first item.
/// Returns true if the item was found.
structure ListDelRequest {
    /// name of list to modify
  @required
  listName: String,
  @required
  value: String
}

/// Deletes a list and its contents
/// input: list name
/// output: true if the list existed and was deleted
operation ListClear {
  input: String,
  output: Boolean
}

/// Retrieves a range of values from a list using 0-based indices.
/// Start and end values are inclusive, for example, (0,10) returns
/// 11 items if the list contains at least 11 items. If the stop value
/// is beyond the end of the list, it is treated as the end of the list.
operation ListRange {
    input: ListRangeRequest,
    output: StringList,
}

structure ListRangeRequest {

    /// name of list
    @required
    list_name: String,

    /// start index of the range, 0-based, inclusive.
    @required
    start: I32,

    /// end index of the range, 0-based, inclusive.
    @required
    stop: I32,
}

/// Add an item into a set. Returns number of items added (1 or 0)
operation SetAdd {
  input: SetAddRequest,
  output: U32,
}

structure SetAddRequest {
    /// name of the set
    @required
    setName: String,
    /// value to add to the set
    @required
    value: String,
}

/// Deletes an item from the set. Returns number of items removed from the set (1 or 0)
operation SetDel {
  input: SetDelRequest,
  output: U32,
}

structure SetDelRequest {
  @required
  setName: String,
  @required
  value: String,
}

/// perform union of sets and returns values from the union
/// input: list of sets for performing union (at least two)
/// output: union of values
operation SetUnion {
  input: StringList,
  output: StringList,
}

/// perform intersection of sets and returns values from the intersection.
/// input: list of sets for performing intersection (at least two)
/// output: values
operation SetIntersection {
  input: StringList,
  output: StringList,
}

/// Retrieves all items from a set
/// input: String
/// output: set members
operation SetQuery {
  input: String,
  output: StringList,
}

/// returns whether the store contains the key
@readonly
operation Contains {
  input: String,
  output: Boolean,
}

/// clears all values from the set and removes it
/// input: set name
/// output: true if the set existed and was deleted
operation SetClear {
    input: String
    output: Boolean
}

