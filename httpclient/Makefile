# interface/httpclient/Makefile

export project_dir=$(abspath $(shell pwd))

include ../build/makefiles/interface.mk

all:: build

build release clean::
	$(MAKE) -C rust $@

.PHONY: all build release clean
