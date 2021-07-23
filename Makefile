# top-level makefile for wasmcloud/interfaces

project_dir = $(abspath $(shell pwd))
subdirs     = core httpclient httpserver numbergen

include build/makefiles/interface.mk


# this expression makes the simplifying assumption that all .smithy files
# are inside immediate subdirs. If they can be somewhere else, update the search
interfaces := $(wildcard $(patsubst %,%/*.smithy,$(subdirs)))
# all interfaces are currently published under /idl/org.wasmcloud
published  := $(patsubst %.smithy,docs/idl/org.wasmcloud/%.smithy,$(notdir $(interfaces)))


# use make publish to copy idl files to docs folder
publish: $(published)


docs/idl/org.wasmcloud/%.smithy:  core/%.smithy
	cp -p $< $@
docs/idl/org.wasmcloud/%.smithy:  httpclient/%.smithy
	cp -p $< $@
docs/idl/org.wasmcloud/%.smithy:  httpserver/%.smithy
	cp -p $< $@
docs/idl/org.wasmcloud/%.smithy:  numbergen/%.smithy
	cp -p $< $@

.PHONY: all publish build release clean lint validate
