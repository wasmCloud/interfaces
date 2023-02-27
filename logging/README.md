[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-logging.svg)](https://crates.io/crates/wasmcloud-interface-logging)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=logging%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/logging/tinygo)
# wasmCloud Builtin Logging Interface
This interface defines the wasmCloud built-in logging interface that comes with each of our supported host runtimes. Actors that use this interface must have the capability contract `wasmcloud:builtin:logging` in their claims list (`wash claims sign --logging`).

## Capability Provider Implementations

There are no external implementations for this provider as they are built directly into the host runtime.

## Example Usage 

### 🦀 Rust

Logging at all available levels:
```rust
use wasmbus_rpc::actor::prelude::RpcResult;
use wasmcloud_interface_logging::{debug, error, info, warn};

// Note: The function you're logging in _must_ be async. This is due to the
// way our logging macros work and is a known limitation of actor logging
async fn log_to_all() -> RpcResult<()> {
    debug!("Watch out for moths");
    info!("This is an info level log!");
    warn!("Some viewers may find the following log disturbing");
    error!("I can't let you do that, Dave");

    Ok(())
}
```

## 🐭 Golang

```go
import (
	"github.com/wasmcloud/actor-tinygo"
	httpserver "github.com/wasmcloud/interfaces/httpserver/tinygo"
	logging "github.com/wasmcloud/interfaces/logging/tinygo"
)

type Actor struct {
	logger *logging.LoggingSender
}

func main() {
	me := Actor{
		logger: logging.NewProviderLogging(),
	}
 
    // The calls to WriteLog require an actor.Context that is
    // accessed from any invocation 
    actor.RegisterHandlers(httpserver.HttpServerHandler(&me))
}

func (e *Actor) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {    
    _ = e.logger.WriteLog(ctx, logging.LogEntry{Level: "debug", Text: "This is a debug log"})
    _ = e.logger.WriteLog(ctx, logging.LogEntry{Level: "info",  Text: "This is an info log"})
    _ = e.logger.WriteLog(ctx, logging.LogEntry{Level: "warn",  Text: "This is a warn log"})
    _ = e.logger.WriteLog(ctx, logging.LogEntry{Level: "error", Text: "This is an error log"})
 
    return nil, nil
}
```
