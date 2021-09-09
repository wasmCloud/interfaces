# top-level makefile for wasmcloud/interfaces

project_dir = $(abspath $(shell pwd))
# Reminder: when adding a subdir below don't forget to update
# Cargo.toml (if it's a rust project), and codegen.toml (if there are smithy files)
subdirs     = core control factorial httpclient httpserver \
			  keyvalue logging messaging numbergen testing

include build/makefiles/interface.mk

# Find all .smithy files in subdirectories
# this expression makes the simplifying assumption that all .smithy files
# are inside immediate subdirs. If they can be somewhere else, update the search
interfaces := $(wildcard $(patsubst %,%/*.smithy,$(subdirs)))
# all interfaces are currently published under docs/idl/org.wasmcloud
published  := $(patsubst %.smithy,docs/idl/org.wasmcloud/%.smithy,$(notdir $(interfaces)))

# use make publish to copy idl files to docs folder
# TODO: add `wash gen` as part of publish (waiting on wash #163)
publish: $(published)

# generate rules for copying interfaces to github pages dir (/docs)
define pub_template
docs/idl/org.wasmcloud/$$(notdir $(1)):  $(1) FORCE
	cp -p $$< $$@
endef
$(foreach iface,$(interfaces),$(eval $(call pub_template,$(iface))))

FORCE:

.PHONY: publish 
