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

## Example Usage (🦀 Rust)

## Example Usage (<img alt="gopher" src="https://i.imgur.com/fl5JozD.png" height="25px"> TinyGo)