# Number Generator
This is the interface definition for the wasmCloud built-in interface that is guaranteed to be supported by all runtime hosts, `wasmcloud:builtin:numbergen`. The number generator interface provides for the creation of things like random numbers, random numbers within a given range, and globally unique identifiers (GUIDs).

## Implementations
There are no external implementations of this provider as all implementations of the `wasmcloud:builtin:numbergen` contract are built directly into the host runtime(s).