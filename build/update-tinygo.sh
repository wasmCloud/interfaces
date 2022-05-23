#!/bin/sh

# Regenerate TinyGo
# requires wash >=0.11

# Usage:  (from inside 'interfaces' folder)
#   build/update-tinygo.sh

set -e
set -x

actor_sdk_version=3893510790129936fa561b386eedcec6c71338a0
tinygo_msgpack_version=192ec93ec5c440e685d61c705c8fe68ff4aa1973

WASH=${WASH:-wash}
$WASH gen -c ./codegen-go.toml

here=$PWD
for i in */tinygo; do
  [ "$i" == "core/tinygo" ] && continue
  echo Updating $i
  cd $here/$i
  go get -u github.com/wasmcloud/actor-tinygo@$actor_sdk_version
  go get -u github.com/wasmcloud/tinygo-msgpack@$tinygo_msgpack_version
  go mod tidy
  go build
done
