#!/bin/sh

# Regenerate TinyGo
# requires wash >=0.11

# Usage:  (from inside 'interfaces' folder)
#   build/update-tinygo.sh

set -e
set -x

actor_sdk_version=v0.1.1
tinygo_msgpack_version=v0.1.4
tinygo_cbor_version=v0.1.0

WASH=${WASH:-wash}
#$WASH gen -c ./codegen-go.toml
../weld/codegen/target/debug/codegen ./codegen-go.toml

here=$PWD
for i in */tinygo; do
  [ "$i" == "core/tinygo" ] && continue
  echo Updating $i
  cd $here/$i
  go get -u github.com/wasmcloud/actor-tinygo@$actor_sdk_version
  go get -u github.com/wasmcloud/tinygo-msgpack@$tinygo_msgpack_version
  go get -u github.com/wasmcloud/tinygo-cbor@$tinygo_cbor_version
  go mod tidy
  go build
done
