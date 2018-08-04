BASE_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SUBDIRS := $(wildcard $(BASE_DIR)/*/.)
SUBDIRS := $(BASE_DIR)/services/. $(BASE_DIR)/gateway/. $(BASE_DIR)/common/.
TARGETS := all test

SUBDIRS_TARGETS := $(foreach target,$(TARGETS),$(addsuffix $(target),$(SUBDIRS)))

.PHONY: $(TARGETS) $(SUBDIRS_TARGETS)

$(TARGETS): % : $(addsuffix %,$(SUBDIRS))
	@echo 'Finished running target "$*"'

$(SUBDIRS_TARGETS):
	$(MAKE) -C $(@D) $(@F:.%=%)

.PHONY: run
run:
	@$(BASE_DIR)/scripts/run.sh

.PHONY: fmt
fmt:
	go fmt ./...
