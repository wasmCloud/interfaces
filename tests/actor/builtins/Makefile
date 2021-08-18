# interfaces/tests/builtins/Makefile

KEYDIR ?= .keys
CARGO ?= cargo
REVISION := 1
RUST_DEPS := Cargo.toml Makefile $(wildcard src/*.rs) 
VERSION := $(shell cargo metadata --no-deps --format-version 1 | jq -r '.packages[0] .version')
# project is a short slug name that will be used in the final file name
PROJECT := builtins_test_actor
WASH := /home/steve/projects/wasmcloud/wash/target/debug/wash

# command to push to registry (without trailing file path)
REG_URL = localhost:5000/v2/$(PROJECT):$(VERSION)
PUSH_REG_CMD := $(WASH) reg push --insecure $(REG_URL)

# list of all contract claims for actor signing (space-separated)
CLAIMS  := wasmcloud:builtin:numbergen wasmcloud:builtin:logging
DIST_WASM := build/$(PROJECT)_s.wasm
TARGET := ../target/wasm32-unknown-unknown

all: release $(DIST_WASM) $(DIST_EXS)

release:
		cargo build --release

$(TARGET)/release/$(PROJECT).wasm: $(RUST_DEPS) Makefile
	$(CARGO) build --release

# sign it
$(DIST_WASM): $(TARGET)/release/$(PROJECT).wasm Makefile
	@mkdir -p $(dir $@)
	$(WASH) claims sign $< \
		$(foreach claim,$(CLAIMS), -c $(claim) ) \
		--name "$(PROJECT)" --ver $(VERSION) --rev $(REVISION) \
		--destination $@

# push to repo
push:
	$(PUSH_REG_CMD) $(DIST_WASM)

# start it
start:
	$(WASH) ctl start actor $(REG_URL) --timeout 3

# update it (should update revision before doing this)
update:
	$(PUSH_REG_CMD) $(DIST_WASM)
	$(WASH) ctl update actor  \
        $(shell $(WASH) ctl get hosts -o json | jq -r ".hosts[0].id") \
	    $(shell make actor_id | tail -1) \
	    $(REG_URL) --timeout 3

test:
	$(WASH) call --test --data test-options.json --rpc-timeout 2 \
	    $(shell make actor_id | tail -1) \
	    Start

check clippy doc:
	$(CARGO) $@

clean:
	$(CARGO) clean
	rm -rf build

claims: $(DIST_WASM)
	$(WASH) claims inspect $(DIST_WASM)

actor_id: $(DIST_WASM)
	$(WASH) claims inspect $(DIST_WASM) -o json | jq -r .module

.PHONY: actor_id all check clean clippy doc release test update
