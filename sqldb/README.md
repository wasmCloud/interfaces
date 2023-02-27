[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-sqldb.svg)](https://crates.io/crates/wasmcloud-interface-sqldb)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=sqldb%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/sqldb/tinygo)
# wasmCloud SQL Database Interface
This interface defines a basic SQL Database 
provider with the capability contract `wasmcloud:sqldb`.

The initial version of this interface (0.1) supports
executing sql queries (inserts, update, create table, etc.)
and fetching data (select).

The api is intended to be independent of any specific relational database implementation
(postgres, mysql, mariadb, sqlite, etc.).

For efficiency, query results are encoded in Compact Binary Object
Representation [CBOR](https://cbor.io), a language-neutral format.
CBOR is designed to be an extensible,  language-neutral,
about 50-70% denser than JSON, and suitable for constrained
environments (low cpu and memory requirements). Parsers are simple to
write, and libraries are available in [several languages](https://cbor.io/impls.html).

This interface is **pre-release and subject to change**.
The following features are currently unsupported:
- nullable fields
- transactions
- prepared statements
- streaming results

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:sqldb` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [sqldb-postgres](https://github.com/wasmCloud/capability-providers/tree/main/sqldb-postgres) | wasmCloud | Implementation of the sqldb contract to interface with Postgres-compatible databases (also works for Azure CosmosDB with a Postgres backend, for example)

## Example Usage 
### 🦀 Rust
The following examples were pulled from the [Todo-sql example actor](https://github.com/wasmCloud/examples/tree/main/actor/todo-sql).
Create a table to store `DbTodo` objects for a TODO list:
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_sqldb::{minicbor, SqlDb, SqlDbError, SqlDbSender};
use minicbor::{decode, Decode, Encode};
#[derive(Encode, Decode)]
struct DbTodo {
    #[n(0)]
    url: String,
    #[n(1)]
    title: String,
    #[n(2)]
    completed: bool,
    #[n(3)]
    priority: i32,
}
/// create an empty table with the proper schema
async fn create_table(ctx: &Context) -> Result<(), SqlDbError> {
    let db = SqlDbSender::new();
    let sql = format!(
        r#"create table if not exists {} (
            id varchar(36) not null,
            url varchar(42) not null,
            title varchar(100) not null,
            priority int4 not null default 0,
            completed bool not null default false
        );"#,
        TABLE_NAME
    );
    let _resp = db.execute(ctx, sql.into()).await?;
    Ok(())
}
```

Fetching a `DbTodo` from a database
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_sqldb::{minicbor, SqlDb, SqlDbError, SqlDbSender};
use wasmcloud_interface_logging::info;
use minicbor::{decode, Decode, Encode};
async fn get_db_todo(ctx: &Context, url: &str) -> Result<DbTodo, SqlDbError> {
    info!("Getting a todo...");
    let db = SqlDbSender::new();
    check_safety("url", url)?;
    let resp = db
        .fetch(
            ctx,
            &format!(
                "select url, title, completed, priority from {} where url='{}'",
                TABLE_NAME, url
            ),
        )
        .await?;
    if resp.num_rows == 0 {
        return Err(SqlDbError::new("notFound", "url not found".to_string()));
    }
    let mut rows: Vec<DbTodo> = decode(&resp.rows)?;
    let db_todo = rows.remove(0);
    Ok(db_todo)
}
```

### 🐭Golang
Create a table to store `DbTodo` objects for a TODO list:
```go
import (
   "github.com/wasmcloud/actor-tinygo"
   sqldb "github.com/wasmcloud/interfaces/sqldb/tinygo"
)

var sql string = `
create table if not exists {} (
    id varchar(36) not null,
    url varchar(42) not null,
    title varchar(100) not null,
    priority int4 not null default 0,
    completed bool not null default false
);
`
func CreateTable(ctx *actor.Context, db string) (*sqldb.ExecuteResult, error) {
   client := sqldb.NewProviderSqlDb()
   return client.Execute(ctx, sqldb.Statement{
      Database: db,
      Sql:      sql,
   })
}
```
