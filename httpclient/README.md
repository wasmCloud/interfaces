[![crates.io](https://img.shields.io/crates/v/wasmcloud-interface-httpclient.svg)](https://crates.io/crates/wasmcloud-interface-httpclient)&nbsp;
[![TinyGo Version](https://img.shields.io/github/go-mod/go-version/wasmcloud/interfaces?label=TinyGo&filename=httpclient%2Ftinygo%2Fgo.mod)](https://pkg.go.dev/github.com/wasmcloud/interfaces/httpclient/tinygo)
# wasmCloud HTTP Client Interface
This is the interface definition for the interface with the contract ID `wasmcloud:httpclient`.

Actors utilizing this interface can make HTTP requests and receive HTTP responses for processing. Since this is just an interface, and not an actual provider, you will need to check the documentation for individual provider implementations for a list of link definition values supported by that provider.

## Capability Provider Implementations
The following is a list of implementations of the `wasmcloud:httpclient` contract. Feel free to submit a PR adding your implementation if you have a community/open source version.

| Name | Vendor | Description |
| :--- | :---: | :--- |
| [HTTPClient](https://github.com/wasmCloud/capability-providers/tree/main/httpclient) | wasmCloud | wasmCloud implementation of the HTTP Client Provider

## Example Usage 
### 🦀 Rust
Retrieve a random XKCD comic and format the response as an HTML page
```rust
use serde::Deserialize;
use wasmbus_rpc::actor::prelude::*;
use wasmcloud_interface_httpclient::*;
use wasmcloud_interface_httpserver::HttpResponse;
use wasmcloud_interface_numbergen::random_in_range;

const MAX_COMIC_ID: u32 = 2500;

#[derive(Deserialize)]
struct XkcdMetadata {
    title: String,
    img: String,
}

async fn get_comic(ctx: &Context) -> RpcResult<HttpResponse> {
    let comic_num = random_in_range(1, MAX_COMIC_ID).await?;

    // make a request to get the json metadata
    let url = format!("https://xkcd.com/{}/info.0.json", comic_num);
    let client = HttpClientSender::new();
    let resp = client.request(ctx, &HttpRequest::get(&url)).await?;

    if !(200..300).contains(&resp.status_code) {
        return Err(format!("HTTP Request error {}", resp.status_code.to_string(),).into());
    }
    let comic = serde_json::from_slice::<XkcdMetadata>(&resp.body)
        .map_err(|_| "Error deserializing response body")?;
    let html = format!(
        r#"<!DOCTYPE html>
        <html>
        <head>
            <title>Your XKCD random comic</title>
        </head>
        <body>
            <h1>{}</h1>
            <img src="{}"/>
        </body>
        </html>
        "#,
        &comic.title, &comic.img
    );
    let resp = HttpResponse {
        body: html.into_bytes(),
        ..Default::default()
    };
    Ok(resp)
}

```

### 🐭 Golang
Fetch IP address
```go
func GetIpAddress(ctx *actor.Context) ([]byte, error) {
	client := httpclient.NewProviderHttpClient()

	resp, err := client.Request(ctx, httpclient.HttpRequest{
		Method: "GET",
		Url:    "https://ifconfig.io/ip",
		// Body can not be blank due to a bug
		Body: []byte("a"),
	})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
```
