# makefile for interfaces/tests/actor/builtins

PROJECT := builtins_test_actor
REVISION := 0
VERSION := $(shell cargo metadata --no-deps --format-version 1 | jq -r '.packages[0] .version')
WASH = wash

REG_URL = localhost:5000/v2/$(PROJECT):$(VERSION)
PUSH_REG_CMD := $(WASH) reg push --insecure $(REG_URL)

# list of all contract claims for actor signing (space-separated)
CLAIMS  := wasmcloud:builtin:numbergen wasmcloud:builtin:logging

include ../../../build/makefiles/actor.mk

