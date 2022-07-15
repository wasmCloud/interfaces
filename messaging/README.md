[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-messaging.svg)](https://crates.io/crates/wasmcloud-interface-messaging)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=messaging%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/messaging/tinygo)
# wasmCloud Messaging Interface
This is the interface for the `wasmcloud:messaging` contract. This contract is a very simple abstraction over the concept of a message broker. This contract does not have controls or knobs or settings for things like delivery guarantees (e.g. "at most once" vs "at least once"), persistence of messages, etc. Such details are the responsibility of individual providers.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:messaging` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [NATS](https://github.com/wasmCloud/capability-providers/tree/main/nats) | wasmCloud | wasmCloud Messaging Provider for the [NATS](https://nats.io) broker

## Example Usage (🦀 Rust)
Implementing the `Messaging.HandleMessage` operation
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_logging::info;
use wasmcloud_interface_messaging::{MessageSubscriber, MessageSubscriberReceiver, SubMessage};

#[derive(Debug, Default, Actor, HealthResponder)]
#[services(Actor, MessageSubscriber)]
struct LogMessagingActor {}

#[async_trait]
impl MessageSubscriber for LogMessagingActor {
    /// Handle a message received on a subscription
    async fn handle_message(&self, _ctx: &Context, msg: &SubMessage) -> RpcResult<()> {
        info!("Received message: {:?}", msg);
        Ok(())
    }
}
```

Sending a message via a `wasmcloud:messaging` provider without expecting a reply
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_messaging::{Messaging, MessagingSender, PubMessage};
async fn publish_message(ctx: &Context, subject: &str, body: &[u8]) -> RpcResult<()> {
    let provider = MessagingSender::new();
    if let Err(e) = provider
        .publish(
            ctx,
            &PubMessage {
                body: body.to_vec(),
                reply_to: None,
                subject: subject.to_owned(),
            },
        )
        .await
    {
        Err(format!("Could not publish message {}", e.to_string()).into())
    } else {
        Ok(())
    }
}
```

Sending a message via a `wasmcloud:messaging` provider, waiting one second for a reply
```rust
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_messaging::{Messaging, MessagingSender, RequestMessage};
async fn message_request(ctx: &Context, subject: &str, body: &[u8]) -> RpcResult<()> {
    let provider = MessagingSender::new();
    if let Err(e) = provider
        .request(
            ctx,
            &RequestMessage {
                body: body.to_vec(),
                subject: subject.to_owned(),
                timeout_ms: 1_000,
            },
        )
        .await
    {
        Err(format!("Could not request message {}", e.to_string()).into())
    } else {
        Ok(())
    }
}
```