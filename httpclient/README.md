# HTTP Client
This is the interface definition for the interface with the contract ID `wasmcloud:httpclient`.

Actors utilizing this interface can make HTTP requests and receive HTTP responses for processing. Since this is just an interface, and not an actual provider, you will need to check the documentation for individual provider implementations for a list of link definition values supported by that provider.

## Implementations
The following is a list of implementations of the HTTP client contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :---: | :---: | :--- |
| [Default Client](https://github.com/wasmCloud/capability-providers/tree/main/httpclient) | wasmCloud | wasmCloud Default HTTP Client Provider