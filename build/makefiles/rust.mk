# common rust rules

build clean clippy test::
	cargo $@

release::
	cargo build --release

.PHONY: build release clean clippy test
