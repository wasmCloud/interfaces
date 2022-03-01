// This file is generated automatically using wasmcloud/weld-codegen 0.4.2

#[allow(unused_imports)]
use async_trait::async_trait;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use std::{borrow::Borrow, borrow::Cow, io::Write, string::ToString};
#[allow(unused_imports)]
use wasmbus_rpc::{
    cbor::*,
    common::{
        deserialize, message_format, serialize, Context, Message, MessageDispatch, MessageFormat,
        SendOpts, Transport,
    },
    error::{RpcError, RpcResult},
    Timestamp,
};

pub const SMITHY_VERSION: &str = "1.0";

/// Metadata about a Column in the result set
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct Column {
    /// column ordinal
    #[serde(default)]
    pub ordinal: u32,
    /// Column name in the result
    #[serde(default)]
    pub name: String,
    /// column data type as reported by the database
    #[serde(rename = "dbType")]
    #[serde(default)]
    pub db_type: String,
}

// Encode Column as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_column<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Column,
) -> RpcResult<()> {
    e.array(3)?;
    e.u32(val.ordinal)?;
    e.str(&val.name)?;
    e.str(&val.db_type)?;
    Ok(())
}

// Decode Column from cbor input stream
#[doc(hidden)]
pub fn decode_column(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Column, RpcError> {
    let __result = {
        let mut ordinal: Option<u32> = None;
        let mut name: Option<String> = None;
        let mut db_type: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct Column, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct Column: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => ordinal = Some(d.u32()?),
                    1 => name = Some(d.str()?.to_string()),
                    2 => db_type = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser("decoding struct Column: indefinite map not supported".to_string())
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "ordinal" => ordinal = Some(d.u32()?),
                    "name" => name = Some(d.str()?.to_string()),
                    "dbType" => db_type = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        Column {
            ordinal: if let Some(__x) = ordinal {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Column.ordinal (#0)".to_string(),
                ));
            },

            name: if let Some(__x) = name {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Column.name (#1)".to_string(),
                ));
            },

            db_type: if let Some(__x) = db_type {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Column.db_type (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// List of columns in the result set returned by a Query operation
pub type Columns = Vec<Column>;

// Encode Columns as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_columns<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Columns,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        encode_column(e, item)?;
    }
    Ok(())
}

// Decode Columns from cbor input stream
#[doc(hidden)]
pub fn decode_columns(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Columns, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<Column> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(decode_column(d).map_err(|e| format!("decoding 'Column': {}", e))?)
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<Column> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => {
                        arr.push(decode_column(d).map_err(|e| format!("decoding 'Column': {}", e))?)
                    }
                }
            }
            arr
        }
    };
    Ok(__result)
}
/// Result of an Execute operation
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct ExecuteResult {
    /// the number of rows affected by the query
    #[serde(rename = "rowsAffected")]
    #[serde(default)]
    pub rows_affected: u64,
    /// optional error information.
    /// If error is included in the QueryResult, other values should be ignored.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub error: Option<SqlDbError>,
}

