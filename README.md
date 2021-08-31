# wasmCloud API Interfaces
This repository contains the wasmCloud contract interface definitions (defined in the _Smithy_ IDL) for those interfaces that are defined and supported by the wasmCloud team. These interfaces are definitely not the _only_ interfaces available, as teams and companies can create their own private or bespoke interfaces as desired.

## Smithy IDLs and Shared Libraries

Each interface is defined in a file with the `.smithy` extension. If
the folder contains a `codegen.toml` file,
a library and/or html documentation can be
automatically generated from the `.smithy` file. 

More information on code
generation and the `codegen.toml` files is in the [weld
crate](https://github.com/wasmcloud/weld)

The `docs` folder in this repository is published to github pages at
[Wasmcloud Interfaces](https://wasmcloud.github.io/interfaces/), and
contains copies of the interfaces available for direct download, and
html generated documentation.

For more information on Smithy, see
 - [Smithy](https://awslabs.github.io/smithy/index.html) A language for
  defining services and SDKs
 - [IDL specification](https://awslabs.github.io/smithy/1.0/spec/core/idl.html)

For more on wasmcloud, see
 - [wasmCloud](https://wasmcloud.dev)

