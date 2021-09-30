//! This library provides the Rust implementation
//! of the wasmcloud SQL database capability contract wasmcloud:sqldb.
//!
//! The initial version of this interface supports
//! executing sql queries (inserts, update, create table, etc.)
//! and fetching data (select).
//!
//! The api is intended to be independent of any specific relational database implementation
//! (postgres, mysql, mariadb, sqlite, etc.).
//!
//! For efficiency, query results are encoded in Compact Binary Object
//! Representation [CBOR](https://cbor.io), a language-neutral format.
//! CBOR is designed to be an extensible,  language-neutral,
//! about 50-70% denser than JSON, and suitable for constrained
//! environments (low cpu and memory requirements). Parsers are simple to
//! write, and libraries are available in [several languages](https://cbor.io/impls.html).
//!
//! This interface currently does not support:
//! - transactions
//! - streaming results
//! - prepared statements
//!

mod sqldb;
pub use sqldb::*;
// re-export minicbor
pub use minicbor;

impl SqlDbError {
    pub fn new<T: ToString>(code: T, message: String) -> SqlDbError {
        SqlDbError {
            code: code.to_string(),
            message,
        }
    }
}

impl std::fmt::Display for SqlDbError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "SqlDbError {}: {}", &self.code, &self.message)
    }
}


impl From<minicbor::decode::Error> for SqlDbError {
    fn from(e: minicbor::decode::Error) -> SqlDbError {
        SqlDbError{ code: "decoding".to_string(), message: e.to_string() }
    }
}

use wasmbus_rpc::RpcError;
impl From<SqlDbError> for RpcError {
    fn from(e: SqlDbError) -> RpcError {
        RpcError::Other(format!("SqlDb error {}: {}", e.code, e.message))
    }
}

impl From<RpcError> for SqlDbError {
    fn from(e: RpcError) -> SqlDbError {
        SqlDbError::new("rpc", e.to_string())
    }
}
