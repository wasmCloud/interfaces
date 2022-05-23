#!/bin/sh

# Regenerate TinyGo
# requires wash >=0.11

# Usage:  (from inside 'interfaces' folder)
#   build/update-tinygo.sh

set -e
set -x

version=latest

wash gen -c ./codegen-go.toml

here=$PWD
for i in */tinygo; do
  [ "$i" == "core/tinygo" ] && continue
  echo Updating $i
  cd $here/$i
  go get -u github.com/wasmcloud/actor-tinygo@$version
  go mod tidy
  go build
done
