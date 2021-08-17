// testing.smithy
// Test service operations
//

// Tell the code generator how to reference symbols defined in this namespace
metadata package = [ { namespace: "org.wasmcloud.interface.testing", crate: "wasmcloud-interface-testing" } ]

namespace org.wasmcloud.interface.testing

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#I32
use org.wasmcloud.model#codegenRust

/// Test api for testable actors and providers
@wasmbus(
    contractId: "wasmcloud:testing",
    providerReceive: true,
    actorReceive: true,
)
service Testing {
  version: "0.0.1",
  operations: [
    Start
  ]
}

/// Begin tests
operation Start {
  input: TestOptions,
  output: TestResults,
}

/// Options passed to all test cases
@codegenRust(deriveDefault:false)
structure TestOptions {

    /// List of regex patterns for test names to run
    /// Default is ".*", to run all tests.
    @required
    patterns: PatternList,

    /// additional test configuration, optional
    /// Keys may be test case names, or other keys meaningful for the test.
    /// Values are serialized json, with contents specific to the test
    @required
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

structure TestResult {

    /// test case name
    @required
    name: String,

    /// true if the test case passed
    @required
    pass: Boolean,

    /// (optional) more detailed results, if available.
    /// data is snap-compressed json
    /// failed tests should have a firsts-level key called "error".
    snapData: Blob
}

list TestResults {
    member: TestResult
}