// Encode ExecuteResult as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_execute_result<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &ExecuteResult,
) -> RpcResult<()> {
    e.array(2)?;
    e.u64(val.rows_affected)?;
    if let Some(val) = val.error.as_ref() {
        encode_sql_db_error(e, val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode ExecuteResult from cbor input stream
#[doc(hidden)]
pub fn decode_execute_result(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<ExecuteResult, RpcError> {
    let __result = {
        let mut rows_affected: Option<u64> = None;
        let mut error: Option<Option<SqlDbError>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct ExecuteResult, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ExecuteResult: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => rows_affected = Some(d.u64()?),
                    1 => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct ExecuteResult: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "rowsAffected" => rows_affected = Some(d.u64()?),
                    "error" => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        ExecuteResult {
            rows_affected: if let Some(__x) = rows_affected {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field ExecuteResult.rows_affected (#0)".to_string(),
                ));
            },
            error: error.unwrap(),
        }
    };
    Ok(__result)
}
/// An optional list of arguments to be used in the SQL statement.
/// When a statement uses question marks '?' for placeholders,
/// the capability provider will replace the specified arguments during execution.
/// The command must have exactly as many placeholders as arguments, or the request will fail.
/// The members are CBOR encoded.
pub type Parameters = Vec<Vec<u8>>;

// Encode Parameters as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_parameters<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Parameters,
) -> RpcResult<()> {
    e.array(val.len() as u64)?;
    for item in val.iter() {
        e.bytes(item)?;
    }
    Ok(())
}

// Decode Parameters from cbor input stream
#[doc(hidden)]
pub fn decode_parameters(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Parameters, RpcError> {
    let __result = {
        if let Some(n) = d.array()? {
            let mut arr: Vec<Vec<u8>> = Vec::with_capacity(n as usize);
            for _ in 0..(n as usize) {
                arr.push(d.bytes()?.to_vec())
            }
            arr
        } else {
            // indefinite array
            let mut arr: Vec<Vec<u8>> = Vec::new();
            loop {
                match d.datatype() {
                    Err(_) => break,
                    Ok(wasmbus_rpc::cbor::Type::Break) => break,
                    Ok(_) => arr.push(d.bytes()?.to_vec()),
                }
            }
            arr
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct PingResult {
    /// Optional error information.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub error: Option<SqlDbError>,
}

// Encode PingResult as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_ping_result<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &PingResult,
) -> RpcResult<()> {
    e.map(1)?;
    if let Some(val) = val.error.as_ref() {
        e.str("error")?;
        encode_sql_db_error(e, val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode PingResult from cbor input stream
#[doc(hidden)]
pub fn decode_ping_result(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<PingResult, RpcError> {
    let __result = {
        let mut error: Option<Option<SqlDbError>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct PingResult, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PingResult: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct PingResult: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "error" => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        PingResult {
            error: error.unwrap(),
        }
    };
    Ok(__result)
}
/// Result of a query
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct QueryResult {
    /// number of rows returned
    #[serde(rename = "numRows")]
    #[serde(default)]
    pub num_rows: u64,
    /// description of columns returned
    pub columns: Columns,
    /// result rows, encoded in CBOR as
    /// an array (rows) of arrays (fields per row)
    #[serde(with = "serde_bytes")]
    #[serde(default)]
    pub rows: Vec<u8>,
    /// optional error information.
    /// If error is included in the QueryResult, other values should be ignored.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub error: Option<SqlDbError>,
}

// Encode QueryResult as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_query_result<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &QueryResult,
) -> RpcResult<()> {
    e.array(4)?;
    e.u64(val.num_rows)?;
    encode_columns(e, &val.columns)?;
    e.bytes(&val.rows)?;
    if let Some(val) = val.error.as_ref() {
        encode_sql_db_error(e, val)?;
    } else {
        e.null()?;
    }
    Ok(())
}

// Decode QueryResult from cbor input stream
#[doc(hidden)]
pub fn decode_query_result(
    d: &mut wasmbus_rpc::cbor::Decoder<'_>,
) -> Result<QueryResult, RpcError> {
    let __result = {
        let mut num_rows: Option<u64> = None;
        let mut columns: Option<Columns> = None;
        let mut rows: Option<Vec<u8>> = None;
        let mut error: Option<Option<SqlDbError>> = Some(None);

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct QueryResult, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct QueryResult: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => num_rows = Some(d.u64()?),
                    1 => {
                        columns = Some(
                            decode_columns(d).map_err(|e| format!("decoding 'Columns': {}", e))?,
                        )
                    }
                    2 => rows = Some(d.bytes()?.to_vec()),
                    3 => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }

                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct QueryResult: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "numRows" => num_rows = Some(d.u64()?),
                    "columns" => {
                        columns = Some(
                            decode_columns(d).map_err(|e| format!("decoding 'Columns': {}", e))?,
                        )
                    }
                    "rows" => rows = Some(d.bytes()?.to_vec()),
                    "error" => {
                        error = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_sql_db_error(d)
                                    .map_err(|e| format!("decoding 'SqlDbError': {}", e))?,
                            ))
                        }
                    }
                    _ => d.skip()?,
                }
            }
        }
        QueryResult {
            num_rows: if let Some(__x) = num_rows {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field QueryResult.num_rows (#0)".to_string(),
                ));
            },

            columns: if let Some(__x) = columns {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field QueryResult.columns (#1)".to_string(),
                ));
            },

            rows: if let Some(__x) = rows {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field QueryResult.rows (#2)".to_string(),
                ));
            },
            error: error.unwrap(),
        }
    };
    Ok(__result)
}
/// Detailed error information from the previous operation
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct SqlDbError {
    /// Type of error.
    /// The list of enum variants for this field may be expanded in the future
    /// to provide finer-granularity failure information
    #[serde(default)]
    pub code: String,
    /// error message
    #[serde(default)]
    pub message: String,
}

