// testing.smithy
// Test service operations
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [{
    namespace: "org.wasmcloud.interface.testing",
    crate: "wasmcloud_interface_testing",
    py_module: "wasmcloud_interface_testing",
}]

namespace org.wasmcloud.interface.testing

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#n
use org.wasmcloud.model#U32
use org.wasmcloud.model#I32
use org.wasmcloud.model#F32
use org.wasmcloud.model#F64
use org.wasmcloud.model#codegenRust

/// Test api for testable actors and providers
@wasmbus(
    contractId: "wasmcloud:testing",
    providerReceive: true,
    actorReceive: true,
    protocol: "2",
)
service Testing {
  version: "0.0.1",
  operations: [
    Start, Foo
  ]
}

/// Begin tests
operation Start {
  input: TestOptions,
  output: TestResults,
}

/// Options passed to all test cases
@codegenRust(noDeriveDefault:true)
structure TestOptions {

    /// List of regex patterns for test names to run
    /// Default is ".*", to run all tests.
    @required
    @n(0)
    patterns: PatternList,

    /// additional test configuration, optional
    /// Keys may be test case names, or other keys meaningful for the test.
    /// Values are serialized json, with contents specific to the test
    @required
    @n(1)
    options: OptMap,
}


/// A map of test options.
/// Keys may be test case names, or other keys meaningful for the test.
/// Values are utf8 strings containing serialized json, with contents specific to the test
map OptMap {
    key: String,
    value: String,
}

/// list of regex patterns
list PatternList {
    member: String,
}

operation Foo {
	output: SampleUnion
}

/// A test of union
@codegenRust(noDeriveEq: true)
union SampleUnion {

    /// first field is a String
    @n(0)
    one: String,

    /// Second field is a TestResult
    @n(1)
    two: TestResult,

    /// Third field is array of f32
    @n(2)
    three: F32Data

    /// Fourth field is array of f64
    @n(3)
    four: F64Data
}

list F32Data {
    member: F32
}

list F64Data {
    member: F64
}

structure TestResult {

    /// test case name
    @required
    @n(0)
    name: String,

    /// true if the test case passed
    @required
    @n(1)
    passed: Boolean,

    /// (optional) more detailed results, if available.
    /// data is snap-compressed json
    /// failed tests should have a firsts-level key called "error".
    @n(2)
    snapData: Blob
}

list TestResults {
    member: TestResult
}
