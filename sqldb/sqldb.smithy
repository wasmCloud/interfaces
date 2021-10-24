// sqldb.smithy
// definition of an sql database capability contract
//
// The interface status is *pre-release* and subject to change.
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

metadata package = [{
    namespace: "org.wasmcloud.interface.sqldb",
    crate: "wasmcloud_interface_sqldb",
    py_module: "wasmcloud_interface_sqldb",
}]

namespace org.wasmcloud.interface.sqldb

use org.wasmcloud.model#wasmbus
use org.wasmcloud.model#n
use org.wasmcloud.model#U32
use org.wasmcloud.model#U64
use org.wasmcloud.model#I64

/// SqlDb - SQL Database connections
/// To use this capability, the actor must be linked
/// with the capability contract "wasmcloud:sqldb"
@wasmbus(
    contractId: "wasmcloud:sqldb",
    providerReceive: true )
service SqlDb {
    version: "0.1",
    operations: [ Execute, Fetch ],
}

/// Execute an sql statement
operation Execute {
    input: Query,
    output: ExecuteResult,
}

/// A query is a non-empty string containing an SQL query or statement,
/// in the syntax of the back-end database.
@length(min:1)
string Query

/// Result of an Execute operation
structure ExecuteResult {
    /// the number of rows affected by the query
    @required
    @n(0)
    rowsAffected: U64,

    /// optional error information.
    /// If error is included in the FetchResult, other values should be ignored.
    @n(1)
    error: SqlDbError,
}

/// Perform select query on database, returning all result rows
operation Fetch {
    input: Query,
    output: FetchResult
}


/// Result of a fetch query
structure FetchResult {
    /// number of rows returned
    @required
    @n(0)
    numRows: U64,

    /// description of columns returned
    @required
    @n(1)
    columns: Columns,

    /// result rows, encoded in CBOR as
    /// an array (rows) of arrays (fields per row)
    @required
    @n(2)
    rows: Blob,

    /// optional error information.
    /// If error is included in the FetchResult, other values should be ignored.
    @n(3)
    error: SqlDbError,
}

/// Detailed error information from the previous operation
structure SqlDbError {

    /// Type of error.
    /// The list of enum variants for this field may be expanded in the future
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
    @n(0)
    code: String,

    /// error message
    @required
    @n(1)
    message: String,
}


/// List of columns in the result set returned by a Fetch operation
list Columns {
    member: Column
}

/// Metadata about a Column in the result set
structure Column {
    /// column ordinal
    @required
    @n(0)
    ordinal: U32,

    /// Column name in the result
    @required
    @n(1)
    name: String,

    /// column data type as reported by the database
    @required
    @n(2)
    dbType: String,
}