// Encode SqlDbError as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_sql_db_error<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &SqlDbError,
) -> RpcResult<()> {
    e.array(2)?;
    e.str(&val.code)?;
    e.str(&val.message)?;
    Ok(())
}

// Decode SqlDbError from cbor input stream
#[doc(hidden)]
pub fn decode_sql_db_error(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<SqlDbError, RpcError> {
    let __result = {
        let mut code: Option<String> = None;
        let mut message: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct SqlDbError, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SqlDbError: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => code = Some(d.str()?.to_string()),
                    1 => message = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct SqlDbError: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "code" => code = Some(d.str()?.to_string()),
                    "message" => message = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        SqlDbError {
            code: if let Some(__x) = code {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SqlDbError.code (#0)".to_string(),
                ));
            },

            message: if let Some(__x) = message {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field SqlDbError.message (#1)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
#[derive(Clone, Debug, Default, Deserialize, Eq, PartialEq, Serialize)]
pub struct Statement {
    /// Optional database in which the statement must be executed.
    /// The value in this field is case-sensitive.
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub database: Option<String>,
    #[serde(default, skip_serializing_if = "Option::is_none")]
    pub parameters: Option<Parameters>,
    /// A sql query or statement that is a non-empty string containing
    /// in the syntax of the back-end database.
    #[serde(default)]
    pub sql: String,
}

// Encode Statement as CBOR and append to output stream
#[doc(hidden)]
pub fn encode_statement<W: wasmbus_rpc::cbor::Write>(
    e: &mut wasmbus_rpc::cbor::Encoder<W>,
    val: &Statement,
) -> RpcResult<()> {
    e.map(3)?;
    if let Some(val) = val.database.as_ref() {
        e.str("database")?;
        e.str(val)?;
    } else {
        e.null()?;
    }
    if let Some(val) = val.parameters.as_ref() {
        e.str("parameters")?;
        encode_parameters(e, val)?;
    } else {
        e.null()?;
    }
    e.str("sql")?;
    e.str(&val.sql)?;
    Ok(())
}

// Decode Statement from cbor input stream
#[doc(hidden)]
pub fn decode_statement(d: &mut wasmbus_rpc::cbor::Decoder<'_>) -> Result<Statement, RpcError> {
    let __result = {
        let mut database: Option<Option<String>> = Some(None);
        let mut parameters: Option<Option<Parameters>> = Some(None);
        let mut sql: Option<String> = None;

        let is_array = match d.datatype()? {
            wasmbus_rpc::cbor::Type::Array => true,
            wasmbus_rpc::cbor::Type::Map => false,
            _ => {
                return Err(RpcError::Deser(
                    "decoding struct Statement, expected array or map".to_string(),
                ))
            }
        };
        if is_array {
            let len = d.array()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct Statement: indefinite array not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match __i {
                    0 => {
                        database = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    1 => {
                        parameters = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_parameters(d)
                                    .map_err(|e| format!("decoding 'Parameters': {}", e))?,
                            ))
                        }
                    }
                    2 => sql = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        } else {
            let len = d.map()?.ok_or_else(|| {
                RpcError::Deser(
                    "decoding struct Statement: indefinite map not supported".to_string(),
                )
            })?;
            for __i in 0..(len as usize) {
                match d.str()? {
                    "database" => {
                        database = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(d.str()?.to_string()))
                        }
                    }
                    "parameters" => {
                        parameters = if wasmbus_rpc::cbor::Type::Null == d.datatype()? {
                            d.skip()?;
                            Some(None)
                        } else {
                            Some(Some(
                                decode_parameters(d)
                                    .map_err(|e| format!("decoding 'Parameters': {}", e))?,
                            ))
                        }
                    }
                    "sql" => sql = Some(d.str()?.to_string()),
                    _ => d.skip()?,
                }
            }
        }
        Statement {
            database: database.unwrap(),
            parameters: parameters.unwrap(),

            sql: if let Some(__x) = sql {
                __x
            } else {
                return Err(RpcError::Deser(
                    "missing field Statement.sql (#2)".to_string(),
                ));
            },
        }
    };
    Ok(__result)
}
/// SqlDb - SQL Database connections
/// To use this capability, the actor must be linked
/// with the capability contract "wasmcloud:sqldb"
/// wasmbus.contractId: wasmcloud:sqldb
/// wasmbus.providerReceive
#[async_trait]
pub trait SqlDb {
    /// returns the capability contract id for this interface
    fn contract_id() -> &'static str {
        "wasmcloud:sqldb"
    }
    /// Execute an sql statement
    async fn execute(&self, ctx: &Context, arg: &Statement) -> RpcResult<ExecuteResult>;
    /// Perform select query on database, returning all result rows
    async fn query(&self, ctx: &Context, arg: &Statement) -> RpcResult<QueryResult>;
}

/// SqlDbReceiver receives messages defined in the SqlDb service trait
/// SqlDb - SQL Database connections
/// To use this capability, the actor must be linked
/// with the capability contract "wasmcloud:sqldb"
#[doc(hidden)]
#[async_trait]
pub trait SqlDbReceiver: MessageDispatch + SqlDb {
    async fn dispatch<'disp__, 'ctx__, 'msg__>(
        &'disp__ self,
        ctx: &'ctx__ Context,
        message: &Message<'msg__>,
    ) -> Result<Message<'msg__>, RpcError> {
        match message.method {
            "Execute" => {
                let value: Statement = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'Statement': {}", e)))?;
                let resp = SqlDb::execute(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "SqlDb.Execute",
                    arg: Cow::Owned(buf),
                })
            }
            "Query" => {
                let value: Statement = wasmbus_rpc::common::deserialize(&message.arg)
                    .map_err(|e| RpcError::Deser(format!("'Statement': {}", e)))?;
                let resp = SqlDb::query(self, ctx, &value).await?;
                let buf = wasmbus_rpc::common::serialize(&resp)?;
                Ok(Message {
                    method: "SqlDb.Query",
                    arg: Cow::Owned(buf),
                })
            }
            _ => Err(RpcError::MethodNotHandled(format!(
                "SqlDb::{}",
                message.method
            ))),
        }
    }
}

