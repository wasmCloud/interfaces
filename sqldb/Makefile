# interface/sqldb/Makefile

export project_dir=$(abspath $(shell pwd))
include ../build/makefiles/interface.mk

# cargo test is run in the included make
# also run these tests on 'make test'
test::
	wash lint \
	&& wash validate \
	&& cd rust \
	&& cargo fmt -- --check \
	&& cargo clippy