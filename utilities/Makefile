BASE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SUBDIRS := $(wildcard $(BASE_DIR)/*/.)
TARGETS := all

SUBDIRS_TARGETS := $(foreach target,$(TARGETS),$(addsuffix $(target),$(SUBDIRS)))

.PHONY: $(TARGETS) $(SUBDIRS_TARGETS)

$(TARGETS): % : $(addsuffix %,$(SUBDIRS))
	@echo 'Finished running target "$*" on services'

$(SUBDIRS_TARGETS):
	$(MAKE) -C $(@D) $(@F:.%=%)

.PHONY: test
test:
