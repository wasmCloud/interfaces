# interface.mak
#
# common rules for building smithy models
# Some of these may depend on GNUMakefile >= 4.0
#

html_target     ?= html
codegen_config  ?= $(project_dir)/codegen.toml

platform_id = $$( uname -s )
platform = $$( \
	case $(platform_id) in \
		( Linux | Darwin | FreeBSD ) echo $(platform_id) ;; \
		( * ) echo Unrecognized Platform;; \
	esac )

# find weld binary: (search order: environment (WELD), top folder debug build, PATH)
ifeq ($(weld),)
	ifeq ($(shell which weld 2>/dev/null),)
		$(error No weld in your PATH. try installing with 'cargo install weld-bin')
	else
		weld:=weld
	endif
endif

all::


clean::
	rm -rf $(html_target)/*.html

# Run smithy model lint or validation checks
lint validate:: $(codegen_config) $(weld)
	$(weld) $@ --config $(codegen_config)


# for debugging - show variables make is using
make-vars:
	@echo "weld:          :$(weld)"
	@echo "codegen_config :$(codegen_config)"
	@echo "platform_id    :$(platform_id)"
	@echo "platform       :$(platform)"


.PHONY: all clean lint validate $(weld)