/// SqlDbSender sends messages to a SqlDb service
/// SqlDb - SQL Database connections
/// To use this capability, the actor must be linked
/// with the capability contract "wasmcloud:sqldb"
/// client for sending SqlDb messages
#[derive(Debug)]
pub struct SqlDbSender<T: Transport> {
    transport: T,
}

impl<T: Transport> SqlDbSender<T> {
    /// Constructs a SqlDbSender with the specified transport
    pub fn via(transport: T) -> Self {
        Self { transport }
    }

    pub fn set_timeout(&self, interval: std::time::Duration) {
        self.transport.set_timeout(interval);
    }
}

#[cfg(target_arch = "wasm32")]
impl SqlDbSender<wasmbus_rpc::actor::prelude::WasmHost> {
    /// Constructs a client for sending to a SqlDb provider
    /// implementing the 'wasmcloud:sqldb' capability contract, with the "default" link
    pub fn new() -> Self {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:sqldb", "default")
                .unwrap();
        Self { transport }
    }

    /// Constructs a client for sending to a SqlDb provider
    /// implementing the 'wasmcloud:sqldb' capability contract, with the specified link name
    pub fn new_with_link(link_name: &str) -> wasmbus_rpc::error::RpcResult<Self> {
        let transport =
            wasmbus_rpc::actor::prelude::WasmHost::to_provider("wasmcloud:sqldb", link_name)?;
        Ok(Self { transport })
    }
}
#[async_trait]
impl<T: Transport + std::marker::Sync + std::marker::Send> SqlDb for SqlDbSender<T> {
    #[allow(unused)]
    /// Execute an sql statement
    async fn execute(&self, ctx: &Context, arg: &Statement) -> RpcResult<ExecuteResult> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "SqlDb.Execute",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: ExecuteResult = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': ExecuteResult", e)))?;
        Ok(value)
    }

    #[allow(unused)]
    /// Perform select query on database, returning all result rows
    async fn query(&self, ctx: &Context, arg: &Statement) -> RpcResult<QueryResult> {
        let buf = wasmbus_rpc::common::serialize(arg)?;
        let resp = self
            .transport
            .send(
                ctx,
                Message {
                    method: "SqlDb.Query",
                    arg: Cow::Borrowed(&buf),
                },
                None,
            )
            .await?;

        let value: QueryResult = wasmbus_rpc::common::deserialize(&resp)
            .map_err(|e| RpcError::Deser(format!("'{}': QueryResult", e)))?;
        Ok(value)
    }
}
