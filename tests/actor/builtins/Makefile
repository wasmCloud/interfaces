# interfaces/tests/builtins/Makefile

PROJECT := builtins_test_actor
VERSION := $(shell cargo metadata --no-deps --format-version 1 | jq -r '.packages[0] .version')
REVISION := 0
RUST_DEPS := Cargo.toml Makefile $(wildcard src/*.rs)

# command to push to registry (without trailing file path)
REG_URL = localhost:5000/v2/$(PROJECT):$(VERSION)
PUSH_REG_CMD := wash reg push --insecure $(REG_URL)

# list of all contract claims for actor signing (space-separated)
CLAIMS  := wasmcloud:builtin:numbergen wasmcloud:builtin:logging

include ../../../build/makefiles/actor.mk
