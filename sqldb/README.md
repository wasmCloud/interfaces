# SQL Database

This interface defines a basic SQL Database 
provider with the capability contract wasmcloud:sqldb.

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

This interface currently does not support:
- transactions
- streaming results
- prepared statements
 