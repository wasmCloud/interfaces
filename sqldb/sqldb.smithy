// sqldb.smithy
// definition of an sql database capability contract
//
// Version 0.1 of this interface has the following features:
//    Execute       - Execute sql operations (insert, update, create table, etc.)
//                    Returns number of rows affected
//    Fetch         - Select 0 or more rows from database
//                    The returned result set is encoded in CBOR,
//                    a language-neutral compact representation.
//
// CBOR is designed to be an extensible format that is language-neutral,
// about 50-70% denser than JSON, and suitable for constrained
// environments (low cpu and memory requirements). Parsers are simple to
// write, and libraries are availble in [several languages](https://cbor.io).
//
// Not currently supported:
// - transactions
// - batch operations (multiple execute or fetch queries in single rpc call)
// - streaming results
// - prepared statements
// - results with NULL values, array column types, or custom column data types
//

metadata package = [
    {
        namespace: "org.wasmcloud.interface.sqldb",
        crate: "wasmcloud-interface-sqldb"
     }
]

namespace org.wasmcloud.interface.sqldb

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64
use org.wasmcloud.model#I64
//use org.wasmcloud.model#codegenRust

/// SqlDb - SQL Database connections
/// To use this capability, the actor must be linked
/// with "wasmcloud:sqldb"
@wasmbus(
    contractId: "wasmcloud:sqldb",
    providerReceive: true )
service SqlDb {
    version: "0.1",
    operations: [ Execute, Fetch ],
}

operation Execute {
    input: Query,
    output: ExecuteResult,
}

@length(min:1)
string Query

structure ExecuteResult {
    /// optional error information.
    /// If error is included in the FetchResult, other values should be ignored.
    error: SqlDbError,

    /// number of rows affected by the query
    @required
    rowsAffected: U64,
}

/// perform select query on database, returning all result rows
operation Fetch {
    input: Query,
    output: FetchResult
}


/// Result of a fetch query
structure FetchResult {
    /// optional error information.
    /// If error is included in the FetchResult, other values should be ignored.
    error: SqlDbError,

    /// number of rows returned
    @required
    numRows: U64,

    /// description of columns returned
    @required
    columns: Columns,

    /// result rows
    @required
    rows: Blob,
}

structure SqlDbError {

    /// Type of error.
    /// The list of error codes below may be expanded in the future
    /// to provide finer-granularity failure information
    @enum([
    { "name": "config",
      "description": "error parsing the connection string or other configuration parameters" },
    { "name": "db",
      "description": "error returned from the database backend" },
    { "name": "io",
      "description": "error communicating with the database backend" },
    { "name": "notFound",
      "description": "No rows returned by a query expected to return at least one row" },
    { "name": "encoding",
      "description": "error encountered encoding result data" },
    { "name": "decoding",
      "description": "error encountered decoding result data" },
    { "name": "provider",
      "description": "the capability provider had an internal error" },
    { "name": "other",
      "description": "some other error that could not be categorized as one of the above" },
      ])
    @required
    code: String,

    /// error message
    @required
    message: String,
}


/// list of columns provided in a result set
list Columns {
    member: Column
}

/// Columns in result set
structure Column {
    /// column ordinal
    @required
    ordinal: U32,

    /// Column name in the result
    @required
    name: String,

    /// Data type of the column
    @required
    ty: String,
}


